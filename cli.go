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
	stdin      *os.File    // 标准输入
}

// NewCli 新建cli
func NewCli(ctx context.Context, cancel context.CancelFunc, prompt string, colorMod bool, stdout, stdin *os.File) *Cli {
	return &Cli{
		prompt: prompt,
		color:  color.Color(colorMod),
		stdout: stdout,
		stdin:  stdin,
		ctx:    ctx,
		cancel: cancel,
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
	c.stdin = file
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

	reader := bufio.NewReader(c.stdin)

	go func() {
		for {
			c.printf(c.color.Purple(c.prompt))
			input, err := reader.ReadString('\n')
			if err != nil {
				c.println(c.color.Red(fmt.Sprintf("%s", err)))
			}
			c.handleInput(input)
		}
	}()

	<-c.ctx.Done()
}

func (c *Cli) println(a ...interface{}) {
	_, _ = fmt.Fprintln(c.stdout, a...)
}

func (c *Cli) printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(c.stdout, format, a...)
}

func (c *Cli) handleInput(input string) {
	input = strings.TrimSpace(input)
	if len(input) == 0{
		c.printf(input)
	}else{
		c.println(input)
	}
}
