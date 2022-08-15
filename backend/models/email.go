package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type Email struct {
	Id       int    `json:"id" xml:"id"`
	From     string `json:"from" xml:"from"`
	Subject  string `json:"subjext" xml:"subject"`
	Body     string `json:"body" xml:"body"`
	SentAt   string `json:"sentAt" xml:"sentAt"`
	Archived bool   `json:"archived" xml:"archived"`
	Read     bool   `json:"read" xml:"read"`
}

func NewEmails(reader io.Reader) ([]Email, error) {
	var emails []Email

	err := json.NewDecoder(reader).Decode(&emails)

	if err != nil {
		err = fmt.Errorf("problem parsing email, %v", err)
	}

	return emails, err
}
