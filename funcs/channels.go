package funcs

import (
	"fmt"
	"time"
)

func TestChannels() {
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
