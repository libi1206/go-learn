package L14_object

import (
	"fmt"
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	//创建一个对象缓存
	pool := sync.Pool{
		//指定默认对象的构造方法
		New: func() interface{} {
			fmt.Println("新增一个对象")
			return 100
		},
	}

	//这个缓存池的调用方式
	fmt.Println(pool.Get())
	pool.Put("你是主")
	v := pool.Get()
	fmt.Println(v)
	fmt.Println(pool.Get())
}
