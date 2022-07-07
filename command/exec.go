package command

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd 只执行命令，不获取结果
func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run failed with %s\n", err)
	}
}

// RunCmdOut 执行命令，并获取结果，不区分stdout和stderr
func RunCmdOut(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run failed with %s\n", err)
	}
	log.Printf("combined out:\n%s\n", string(output))
}

// RunCmdStd 执行命令，获取结果区分stdout和stderr
func RunCmdStd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// RunCmdForPipe 使用管道执行 <==> grep ERROR /var/log/messages | wc -l
func RunCmdForPipe() {
	c1 := exec.Command("grep", "ERROR", "/var/log/messages")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}

// RunCommandWithCtx 带context
func RunCommandWithCtx(ctx context.Context, name string, arg ...string) {
	cmd := exec.CommandContext(ctx, name, arg...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run failed with %s\n", err)
	}
}
