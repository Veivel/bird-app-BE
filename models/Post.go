package models

import "time"

type Post struct {
	Text           string    `json:"text"`
	Author         string    `json:"author"`
	AuthorAvatar   string    `json:"author_avatar"`
	IsCloseFriends bool      `json:"is_close_friends"` // later
	CreatedAt      time.Time `json:"created_at"`
	Likes          int       `json:"likes"` // later
}
