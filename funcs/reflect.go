package funcs

import (
	"fmt"
	"reflect"
)

func TestReflect() {
	//反射结构体
	s := Stu{Name: "joe", Age: 11}
	testReflect(s)
	//反射Map
	m := map[string]string{"city": "北京", "code": "010"}
	testReflect(m)
}

type Stu struct {
	Name string `no:"第一" color:"red"`
	Age  int
}

func (Stu) Say(name string) {
	fmt.Println("Hello " + name)
}

func testReflect(val interface{}) {
	v := reflect.ValueOf(val)
	t := reflect.TypeOf(val)
	kind := v.Kind()
	fmt.Printf("类型是：%s \n", kind)
	//反射结构体
	if kind == reflect.Struct {
		//获取值
		fmt.Printf("Name的值是%s,Age的值是%v \n", v.FieldByName("Name").String(), v.FieldByName("Age").Interface())
		//调用方法
		fmt.Print("调用方法Say:")
		v.MethodByName("Say").Call([]reflect.Value{reflect.ValueOf("lucy")})
		//反射结构体标签
		field := t.Field(0)
		tag := field.Tag
		fmt.Printf("第一个元素是：'%s',No是 '%s', color是 '%s' \n", field.Name, tag.Get("no"), tag.Get("color"))
		//设置值
		valStruct, ok := val.(Stu) //类型断言
		if ok {
			fmt.Println("断言成功")
		}
		reflect.ValueOf(&valStruct).Elem().FieldByName("Age").SetInt(30)
		fmt.Println("Age的新值是:", reflect.ValueOf(valStruct).FieldByName("Age").Int())
	}
	//反射Map
	if kind == reflect.Map {
		fmt.Println("元素个数是:", v.Len())
		for _, key := range v.MapKeys() {
			fmt.Printf("%s的值是:%v", key, v.MapIndex(key))
		}
	}
}
