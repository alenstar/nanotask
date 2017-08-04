package pool

import (
	"runtime"
	//"sync"
)

type Runner interface {
	Run()
}

var (
	defaultObjectPools    []*ObjectPool
	defaultCoroutinePools []*CoroutinePool
)

func init() {
	defaultObjectPools = make([]*ObjectPool, runtime.NumCPU()*4)
	defaultCoroutinePools = make([]*CoroutinePool, runtime.NumCPU()*4)
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
