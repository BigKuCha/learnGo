package packages

import (
	"flag"
	"fmt"
)

func TestFlag() {
	basicUseOfFlag()
	//flagSet()
}

//基础用法
//e.g. go run main.go -name=lucy
func basicUseOfFlag() {
	var name = flag.String("name", "Joe", "请输入姓名")
	var salary int
	flag.IntVar(&salary, "salary", 2000, "工资")
	flag.Parse() //解析命令行参数到定义的flag
	fmt.Printf("姓名：%s, 工资：%d \n", *name, salary)
	otherFlags := flag.Args()
	fmt.Printf("其他参数：%v \n", otherFlags)
}

//go标准库中，经常这么做：
//定义了一个类型，提供了很多方法；为了方便使用，会实例化一个该类型的实例（通用），这样便可以直接使用该实例调用方法。
//比如：encoding/base64中提供了StdEncoding和URLEncoding实例，使用时：base64.StdEncoding.Encode()
func flagSet() {
	fset := flag.NewFlagSet("args", flag.ContinueOnError)
	city := fset.String("city", "BeiJing", "城市名称")
	fmt.Println(*city)
}
