package shu_cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/kainhuck/shu-cli/color"
	"os"
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

	cmds map[string]*Command
}

// NewCli 新建cli
func NewCli(ctx context.Context, cancel context.CancelFunc, prompt string, colorMod bool, stdout, stdin *os.File) *Cli {
	return &Cli{
		prompt: prompt,
		color:  color.Color(colorMod),
		stdout: stdout,
		reader: bufio.NewReader(stdin),
		ctx:    ctx,
		cancel: cancel,
		cmds:   make(map[string]*Command),
	}
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
			c.handleInput(c.readInput())
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
func (c *Cli) Register(cmd *Command) {
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
	cmd.Handler(c.readInputFunc, c.printf, cmdArgs[1:]...)
	c.prompt = prompt
}

func (c *Cli) readInputFunc(prompt string) []string {
	c.println(prompt)
	return c.readInput()
}
