package L7_map

import "testing"

func TestMap(t *testing.T) {
	// 声明一个Map
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m["one"]) //1

	//使用make声明一个map
	m2 := make(map[string]int, 10)
	//往map里添加键值对
	m2["add"] = 1 //1
	t.Log(len(m2))

	value, ok := m2["one"]
	t.Log(value, ok) // 0,false
	value, ok = m2["add"]
	t.Log(value, ok) // 1,true

	for k, v := range m {
		t.Log(k, ":", v)
	}
}

func TestMapFunc(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2),m[2](2),m[3](2))
}
