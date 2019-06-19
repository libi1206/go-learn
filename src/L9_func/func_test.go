package L9_func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//入口函数
func TestFunc(t *testing.T) {
	t.Log(add(1,2,3,4,5,6,7,8,9,9,9,9,10))
}

//需要计时的函数
func funcDemo()(int,int){
	//返回两个随机数
	return rand.Intn(10),rand.Intn(20)
}

//计时函数的函数，需要传入一个函数（但是现在这个函数的定义必须要符合规则）
func timeSpent(inner func()(int, int)) (int,int) {
	start := time.Now()
	ret1,ret2 := inner()
	fmt.Println("用时：",time.Since(start).Nanoseconds())
	return ret1,ret2
}

func add (param ...int) int {
	result := 0;
	for v,_ := range param {
		result += v
	}
	return result
}

func TestDefer(t *testing.T) {
	defer finallyFunc()
	fmt.Println("开始执行")
	//执行一次异常, 发现defer语句块仍然执行了
	panic("err")
}

func finallyFunc() {
	fmt.Println("一定会执行的语句块")
}