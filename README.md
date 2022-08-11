# trace-func-call-chain

## 功能特性

追踪 Go 函数调用链

[参考 Tony Bai 老师的博客](https://zhuanlan.zhihu.com/p/335592499)

1. 利用 defer 实现函数出入口的跟踪

```go
func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	fmt.Printf("enter: %s\n", name)
	return func() { fmt.Printf("exit: %s\n", name) }
}
```

2. 利用 `go build tags` 添加 trace 开关

3. 解析 Go 源码为 AST，向 AST 的所有函数声明中注入 Trace() 函数

## 快速开始

1. 克隆源码
```
git clone https://github.com/myyppp/functrace.git
```

2. 编译
```
cd ./cmd/gen
go build -tags trace .
```

## 使用指南

```
gen [-w] xxx.go // -w：写入源文件
```
