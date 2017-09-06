package pool

import (
	"gopkg.in/go-playground/assert.v1"
	"fmt"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {

	o := NewObjectPool()
	now := time.Now()
	o.RegisterType("time", &now)
	tm, err := o.Obtain("time")
	assert.NotEqual(t, *tm.(*time.Time), now)
	t.Log("Object:",tm.(*time.Time),  err)
	*tm.(*time.Time) = now
	o.Release(tm)
	tm, err = o.Obtain("time")
	assert.Equal(t, *tm.(*time.Time), now)
	t.Log("Object:",tm.(*time.Time),  err)
}

func TestRoutinePool(t *testing.T) {
	cp := NewRoutinePool(32)
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

