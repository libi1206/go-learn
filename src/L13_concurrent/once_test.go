package L13_concurrent

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singletonInstance *Singleton
var once sync.Once

func getSingletonInstance() *Singleton {
	once.Do(func() {
		fmt.Println("开始创建实例，这个方法只会执行一次")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

//程序入口
func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	//使用多个线程去获取实例，验证真的只执行了一次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := getSingletonInstance()
			//用于输出这个对象的地址，证明真的只执行一次
			fmt.Println("这个对象的地址是:",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
