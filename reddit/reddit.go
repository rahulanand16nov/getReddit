package reddit

import ()

// This format will be extracted from the heavy JSON we get from the API

type Post struct {
	Data struct {
		Subreddit string `json: "subreddit"`
		Title     string `json: "title"`
		Author    string `json: "author"`
		Permalink string `json: "permalink"`
		Url       string `json: "url"`
	} `json: "data"`
}

type Listing struct {
	Data struct {
		Childrens []Post `json: "children"`
	} `json : "data"`
}
