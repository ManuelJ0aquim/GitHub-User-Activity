package github

import (
	"encoding/json"
	"net/http"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/model"
)

func FetchFollowers(username string) ([]model.User, error) {
	url := "https://api.github.com/users/" + username + "/followers"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var followers []model.User
	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		return nil, err
	}

	return followers, nil
}
