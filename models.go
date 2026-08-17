package main

import (
	
	"time"
	"github.com/google/uuid"
	"github.com/BuildwithY26/go/internal/database"
)

type User struct {
	ID  uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	APIKey string `json:"api_key"`
	
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Username,
		APIKey: dbUser.ApiKey,
	}
}

type Feed struct {
	ID  uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	URL string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		URL: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = databaseFeedToFeed(dbFeed)
	}
	return feeds
}

type FeedFollow struct {
	ID  uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}

}

type Post struct{
	ID  uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
	Description *string `json:"desc"`
	PublishedAt time.Time `json:"published_at"`
	Url string `json:"url"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = databaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}