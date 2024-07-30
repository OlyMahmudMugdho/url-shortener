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
	var query = `SELECT * FROM "urls" WHERE user_id=$1`
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
		_ = rows.Close()
	}(rows)

	return links, err
}

func (s *Store) GetLink(urlId int) (*models.Link, error) {
	link := new(models.Link)
	var query = `SELECT * FROM "urls" where url_id=$1`
	row := s.db.QueryRow(query, urlId)
	err := row.Scan(&link.Id, &link.UserId, &link.FullUrl, &link.ShortUrl, &link.UpdatedAt, &link.CreatedAt)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return link, nil
}

func (s *Store) GetPublicLink(shortUrl string) (*models.Link, error) {
	link := new(models.Link)
	var query = `SELECT full_url FROM "urls" where short_url=$1`

	row := s.db.QueryRow(query, shortUrl)
	err := row.Scan(&link.FullUrl)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return link, nil
}

func (s *Store) UpdateLink(link *models.Link) (*models.Link, error) {
	var query = `UPDATE "urls" SET full_url=$1, short_url=$2, updated_at=current_timestamp WHERE url_id=$3 AND user_id=$4 RETURNING url_id, user_id, full_url, short_url, updated_at, created_at`
	row := s.db.QueryRow(query, link.FullUrl, link.ShortUrl, link.Id, link.UserId)
	err := row.Scan(&link.Id, &link.UserId, &link.FullUrl, &link.ShortUrl, &link.UpdatedAt, &link.CreatedAt)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (s *Store) DeleteLink(urlId int, userId string) error {
	var query = `DELETE FROM "urls" WHERE url_id=$1 AND urls.user_id=$2`
	_, err := s.db.Query(query, urlId, userId)
	if err != nil {
		return err
	}
	return nil
}
