package packages

import (
	"fmt"
	"runtime"
	"time"
)

func TestG() {
	names := []string{"Jack", "Lucy", "Lily", "Jone", "Mark"}

	for _, name := range names {
		go func(name string) {
			fmt.Println(name)
		}(name)
	}
	runtime.Gosched()
	time.Sleep(1e9)
}
