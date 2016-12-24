package funcs

import (
	"fmt"
)

type car struct {
	Brand string
	Price float32
}

func TestFmt() {
	carBMW := car{Brand: "BMW", Price: 500000}
	fmt.Printf("%v\n", carBMW)  //默认格式
	fmt.Printf("%+v\n", carBMW) //打印struct（包含字段名）
	fmt.Printf("%#v\n", carBMW) //值的Go语法表示
	fmt.Printf("%T\n", carBMW)  //类型的Go语法表示

	fmt.Printf("18的二进制是： %b\n", 18) //二进制
	varStr := '我'
	fmt.Printf("varStr的原始值(十进制)是 '%d'\n", varStr)
	fmt.Printf("varStr的原始值(十六进制)是 '%x'\n", varStr)
	fmt.Printf("varStr对应的unicode码值是 '%c'\n", varStr) //unicode码值

	s := `<\br>`
	fmt.Printf("%v\n", s) //<\br>
	fmt.Printf("%q\n", s) //"<\\br>"

	f := 3.2285657
	fmt.Printf("%.2f\n", f)  //两位小数
	fmt.Printf("%9.3f\n", f) //宽度为9，小数位为3
}
