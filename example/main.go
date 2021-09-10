package main

import (
	"fmt"
	CLI "github.com/kainhuck/shu-cli"
	"strconv"
	"strings"
)

func main() {
	cli := CLI.DefaultCli()
	cli.SetWelcomeMsg(`
███████╗██╗  ██╗██╗   ██╗
██╔════╝██║  ██║██║   ██║
███████╗███████║██║   ██║
╚════██║██╔══██║██║   ██║
███████║██║  ██║╚██████╔╝
╚══════╝╚═╝  ╚═╝ ╚═════╝`)

	cli.Register(&CLI.Command{
		Cmd:     "echo",
		Usage:   "echo args...",
		Desc:    "echo args...",
		Handler: func(readInput func(prompt string) []string, printf func(format string, a ...interface{}), args ...string) {
			printf("%s\n", strings.Join(args, " "))
		},
	})

	cli.Register(&CLI.Command{
		Cmd: "install",
		Usage: "install remote/local",
		Desc: "install remote/local",
		Handler: func(readInput func(prompt string) []string, printf func(format string, a ...interface{}), args ...string) {
			if len(args) == 0 {
				printf("缺少安装对象，退出安装!\n")
				return
			}
			numbers := readInput("请输入要安装的服务的个数...")
			if len(numbers) == 0{
				printf("未输入安装个数，退出安装!\n")
				return
			}
			number, _ := strconv.Atoi(numbers[0])
			if number == 0 {
				printf("至少安装一个服务，退出安装!\n")
				return
			}
			for i := 0; i < number;i++{
				port := readInput(fmt.Sprintf("请输入第%d服务的端口", i+1))
				method := readInput(fmt.Sprintf("请输入第%d服务的加密方法", i+1))
				key := readInput(fmt.Sprintf("请输入第%d服务的秘钥", i+1))
				printf("第%d个服务安装成功: port: %s, method: %s, key: %s\n", i+1, port[0], method[0], key[0])
			}
			printf("服务安装成功!\n")
		},
	})

	cli.Run()
}
