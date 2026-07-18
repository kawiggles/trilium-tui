package main

import(
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Session struct {
	token string
	endpoint string
	client *http.Client
}

func StartSession() Session {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading environmental variables")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return Session{
		token: os.Getenv("TOKEN"),
		endpoint: fmt.Sprintf("http://%s:%s/etapi", host, port),
		client: client,
	}
}

func (s *Session) GetRoot() string {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/notes/root/content", s.endpoint), nil)
	if err != nil {
		log.Fatal("Error generating api request headers")
	}

	req.Header.Set("Authorization", s.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "applciation/json")

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("Error with API request: %v\n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
