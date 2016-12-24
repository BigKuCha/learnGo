package funcs

import (
	"fmt"
	"time"
)

func TestChannels() {
	/*基础用法*/
	//basicUseOfChan()

	/*rangeChannel*/
	//rangeChan()

	/*带缓存的chan*/
	cacheChan()
}

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
	fmt.Println("程序运行结束")
}

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
