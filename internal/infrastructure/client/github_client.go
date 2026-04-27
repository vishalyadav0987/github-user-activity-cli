package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "github.com/vishalyadav0987/github-activity/internal/domain/github-activity"
)

type GitHubClient struct{}

func NewGitHubClient() *GitHubClient {
	return &GitHubClient{}
}

func (c *GitHubClient) FetchUserEvents(username string) ([]*domain.Activity, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user not found or API error: %d", resp.StatusCode)
	}

	var activities []*domain.Activity

	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return activities, nil
}
