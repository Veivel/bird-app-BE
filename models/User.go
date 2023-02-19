package models

import "time"

type User struct {
	Username     string    `json:"username"`
	Avatar       string    `json:"avatar"` // later
	CreatedAt    time.Time `json:"created_at"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CloseFriends []string  `json:"close_friends"` // later
}

type UserAuth struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}
