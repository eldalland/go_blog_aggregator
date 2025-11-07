package main 
import("fmt"
"database/sql"
"os"
"log"
"github.com/eldalland/go_blog_aggregator/internal/config"
"github.com/eldalland/go_blog_aggregator/internal/database"
_ "github.com/lib/pq"
)


func main(){
	//reads config.json data
	jsonData,err :=config.Read()
	if err != nil{
		fmt.Printf("error reading config.json")
		os.Exit(1)
	}
	//creates current config based on json data
	currState := state{
		cfg: &jsonData,
	}
	//creates connection to database
	db, err := sql.Open("postgres", currState.cfg.DBURL)
	if err != nil{
		fmt.Printf("failed to create connection to database: %v\n", err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
    log.Fatalf("db ping failed: %v", err)
	}
	fmt.Println("DBURL:", currState.cfg.DBURL)
	if err := db.Ping(); err != nil {
    	log.Fatalf("db ping failed: %v", err)
	}


	dbQueries := database.New(db)//returns pointer to queries struct
	currState.db = dbQueries //sets current states database value to queries struct
	//list of possible cli commands, then dynamically registers new commands
	currCommands := commands{
		commandMap: map[string]func(*state, command) error{},
	}
	currCommands.register("login",handlerLogin)
	currCommands.register("register",handlerRegister)
	currCommands.register("reset",handlerReset)
	currCommands.register("users",handlerUsers)
	currCommands.register("agg",handlerAgg)
	currCommands.register("addfeed",handlerFeed)
	currCommands.register("feeds",handlerFeeds)
	currCommands.register("follow",handlerFollow)
	userInput := os.Args
	//checks user input for command and args, then calls corresponding handler through run function
	if len(userInput) < 2{
		fmt.Printf("not enough arguments provided")
		os.Exit(1)
		}
	currCommand := command{
		name: userInput[1],
		args: userInput[2:],
	}

	currCommands.run(&currState,currCommand)
	
}