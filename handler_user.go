package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dandytron/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>\n", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("could not create new user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	s.cfg.SetUser(user.Name)

	fmt.Println("User created successfully:")
	// see helper func printUser below
	printUser(user)
	return nil

}

// This is the function signature of all command handlers
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>\n", cmd.Name)
	}

	username := cmd.Args[0]

	// checks to see if user is already in database, exits if it isn't
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("could not find user: %w\n", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w\n", err)
	}

	fmt.Printf("Username has been set to: %v\n", username)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	names, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not retrieve names: %w", err)
	}

	if len(names) == 0 {
		return fmt.Errorf("There are no registered users: %w", err)
	}

	for _, name := range names {
		if name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", name)
			continue
		}
		fmt.Printf("* %v\n", name)
	}

	return nil
}

// Helper function to print the fields of a user
func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)
}
