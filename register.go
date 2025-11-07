package main
//function to add new handlers to commandMap
func(c *commands) register(name string, f func(*state,command)error){
	c.commandMap[name]=f
}