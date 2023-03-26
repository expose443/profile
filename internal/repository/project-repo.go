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

func (r *Repository) GetAllProjects() ([]model.Project, error) {
	rows, err := r.db.Query(`SELECT projects.project_id, projects.title, projects.description, projects.github_link, projects.image, projects.created_at, projects.updated_at, users.username 
	FROM projects 
	JOIN users ON projects.user_id = users.user_id`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var result []model.Project
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.GithubLink, &project.Image, &project.Created, &project.Updated, &project.Author); err != nil {
			return nil, err
		}
		result = append(result, project)
	}
	return result, nil
}
