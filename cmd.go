package shu_cli

import (
	"fmt"
	"os"
)

const (
	CmdHelp = "help"
	CmdExit = "exit"
)

func handleHelp(args ...string) {
	fmt.Println("shu-cli help info:")
	fmt.Println(" help: print help info")
	fmt.Println(" exit: exit program")
}

func handleExit(args ...string) {
	fmt.Println("Bye~")
	os.Exit(0)
}
