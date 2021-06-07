// 生成进程: 调用命令进程

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	datecmd := exec.Command("date")

	dateOut, err := datecmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIN, _ := grepCmd.StdinPipe()
	grepOUT, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIN.Write([]byte("hello grep\ngoodbye grep"))
	grepIN.Close()
	grepBytes, _ := ioutil.ReadAll(grepOUT)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -alh")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -alh")
	fmt.Println(string(lsOut))

}