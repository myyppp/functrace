//go:build trace

// 编译时指定 -tags trace 参数
// go build -tags trace
package functrace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

var (
	mu sync.Mutex
	m  = make(map[uint64]int)
)

func printTrace(id uint64, name, typ string, count int) {
	fmt.Printf("g[%02d]:%s%s%s\n", id, strings.Repeat("\t", count), typ, name)
}

func Trace() func() {
	// 获得当前 G 的函数调用栈上的信息
	// pc 程序计数器
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	id := getGID()
	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	mu.Lock()
	cnt := m[id]
	m[id] = cnt + 1 // 获取当前 gid 对应的缩进层次，并保存
	mu.Unlock()
	printTrace(id, name, "->", cnt+1)

	return func() {
		mu.Lock()
		cnt := m[id]
		m[id] = cnt - 1
		mu.Unlock()
		printTrace(id, name, "<-", cnt)
	} // deferred
}
