package L8_string

import "testing"

func TestString(t *testing.T)  {
	s := "中"
	c := []rune(s)
	t.Log(len(c),len(s)) // 1
	t.Logf("中 unicode %x",c[0]) // 中 unicode 4e2d
	t.Logf("中 UTF-8 %x",s) // 中 UTF-8 e4b8ad
}