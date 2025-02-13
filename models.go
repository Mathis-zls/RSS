package main

import (
	"time"

	"github.com/Mathis-zls/RSS/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"lastupdate"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"lastupdate"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userid"`
}

func databaseFeedToFeed(db database.Feed) Feed {
	return Feed{
		ID:        db.ID,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		Name:      db.Name,
		Url:       db.Url,
		UserID:    db.UserID,
	}
}

func databaseFeedsToFeeds(db []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range db {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"lastupdate"`
	UserID    uuid.UUID `json:"userid"`
	FeedID    uuid.UUID `json:"feedid"`
}

func databaseFeedFollowToFeedFollow(db database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        db.ID,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		UserID:    db.IDUser,
		FeedID:    db.IDFeed,
	}
}
func databaseFeedFollowsToFeedFollow(db []database.FeedFollow) []FeedFollow {
	feedfollows := []FeedFollow{}
	for _, feedfollow := range db {
		feedfollows = append(feedfollows, databaseFeedFollowToFeedFollow(feedfollow))
	}
	return feedfollows
}

type Posts struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(db database.Post) Posts {
	var desc *string
	if db.Description.Valid {
		desc = &db.Description.String
	}
	return Posts{
		ID:          db.ID,
		CreatedAt:   db.CreatedAt,
		UpdatedAt:   db.UpdatedAt,
		Title:       db.Title,
		Description: desc,
		PublishedAt: db.PublishedAt,
		Url:         db.Url,
		FeedID:      db.FeedID,
	}

}
func databasePostsToPosts(db []database.Post) []Posts {
	posts := []Posts{}
	for _, post := range db {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
