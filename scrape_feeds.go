package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/eldalland/go_blog_aggregator/internal/database"
	"github.com/eldalland/go_blog_aggregator/internal/rssfeed"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	fmt.Printf("%s",feed.Url)
	if err != nil {
		fmt.Printf("error getting feed: %s", err)
		os.Exit(1)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("error marking feed: %s", err)
		os.Exit(1)
	}
	feedData, err := rssfeed.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("error getting feed data: %s", err)
		os.Exit(1)
	}
	for _, value := range feedData.Channel.Item {
		description := sql.NullString{}
		if value.Description != "" {
			description.String = value.Description
			description.Valid = true
		}
		pubTime, err := time.Parse(time.RFC1123Z, value.PubDate)
		if err != nil{
			fmt.Printf("error parsing pubDate: %s", err)
			os.Exit(1)
		}

		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       value.Title,
			Url:         value.Link,
			Description: description,
			PublishedAt: pubTime, 
			FeedID:      feed.ID,
		})
		if err != nil{
			fmt.Printf("error creating post: %s", err)
			os.Exit(1)
		}
	}
	return nil
}
