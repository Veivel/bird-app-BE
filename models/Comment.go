package models

import "time"

type Comment struct {
	Text         string    `json:"text"`
	Author       string    `json:"author"`
	AuthorAvatar string    `json:"author_avatar"`
	CreatedAt    time.Time `json:"created_at"`
}
