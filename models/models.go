package models

import "time"

// コメント構造体
type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// 記事構造体
type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}
