package repository

import "github.com/chirag3003/go-backend-template/db"

// Repository aggregates all data access repositories.
type Repository struct {
	User  UserRepository
	Media MediaRepository
	S3    S3Repository
	Link  LinkRepository
	Click ClickRepository
}

// NewRepository creates all repositories with the given database connection.
func NewRepository(conn db.Connection) *Repository {
	return &Repository{
		User:  NewUserRepository(conn),
		Media: NewMediaRepository(conn),
		S3:    NewS3Repository(),
		Link:  NewLinkRepository(conn),
		Click: NewClickRepository(conn),
	}
}
