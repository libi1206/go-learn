package L13_concurrent

import (
	"fmt"
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	//导入一个锁对象
	var mut sync.Mutex
	//导入一个需要等待的携程数，用于主线程精准的等待进行累加的携程
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		//告诉等待队列需要等待的
		wg.Add(1)
		go func() {
			//使用defer来保证这个锁一定会被释放
			defer func() {
				mut.Unlock()
			}()
			//获得锁
			mut.Lock()
			counter++
			//告诉等待队列这个任务结束
			wg.Done()
		}()
	}
	//等待所有携程结束
	wg.Wait()
	fmt.Println("counter = ",counter)
}
