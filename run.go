package main
import ("fmt")
//function that calls handlers with info from passed in command
func (c *commands) run(s *state, cmd command) error{
	handler, ok := c.commandMap[cmd.name]
	if !ok{
		return fmt.Errorf("unkown command: %s",cmd.name)
	}
	handler(s,cmd)
	return nil
}