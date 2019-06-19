package L13_concurrent

import (
	"fmt"
	"sync"
	"testing"
)

//生产者,把0-9依次放在 channel里
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//关闭通道
		close(ch)
		wg.Done()
	}()

}

//消费者，拿出channel里的数据
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for true {
			data,ok := <-ch
			//观察channel是否关闭
			if !ok {
				break
			}
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func Test1(t *testing.T)  {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	dataProducer(ch,wg)
	dataReceiver(ch,wg)
	wg.Wait()
}
