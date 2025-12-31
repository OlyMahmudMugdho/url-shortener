package redirector

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OlyMahmudMugdho/url-shortener/models"
)

type MockShortenerStore struct {
	SaveLinkFunc      func(link *models.Link) (*models.Link, error)
	GetAllLinksFunc   func(userId string) ([]models.Link, error)
	GetLinkFunc       func(urlId int) (*models.Link, error)
	GetPublicLinkFunc func(shortUrl string) (*models.Link, error)
	UpdateLinkFunc    func(link *models.Link) (*models.Link, error)
	DeleteLinkFunc    func(urlId int, userId int) error
}

func (m *MockShortenerStore) SaveLink(link *models.Link) (*models.Link, error) { return nil, nil }
func (m *MockShortenerStore) GetAllLinks(userId string) ([]models.Link, error) { return nil, nil }
func (m *MockShortenerStore) GetLink(urlId int) (*models.Link, error)          { return nil, nil }
func (m *MockShortenerStore) GetPublicLink(shortUrl string) (*models.Link, error) {
	return m.GetPublicLinkFunc(shortUrl)
}
func (m *MockShortenerStore) UpdateLink(link *models.Link) (*models.Link, error) { return nil, nil }
func (m *MockShortenerStore) DeleteLink(urlId int, userId int) error             { return nil }

func TestOpenLink(t *testing.T) {
	mockStore := &MockShortenerStore{
		GetPublicLinkFunc: func(shortUrl string) (*models.Link, error) {
			if shortUrl == "valid" {
				return &models.Link{FullUrl: "https://google.com"}, nil
			}
			return nil, http.ErrNoCookie // Just an error
		},
	}
	handler := NewRedirectorHandler(mockStore)

	t.Run("Valid short link", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/app/valid", nil)
		rr := httptest.NewRecorder()

		handler.OpenLink(rr, req)

		if rr.Code != http.StatusSeeOther {
			t.Errorf("Expected status SeeOther, got %v", rr.Code)
		}
		if rr.Header().Get("Location") != "https://google.com" {
			t.Errorf("Expected redirect to https://google.com, got %s", rr.Header().Get("Location"))
		}
	})

	t.Run("Invalid short link", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/app/invalid", nil)
		rr := httptest.NewRecorder()

		handler.OpenLink(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected status NotFound, got %v", rr.Code)
		}
	})
}
