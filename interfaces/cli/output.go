package cli

import (
	"fmt"

	domain "github.com/vishalyadav0987/github-activity/internal/domain/github-activity"
)

func FormatActivity(a *domain.Activity) string {
	repoName := "unknown repo"

	if a.Repo != nil {
		if name, ok := a.Repo["name"].(string); ok {
			repoName = name
		}
	}

	switch a.Type {
	case "PushEvent":
		return fmt.Sprintf("Pushed commits to %s", repoName)
	case "IssuesEvent":
		return fmt.Sprintf("Opened an issue in %s", repoName)
	case "WatchEvent":
		return fmt.Sprintf("Starred %s", repoName)
	default:
		return fmt.Sprintf("Did %s on %s", a.Type, repoName)
	}
}
