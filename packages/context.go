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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	wg.Add(2)
	go work(ctx, "work1")
	time.Sleep(2 * time.Second)
	go work(ctx, "work2")
	wg.Wait()
}

func work(ctx context.Context, str string) {
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
