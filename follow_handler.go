package main
import ("fmt"
		"context"
		"time"
		"github.com/google/uuid"
		"os"
		"github.com/eldalland/go_blog_aggregator/internal/database")

func handlerFollow(s *state, cmd command)error{
	currUser := s.cfg.CurrentUsername
	currUrl := cmd.args[0]
	user, err := s.db.GetUser(context.Background(),currUser)
	if err != nil{
		fmt.Printf("error getting user: %s",err)
		os.Exit(1)
	}
	feed,err := s.db.GetFeedFromURL(context.Background(),currUrl)
	if err != nil{
		fmt.Printf("error getting feed: %s", err)
		os.Exit(1)
	}
	feedFollow , err := s.db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil{
		fmt.Printf("error inserting feedfollow: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%s %s",feedFollow[7], feedFollow[8])
	return nil
}