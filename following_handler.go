package main

import (
	"context"
	"fmt"
"github.com/eldalland/go_blog_aggregator/internal/database"
	"os"
)
//returns what feed the current user is following
func handleFollowing(s *state, cmd command, user database.User) error {

	following, err := s.db.GetUserFeeds(context.Background(), user.ID)
	if err != nil {
		fmt.Printf("error getting user: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%s is following these feeds: ", s.cfg.CurrentUsername)
	for _,entry := range following {
		fmt.Printf("\n%s", entry.Name)
	}
	return nil
}
