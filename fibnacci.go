package fibnacci

import (
	"fmt"
	"github.com/xmufeng/mprint"
)

func fib() func() int {
	var a, b int = 0, 1
	return func() int {
		tmp := a
		a = b
		b = tmp + b
		return a
	}
}

func Fibnacci(count int) {

	mprint.Println()

	f := fib()
	for i := 0; i < count; i++ {
		fmt.Printf("%d - ", f())
	}
}
