package color

import "fmt"

const (
	Red    = 31
	Green  = 32
	Yellow = 33
	Blue   = 34
	Purple = 35
)

type Color bool

func (c Color) Format(color int, str string) string {
	if c {
		return fmt.Sprintf("\u001B[%dm%s\u001B[0m", color, str)
	}

	return str
}

func (c Color) Red(str string) string {
	return c.Format(Red, str)
}

func (c Color) Green(str string) string {
	return c.Format(Green, str)
}

func (c Color) Yellow(str string) string {
	return c.Format(Yellow, str)
}

func (c Color) Blue(str string) string {
	return c.Format(Blue, str)
}

func (c Color) Purple(str string) string {
	return c.Format(Purple, str)
}
