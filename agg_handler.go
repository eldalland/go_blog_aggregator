package main
import ("fmt"
"context"
"github.com/eldalland/go_blog_aggregator/internal/rssfeed")
//returns RSSfeed data from supplied url
func handlerAgg(s *state, cmd command) error{
	feedUrl := "https://www.wagslane.dev/index.xml"

	feed, err := rssfeed.FetchFeed(context.Background(),feedUrl)
	if err != nil{
		return fmt.Errorf("error fetching rss feed: %s",err)
	}
	fmt.Printf("RSSFeed: %s",feed)
	return nil
}