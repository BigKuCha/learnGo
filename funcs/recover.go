package funcs

import (
	"fmt"
	"runtime"
)

func RunRecover() {
	protectRun(func() {
		fmt.Println("手动panic之前")
		panic("手动panic")
		fmt.Println("手动panic之后") // 不会执行
	})

	protectRun(func() {
		fmt.Println("运行错误之前")
		var m map[int]int
		m[1] = 1              // assignment to entry in nil map
		fmt.Println("运行错误之后") // 不会执行
	})
}

func protectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("运行时错误", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}
