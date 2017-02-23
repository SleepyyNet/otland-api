package otland_api

import "testing"

func TestGetForumThreads(t *testing.T) {
	// Get first page of discussion forum
	threads, err := GetForumThreads(251, 0)

	if err != nil {
		t.Fatal(err)
	}

	// Check for any information
	if len(threads.List) <= 0 {
		t.Fatal("Number of threads cant be empty")
	}
}