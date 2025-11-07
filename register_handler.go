package main

import (
	"context"
	"fmt"
	"time"
	"os"
	"errors"
	"database/sql"
	"github.com/google/uuid"
	"github.com/eldalland/go_blog_aggregator/internal/database"
)
//adds new user to users table
func handlerRegister(s *state, c command)error{
	if len(c.args) == 0{
		return fmt.Errorf("expecting a name, no name entered")
	}
	
	//does user exist?
	u, err :=s.db.GetUser(context.Background(),c.args[0])
	if err == nil{ 
		fmt.Printf("user already exists: %s",u.Name)
		os.Exit(1)
	} else if !errors.Is(err, sql.ErrNoRows) {
        return fmt.Errorf("lookup failed: %w", err)
    }
	//create user
	user,err := s.db.CreateUser(context.Background(),database.CreateUserParams{
	ID:        uuid.New(),
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Name:       c.args[0],
	})
	
	
	if err != nil{
		return fmt.Errorf("create failed: %w",err)
	}


	fmt.Printf("user created: %s (%s)\n",user.Name, user.ID.String())
	s.cfg.SetUser(c.args[0])
	fmt.Print("User has been set.")
	return nil
}