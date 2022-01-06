package models

import "time"

type Article struct {
	Id         int    `json:"id"`
	Author_id  int    `json:"author_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Created_at time.Time
}

type Articles struct {
	Articles []Articles `json:"articles"`
}
