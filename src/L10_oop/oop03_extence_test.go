package L10_oop

import (
	"fmt"
	"testing"
)

//======父类 宠物类======
type Pat struct {

}

func (p *Pat) speak() {
	fmt.Println("...")
}

func (p *Pat) speakTo(host string)  {
	p.speak()
	fmt.Println(host)
}

//======子类 狗🐶=======
type Dog struct {
	Pat
}

func (d *Dog) speak() {
	fmt.Println("旺！")
}



//======程序入口======
func TestDog(t *testing.T)  {
	//var p Pat  = new(Dog)
	d := new(Dog)
	d.speakTo("123")
}