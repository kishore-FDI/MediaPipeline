package db

import (
	"crypto/rand"
	"encoding/hex"
)

type Business struct {
	ID        int
	Name      string
	Email     string
	APIKey    string
	CreatedAt string
}

// Generate a random API key
func GenerateAPIKey() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Insert a new business
func CreateBusiness(name, email string) (*Business, error) {
	apiKey := GenerateAPIKey()
	res, err := SQLDB.Exec("INSERT INTO business (name, email, api_key) VALUES (?, ?, ?)", name, email, apiKey)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return &Business{ID: int(id), Name: name, Email: email, APIKey: apiKey}, nil
}
