package color

import (
	"testing"
)

func TestColor_Format(t *testing.T) {
	//c := Color(false)
	c := Color(true)
	str := "hello world"

	t.Log(c.Black(str))
	t.Log(c.Red(str))
	t.Log(c.Green(str))
	t.Log(c.Yellow(str))
	t.Log(c.Blue(str))
	t.Log(c.Purple(str))
	t.Log(c.Cyan(str))
	t.Log(c.Gray(str))

	t.Log(c.Black(str, Background))
	t.Log(c.Red(str, Background))
	t.Log(c.Green(str, Background))
	t.Log(c.Yellow(str, Background))
	t.Log(c.Blue(str, Background))
	t.Log(c.Purple(str, Background))
	t.Log(c.Cyan(str, Background))
	t.Log(c.Gray(str, Background))

	t.Log(c.Red(c.Cyan(str), Background))
}
