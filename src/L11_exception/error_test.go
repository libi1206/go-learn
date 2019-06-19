package L11_exception

import (
	"errors"
	"fmt"
	"testing"
)

// error接口的示例：斐波纳切数列
func GetFib(n int) ([]int, error) {
	if n<0 || n>100{
		return nil,errors.New("输入的数不符合规定")
	}
	fibList := []int {1,1}
	for i:=2;i<n;i++{
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}

func TestError(t *testing.T)  {
	defer func () {
		if err := recover();err != nil {
			fmt.Println("异常被掩盖：",err)
		}
		fmt.Println("运行defer")
	}()
	fmt.Println("开始执行")
	//出现不可调和的错误，终止程序
	panic(errors.New("出现异常"))
	fmt.Printf("执行结束")
}