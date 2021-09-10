# shu-cli

![GitHub](https://img.shields.io/github/license/kainhuck/shu-cli) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kainhuck/shu-cli) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/kainhuck/shu-cli)

## 介绍

shu-cli 是一个`终端可交互的`cli，支持彩色输出

## 使用说明

[example](./example)

1. go get 

   ```
   go get github.com/kainhuck/shu-cli
   ```

2. 新建一个Cli对象

   ```go
   import CLI "github.com/kainhuck/shu-cli"
   
   ...
   
   cli := CLI.DefaultCli() // or use CLI.NewCli()
   ```

3. 定义并注册命令

   ```go
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
   ```

   ```golang
   cli.Register(cmd.InstallCmd)
   ```

4. just run it

   ```golang
   cli.Run()
   ```



5. 全局方法

   `shu-cli`提供了一些方法，可以方便用户使用

   ```golang
   // ReadInput 从cli.stdin读取用户输入 会先打印提示语句 prompt
   func ReadInput(prompt string) []string 
   // ReadOne 只读取第一个输入，如果不存在则输出 空字符串
   func ReadOne(prompt string) string 
   // ReadInt 将读到的slice转成int slice，如果转换失败 则置为0
   func ReadInt(prompt string) []int 
   // ReadOneInt 只读取第一个int 如果不存在则输出 0
   func ReadOneInt(prompt string) int
   // Printf 格式化输出到cli.stdout
   func Printf(format string, a ...interface{}) 
   // Println 输出到cli.stdout
   func Println(a ...interface{}) 
   // Store 用于存放数据到内存中
   func Store(key string, value interface{}) 
   // Load 从内存数据库中取数据
   func Load(key string) (value interface{}, ok bool) 
   // Black 黑色输出
   func Black(s string) string 
   // Red 红色输出
   func Red(s string) string 
   // Green 绿色输出
   func Green(s string) string 
   // Yellow 黄色输出
   func Yellow(s string) string 
   // Blue 蓝色输出
   func Blue(s string) string 
   // Purple 紫色输出
   func Purple(s string) string 
   // Cyan 青色输出
   func Cyan(s string) string 
   // Gray 灰色输出
   func Gray(s string) string 
   ```

6. 内置方法

   `shu-cli`提供了一些内置的命令

   - exit

     输入`exit`可退出cli

   - help

     输入`help`可以输出所有的命令已经命令的使用方法

     `help <cmd>...`help后面跟命令可以查看该命令的详细信息

## 屏幕截图

![img](./imgs/img.png)