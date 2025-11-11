package main

import (
	"fmt"
	"os"
	"time"
)

//begins aggregating posts from supplied feeds at an interval
func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Printf("enter a duration")
		os.Exit(1)
	}
	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		fmt.Printf("error parsing duration: %s", err)
		os.Exit(1)
	}
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}
