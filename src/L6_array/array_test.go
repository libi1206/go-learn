package L6_array

import "testing"

func TestArray(t *testing.T) {
	//声明一个长度是3的数组，初始话为0值
	var a [3]int
	a[0] = 1

	//声明数组的同时初始化
	b := [3]int{1,2,3}
	c := [2][2] int {{1,2},{3,4}}
	t.Log(b)
	t.Log(c)

	//使用for进行遍历（这里也可以进行普通的下标进行遍历）
	for _ /*索引值，下划线表示占位，不使用*/,e /*元素值*/:= range a{
		t.Log("占位循环：",e)
	}

}

func TestSlice(t *testing.T)  {
	//声明一个切片
	var s0 []int;
	t.Log(cap(s0),len(s0))

	//往切片里添加一个元素
	s0 = append(s0, 1);
	t.Log(cap(s0),len(s0))

	//创建一个数组长度是2，容量是3的切片
	s1 := make([]int,2,3)
	t.Log(cap(s1),len(s1))
}

func TestShare(t *testing.T) {
	month := []int {1,2,3,4,5,6,7,8,9,10,11,12}
	q2 := month[3:6]
	t.Log(len(q2),cap(q2))
	summer := month[4:7]
	t.Log(len(summer), cap(summer))
}