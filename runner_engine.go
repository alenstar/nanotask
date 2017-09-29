package nano


type Engine interface{
	Run()
	Shutdown()
}