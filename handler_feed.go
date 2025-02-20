package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dandytron/gator/internal/database"
	"github.com/google/uuid"
)

const URL = "https://www.wagslane.dev/index.xml"

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed name> <url>\n", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("==========================")

	return nil
}

// Helper function to print the fields of a feed
func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:		%v\n", feed.ID)
	fmt.Printf(" * Name:	%v\n", feed.Name)
	fmt.Printf(" * URL:		%v\n", feed.Url)
}

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s\n", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Could not retrieve feeds: %w", err)
	}

	if len(feeds) == 0 {
		return errors.New("No feeds found in database.")
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		fmt.Printf("%v, %v, created by: %v\n", feed.Name, feed.Url, feed.User)
	}

	return nil
}
