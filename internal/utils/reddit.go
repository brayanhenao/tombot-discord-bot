package utils

import (
	"encoding/json"
	"github.com/brayanhenao/tombot-discord-bot/internal/config"
	"log"
	"net/http"
	"time"
)

func GetRedditResponse(url string) (RedditResponse, error) {
	userAgent := "Reddit JSON API"

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RedditResponse{}, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("API_KEY", config.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	redditResponse := RedditResponse{}
	err = json.NewDecoder(resp.Body).Decode(&redditResponse)
	if err != nil {
		return RedditResponse{}, err
	}

	return redditResponse, nil
}

type RedditResponse struct {
	Data struct {
		Children []ResponseData `json:"children,omitempty"`
	} `json:"data"`
}

type ResponseData struct {
	Data map[string]interface{} `json:"data,omitempty"`
}
