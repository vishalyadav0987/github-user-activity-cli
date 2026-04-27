package githubactivity

type Activity struct {
	Id        string                 `json:"id"`
	Type      string                 `json:"type"`
	Actor     map[string]interface{} `json:"actor"`
	Repo      map[string]interface{} `json:"repo"`
	Payload   map[string]interface{} `json:"payload"`
	CreatedAt string                 `json:"created_at"`
}
