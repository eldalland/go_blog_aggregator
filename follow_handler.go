package main
import ("fmt"
		"context"
		"time"
		"github.com/google/uuid"
		"os"
		"github.com/eldalland/go_blog_aggregator/internal/database")

func handlerFollow(s *state, cmd command, user database.User)error{
	currUrl := cmd.args[0]

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
	fmt.Printf("%s is following %s",feedFollow[0].UserName,feedFollow[0].FeedName)
	return nil
}