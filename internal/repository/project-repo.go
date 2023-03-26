package repository

import (
	"github.com/with-insomnia/profile/internal/model"
)

func (r *Repository) InsertNewProject(project model.Project) error {
	query := `INSERT INTO projects(title, description, user_id, github_link, image, created_at, updated_at)
	VALUES
	(
		$1,
		$2,
		(SELECT user_id FROM users WHERE username=$3),
		$4,
		$5,
		$6,
		$7
	);
	`
	_, err := r.db.Exec(query, project.Title, project.Description, project.Author, project.GithubLink, project.Image, project.Created, project.Updated)
	if err != nil {
		return err
	}
	return nil
}
