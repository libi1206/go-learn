package L5_flow

import "testing"

func TestIfFor(t *testing.T) {
	i := 0
	for i < 5{
		t.Log(i)
		i++
	}

	for n:=0; n<5;n++  {
		t.Log(n)
	}

	//普通的
	if 1 == i {
		t.Log(true)
	} else if false {
		t.Log(false)
	} else {

	}

	if a:=1==1;a{
		t.Log("a变量在赋值后可以直接拿来用")
	}

	x := "1"
	switch x {
	case "1":
		t.Log(1)
	case "2" , "3":
		t.Log(2)
	default:
		t.Log("unknow")
	}
	n := 3

	switch {
	case n%2==0:
		t.Log("偶数")
	case n%2==1:
		t.Log("奇数")
	default:
		t.Log("未知")
	}
}
