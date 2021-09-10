package cmd

import (
	"fmt"
	CLI "github.com/kainhuck/shu-cli"
	"github.com/kainhuck/shu-cli/cmd"
	"strconv"
)

var InstallCmd = &cmd.Command{
	Cmd: "install",
	Usage: "install remote/local",
	Desc: "install remote/local",
	Handler: func(args ...string) {
		if len(args) == 0 {
		CLI.Println("缺少安装对象，退出安装!")
		return
		}
		numbers := CLI.ReadInput("请输入要安装的服务的个数...")
		if len(numbers) == 0{
			CLI.Println("未输入安装个数，退出安装!")
		return
		}
		number, _ := strconv.Atoi(numbers[0])
		if number == 0 {
			CLI.Println("至少安装一个服务，退出安装!")
		return
		}
		for i := 0; i < number;i++{
		port := CLI.ReadInput(fmt.Sprintf("请输入第%d服务的端口", i+1))
		method :=  CLI.ReadInput(fmt.Sprintf("请输入第%d服务的加密方法", i+1))
		key :=  CLI.ReadInput(fmt.Sprintf("请输入第%d服务的秘钥", i+1))
			CLI.Printf("第%d个服务安装成功: port: %s, method: %s, key: %s\n", i+1, port[0], method[0], key[0])
		}
		CLI.Println("服务安装成功!")
	},
}
