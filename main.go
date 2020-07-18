package main

import (
	"fmt"
	"getReddit/http"
	"getReddit/reddit"
)

func main() {
	fmt.Println("getReddit - v0.1")
	mainListing := reddit.Listing{}
	http.GetJSON("https://www.reddit.com/.json", &mainListing)
	for i := 0; i < len(mainListing.Data.Children); i++ {
		url_len := len(mainListing.Data.Children[i].Data.Media_URL)
		if url_len > 3 && mainListing.Data.Children[i].Data.Media_URL[url_len-3:] == "jpg" {
			err := http.DownloadImages(mainListing.Data.Children[i].Data.Name, mainListing.Data.Children[i].Data.Media_URL)
			mainListing.Data.Children[i].Data.Media_type = "Img";
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("DOWNLOADED IMAGE!")
		}
	}
	err := http.SaveJSON(&mainListing.Data.Children)
	if err == nil {
		fmt.Println("Saved data to data.json file")
	}
}
