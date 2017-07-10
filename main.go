package main

import "learnGo/utils"

func main() {
	//testFuncs()
	testUtils()
}

/*功能*/
func testUtils() {
	/*消息队列*/
	//utils.TestRabbitMQ()

	/*Excel操作*/
	//utils.TestExcel()
	//utils.ReadXlsx()
	utils.WriteXlsx()
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
