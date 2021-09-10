package shu_cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/kainhuck/shu-cli/cmd"
	"github.com/kainhuck/shu-cli/color"
	"os"
	"strconv"
	"strings"
)

type Cli struct {
	ctx        context.Context
	cancel     context.CancelFunc
	welcomeMsg string      // 欢迎语句
	prompt     string      // 输入提示符
	color      color.Color // 颜色输出模块
	stdout     *os.File    // 标准输出
	reader     *bufio.Reader
	database   map[string]interface{} // 内存数据库，可以存放一些数据

	cmds map[string]*cmd.Command
}

var cli *Cli

// NewCli 新建cli
func NewCli(ctx context.Context, cancel context.CancelFunc, prompt string, colorMod bool, stdout, stdin *os.File) *Cli {
	cli = &Cli{
		prompt:   prompt,
		color:    color.Color(colorMod),
		stdout:   stdout,
		reader:   bufio.NewReader(stdin),
		ctx:      ctx,
		cancel:   cancel,
		cmds:     make(map[string]*cmd.Command),
		database: make(map[string]interface{}),
	}

	cli.Register(exitCmd)
	cli.Register(helpCmd)

	return cli
}

// DefaultCli 新建默认的cli
func DefaultCli() *Cli {
	ctx, cancel := context.WithCancel(context.Background())
	return NewCli(ctx, cancel, ">>> ", true, os.Stdout, os.Stdin)
}

// SetWelcomeMsg 设置欢迎语句
func (c *Cli) SetWelcomeMsg(msg string) {
	c.welcomeMsg = msg
}

// SetPrompt 设置输入提示符
func (c *Cli) SetPrompt(prompt string) {
	c.prompt = prompt
}

// SetColorMod 设置色彩模式
func (c *Cli) SetColorMod(colorMod bool) {
	c.color = color.Color(colorMod)
}

// SetStdOut 设置标准输出
func (c *Cli) SetStdOut(file *os.File) {
	c.stdout = file
}

// SetStdIn 设置标准输入
func (c *Cli) SetStdIn(file *os.File) {
	c.reader = bufio.NewReader(file)
}

// CancelFunc 返回cancel()
func (c *Cli) CancelFunc() context.CancelFunc {
	return c.cancel
}

// Run 运行cli，调用 cancel 可退出
func (c *Cli) Run() {
	if len(c.welcomeMsg) > 0 {
		c.println(c.color.Yellow(c.welcomeMsg))
	}

	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				c.handleInput(c.readInput())
			}
		}
	}()

	<-c.ctx.Done()
}

// readInput 读取输入
func (c *Cli) readInput() []string {
	c.printf(c.color.Purple(c.prompt))
	input, _ := c.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}

	// 分离命令和参数
	cmdArgs := strings.Split(input, " ")
	for i := 0; i < len(cmdArgs); i++ {
		cmdArgs[i] = strings.TrimSpace(cmdArgs[i])
	}

	return cmdArgs
}

// Register 注册命令
func (c *Cli) Register(cmd *cmd.Command) {
	c.cmds[cmd.Cmd] = cmd
}

func (c *Cli) println(a ...interface{}) {
	_, _ = fmt.Fprintln(c.stdout, a...)
}

func (c *Cli) printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(c.stdout, format, a...)
}

func (c *Cli) handleInput(cmdArgs []string) {
	if len(cmdArgs) == 0 {
		return
	}

	cmd, ok := c.cmds[cmdArgs[0]]
	if !ok {
		c.printf(c.color.Red("unknown cmd: `%v`\n"), cmdArgs[0])
		return
	}

	prompt := c.prompt
	c.prompt = fmt.Sprintf("[%s"+c.color.Purple("]%s"), c.color.Cyan(cmdArgs[0]), prompt)
	cmd.Handler(cmdArgs[1:]...)
	c.prompt = prompt
}

// ======================定义一些全局方法供使用======================

// ReadInput 从cli.stdin读取用户输入 会先打印提示语句 prompt
func ReadInput(prompt string) []string {
	cli.println(prompt)
	return cli.readInput()
}

// ReadOne 只读取第一个输入，如果不存在则输出 空字符串
func ReadOne(prompt string) string {
	all := ReadInput(prompt)
	if len(all) == 0 {
		return ""
	}

	return all[0]
}

// ReadInt 将读到的slice转成int slice，如果转换失败 则置为0
func ReadInt(prompt string) []int {
	strSlice := ReadInput(prompt)
	intSlice := make([]int, len(strSlice))

	for i := 0; i < len(strSlice); i++ {
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}

	return intSlice
}

// ReadOneInt 只读取第一个int 如果不存在则输出 0
func ReadOneInt(prompt string) int{
	all := ReadInt(prompt)
	if len(all) == 0{
		return 0
	}

	return all[0]
}

// Printf 格式化输出到cli.stdout
func Printf(format string, a ...interface{}) {
	cli.printf(format, a...)
}

// Println 输出到cli.stdout
func Println(a ...interface{}) {
	cli.println(a...)
}

// Store 用于存放数据到内存中
func Store(key string, value interface{}) {
	cli.database[key] = value
}

// Load 从内存数据库中取数据
func Load(key string) (value interface{}, ok bool) {
	value, ok = cli.database[key]
	return
}

// Black 黑色输出
func Black(s string) string {
	return cli.color.Black(s)
}

// Red 红色输出
func Red(s string) string {
	return cli.color.Red(s)
}

// Green 绿色输出
func Green(s string) string {
	return cli.color.Green(s)
}

// Yellow 黄色输出
func Yellow(s string) string {
	return cli.color.Yellow(s)
}

// Blue 蓝色输出
func Blue(s string) string {
	return cli.color.Blue(s)
}

// Purple 紫色输出
func Purple(s string) string {
	return cli.color.Purple(s)
}

// Cyan 青色输出
func Cyan(s string) string {
	return cli.color.Cyan(s)
}

// Gray 灰色输出
func Gray(s string) string {
	return cli.color.Gray(s)
}

// ======================定义一些默认的命令======================

var exitCmd = &cmd.Command{
	Cmd:   "exit",
	Usage: "exit",
	Desc:  "exit the cli",
	Handler: func(args ...string) {
		cli.cancel()
		cli.println(cli.color.Yellow("Bye~"))
	},
}

var helpCmd = &cmd.Command{
	Cmd:   "help",
	Usage: "help or help <cmd> ...",
	Desc:  "print the help info for cmds",
	Handler: func(args ...string) {
		if len(args) == 0 {
			cli.println("All Commands [cmd] [usage]:")
			for k, v := range cli.cmds {
				cli.printf("	%s	%s\n", cli.color.Green(k), v.Usage)
			}
			return
		}

		for _, a := range args {
			cmd, ok := cli.cmds[a]
			if !ok {
				cli.printf(cli.color.Red("unknown cmd: `%s`\n"), a)
			}
			cli.printf("	cmd:	%s\n", cmd.Cmd)
			cli.printf("	Usage:	%s\n", cmd.Usage)
			cli.printf("	Desc:	%s\n", cmd.Desc)
		}
	},
}
