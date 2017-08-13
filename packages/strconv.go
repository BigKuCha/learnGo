package packages

import (
	"fmt"
	"reflect"
	"strconv"
)

func TestStrconv() {
	/*布尔类型转换*/
	st := "true"
	isBool, err := strconv.ParseBool(st)
	if err == nil && isBool {
		fmt.Println("true")
	}
	/*16进制转换*/
	sint := "0x11"
	i, err := strconv.ParseInt(sint, 0, 0)
	if err == nil {
		fmt.Printf("类型是：%v,值是：%v \n", reflect.TypeOf(i).Kind(), i)
	} else {
		fmt.Println(err)
	}
	/*字符串转数字*/
	sint2 := "998"
	i2, err := strconv.Atoi(sint2)
	if err == nil {
		fmt.Printf("类型是:%v, 值是：%v \n", reflect.TypeOf(i2).Kind(), i2)
	}
	/*数字转字符串*/
	num := 100
	s := strconv.Itoa(num)
	fmt.Printf("类型是：%v, 值是 %v, \n", reflect.ValueOf(s).Kind(), s)
}
