package main

import (
	"fmt"
	"getReddit/http"
	"getReddit/reddit"
)

func main() {
	fmt.Println("getReddit - v0.1")
	mainListing := reddit.Listing{}
	http.GetJSON("https://www.reddit.com/r/askReddit/.json", &mainListing)
	fmt.Printf("Result: \n %+v", mainListing.Data.Children)
}
