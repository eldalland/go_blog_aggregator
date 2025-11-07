package main

import (
	"context"
	"fmt"
	"os"
)

//function to set the login field within config.json
func handlerLogin(s *state, cmd command) error{
	if len(cmd.args)==0{
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	if _, err := s.db.GetUser(context.Background(),cmd.args[0]); err != nil{
		fmt.Printf("user not in database: %s",err)
		os.Exit(1)
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Print("User has been set.")
	return nil
}

