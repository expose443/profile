package model

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	GithubLink  string `json:"github_link"`
	Image       string `json:"image"`
}
