package cli

import (
	"context"
	"fmt"
	"os"

	app "github.com/vishalyadav0987/github-activity/internal/infrastructure/github-activity"
)

type Handler struct {
	service *app.ActivityService
}

func NewHandler(service *app.ActivityService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Error: username is required")
		return
	}

	cmd, err := Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ctx := context.Background()

	activities, err := h.service.GetUserActivity(ctx, cmd.Name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, a := range activities {
		fmt.Println("-", FormatActivity(a))
	}
}
