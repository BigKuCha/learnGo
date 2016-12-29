package funcs

import (
	"os"
	"syscall"
	"fmt"
	"os/signal"
	"time"
	"sync"
)

func TestSignal() {
	//testNotify()
	testNotify2()
}

//多个通道接收同一个信号
func testNotify2() {
	receiverOne := make(chan os.Signal, 1)
	receiverTwo := make(chan os.Signal, 2)

	sigs1 := []os.Signal{syscall.SIGINT}
	sigs2 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(receiverOne, sigs1...)
	signal.Notify(receiverTwo, sigs2...)
	fmt.Println("receiverOne 只处理中断信号")
	fmt.Println("receiverTwo 处理中断信号和退出信号")
	go func() {
		time.Sleep(10e9)
		fmt.Println("程序运行了10s, receiverOne 不再处理信号")
		signal.Stop(receiverOne)
		time.Sleep(5e9)
		fmt.Println("程序运行了15s, 不再处理任何信号") //按ctrl c 会结束运行
		signal.Stop(receiverTwo)
	}()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range receiverOne {
			if sig == syscall.SIGINT {
				fmt.Println("receiverOne 接到了中断信号，程序继续处理")
				continue
			}
			fmt.Println("receiverOne 接到了其他信号", sig)
		}
		wg.Done()
	}()

	go func() {
		for sig := range receiverTwo {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("receiverTwo 接到了中断信号，程序继续运行")
			case syscall.SIGQUIT:
				fmt.Println("receiverTwo 接到了退出信号，程序继续运行")
			default:
				fmt.Println("receiverTwo 接到了其他信号", sig)
			}

		}
		wg.Done()
	}()
	wg.Wait()
}
func testNotify() {
	receiver := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("给接受者设置的信号为%s...[reciver]\n", sigs)
	signal.Notify(receiver, sigs...)
	go func() {
		//10秒之后停止自行处理的信号行为
		time.Sleep(10e9)
		signal.Stop(receiver)
		close(receiver) //关闭通道，防止后面的range持续阻塞
	}()
	for sg := range receiver {
		switch sg {
		case syscall.SIGINT:
			fmt.Println("你按下了ctrl c中断信号！")
		case syscall.SIGQUIT:
			fmt.Println(`你按下了ctrl \，程序继续运行`)
		}
	}
}
