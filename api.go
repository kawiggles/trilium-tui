package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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

func (s *Session) GetNote(noteId string) (string, string) {
	target := fmt.Sprintf("%s/notes/%s", s.endpoint, noteId)
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		log.Fatal("Error generating api request headers")
	}

	req.Header.Set("Authorization", s.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "applciation/json")

	resp1, err := s.client.Do(req)
	if err != nil {
		log.Printf("Error getting note metadata: %v\n", err)
		return "<p>error</p>", "nil"
	}
	defer resp1.Body.Close()
	metadata, _ := io.ReadAll(resp1.Body)

	url, _ := url.Parse(fmt.Sprintf("%s/content", target))
	req.URL = url
	resp2, err := s.client.Do(req)
	if err != nil {
		log.Printf("Error getting note content: %v\n", err)
		return string(metadata), "nil"
	}
	defer resp2.Body.Close()
	body, _ := io.ReadAll(resp2.Body)

	return string(metadata), string(body)
}
