package packages

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func TestString() {
	/*相等判断*/
	if strings.EqualFold("Hello", "hellO") {
		fmt.Println("二者俩等")
	}

	/*前缀*/
	s1 := "t_users"
	if strings.HasPrefix(s1, "t_") {
		fmt.Printf("Yes,%v is a table's name \n", s1)
	}

	/*是否包含字符串*/
	s2 := "我哒哒的马蹄声，是个美丽的错误，我不是归人，是个嫖客"
	if strings.Contains(s2, "美丽") {
		fmt.Println("字符串s2包含'美丽'")
	}

	/*第一次出现的位置*/
	s3 := "google book food good book"
	fmt.Printf("字符串s3包含'oo'%d次 \n", strings.Count(s3, "oo"))
	fmt.Printf("'book'在s3中第一次出现的位置是%d \n", strings.Index(s3, "book"))

	/*首字母大写*/
	s4 := "Hello guys, i'm your dad"
	fmt.Println(strings.Title(s4))

	/*不知道这俩货有啥区别*/
	fmt.Println(strings.ToUpper("gopher phper"))
	fmt.Println(strings.ToTitle("gopher phper"))

	/*Trim*/
	s5 := ""
	for i := 0; i < 10; i++ {
		s5 += strconv.Itoa(i) + ","
	}
	fmt.Println(strings.TrimRight(s5, ","))

	/*按照不是字母和数字的符号来分割字符串*/
	s6 := "foo;bar-abc*hehe"
	sl := strings.FieldsFunc(s6, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	fmt.Println(sl)

	/*(explode)指定字符串分割*/
	s7 := "1,2,3,4,5"
	fmt.Println(strings.Split(s7, ","))     // [1 2 3 4 5]
	fmt.Println(strings.SplitN(s7, ",", 3)) // [1 2 3,4,5]

	/*(implode)分解切片，用指定字符串链接*/
	sl1 := []string{"foo", "bar", "func", "var"}
	fmt.Println(strings.Join(sl1, "-"))

	/*Reader*/
	s8 := "还君明珠双泪垂，恨不相逢未嫁时"
	reader := strings.NewReader(s8)
	fmt.Printf("字符串s8的长度为%d", reader.Len())
}

func IsContains(s, t string) bool {
	return strings.Contains(s, t)
}
