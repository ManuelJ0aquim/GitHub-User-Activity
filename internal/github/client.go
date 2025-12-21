package github

import (
	"encoding/json"
	"net/http"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/model"
)

func FetchEvents(username string) ([]model.Event, error) {
	url := "https://api.github.com/users/" + username + "/events"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []model.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}
