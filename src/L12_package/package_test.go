package L12_package

import (
	"L12_package/service"
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestPackage (t *testing.T)  {
	t.Log(service.Add(1,2))
	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("321"),"321")
	t.Log(m.Get(cm.StrKey("321")))
}
