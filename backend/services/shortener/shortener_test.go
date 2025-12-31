package shortener

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OlyMahmudMugdho/url-shortener/models"
	"github.com/OlyMahmudMugdho/url-shortener/types"
)

type MockShortenerStore struct {
	SaveLinkFunc      func(link *models.Link) (*models.Link, error)
	GetAllLinksFunc   func(userId string) ([]models.Link, error)
	GetLinkFunc       func(urlId int) (*models.Link, error)
	GetPublicLinkFunc func(shortUrl string) (*models.Link, error)
	UpdateLinkFunc    func(link *models.Link) (*models.Link, error)
	DeleteLinkFunc    func(urlId int, userId int) error
}

func (m *MockShortenerStore) SaveLink(link *models.Link) (*models.Link, error) {
	return m.SaveLinkFunc(link)
}

func (m *MockShortenerStore) GetAllLinks(userId string) ([]models.Link, error) {
	return m.GetAllLinksFunc(userId)
}

func (m *MockShortenerStore) GetLink(urlId int) (*models.Link, error) {
	return m.GetLinkFunc(urlId)
}

func (m *MockShortenerStore) GetPublicLink(shortUrl string) (*models.Link, error) {
	return m.GetPublicLinkFunc(shortUrl)
}

func (m *MockShortenerStore) UpdateLink(link *models.Link) (*models.Link, error) {
	return m.UpdateLinkFunc(link)
}

func (m *MockShortenerStore) DeleteLink(urlId int, userId int) error {
	return m.DeleteLinkFunc(urlId, userId)
}

func TestAddUrl(t *testing.T) {
	mockStore := &MockShortenerStore{
		SaveLinkFunc: func(link *models.Link) (*models.Link, error) {
			link.Id = 1
			return link, nil
		},
	}
	handler := NewShortenerHandler(mockStore)

	link := models.Link{
		FullUrl: "https://google.com",
	}
	body, _ := json.Marshal(link)
	req, _ := http.NewRequest("POST", "/add-url", bytes.NewBuffer(body))

	// Add userId to context
	ctx := context.WithValue(req.Context(), types.ContextKey("userId"), "123")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.AddUrl(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.Link
	json.NewDecoder(rr.Body).Decode(&response)

	if response.Id != 1 {
		t.Errorf("Expected link ID 1, got %v", response.Id)
	}
	if response.FullUrl != link.FullUrl {
		t.Errorf("Expected FullUrl %s, got %s", link.FullUrl, response.FullUrl)
	}
	if response.ShortUrl == "" {
		t.Errorf("ShortUrl should not be empty")
	}
}

func TestGetAllLinks(t *testing.T) {
	mockStore := &MockShortenerStore{
		GetAllLinksFunc: func(userId string) ([]models.Link, error) {
			return []models.Link{
				{Id: 1, FullUrl: "https://google.com", ShortUrl: "abc"},
				{Id: 2, FullUrl: "https://github.com", ShortUrl: "def"},
			}, nil
		},
	}
	handler := NewShortenerHandler(mockStore)

	req, _ := http.NewRequest("GET", "/links", nil)
	ctx := context.WithValue(req.Context(), types.ContextKey("userId"), "123")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.GetAllLinks(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.Link
	json.NewDecoder(rr.Body).Decode(&response)

	if len(response) != 2 {
		t.Errorf("Expected 2 links, got %d", len(response))
	}
}

func TestGetLink(t *testing.T) {
	mockStore := &MockShortenerStore{
		GetLinkFunc: func(urlId int) (*models.Link, error) {
			return &models.Link{Id: urlId, FullUrl: "https://google.com"}, nil
		},
	}
	handler := NewShortenerHandler(mockStore)

	req, _ := http.NewRequest("GET", "/links/1", nil)
	ctx := context.WithValue(req.Context(), types.ContextKey("userId"), "123")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.GetLink(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.Link
	json.NewDecoder(rr.Body).Decode(&response)

	if response.Id != 1 {
		t.Errorf("Expected link ID 1, got %d", response.Id)
	}
}

func TestDeleteLink(t *testing.T) {
	mockStore := &MockShortenerStore{
		DeleteLinkFunc: func(urlId int, userId int) error {
			return nil
		},
	}
	handler := NewShortenerHandler(mockStore)

	req, _ := http.NewRequest("DELETE", "/links/1", nil)
	ctx := context.WithValue(req.Context(), types.ContextKey("userId"), "123")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.DeleteLink(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]any
	json.NewDecoder(rr.Body).Decode(&response)

	if response["ok"] != true {
		t.Errorf("Expected ok to be true")
	}
}
