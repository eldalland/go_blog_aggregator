package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/eldalland/go_blog_aggregator/internal/database"
	
)
//displays all aggregated posts from feeds the user is following
func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		newlimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Printf("error parsing limit, please enter an integer: %s", err)

		}
		limit = newlimit
	}
	feeds, err := s.db.GetUserFeeds(context.Background(), user.ID)
	if err != nil {
		fmt.Printf("error getting feeds for current user: %s", err)
		os.Exit(1)
	}


	feedIdStrings := []string{}
	for _, value := range feeds {

		feedIdStrings = append(feedIdStrings, value.ID.String())

	}



	// Test the query manually first
	

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		FeedID: feedIdStrings,
		Limit:  int32(limit),
	})
	if err != nil {
		fmt.Printf("error getting posts: %s", err)
		os.Exit(1)
	}


	fmt.Printf("posts: %v", posts)
	for _, post := range posts {
		fmt.Printf("\nTitle: %s\n%v\npublication date: %s", post.Title, post.Description.String, post.PublishedAt)
	}
	return nil
}
