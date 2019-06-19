package L10_oop

import "testing"

//=========接口定义==========
type Programmer interface {
	writeCode() string
}

//=========接口实现==========
type GoProgrammer struct {

}

func (g *GoProgrammer) writeCode() string {
	return "fmt.println(\"hello world\")"
}

func TestInterface(t *testing.T)  {
	var programmer Programmer
	programmer = new(GoProgrammer)
	t.Log(programmer.writeCode())
}

