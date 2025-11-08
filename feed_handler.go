package main
import ("fmt"
		"context"
		"time"
		"github.com/google/uuid"
		"os"
		"github.com/eldalland/go_blog_aggregator/internal/database")
		//Inserts a new feed into feeds table
func handlerFeed(s *state, cmd command,user database.User) error{
	if len(cmd.args) < 2{
		fmt.Printf("please enter the name of the feed as well as the url")
		os.Exit(1)
	}
	//retrieves current username to use as field data

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
	


	



	
	currUrl := cmd.args[1]
	
	if err != nil{
		fmt.Printf("error getting user: %s",err)
		os.Exit(1)
	}
	feed,err = s.db.GetFeedFromURL(context.Background(),currUrl)
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