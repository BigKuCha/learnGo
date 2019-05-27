package packages

import (
	"context"
	"fmt"
	"github.com/urfave/cli"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Context(c *cli.Context) {
	switch c.GlobalInt("m") {
	case 1:
		withTimeout()
		break
	case 2:
		withValue()
		break
	case 3:
		withCancel()
	}
}

func withTimeout() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	wg.Add(2)
	go work(ctx, "work1")
	time.Sleep(2 * time.Second)
	go work(ctx, "work2")
	wg.Wait()
}

func withValue() {
	fmt.Println("with value")
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctx := context.WithValue(ctxWithTimeout, "k", "bbb")
	wg.Add(1)
	go work(ctx, "work with value")
	time.Sleep(2 * time.Second)
	wg.Wait()
}

func withCancel() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(1)
	go work(ctx, "work with cancel")
	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc()
	}()
	wg.Wait()
}

func work(ctx context.Context, str string) {
	fmt.Println("key:", ctx.Value("k"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println(str, "done")
			wg.Done()
			return
		default:
			fmt.Println(str, "work")
		}
		time.Sleep(time.Second)
	}
}
