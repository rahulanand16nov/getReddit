package main

import (
	"fmt"
	"getReddit/http"
	"getReddit/reddit"
)

func main() {
	fmt.Println("getReddit - v0.1")
	mainListing := reddit.Listing{}
	http.GetJSON("https://www.reddit.com/r/pics/.json", &mainListing)
	fmt.Print("Result: \n")
	for i := 1; i <= len(mainListing.Data.Children); i++ {
		fmt.Printf("\t Title: %s \n", mainListing.Data.Children[i-1].Data.Title)
		fmt.Printf("\t Author: %s \n", mainListing.Data.Children[i-1].Data.Author)
		fmt.Printf("\t Link: %s \n", mainListing.Data.Children[i-1].Data.URL)
		fmt.Printf("\t Type: %s \n", mainListing.Data.Children[i-1].Data.Category)
		if mainListing.Data.Children[i-1].Data.Category[0] == "photography" {
			err := http.DownloadFile(mainListing.Data.Children[i-1].Data.Name, mainListing.Data.Children[i-1].Data.URL)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("DOWNLOADED!")
		}
		fmt.Print("\n")
	}
	err := http.RemoveImages()
	if err == nil {
		fmt.Println("Removed all images")
	}
}
