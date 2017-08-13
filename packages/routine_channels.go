package packages

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func TestRouChan() {
	files := []int{1, 2, 3, 4, 5, 6, 7, 8}
	succ, fail := makeThumbs(files)
	fmt.Printf("总数%d,成功%d,失败%d", len(files), succ, fail)
}

func TestExport() {
	start, _ := time.Parse("01/2006", "06/2015")
	end, _ := time.Parse("01/2006", "12/2016")
	ch := make(chan string)
	var wg sync.WaitGroup
	for start.Before(end) {
		wg.Add(1)
		go func(s time.Time) {
			defer wg.Done()
			export(s)
			ch <- s.Format("2006-01")
		}(start)
		start = start.AddDate(0, 1, 0)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for range ch {

	}
}

func export(tm time.Time) {
	fmt.Printf("开始处理%s \n", tm.Format("2006-01"))
	t := tm.Format("1") + "s"
	d, _ := time.ParseDuration(t)
	time.Sleep(d)
	fmt.Println(tm.Format("200601"))
}

func makeThumbs(files []int) (succ, fail int) {
	ch := make(chan error, len(files))
	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(f int) {
			time.Sleep(1 * time.Second)
			defer wg.Done()
			_, err := longTimeFunc(f)
			ch <- err
		}(f)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for err := range ch {
		if err != nil {
			fail++
			continue
		}
		succ++
	}
	fmt.Println("完成了")
	return succ, fail
}

func longTimeFunc(f int) (string, error) {
	fmt.Println(f)
	if f == 3 {
		return "", errors.New("错误3")
	}
	return "ok", nil
}
