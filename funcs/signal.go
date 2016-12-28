package funcs

import (
	"os"
	"syscall"
	"fmt"
	"os/signal"
	"time"
)

func TestSignal() {
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
			fmt.Println("你按下了Esc，程序继续运行")
		}
	}
}
