package main
//struct to hold info to identify handler functions within commandMap, and args to be passed into handler functions
type command struct{
	name string
	args []string
}