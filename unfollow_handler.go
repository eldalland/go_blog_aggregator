package main
import ("fmt"
		"context"
	
		"os"
		"github.com/eldalland/go_blog_aggregator/internal/database")

func handleUnfollow(s *state, cmd command, user database.User) error{
feed, err := s.db.GetFeedFromURL(context.Background(),cmd.args[0])
if err != nil{
	fmt.Printf("error getting feed from url: %s", err)
	os.Exit(1)
}
err = s.db.DeleteFeedFollow(context.Background(), 
database.DeleteFeedFollowParams{
UserID: user.ID,
FeedID: feed.ID,
})
if err != nil{
	fmt.Printf("error deleting feedfollow: %s",err)
}
return nil
}