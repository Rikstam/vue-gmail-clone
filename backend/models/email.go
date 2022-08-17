package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type Email struct {
	Id       int    `json:"id" xml:"id"`
	From     string `json:"from" xml:"from"`
	Subject  string `json:"subject" xml:"subject"`
	Body     string `json:"body" xml:"body"`
	SentAt   string `json:"sentAt" xml:"sentAt"`
	Archived bool   `json:"archived" xml:"archived"`
	Read     bool   `json:"read" xml:"read"`
}

type Inbox []Email

func (inbox Inbox) Find(id int) *Email {
	for index, email := range inbox {
		if email.Id == id {
			return &inbox[index]
		}
	}
	return nil
}

func NewEmails(reader io.Reader) (Inbox, error) {
	var emails Inbox

	err := json.NewDecoder(reader).Decode(&emails)

	if err != nil {
		err = fmt.Errorf("problem parsing email, %v", err)
	}

	return emails, err
}
