package data_type_demo

import (
	"math"
	"testing"
)

func TestDataType(t *testing.T)  {
	uintMax := math.MaxUint32
	var intMax int64 = math.MaxInt64
	var floatMax float64 = math.MaxFloat64
	t.Log(intMax)
	t.Log(floatMax)
	t.Log(uintMax)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}

func TestString(t *testing.T)  {
	//字符串的初始化会初始化为空串
	var s string
	t.Log("*"+s+"*")
}