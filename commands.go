package main
//holds all functions available through cli prompts
type commands struct{
	commandMap map[string]func(*state,command)error
}

