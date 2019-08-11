package L13_concurrent

import (
	"fmt"
	"testing"
	"time"
)


func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("结果是%d", id)
}

//当第一个任务返回时整个任务返回
func FirstResponse() string  {
	number := 10
	ch := make(chan string,number)
	for i:=0;i<number;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

//所有任务完成才返回
func AllResponse() string {
	number := 10
	ch := make(chan string,number)
	for i:=0;i<number;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalRet := ""
	for j:=0;j<number;j++{
		finalRet += <-ch+"\n"
	}
	return finalRet

}

func TestFirst(t *testing.T) {
	t.Log(FirstResponse())
	t.Log(AllResponse())
}