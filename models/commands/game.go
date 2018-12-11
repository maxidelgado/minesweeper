package commands

type CreateGameCommand struct {
	Email string `json:"email"`
	Size  int    `json:"size"`
}
