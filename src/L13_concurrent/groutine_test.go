package L13_concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T)  {
	for i :=0 ; i<10 ; i++{
		go func(s int) {
			fmt.Println(s)
			time.Sleep(time.Millisecond*100)
		}(i)
	}
}