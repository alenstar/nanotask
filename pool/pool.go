package pool

import (
	"runtime"
)

type Runner interface {
	Run()
}

var (
	defaultObjectPools    []*ObjectPool
	defaultCoroutinePools []*RoutinePool
)

func init() {
	defaultObjectPools = make([]*ObjectPool, runtime.NumCPU()*4)
	defaultCoroutinePools = make([]*RoutinePool, runtime.NumCPU()*4)
}

func Start() {

}

func Shutdown() {

}

func AddObject() {

}
func GetObject() {

}

func AddWorker(fn func()) {

}
