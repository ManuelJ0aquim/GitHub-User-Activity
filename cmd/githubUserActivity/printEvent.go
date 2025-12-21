package main

import (
	"encoding/json"
	"fmt"

	"github.com/ManuelJ0aquim/GitHub-User-Activity/internal/model"
)

func PrintEvent(event model.Event) {
	switch event.Type {

	case "PushEvent":
		var payload model.PushPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return
		}
		fmt.Printf("Pushed %d commits to %s\n", len(payload.Commits), event.Repo.Name)

	case "PullRequestEvent":
		var payload model.PullRequestPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return
		}
		fmt.Printf("Pull request #%d in %s\n", payload.PullRequest.Number, event.Repo.Name)

	case "IssuesEvent":
		var payload model.IssuePayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return
		}
		fmt.Printf("Opened issue #%d in %s\n", payload.Issue.Number, event.Repo.Name)

	case "IssueCommentEvent":
		var payload model.IssuePayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return
		}
		fmt.Printf("Commented on issue #%d in %s\n", payload.Issue.Number, event.Repo.Name)

	case "WatchEvent":
		fmt.Printf("Starred %s\n", event.Repo.Name)
	}
}
