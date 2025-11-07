package main
import ("fmt"
		"context"
		"time"
		"github.com/google/uuid"
		"os"
		"github.com/eldalland/go_blog_aggregator/internal/database")
		//Inserts a new feed into feeds table
func handlerFeed(s *state, cmd command) error{
	if len(cmd.args) < 2{
		fmt.Printf("please enter the name of the feed as well as the url")
		os.Exit(1)
	}
	//retrieves current username to use as field data
	currentUser := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(),currentUser)
	if err != nil{
		fmt.Printf("error getting current user from database: %s",err)
		os.Exit(1)
	}
	userId := user.ID
	feed, err := s.db.CreateFeed(context.Background(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:		cmd.args[0],
		Url: 		cmd.args[1],
		UserID: 	userId,	
	})
	if err != nil{
		fmt.Printf("error inserting feed: %s",err)
		os.Exit(1)
	}
	fmt.Printf("feed created: %s (%s)\n",feed.Name, feed.ID.String())
	return nil
}