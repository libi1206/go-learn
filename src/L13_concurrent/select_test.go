package L13_concurrent

//func TestSelect1(t *testing.T) {
//	s := asyncService()
//	time.Sleep(time.Millisecond*1000)
//	select {
//	//这一个语句块是为了做超时处理，10s后如果没有结果他就会返回结果
//	//（当然有了default语句块这个语句块也就没有意义了）
//	case <-time.After(10*time.Second):
//		print("10s")
//	case ret := <-s:
//		print("result:",ret)
//	default:
//		t.Error("error")
//	}
//}
