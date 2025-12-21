package github

import (
	"encoding/json"
	"net/http"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/model"
)

func FetchRepos(username string) ([]model.Repo, error) {
	url := "https://api.github.com/users/" + username + "/repos"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []model.Repo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}
