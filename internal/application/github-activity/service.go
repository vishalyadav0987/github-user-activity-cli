package githubactivity

import (
	"context"

	domain "github.com/vishalyadav0987/github-activity/internal/domain/github-activity"
)

type ActivityService interface {
	GetUserActivity(ctx context.Context, username string) ([]*domain.Activity, error)
}
