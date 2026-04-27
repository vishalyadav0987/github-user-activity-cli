package githubactivity

import (
	"context"
	"strings"

	domain "github.com/vishalyadav0987/github-activity/internal/domain/github-activity"
)

type ActivityFetcher interface {
	FetchUserEvents(username string) ([]*domain.Activity, error)
}

// Implementation
type ActivityService struct {
	fetcher ActivityFetcher
}

// Constructor
func NewActivityService(fetcher ActivityFetcher) *ActivityService {
	return &ActivityService{
		fetcher: fetcher,
	}
}

func (s *ActivityService) GetUserActivity(ctx context.Context, username string) ([]*domain.Activity, error) {

	// 🔹 Validation (domain rule)
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, domain.ErrInvalidUsername
	}

	// 🔹 Call infrastructure via abstraction
	activities, err := s.fetcher.FetchUserEvents(username)
	if err != nil {
		return nil, err // ideally map kar sakta hai (advanced)
	}

	// 🔹 Business rule
	if len(activities) == 0 {
		return nil, domain.ErrNoActivityFound
	}

	return activities, nil
}
