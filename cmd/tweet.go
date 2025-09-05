package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dghubble/oauth1"
)

var consumerKey = os.Getenv("CONSUMER_KEY")
var consumerSecret = os.Getenv("CONSUMER_SECRET")
var accessToken = os.Getenv("ACCESS_TOKEN")
var accessSecret = os.Getenv("ACCESS_SECRET")

var config = oauth1.NewConfig(consumerKey, consumerSecret)
var token = oauth1.NewToken(accessToken, accessSecret)
var httpClient = config.Client(oauth1.NoContext, token)

type Body struct {
	Text string `json:"text"`
}

func tweet(m *model) tea.Model {
	m.done = true

	tweetText := m.textarea.Value()

	body := Body{
		Text: tweetText,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		m.err = err
		return m
	}

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(jsonBody))
	if err != nil {
		m.err = err
		return m
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		m.err = err
		return m
	}
	defer resp.Body.Close()

	return m
}
