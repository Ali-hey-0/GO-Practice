package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/wags-lane/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Name      string	`json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User {
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Name      string	`json:"name"`
	Url       string	`json:"url"`
	UserID    uuid.UUID `json:"user_id"`

}



func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed {
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbFeed := range dbFeeds {
		feeds = append(feeds,databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}



func databaseFeedsFollowToFeeds(dbFeedFollows database.FeedFollow) FeedFollow {
	return FeedFollow { 
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserID:    dbFeedFollows.UserID,
		FeedID:    dbFeedFollows.FeedID,
	}
}


func databaseFeedsToFollowsFeedsFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	FeedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		FeedFollows = append(FeedFollows, databaseFeedsFollowToFeeds(dbFeedFollow))
	}
	return FeedFollows
}




type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: nullStringToStringPtr(post.Description),
		PublishedAt: nullTimeToTimePtr(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
