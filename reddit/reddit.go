package reddit

// This format will be extracted from the heavy JSON we get from the API

type post struct {
	Data struct {
		Subreddit string   `json:"subreddit"`
		Title     string   `json:"title"`
		Author    string   `json:"author"`
		Name	  string   `json:"name"`
	} `json:"data"`
}

type Listing struct {
	Data struct {
		Children []post `json:"children"`
	} `json:"data"`
}
