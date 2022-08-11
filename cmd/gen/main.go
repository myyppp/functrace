package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/myyppp/functrace/pkg/generator"
)

var wrote bool

func init() {
	flag.BoolVar(&wrote, "w", false, "write result to (source) file instead of stdout")
}

func usage() {
	fmt.Println("gen [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		usage()
		return
	}

	var file string
	if len(os.Args) == 3 {
		// os.Args[1] -w 参数
		file = os.Args[2]
	}
	if len(os.Args) == 2 {
		file = os.Args[1]
	}

	if filepath.Ext(file) != ".go" {
		usage()
		return
	}

	newSrc, err := generator.Rewrite(file)
	if err != nil {
		panic(err)
	}
	if newSrc == nil {
		fmt.Printf("on trace added for %s\n", file)
		return
	}
	if !wrote {
		fmt.Println(string(newSrc))
		return
	}
	// 写入源文件
	if err := os.WriteFile(file, newSrc, 0666); err != nil {
		fmt.Printf("write %s error: %v\n", file, err)
		return
	}
	fmt.Printf("add trace for %s ok\n", file)
}
