# getReddit
Concurrency based Reddit Scraper in Go
Note: This repo is meant to be used in conjunction [reddit-frame](https://github.com/rahulanand16nov/reddit-frame)

## See this in action
[![Watch getReddit in Action](https://i.imgur.com/ZwUbcxW.png)](https://drive.google.com/file/d/1v0s0y4S7AJLRPDZdbq81jYuYTkyF3Afu/view)

## What's the point?
* I wanted to learn Golang and why it's praised a lot in distributed & parallel computing community.
* I hated the fact that reddit sometimes consumes more data than youtube (hats off to youtube engineers!).
* Showcase my skills as an aspiring software engineer.

## Running the project
Make sure you have downloaded and install GoLang from [official website](https://golang.org/dl/)

### Steps
* Clone the project
* CD into the project
* Run `go run main.go *subreddit_names_with_space*`. Eg: `go run main.go pics askreddit meme`

You will get `postData.json` and `images` directory containing images scraped from given subreddits

## Why reddit?
* First, I love reddit!
* Second, Reddit has a very friendly APIs which allowed me to focus on my main objectives.