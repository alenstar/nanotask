package pool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Time struct {
	Name string
}

func TestObjectPool(t *testing.T) {

	o := NewObjectPool()
	now := time.Now()
	o.RegisterType("pool.Time", (*Time)(nil))
	o.RegisterType("time.Time", &now)
	tm, err := o.Obtain("time.Time")
	assert.NotEqual(t, *tm.(*time.Time), now)
	t.Log("Object:", tm.(*time.Time), err)
	*tm.(*time.Time) = now
	o.Release(tm)
	tm, err = o.Obtain("time.Time")
	assert.Equal(t, *tm.(*time.Time), now)
	t.Log("Object:", tm.(*time.Time), err)

	tm2, err := o.Obtain("pool.Time")
	assert.Equal(t, err, nil)
	assert.Equal(t, tm2.(*Time).Name, "")

	tm2.(*Time).Name = "time2"
	o.Release(tm2)
	tm2, err = o.Obtain("pool.Time")
	assert.Equal(t, err, nil)
	assert.Equal(t, tm2.(*Time).Name, "time2")
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
