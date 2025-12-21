package github

import (
	"encoding/json"
	"net/http"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/model"
)

func FetchFollowing(username string) ([]model.User, error) {
	url := "https://api.github.com/users/" + username + "/following"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var following []model.User
	if err := json.NewDecoder(resp.Body).Decode(&following); err != nil {
		return nil, err
	}

	return following, nil
}
