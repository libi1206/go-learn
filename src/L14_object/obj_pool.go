package L14_object

/**
  对象池
 */
import (
	"errors"
	"time"
)

//可以存入对象池的对象（也可以换成接口）
type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(poolSize int) *ObjPool {
	objPool := new(ObjPool)
	objPool.bufChan = make(chan *ReusableObj, poolSize)
	for i := 0; i < poolSize; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

func (p *ObjPool) GetObj()(*ReusableObj, error) {
	select {
	case ret := <- p.bufChan:
		return ret,nil
	case <- time.After(time.Duration(time.Second*1)):
		return nil,errors.New("time out")
	}
}

func (p *ObjPool) PutObj(obj *ReusableObj)error{
	select {
	case p.bufChan<-obj:
		return nil
	case <-time.After(time.Duration(time.Second*1)):
		return errors.New("over flow")
	}
}
