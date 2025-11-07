package main

import ("context"
		"fmt"
		"os")
//drops users and feeds tables then re-creates them
func handlerReset(s *state, c command)error{
	err := s.db.DropFeeds(context.Background())
	if err != nil{
		fmt.Printf("failed to drop feeds table: %s\n",err)
		os.Exit(1)
	}
	err = s.db.DropUsers(context.Background())
	if err != nil{
		fmt.Printf("failed to drop users table: %s\n",err)
		os.Exit(1)
	}

	err = s.db.CreateUsers(context.Background())
	if err != nil{
		fmt.Printf("failed to re-initialize users table: %s\n",err)
		os.Exit(1)
	}
	err = s.db.CreateFeeds(context.Background())
		if err != nil{
		fmt.Printf("failed to re-initialize feeds table: %s\n",err)
		os.Exit(1)
	}
	fmt.Printf("users and feeds tables successful reset")
	
	return nil
}