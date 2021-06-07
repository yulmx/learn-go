// 执行进程

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	bin, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	fmt.Println(bin)

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	err = syscall.Exec(bin, args, env)
	if err != nil {
		panic(err)
	}

}