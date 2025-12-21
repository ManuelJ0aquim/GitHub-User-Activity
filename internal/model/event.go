package model

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Repo    Repo            `json:"repo"`
	Payload json.RawMessage `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type PushPayload struct {
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Message string `json:"message"`
}

type IssuePayload struct {
	Issue Issue `json:"issue"`
}

type Issue struct {
	Number int `json:"number"`
}

type PullRequestPayload struct {
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Number int `json:"number"`
}

type WatchPayload struct {
	Action string `json:"action"`
}
