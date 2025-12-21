package main

import (
	"fmt"
	"os"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/github"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./githubUserActivity <github-username> <events | followers | following | repos>")
		return
	}

	username := os.Args[1]
	command := os.Args[2]

	userExists, err := github.UserExists(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !userExists {
		fmt.Println("User does not exist:", username)
		return
	}

	switch command {

	case "followers":
		followers, err := github.FetchFollowers(username)
		if err != nil {
			fmt.Println("Error fetching followers:", err)
			return
		}

		fmt.Println("Followers:")
		for _, follower := range followers {
			fmt.Println("-", follower.Login)
		}

	case "events":
		events, err := github.FetchEvents(username)
		if err != nil {
			fmt.Println("Error fetching events:", err)
			return
		}

		fmt.Println("Events:")
		for _, event := range events {
			PrintEvent(event)
		}

	case "following":
		following, err := github.FetchFollowing(username)
		if err != nil {
			fmt.Println("Error fetching following:", err)
			return
		}
		fmt.Println("Following:")
		for _, following := range following {
			fmt.Println("-", following.Login)
		}

	case "repos":
		repos, err := github.FetchRepos(username)
		if err != nil {
			fmt.Println("Error fetching repos:", err)
			return
		}
		fmt.Println("Repos:")
		for _, repo := range repos {
			fmt.Println("-", repo.Name)
		}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: events, followers, following, repos")
	}
}
