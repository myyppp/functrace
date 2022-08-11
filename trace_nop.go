//go:build !trace

package functrace

func trace() func() {
	return func() {}
}
