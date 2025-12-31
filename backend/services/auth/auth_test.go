package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OlyMahmudMugdho/url-shortener/models"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

type MockUserStore struct {
	SaveUserFunc           func(user models.User) error
	FindUserByUserNameFunc func(username string) (models.User, error)
}

func (m *MockUserStore) SaveUser(user models.User) error {
	return m.SaveUserFunc(user)
}

func (m *MockUserStore) FindUserByUserName(username string) (models.User, error) {
	return m.FindUserByUserNameFunc(username)
}

func TestRegister(t *testing.T) {
	mockStore := &MockUserStore{
		SaveUserFunc: func(user models.User) error {
			return nil
		},
	}
	handler := NewAuthHandler(mockStore)

	userReq := models.UserRequestBody{
		Username:  "testuser",
		Password:  "password123",
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	body, _ := json.Marshal(userReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Register(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.UserResponseBody
	json.NewDecoder(rr.Body).Decode(&response)

	if response.Username != userReq.Username {
		t.Errorf("Expected username %s, got %s", userReq.Username, response.Username)
	}
}

func TestLogin(t *testing.T) {
	hashedPassword, _ := utils.HashPassword("password123")
	mockStore := &MockUserStore{
		FindUserByUserNameFunc: func(username string) (models.User, error) {
			return models.User{
				Id:       1,
				Username: "testuser",
				Password: string(hashedPassword),
			}, nil
		},
	}
	handler := NewAuthHandler(mockStore)

	loginReq := models.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	body, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Login(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.LoginResponse
	json.NewDecoder(rr.Body).Decode(&response)

	if !response.Ok {
		t.Errorf("Expected login response Ok to be true")
	}
	if response.Token == "" {
		t.Errorf("Expected token in login response")
	}
}

func TestLogOut(t *testing.T) {
	handler := NewAuthHandler(&MockUserStore{})

	req, _ := http.NewRequest("GET", "/logout", nil)
	rr := httptest.NewRecorder()
	handler.LogOut(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	found := false
	for _, cookie := range rr.Result().Cookies() {
		if cookie.Name == "token" && cookie.Value == "" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected token cookie to be cleared")
	}
}
