package models

import "time"

type Link struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	FullUrl   string    `json:"fullUrl"`
	ShortUrl  string    `json:"shortUrl"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
