package shu_cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/kainhuck/shu-cli/internal/color"
	"os"
	"strings"
)

type HandlerFunc func(args ...string)

type cli struct {
	ctx        context.Context
	cancel     context.CancelFunc
	welcomeMsg string
	container  map[string]HandlerFunc
	prompt     string
	color      color.Color
	stdout     *os.File
	stdin      *os.File
	stderr     *os.File
}

func DefaultCli() *cli {
	c := &cli{
		prompt: ">>> ",
		color:  true,
		stdout: os.Stdout,
		stdin:  os.Stdin,
		stderr: os.Stderr,
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	container := make(map[string]HandlerFunc)
	container[CmdHelp] = handleHelp
	container[CmdExit] = c.handleExit
	c.container = container

	return c
}

func (c *cli) SetPrompt(prompt string) {
	c.prompt = prompt
}

func (c *cli) SetColorMod(colorMod bool) {
	c.color = color.Color(colorMod)
}

func (c *cli) SetWelcomeMsg(msg string) {
	c.welcomeMsg = msg
}

func (c *cli) SetStdOut(file *os.File) {
	c.stdout = file
}

func (c *cli) SetStdIn(file *os.File) {
	c.stdin = file
}

func (c *cli) SetStdErr(file *os.File) {
	c.stderr = file
}

func (c *cli) Register(cmd string, handler HandlerFunc) {
	c.container[cmd] = handler
}

func (c *cli) Run() {
	if len(c.welcomeMsg) > 0 {
		c.println(c.color.Yellow(c.welcomeMsg))
	}

	reader := bufio.NewReader(c.stdin)

	go func() {
		for {
			c.printf(c.color.Purple(c.prompt))
			input, err := reader.ReadString('\n')
			if err != nil {
				c.errorln(err)
			}
			c.handleInput(input)
		}
	}()

	<-c.ctx.Done()
	c.printf(c.color.Blue("Bye~"))
}

func (c *cli) handleInput(input string) {
	input = strings.TrimSpace(input)
	cmds := strings.Split(input, " ")
	if len(cmds[0]) == 0 {
		return
	}

	if f, ok := c.container[cmds[0]]; ok {
		f(cmds[1:]...)
		return
	}
	c.errorf(c.color.Red("unSupport cmd `%v`\n"), cmds[0])
}

func (c *cli) handleExit(args ...string) {
	c.cancel()
}

func (c *cli) println(a ...interface{}) {
	_, _ = fmt.Fprintln(c.stdout, a...)
}

func (c *cli) printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(c.stdout, format, a...)
}

func (c *cli) errorln(a ...interface{}) {
	_, _ = fmt.Fprintln(c.stderr, a...)
}

func (c *cli) errorf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(c.stderr, format, a...)
}
