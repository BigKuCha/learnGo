package packages

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"time"
)

func Exec(ctx *cli.Context) {
	switch ctx.GlobalInt("m") {
	case 1:
		getCmdPath("node")
		break
	case 2:
		simpleUseOfExec()
		break
	case 3:
		execCmd()
		break
	case 4:
		execPipe()
		break
	case 5:
		getThread("mysql")
		break
	case 6:
		exportExcel()
		break
	default:

	}
}

// 1 获取命令路径
func getCmdPath(cmd string) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println("你好没有安装命令:", cmd)
	}
	fmt.Printf("命令'%s'的安装目录是：%s", cmd, path)
}

// 2 简单用法
func simpleUseOfExec() {
	out, _ := exec.Command("date").Output()
	fmt.Printf("%s", out)

	cmd := exec.Command("npm", "-v")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// 3 执行命令
func execCmd() {
	//查找php扩展中是否含有curl
	//php -m | grep curl
	cmdPhp := exec.Command("php", "-m")
	cmdGrep := exec.Command("grep", "curl")
	cmdGrep.Stdin, _ = cmdPhp.StdoutPipe()

	cmdGrep.Stdout = os.Stdout
	cmdGrep.Stderr = os.Stderr

	cmdGrep.Start()
	cmdPhp.Run()
	cmdGrep.Wait()
}

// 4 测试管道
func execPipe() {
	cmd := exec.Command("node", "-v")
	stdOut, err := cmd.StdoutPipe()
	defer stdOut.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatalf("start出错了，错误内容是%v\n", err)
	}
	r := bufio.NewReader(stdOut)
	line, err := r.ReadString('\n')
	if err == nil {
		fmt.Println(line)
	}
}

// 5 通过grep查找进程
func getThread(thread string) {
	cmdPs := exec.Command("ps", "aux")
	cmdGrep := exec.Command("grep", thread)
	var grepRzt bytes.Buffer
	cmdGrep.Stdout = &grepRzt

	cmdGrep.Stdin, _ = cmdPs.StdoutPipe()
	cmdPs.Start()
	cmdGrep.Run()
	cmdPs.Wait()
	fmt.Println(grepRzt.String())
}

// 6
func exportExcel() {
	ch := make(chan string)
	defer close(ch)
	go func() {
		start, _ := time.Parse("01/2006", "06/2015")
		end, _ := time.Parse("01/2006", "12/2016")
		for start.Before(end) {
			month := start.Format("200601")
			ch <- month
			start = start.AddDate(0, 1, 0)
		}
	}()

	for month := range ch {
		//month的地址相同，而goroutine调用的顺序是不确定的，所以进入线程的month值不稳定
		//_m在每次循环的时候重新分配内存地址，再传入goroutine的时候是确定值的！
		_m := month
		go func() {
			fmt.Printf("导出%s月份的报表\n", _m)
			cmd := exec.Command("php", "/Users/bigkucha/Work/Admin/artisan", "finance:export", _m)
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	}
}
