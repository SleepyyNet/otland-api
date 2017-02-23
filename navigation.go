package otland_api

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/json"
)

type forumList struct {
	List []forum `json:"elements"`
	Count uint `json:"element_count"`
}

type forum struct {
	ID uint `json:"forum_id"`
	Title string `json:"forum_title"`
	Description string `json:"forum_description"`
	ThreadCount uint `json:"forum_thread_count"`
	PostCount uint `json:"forum_post_count"`
	Links forumLinks `json:"links"`
}

type forumLinks struct {
	PermaLink string `json:"permalink"`
	Detail string `json:"detail"`
	Threads string `json:"threads"`
	Followers string `json:"followers"`
}

// GetForumList returns a complete list with all the forums
func GetForumList() (*forumList, error) {
	// Create request URL
	baseUrl, err := url.Parse(APIUrl + "/navigation")

	if err != nil {
		return nil, err
	}

	// Make GET request
	response, err := http.Get(baseUrl.String())

	if err != nil {
		return nil, err
	}

	// Close response body
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	// Forums data holder
	forums := forumList{}

	// Unmarshal response
	if err := json.Unmarshal(body, &forums); err != nil {
		return nil, err
	}

	return &forums, nil
}