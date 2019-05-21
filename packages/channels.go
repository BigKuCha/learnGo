package packages

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

func Channel(ctx *cli.Context) {
	switch ctx.Int("m") {
	case 1:
		basicUseOfChan()
		break
	case 2:
		rangeChan()
		break
	case 3:
		cacheChan()
		break
	default:
		test()
	}
}

// default
func test() {
	var ch = make(chan int)
	num := 11
	go func() {
		for i := 1; i <= num; i++ {
			ch <- i
			//fmt.Println("++", i)
		}
		close(ch)
	}()
	for ii := range ch {
		val := ii
		go func() {
			fmt.Println("--", val)
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("程序运行结束")
}

// 1
func basicUseOfChan() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		fmt.Println(<-ch)
	}
}

// 2
func rangeChan() {
	chInt := make(chan int)
	chSquare := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			chInt <- i
		}
		close(chInt)
	}()

	go func() {
		for num := range chInt {
			chSquare <- num * num
		}
		close(chSquare)
	}()

	for rzt := range chSquare {
		fmt.Println(rzt)
	}
}

// 3
func cacheChan() {
	chInt := make(chan int, 3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("\n%d 进入队列", i)
			chInt <- i
		}
	}()
	for v := range chInt {
		fmt.Printf("\n%d 出队列", v)
		time.Sleep(1 * time.Second)
	}
}
