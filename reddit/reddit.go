package reddit

import (
	"getReddit/http"
	
)

type Post struct {
	Data struct {
	  Domain string `json: "domain"`
	  Subreddit string `json: "subreddit"`
	  Title string `json: "title"`
	  Permalink string `json: "permalink"`
	  Url string `json: "url"`
	} `json: "data"`
}