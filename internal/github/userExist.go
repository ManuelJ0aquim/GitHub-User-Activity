package github

import (
	"net/http"
)

func UserExists(username string) (bool, error) {
	url := "https://api.github.com/users/" + username

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, nil
}
