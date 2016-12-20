package funcs

import (
	"strings"
	"bufio"
	"fmt"
)

func TestBufio() {
	//readAndPeek()
	readLine()
}

func readAndPeek() {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890S"
	strReader := strings.NewReader(str)

	bufReader := bufio.NewReaderSize(strReader, 0)
	//fmt.Printf("缓存中未读取的数据长度是：%v \n", bufReader.Buffered()) //0
	//
	//s, _ := bufReader.Peek(5)
	//fmt.Printf("%q\n", s)
	//s[1] = 'a'
	//for _, v := range s {
	//	fmt.Printf("%c \n", v)
	//}

	b := make([]byte, 2)
	for n, err := 0, error(nil); err == nil; {
		n, err = bufReader.Read(b)
		fmt.Printf("未读取的数据长度%d,读入b的字节数%d,读取的值%c, 错误%v \n",
			bufReader.Buffered(), n, b[:n], err)
	}
}

func readLine() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")
	bufReader := bufio.NewReaderSize(sr, 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = bufReader.ReadLine()
		fmt.Printf("读取的行数：%d \n", len(line))
		fmt.Printf("一行的内容是%q, %t, 错误是：%v", line, isPrefix, err)
	}
}
