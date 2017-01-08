package funcs

import (
	"fmt"
	"time"
)

func TestSelect() {
	//testSelect()
	testRange()
}

func testRange() {
	ticker := time.Tick(1e9)
	//for {
	//	select {
	//	case t := <-ticker:
	//		fmt.Println(t.Format("2006.01.02 15:04:05"))
	//	case <-time.After(time.Second * 3):
	//		fmt.Println("过去10秒了，不玩了")
	//		return
	//	}
	//}
	for t := range ticker {
		fmt.Println(t.Format("2006.01.02 15:04:05"))
	}
}
func testSelect() {
	chanInt := make(chan int)
	chanStr := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			chanInt <- i
			time.Sleep(1e9)
		}
		close(chanInt)
	}()
	go func() {
		start := time.Now()
		end := time.Now().Add(5 * time.Second)
		for start.Before(end) {
			chanStr <- start.Format("2006-01-02 15:04:05")
			start = start.Add(1e9)
			time.Sleep(1e9)
		}
		close(chanStr)
	}()
	for {
		select {
		case i := <-chanInt:
			fmt.Printf("数字队列%d \n", i)
		case s := <-chanStr:
			fmt.Printf("字符串队列%s \n", s)
		}
	}
	fmt.Println("结束了")
}
