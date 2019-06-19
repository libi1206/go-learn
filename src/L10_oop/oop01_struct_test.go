package L10_oop

import "testing"

type Employee struct {
	Id string
	Name string
	Age string
}

//使用这种方法定义的行为在使用自己（this）的时候会进行值的复制
func (e Employee) getString() string {
	return e.Id+","+e.Name+","+e.Age
}

//使用指针传入的话就会避免内存拷贝
func (e *Employee) getStringPtr() string {
	return e.Id+","+e.Name+","+e.Age
}

func TestStruck (t *testing.T) {
	//初始化一个结构体
	e := Employee{"1","李比","20"}
	e1 := Employee{Id:"2",Age:"23"}
	//使用new关键字返回一个指向这个结构体的指针，相当于使用了取地址符"&"
	e2 := new(Employee)
	e2.Age = "30"
	t.Log(e.getString())
	t.Log(e1.getStringPtr())
	t.Log(e2.getStringPtr())
}