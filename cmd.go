package shu_cli

type HandleFunc = func(readInput func(prompt string) []string, printf func(format string, a ...interface{}), args ...string)

// Command 命令
type Command struct {
	Cmd     string     // 命令名称
	Usage   string     // 命令使用
	Desc    string     // 命令描述
	Handler HandleFunc // 处理方法
}
