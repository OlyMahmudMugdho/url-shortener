package shortener

import (
	"database/sql"
	"log"

	"github.com/OlyMahmudMugdho/url-shortener/models"
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
	var query = `INSERT INTO "urls" (user_id, full_url, short_url, created_at) VALUES ($1,$2,$3,$4) RETURNING url_id, user_id, full_url, short_url, created_at`
	row := s.db.QueryRow(query, link.UserId, link.FullUrl, link.ShortUrl, link.CreatedAt)
	err := row.Scan(&link.Id, &link.UserId, &link.FullUrl, &link.ShortUrl, &link.CreatedAt)

	if err != nil {
		log.Println(err)
		return link, err
	}

	return link, nil
}

func (s *Store) GetAllLinks(userId string) ([]models.Link, error) {
	var links []models.Link
	var query = `SELECT * FROM "urls" where user_id=$1`
	rows, err := s.db.Query(query, userId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var link models.Link
		err := rows.Scan(&link.Id, &link.UserId, &link.FullUrl, &link.ShortUrl, &link.UpdatedAt, &link.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		links = append(links, link)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	return links, err
}
