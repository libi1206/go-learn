package L13_concurrent

import (
	"fmt"
	"testing"
	"time"
)

//这个方法做一些耗时的操作，需要异步调用，
func service() string {
	time.Sleep(time.Second*1)
	return "do something"
}

//使用一个异步调用包装上面的方法
func asyncService() chan string {
	//创建一个string类型的频道
	returnChannel := make(chan string)
	//异步调用
	go func() {
		fmt.Println("开始异步调用")
		returnChannel <- service()
		fmt.Println("异步调用结束")
	}()
	return returnChannel
}

//在程序入口的线程里也要做一些操作
func TestTask(t *testing.T)  {
	//开始异步调用并且等待结果
	result := asyncService()

	//开始做别的事情
	time.Sleep(500*time.Millisecond)
	fmt.Println("do something else")

	//这里再等待结果
	fmt.Println(<-result)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		print(ret)
		println("DONE")
	//这里time.After也会返回一个channel，这个channel在给定的时间后就会给出结果
	case <-time.After(100*time.Millisecond):
		println("TIME_OUT")
	}
}