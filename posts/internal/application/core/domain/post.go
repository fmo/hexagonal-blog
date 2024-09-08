package domain

import "time"

type Post struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"created_at"`
}

func NewPost(title string, body string) Post {
	return Post{
		CreatedAt: time.Now().Unix(),
		Title:     title,
		Body:      body,
	}
}
