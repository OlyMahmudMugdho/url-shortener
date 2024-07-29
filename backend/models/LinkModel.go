package models

import (
	"database/sql"
	"time"
)

type Link struct {
	Id        int          `json:"id"`
	UserId    string       `json:"userId"`
	FullUrl   string       `json:"fullUrl"`
	ShortUrl  string       `json:"shortUrl"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	CreatedAt time.Time    `json:"createdAt"`
}
