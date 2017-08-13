package funcs

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func TestChannel() {
	badminton()
}

/*
《Go In Action》 demo
*/
/*模拟打羽毛球*/
func badminton() {
	// 指定随机数种子(种子相同，随机数也相同)
	rand.Seed(time.Now().UnixNano())
	//创建一个无缓冲的通道 用来存放'羽毛球'
	ch := make(chan int)
	// 计数器加2，标示要等待2个gorountine
	wg.Add(2)
	go player("Joe", ch)
	go player("Lucy", ch)
	//发球
	ch <- 1
	// 等待所有goroutine结束
	wg.Wait()
}

func player(name string, ch chan int) {
	// goroutine运行结束计数器-1
	defer wg.Done()
	for {
		// 等待球被击打过来（阻塞）
		ball, ok := <-ch
		if !ok {
			fmt.Printf("Player %s win the game！\n", name)
			return
		}
		// 随机一个100以内的数，当这个数是13的倍数的时候，表示没接到球，关闭通道结束goroutine
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n", name)
			close(ch)
			return
		}
		fmt.Printf("Player %s hit the ball, the %d time \n", name, ball)
		ball++
		ch <- ball
	}
}
