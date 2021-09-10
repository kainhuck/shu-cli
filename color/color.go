package color

import "fmt"

type TypeColor uint8
type Prospect bool

const (
	foreground = 30
	background = 40

	Foreground Prospect = true
	Background Prospect = false
)

const (
	Black TypeColor = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
)

type Color bool

func (c Color) format(color TypeColor, str string, p ...Prospect) string {
	if len(p) == 0 {
		color += foreground
	} else {
		switch p[0] {
		case Foreground:
			color += foreground
		case Background:
			color += background
		}
	}

	if c {
		return fmt.Sprintf("\u001B[%dm%s\u001B[0m", color, str)
	}

	return str
}

func (c Color) Black(str string, p ...Prospect) string {
	return c.format(Black, str, p...)
}

func (c Color) Red(str string, p ...Prospect) string {
	return c.format(Red, str, p...)
}

func (c Color) Green(str string, p ...Prospect) string {
	return c.format(Green, str, p...)
}

func (c Color) Yellow(str string, p ...Prospect) string {
	return c.format(Yellow, str, p...)
}

func (c Color) Blue(str string, p ...Prospect) string {
	return c.format(Blue, str, p...)
}

func (c Color) Purple(str string, p ...Prospect) string {
	return c.format(Purple, str, p...)
}

func (c Color) Cyan(str string, p ...Prospect) string {
	return c.format(Cyan, str, p...)
}

func (c Color) Gray(str string, p ...Prospect) string {
	return c.format(Gray, str, p...)
}
