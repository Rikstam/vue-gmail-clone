package filesystem_store

import (
	"gmail-clone-backend/models"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{
				"id": "1",
				"from": "team@vuemastery.com",
				"subject": "What's up with Vue 3.0? Here's how to find out from Evan You",
				"body": "The opening keynote of VueConf US this year was Evan You (the creator of Vue), giving his State of the Vuenion address. He walked us through the journey of getting Vue 3 from a prototype to a reality the past year. He also dove into Vue's overall growth in the community.",
				"sentAt": "2020-03-27T18:25:43.511Z",
				"archived": false,
				"read": false
			}]`)

		defer cleanDatabase()

		store := NewFileSystemEmailStore(database)

		got := store.GetEmails()

		want := []models.Email{
			{
				Id:       1,
				From:     "team@vuemastery.com",
				Subject:  "What's up with Vue 3.0? Here's how to find out from Evan You",
				Body:     "The opening keynote of VueConf US this year was Evan You (the creator of Vue), giving his State of the Vuenion address. He walked us through the journey of getting Vue 3 from a prototype to a reality the past year. He also dove into Vue's overall growth in the community.",
				SentAt:   "2020-03-27T18:25:43.511Z",
				Archived: false,
				Read:     false,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
