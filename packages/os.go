package packages

import (
	"fmt"
	"os"
	"strconv"
)

func TestOs() {
	/*获取系统相关信息*/
	//getOsInfo()
	/*获取文件相关信息*/
	//getFileInfo()
	/*写文件*/
	writeFile()
}
func writeFile() {
	file := "./app.log"
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for i := 1; i < 10; i++ {
		si := strconv.Itoa(i)
		f.WriteString(si)
	}
	f.Write([]byte("go"))
}
func getFileInfo() {
	/*目录判断*/
	path := "/usr/local"
	pathStat, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	if pathStat.IsDir() {
		fmt.Println(path, "是目录,目录下的文件有：")
		fo, err := os.Open(path)
		if err == nil {
			rd, err := fo.Readdir(0)
			if err == nil {
				for i := 0; i < len(rd); i++ {
					fmt.Println(rd[i].Name())
				}
			}
		}
	}
	/*文件判断*/
	rootPath, _ := os.Getwd()
	file := rootPath + "/funcs/os.go"
	fileStat, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(file, "文件不存在")
		} else {
			fmt.Println("其他错误", err)
		}
	} else {
		fmt.Printf("文件名称：%s,文件大小：%v,文件权限：%s,修改时间:%v \n",
			fileStat.Name(),
			fileStat.Size(),
			fileStat.Mode(),
			fileStat.ModTime().Format("2006-01-02 15:04:05"))
	}
	/*创建文件、删除文件*/
	newFileName := rootPath + "/funcs/test.log"
	_, err = os.Create(newFileName)
	if err != nil {
		fmt.Println(newFileName, "文件创建失败")
	} else {
		fmt.Println("文件创建成功")
		if os.Remove(newFileName) == nil {
			fmt.Println("文件删除成功")
		}
	}
}

func getOsInfo() {
	//读取主机名称
	hostName, err := os.Hostname()
	if err == nil {
		fmt.Println("主机名称", hostName)
	} else {
		fmt.Println(err)
	}
	//获取环境变量
	fmt.Println("环境变量", os.Environ())
	fmt.Println("GOPATH", os.Getenv("GOPATH"))
	//获取运行用户的id
	fmt.Println("运行用户的id", os.Getuid())
	//获取进程ID
	fmt.Println("进程ID", os.Getpid())
}
