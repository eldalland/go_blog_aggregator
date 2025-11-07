package main

import ("context"
		"fmt"
		"os")
//returns users table data, then checks names against current config users and prints accordingly
func handlerUsers(s *state, c command)error{
	users,err := s.db.GetUsers(context.Background())
	if err != nil{
		fmt.Printf("error getting users: %s",err)
		os.Exit(1)
	}
	for _,user := range users{
		if user.Name == s.cfg.CurrentUsername{
			fmt.Printf("* %s (current)\n",user.Name)
		} else{
			fmt.Printf("* %s\n",user.Name)
		}
	}
	return nil
}