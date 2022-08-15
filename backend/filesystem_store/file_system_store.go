package filesystem_store

import (
	"gmail-clone-backend/models"
	"io"
)

type FileSystemEmailStore struct {
	database io.ReadWriteSeeker
	inbox    models.Inbox
}

func (f *FileSystemEmailStore) GetEmails() models.Inbox {
	return f.inbox
}

func (f *FileSystemEmailStore) GetEmail(id int) *models.Email {
	email := f.inbox.Find(id)

	if email != nil {
		return email
	}
	return nil
}

func NewFileSystemEmailStore(database io.ReadWriteSeeker) *FileSystemEmailStore {
	database.Seek(0, 0)
	emails, _ := models.NewEmails(database)

	return &FileSystemEmailStore{
		database: database,
		inbox:    emails,
	}
}
