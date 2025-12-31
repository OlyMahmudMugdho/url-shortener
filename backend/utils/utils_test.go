package utils

import (
	"testing"

	"github.com/OlyMahmudMugdho/url-shortener/models"
)

func TestHashPassword(t *testing.T) {
	password := "mypassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !IsPassWordValid(password, string(hash)) {
		t.Errorf("Password validation failed for correct password")
	}

	if IsPassWordValid("wrongpassword", string(hash)) {
		t.Errorf("Password validation succeeded for incorrect password")
	}
}

func TestGenerateUserResponseFromUser(t *testing.T) {
	user := models.User{
		Username:  "testuser",
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Password:  "secret",
	}

	response := GenerateUserResponseFromUser(user)

	if response.Username != user.Username {
		t.Errorf("Expected username %s, got %s", user.Username, response.Username)
	}
	if response.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, response.Email)
	}
	if response.FirstName != user.FirstName {
		t.Errorf("Expected first name %s, got %s", user.FirstName, response.FirstName)
	}
	if response.LastName != user.LastName {
		t.Errorf("Expected last name %s, got %s", user.LastName, response.LastName)
	}
}

func TestGetUserFromUserRequest(t *testing.T) {
	request := &models.UserRequestBody{
		Username:  "testuser",
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Password:  "secret",
	}

	user := GetUserFromUserRequest(request)

	if user.Username != request.Username {
		t.Errorf("Expected username %s, got %s", request.Username, user.Username)
	}
	if user.Email != request.Email {
		t.Errorf("Expected email %s, got %s", request.Email, user.Email)
	}
	if user.Password != request.Password {
		t.Errorf("Expected password %s, got %s", request.Password, user.Password)
	}
}

func TestExtractParamFromUrl(t *testing.T) {
	tests := []struct {
		url      string
		prefix   string
		expected string
		found    bool
	}{
		{"/api/v1/urls/123", "/api/v1/urls/", "123", true},
		{"/api/v1/urls/abc", "/api/v1/urls/", "abc", true},
		{"/other/path", "/api/", "/other/path", false},
	}

	for _, tt := range tests {
		result, found := ExtractParamFromUrl(tt.url, tt.prefix)
		if result != tt.expected || found != tt.found {
			t.Errorf("ExtractParamFromUrl(%s, %s) = (%s, %v), expected (%s, %v)",
				tt.url, tt.prefix, result, found, tt.expected, tt.found)
		}
	}
}

func TestBase62Encode(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "a"},
		{1, "b"},
		{61, "9"},
		{62, "ba"},
	}

	for _, tt := range tests {
		result := Base62Encode(tt.input)
		if result != tt.expected {
			t.Errorf("Base62Encode(%d) = %s, expected %s", tt.input, result, tt.expected)
		}
	}
}

func TestGenerateShortUrl(t *testing.T) {
	url := "https://example.com"
	short1 := GenerateShortUrl(url)
	short2 := GenerateShortUrl(url)

	if short1 == "" {
		t.Errorf("Generated short URL is empty")
	}

	// Since we add time.Now().String(), they should be different even for same URL
	if short1 == short2 {
		t.Errorf("Generated short URLs should be unique even for the same input due to timestamp")
	}

	if len(short1) > 7 {
		t.Errorf("Generated short URL length %d exceeds 7 characters", len(short1))
	}
}
