package main
import ("github.com/eldalland/go_blog_aggregator/internal/config"
"github.com/eldalland/go_blog_aggregator/internal/database"
)
//struct that has access to config data
type state struct{
	db  *database.Queries
	cfg *config.Config
}