package main

import (
	CLI "github.com/kainhuck/shu-cli"
)

func main() {
	cli := CLI.DefaultCli()
	cli.SetWelcomeMsg("welcome to use the `shu-cli`")
	cli.Run()
}
