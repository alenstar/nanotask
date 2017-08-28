package pool

import (
	//. "gopkg.in/go-playground/assert.v1"
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {

}

func TestCoPool(t *testing.T) {
	cp := NewCoroutinePool(32)
	cp.Add(func() {
		time.Sleep(time.Second * 1)
		fmt.Println(time.Now(), "\tTest CoPool 1")
	})
	cp.Add(func() {
		time.Sleep(time.Second * 1)
		fmt.Println(time.Now(), "\tTest CoPool 2")
	})
	time.Sleep(time.Second * 2)
	cp.Shutdown()
}
