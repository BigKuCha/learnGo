package main

import (
	"learnGo/utils"
	"crypto/des"
	"log"
	"fmt"
	"crypto/aes"
)

func main() {
	k := "1234567812345678"
	b, e := aes.NewCipher([]byte(k))
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println(b.BlockSize())
	return
	//testFuncs()
	//testUtils()
	key := "12345678"
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(block.BlockSize())
	fmt.Printf("%+v", block)
}

/*功能*/
func testUtils() {
	/*消息队列*/
	//utils.TestRabbitMQ()

	/*Excel操作*/
	//utils.TestExcel()
	//utils.ReadXlsx()
	//utils.WriteXlsx()
	//utils.WriteXlsxFromOrm()

	/*Aes加解密*/
	utils.TestAes()
}

/*包*/
func testFuncs() {
	/*goroutine more*/
	//funcs.TestG()

	/*反射*/
	//funcs.TestReflect()

	/*包os*/
	//funcs.TestOs()

	/*包string*/
	//funcs.TestString()

	/*包strconv*/
	//funcs.TestStrconv()

	/*包log*/
	//funcs.TestLogger()

	/*goroutine*/
	//funcs.TestGoRoutineFib()
	//funcs.TestGoroutineLock()

	/*channels*/
	//funcs.TestChannels()

	/*包bufio*/
	//funcs.TestBufio()

	/*包flag*/
	//funcs.TestFlag()

	/*包fmt*/
	//funcs.TestFmt()

	/*包os/exec*/
	//funcs.TestExec()

	/*Goroutine with Channel*/
	//funcs.TestRouChan()
	//funcs.TestExport()

	/*Select*/
	//funcs.TestSelect()

	/*信号*/
	//funcs.TestSignal()
}
