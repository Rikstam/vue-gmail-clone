package filesystem_store

import (
	"gmail-clone-backend/models"
	"io"
)

type FileSystemEmailStore struct {
	database io.ReadWriteSeeker
	emails   []models.Email
}

func (f *FileSystemEmailStore) GetEmails() []models.Email {
	return f.emails
}

func NewFileSystemEmailStore(database io.ReadWriteSeeker) *FileSystemEmailStore {
	database.Seek(0, 0)
	emails, _ := models.NewEmails(database)

	return &FileSystemEmailStore{
		database: database,
		emails:   emails,
	}
}
