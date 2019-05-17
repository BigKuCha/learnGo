package funcs

import (
	"fmt"
	"runtime"
)

func RunRecover() {
	err()
	fmt.Println("手动错误之后，主程序运行-----")
	runerr()
	fmt.Println("运行错误之后，主程序运行++++++")

}
func err() {
	defer catch()
	fmt.Println("手动错误之前")
	panic("手动错误")
	fmt.Println("手动错误之后")
}

func runerr() {
	defer catch()
	fmt.Println("运行错误之前")
	var m map[int]string
	m[9] = "123"
	fmt.Println("运行错误之后")
}

func catch() {
	err := recover()
	switch err.(type) {
	case runtime.Error:
		fmt.Println("运行时错误", err)
	default:
		fmt.Println("error", err)
	}
}
