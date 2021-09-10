package main

import (
	"example/cmd"
	CLI "github.com/kainhuck/shu-cli"
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
		Handler: func(args ...string) {
			CLI.Printf("%s\n", strings.Join(args, " "))
		},
	})

	cli.Register(cmd.InstallCmd)

	cli.Run()
}
