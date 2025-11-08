package main
import ("fmt"
		"context"
		"github.com/eldalland/go_blog_aggregator/internal/database")
func loginMiddleware(handler func(s *state, cmd command, user database.User)error) func(*state,command)error{
	return func(s *state, cmd command)error{
		currUser, err := s.db.GetUser(context.Background(),s.cfg.CurrentUsername)
		if err != nil{
			fmt.Printf("error getting user: %s", err)
		}
		err = handler(s,cmd,currUser)
		if err != nil{
			fmt.Printf("error calling handler: %s", err)
		}
		return nil
	}
}