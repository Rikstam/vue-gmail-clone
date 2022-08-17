package config

import (
	"gmail-clone-backend/filesystem_store"
	"log"
	"os"
)

const dbFileName = "db.json"

func InitStore() *filesystem_store.FileSystemEmailStore {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := filesystem_store.NewFileSystemEmailStore(db)

	return store
}
