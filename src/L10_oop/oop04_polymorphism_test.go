package L10_oop

import (
	"fmt"
	"testing"
)

//=======我们使用下面的接口定义，再自己添加一个javaProgrammer
//=========接口定义==========
//type Programmer interface {
//	writeCode() string
//}
//
////=========接口实现==========
//type GoProgrammer struct {
//
//}
//
//func (g *GoProgrammer) writeCode() string {
//	return "fmt.println(\"hello world\")"
//}

//======新的Programmer类型实现=====
type JavaProgrammer struct {

}

func (j *JavaProgrammer) writeCode() string {
	return "System.out.println(\"hello world\");"
}

//=======一个普通的方法，传入Programmer类型的数据
func work(programmer Programmer)  {
	fmt.Printf("Type:%T,Code:%s\n",programmer,programmer.writeCode())
}

//======程序入口======
func TestPolymorphism(t *testing.T) {
	java := new(JavaProgrammer)
	goo := new(GoProgrammer)
	work(java)
	work(goo)
}