package packages

import (
	"fmt"
	"github.com/urfave/cli"
	"math/rand"
	"sync"
	"time"
)

var (
	locker = new(sync.Mutex)
	cond   = sync.NewCond(locker)
)

const (
	consumerNum = 3  // 消费者个数
	producerNum = 3  // 生产者个数
	capacity    = 10 // 容量
)

func Cond(ctx *cli.Context) {
	ch := make(chan int, capacity)
	Producer(ch)
	Consumer(ch)
	select {}
}

func Producer(ch chan<- int) {
	for i := 0; i < producerNum; i++ {
		go func(index int) {
			for {
				cond.L.Lock()
				if len(ch) >= capacity {
					fmt.Println("触发阈值，-----------停止生产------------------")
					cond.Wait()
				}
				intn := rand.Intn(100)
				ch <- intn
				fmt.Printf("生产协程%d被唤醒，生产：%d, 当前队列长度:%d\n", index, intn, len(ch))
				cond.L.Unlock()
				time.Sleep(time.Second * 2)
				cond.Signal() // 唤醒其他协程，生产或消费
			}
		}(i)
	}
}

func Consumer(ch <-chan int) {
	for i := 0; i < consumerNum; i++ {
		go func(index int) {
			index += 1000
			for {
				cond.L.Lock()
				if len(ch) == 0 {
					fmt.Printf("[%d]队列为空，++++++++++++++++++停止消费", index)
					cond.Wait()
				}
				cond.L.Unlock()
				fmt.Printf("消费协程%d被唤醒，消费%d\n", index, <-ch)
				time.Sleep(time.Second * 2)
				cond.Signal() // 发送信号，唤醒等待的协程
			}
		}(i)
	}
}
