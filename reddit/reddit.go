package reddit

import "encoding/json"
// This format will be extracted from the heavy JSON we get from the API

type post struct {
	Data struct {
		Subreddit string   `json:"subreddit"`
		Title     string   `json:"title"`
		Author    string   `json:"author"`
		Name	  string   `json:"name"`
		Up_votes  int32    `json:"ups"`
		Comments  int32    `json:"num_comments"`
		Created    json.Number `json:"created_utc"`
		Media_URL string   `json:"url_overridden_by_dest"`
		Media_type string
	} `json:"data"`
}

type Listing struct {
	Data struct {
		Children []post `json:"children"`
	} `json:"data"`
	After string `json:"after"`
}

type MainListing struct {
	Posts []post
}

type ImageTask struct {
	Url string
	Name string
}