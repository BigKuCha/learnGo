package main

import (
	"learnGo/packages"
	"learnGo/plugins"
)

func main() {
	//gopackages()
	//gofuncs()
	goplugins()
}

func goplugins() {
	plugins.RunTest2go()
}

/*功能*/
func gofuncs() {
	/*消息队列*/
	//funcs.TestRabbitMQ()
	//funcs.TestRabbitMQ()

	/*Excel操作*/
	//funcs.TestExcel()
	//funcs.ReadXlsx()
	//funcs.WriteXlsx()
	//funcs.WriteXlsxFromOrm()

	/*Aes加解密*/
	//funcs.TestAes()

	/*demo channel*/
	//funcs.TestChannel()
}

/*包*/
func gopackages() {
	/*goroutine more*/
	//packages.TestG()

	/*反射*/
	//packages.TestReflect()

	/*包os*/
	//packages.TestOs()

	/*包string*/
	//packages.TestString()

	/*包strconv*/
	//packages.TestStrconv()

	/*包log*/
	//packages.TestLogger()

	/*goroutine*/
	//packages.TestGoRoutineFib()
	//packages.TestGoroutineLock()

	/*channels*/
	//packages.TestChannels()

	/*包bufio*/
	//packages.TestBufio()

	/*包flag*/
	//packages.TestFlag()

	/*包fmt*/
	//packages.TestFmt()

	/*包os/exec*/
	//packages.TestExec()

	/*Goroutine with Channel*/
	//funcs.TestRouChan()
	//packages.TestExport()

	/*Select*/
	//packages.TestSelect()

	/*信号*/
	//packages.TestSignal()

	/*Http*/
	packages.TestHttp()
}
