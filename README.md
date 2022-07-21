# trace-func-call-chain
追踪 Go 函数调用链

[参考 Tony Bai 老师的博客](https://zhuanlan.zhihu.com/p/335592499)

1. 利用 defer 实现函数出入口的跟踪

```go
func trace() func() {
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