package packages

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func Bufio(ctx *cli.Context) {
	switch ctx.GlobalInt("m") {
	case 1:
		readFromFile() // 从文件读
		break
	case 2:
		readFromStdin() // 从标准输入读
		break
	case 3:
		scanWords()
		break
	case 4:
		readAndPeek()
		break
	case 5:
		readLine()
		break
	}
}

type MyReader struct {
}

func (m MyReader) Read(a []byte) (int, error) {
	return 1, nil
}

// m=1
func readFromFile() {
	f, err := os.Open("packages/bufio.go")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}
}

// m=2
func readFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入文字>")
	for {
		fmt.Print("Command>")
		line, _, _ := reader.ReadLine()
		fmt.Println(line)
	}
}

//m=3
func scanWords() {
	str := "Hi,I am Poly!"
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // 逐行输出 H i , I...
	}
}

// m=5
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

// m=6
func readLine() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890\nabcdefgd\n!@#$%a&")
	bufReader := bufio.NewReaderSize(sr, 0) // 缓冲区最小16字节 bufio.minReadBufferSize
	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = bufReader.ReadLine()
		//line, err = bufReader.ReadSlice('\n')
		fmt.Printf("当前buffer还可读取的字节数：%d \n", bufReader.Buffered()) //写入位置减去读取位置
		fmt.Printf("读取的字节长度是%d,内容是%q, 是否是行首%t, 错误是：%v\n", len(line), line, isPrefix, err)
	}
}
