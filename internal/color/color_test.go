package color

import (
	"fmt"
	"testing"
)

func TestColor_Format(t *testing.T) {
	//c := Color(false)
	c := Color(true)
	str := "hello world"

	fmt.Println(c.Format(Red, str))
	fmt.Println(c.Format(Green, str))
	fmt.Println(c.Format(Yellow, str))
	fmt.Println(c.Format(Blue, str))
	fmt.Println(c.Format(Purple, str))
}
