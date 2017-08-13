package plugins

import (
	"github.com/muesli/cache2go"
	"fmt"
	"time"
)

type mystruct struct {
	name   string
	salary int
}

func RunTest2go() {
	//writeAndRead()
	//write()
	read()
}

func read() {
	cache := cache2go.Cache("stu")
	val, err := cache.Value("lucy")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val.Data())
}
func write() {
	cache := cache2go.Cache("stu")
	val := mystruct{name: "lucy", salary: 3300}
	cache.Add("lucy", 0, val)
}
func writeAndRead() {
	cache := cache2go.Cache("mycache")
	val := mystruct{name: "Joe", salary: 2000}
	cache.Add("mykey", 5e9, val)
	item, _ := cache.Value("mykey")
	fmt.Println(item.Data())
	time.Sleep(6e9)
	item, err := cache.Value("mykey")
	if err == nil {
		fmt.Println(item.Data())
		return
	} else {
		fmt.Println(err)
		fmt.Println("nothing!")
	}
}
