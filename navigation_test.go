package otland_api

import "testing"

func TestGetForumList(t *testing.T) {
	// Get forum list
	forums, err := GetForumList()

	if err != nil {
		t.Fatal(err)
	}

	// Check for any information
	if len(forums.List) <= 0 {
		t.Fatal("Number of forums cant be empty")
	}
}
