package cli

import (
	domain "github.com/vishalyadav0987/github-activity/internal/domain/github-activity"
)

type Command struct {
	Name string
}

func Parse(username string) (*Command, error) {
	if len(username) < 3 {
		return nil, domain.ErrInvalidUsername
	}

	cmd := &Command{
		Name: username,
	}

	return cmd, nil
}
