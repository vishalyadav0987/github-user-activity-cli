package githubactivity

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidUsername        = errors.New("invalid username")
	ErrExternalServiceFailure = errors.New("failed to fetch data from github")
	ErrRateLimitExceeded      = errors.New("github api rate limit exceeded")
	ErrNoActivityFound        = errors.New("no recent activity found")
)
