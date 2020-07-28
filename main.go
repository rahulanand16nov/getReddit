package main

import (
	"fmt"
	"os"
	"getReddit/http"
	"getReddit/reddit"
	"sync"
)

const MAX_IMAGE_WORKERS int = 10

func main() {
	fmt.Println("getReddit - v1.0")
	subreddits := os.Args[1:]
	mainListing := reddit.MainListing{}

	// Get the JSON data of subreddits and join them into main listing
	for _, subreddit := range subreddits {
		listing := reddit.Listing{}
		url := "https://reddit.com/r/"+subreddit+"/.json?limit=100"
		http.GetJSON(url, &listing)

		mainListing.Posts = append(mainListing.Posts[:], listing.Data.Children...)
	}

	// Remove already downloaded images
	http.RemoveImages()

	// Semaphores
	var wg sync.WaitGroup
	wg.Add(1)

	// Adding image urls into the channel to be downloaded by workers
	imageTasks := make(chan reddit.ImageTask)
	go func() {
		for _, post := range mainListing.Posts {
			url_len := len(post.Data.Media_URL)
			if url_len > 3 && post.Data.Media_URL[url_len-3:] == "jpg" {
				post.Data.Media_type = "Img"
				task := reddit.ImageTask{Url: post.Data.Media_URL, Name: post.Data.Name}
				imageTasks <- task
			}
		}
		close(imageTasks)
		wg.Done()
	}()

	for i:=1; i <= MAX_IMAGE_WORKERS; i++ {
		wg.Add(1)
		go http.ExecuteImageWorker(&wg,imageTasks)
	}

	wg.Wait()

	err := http.SaveJSON(&mainListing)
	if err == nil {
		fmt.Println("Saved data to data.json file")
	}
}
