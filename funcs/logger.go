package funcs

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

func TestLogger() {
	/*输出log*/
	var buf bytes.Buffer
	logger := log.New(&buf, "Log: ", log.LstdFlags|log.Lshortfile)
	logger.Print("我是日志信息")
	fmt.Print(&buf)

	fmt.Println(logger.Flags()) //19 (ob1 | 0b10 | 0b10000)

	/*改变输出格式*/
	buf.Truncate(0)
	logger.SetFlags(log.Llongfile)
	logger.Print("改变输出方式了")
	fmt.Println(&buf)

	/*设置前缀*/
	buf.Truncate(0)
	logger.SetPrefix("warning: ")
	logger.Print("警告")
	logger.Print("再次警告")
	fmt.Println(&buf)

	/*log写入文件*/
	buf.Truncate(0)
	logger.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	logger.Println("警告")

	var fw *os.File
	rootPath, _ := os.Getwd()
	file := rootPath + "/logs/app" + time.Now().Format("2006-01-02") + ".log"
	fw, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	_, err = fw.WriteString(buf.String())
	if err != nil {
		fmt.Println(err)
	}
	defer fw.Close()
}
