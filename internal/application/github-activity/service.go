package githubactivity

import (
	domain "github.com/vishalyadav0987/github-activity/internal/domain/github-activity"
)

type ActivityService interface {
	GetUserActivity(username string) ([]*domain.Activity, error)
}
