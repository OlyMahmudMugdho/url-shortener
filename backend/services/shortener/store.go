package shortener

import (
	"database/sql"
	"github.com/OlyMahmudMugdho/url-shortener/models"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewShortenerStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) SaveLink(link *models.Link) (*models.Link, error) {
	var query = `INSERT INTO "urls" (user_id, full_url, short_url, updated_at, created_at) VALUES ($1,$2,$3,$4,$5)`
	row := s.db.QueryRow(query, link.UserId, link.FullUrl, link.ShortUrl, link.UpdatedAt, link.CreatedAt)
	err := row.Scan(&link.Id, &link.UserId, &link.FullUrl, &link.ShortUrl, &link.UpdatedAt, &link.CreatedAt)

	if err != nil {
		log.Println(err)
		return link, err
	}

	return link, nil
}
