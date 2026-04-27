package main

import (
	"fmt"

	"github.com/vishalyadav0987/github-activity/interfaces/cli"
	"github.com/vishalyadav0987/github-activity/internal/infrastructure/client"
	githubactivity "github.com/vishalyadav0987/github-activity/internal/infrastructure/github-activity"
)

func main() {
	fmt.Println("Wokring on github user activity CLI tool.")

	githubClient := client.NewGitHubClient()

	activityService := githubactivity.NewActivityService(githubClient)

	handler := cli.NewHandler(activityService)

	handler.Run()
}
