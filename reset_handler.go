package main

import ("context"
		"fmt"
		"os")
//drops users and feeds tables then re-creates them
func handlerReset(s *state, c command)error{
	err:= s.db.DropFeedFollows((context.Background()))
	if err != nil{
		fmt.Printf("failed to drop feeds_follows table: %s\n",err)
		os.Exit(1)
	}

	err = s.db.DropPosts(context.Background())
		if err != nil{
		fmt.Printf("failed to drop posts table: %s\n",err)
		os.Exit(1)
	}

	err = s.db.DropFeeds(context.Background())
	if err != nil{
		fmt.Printf("failed to drop feeds table: %s\n",err)
		os.Exit(1)
	}
	err = s.db.DropUsers(context.Background())
	if err != nil{
		fmt.Printf("failed to drop users table: %s\n",err)
		os.Exit(1)
	}

	fmt.Printf("All tables successfully dropped")
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

	err = s.db.CreateFeedFollows(context.Background())
		if err != nil{
		fmt.Printf("failed to re-initialize feed_follows table: %s\n",err)
		os.Exit(1)
	}

	err = s.db.AddLastFetchedAt(context.Background())
		if err != nil{
		fmt.Printf("failed to re-add last_fetched_at column: %s\n",err)
		os.Exit(1)
	}

	err = s.db.CreatePosts(context.Background())
	if err != nil{
		fmt.Printf("failed to re-initialize posts table: %s\n",err)
		os.Exit(1)
	}

	fmt.Printf("\nAll tables successful reset")
	return nil
}