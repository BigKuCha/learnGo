package funcs

import (
	"os/exec"
	"fmt"
	"os"
	"log"
)

func TestExec() {
	/*获取命令的绝对路径*/
	//getCmdPath("node")

	//execCmd()

	execPipe()
}

func execPipe() {
	cmd := exec.Command("node", "-v")
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatalf("start出错了，错误内容是%v\n", err)
	}
	var p []byte
	stdOut.Read(p)
	if err = cmd.Wait(); err != nil {
		log.Fatalf("wait出错了，错误内容是%v\n", err)
	}
	fmt.Println(p)
}
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
func getCmdPath(cmd string) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println("你好没有安装命令:", cmd)
	}
	fmt.Printf("命令'%s'的安装目录是：%s", cmd, path)
}
