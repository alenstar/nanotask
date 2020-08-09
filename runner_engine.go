package goworker


type Engine interface{
	Run()
	Shutdown()
}