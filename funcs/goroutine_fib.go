package funcs

import (
	"fmt"
	"time"
)

func TestGoRoutineFib() {
	go waiting()
	fibNum := fib(45)
	fmt.Printf("\r 菲波那切数列第45个值是：%v \n", fibNum)
}

func waiting() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
