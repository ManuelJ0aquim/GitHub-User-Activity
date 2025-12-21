package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/github"
	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/model"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./githubUserActivity <github-username>")
		return
	}

	username := os.Args[1]
	userExists, err := github.UserExists(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !userExists {
		fmt.Println("User does not exist:", username)
		return
	}

	events, err := github.FetchEvents(username)
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			var payload model.PushPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				fmt.Println("Error decoding PushEvent payload:", err)
				continue
			}
			fmt.Printf("Pushed %d commits to %s\n", len(payload.Commits), event.Repo.Name)

		case "PullRequestEvent":
			var payload model.PullRequestPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				fmt.Println("Error decoding PullRequestEvent payload:", err)
				continue
			}
			fmt.Printf("Pull request #%d in %s\n", payload.PullRequest.Number, event.Repo.Name)

		case "IssuesEvent":
			var payload model.IssuePayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				fmt.Println("Error decoding IssuesEvent payload:", err)
				continue
			}
			fmt.Printf("Opened issue #%d in %s\n", payload.Issue.Number, event.Repo.Name)

		case "IssueCommentEvent":
			var payload model.IssuePayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				fmt.Println("Error decoding IssueCommentEvent payload:", err)
				continue
			}
			fmt.Printf("Commented on issue #%d in %s\n", payload.Issue.Number, event.Repo.Name)

		case "WatchEvent":
			fmt.Printf("Starred %s\n", event.Repo.Name)

		default:
			fmt.Println("Other event:", event.Type)
		}
	}
}
