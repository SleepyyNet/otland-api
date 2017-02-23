package otland_api

import (
	"net/url"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type forumThreads struct {
	Threads []thread `json:"threads"`
	ThreadsTotal uint `json:"threads_total"`
	Links links
}

type links struct {
	Pages uint `json:"pages"`
	Page uint `json:"page"`
	Previous string `json:"prev"`
}

type thread struct {
	ThreadID uint `json:"thread_id"`
	ForumID uint `json:"forum_id"`
	ThreadTitle string `json:"thread_title"`
	ThreadViewCount uint `json:"thread_view_count"`
	CreatorUserID uint `json:"creator_user_id"`
	CreatorUsername string `json:"creator_username"`
	ThreadCreateDate uint64 `json:"thread_create_date"`
	ThreadUpdateDate uint64 `json:"thread_update_date"`
	ThreadPostCount uint `json:"thread_post_count"`
	ThreadTags map[string]string `json:"thread_tags"`
	Links threadLinks `json:"links"`
}

type threadLinks struct {
	PermaLink string `json:"permalink"`
	Detail string `json:"detail"`
	Followers string `json:"followers"`
	Forum string `json:"forum"`
	Posts string `json:"posts"`
	FirstPost string `json:"first_post"`
	LastPost string `json:"last_post"`
}

// GetForumThreads retrieves the given forum threads
func GetForumThreads(forum, page int) (*forumThreads, error) {
	// Create request URL
	baseUrl, err := url.Parse(APIUrl + "/threads")

	if err != nil {
		return nil, err
	}

	// Get base URL query
	baseUrlQuery := baseUrl.Query()

	// Set forum id and page
	baseUrlQuery.Set("forum_id", strconv.Itoa(forum))
	baseUrlQuery.Set("page", strconv.Itoa(page))

	// Set base url query
	baseUrl.RawQuery = baseUrlQuery.Encode()

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

	// Thread data holder
	threads := forumThreads{}

	// Unmarshal response
	if err := json.Unmarshal(body, &threads); err != nil {
		return nil, err
	}

	return &threads, nil
}