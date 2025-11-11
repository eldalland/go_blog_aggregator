package main
import ("context"
		"fmt"
		"os")
//returns name, url, and creator name all feeds in feeds table
func handlerFeeds(s *state, cmd command) error{
	feeds, err := s.db.GetFeeds(context.Background())
	if err !=nil{
		fmt.Printf("error getting feeds: %s", err)
		os.Exit(1)
	}
	for _,feed := range feeds{
		//gets creator name by using users_id foreign key to search users
		user, err := s.db.GetUserById(context.Background(),feed.UserID)
		userName := user.Name
		if err != nil{
			fmt.Printf("error with retrieving username: %s",err)
			os.Exit(1)
		}
		fmt.Printf("%s\n",feed.Name)
		fmt.Printf("%s\n",feed.Url)
		fmt.Printf("uploaded by: %s\n",userName)
	}
	return nil
}