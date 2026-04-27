package githubactivity

type Activity struct {
	Id        string
	Type      string
	Actor     map[string]interface{}
	Repo      map[string]interface{}
	Payload   map[string]interface{}
	CreatedAt string
}
