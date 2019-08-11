package L11_exception

import (
	obj_pool "L14_object"
	"fmt"
	"testing"
	"unsafe"
)

func TestPool(t *testing.T) {
	pool := obj_pool.NewObjPool(10)
	for i := 0; i<11;i++  {
		ret, err := pool.GetObj()
		if(err!=nil){
			fmt.Println("err!",err,"count:",i)
			return
		}
		fmt.Println("addr:",unsafe.Pointer(ret)," count:",i)
	}
}
