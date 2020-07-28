package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"fmt"
	"sync"
	"getReddit/reddit"
)

// It contains all the stuff that is needed for preparing and formatting data which is fetched from reddit APIs
func GetJSON(url string, target interface{}) error {
	// Getting data
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Setting up the header
	req.Header.Set("User-Agent", "getReddit by Rahul")

	// Sending the http request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Body must be closed to prevent resource leak.
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func SaveJSON(data interface{}) error {
	dataJSON, _ := json.Marshal(data)
	err := ioutil.WriteFile("postData.json", dataJSON,0644)
	if err != nil {
		return nil
	}
	return nil
}

func ExecuteImageWorker(wg *sync.WaitGroup, imageTasks chan reddit.ImageTask) {

	for task := range imageTasks {
		// Create the directory
		_ = os.Mkdir("images", os.ModeDir)
		// Create the file
		out, err := os.Create("images/"+task.Name + ".jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer out.Close()
		
		// Get the data
		resp, err := http.Get(task.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Status code not ok")
			continue
		}
		fmt.Println("IMAGE DOWNLOADED!")
		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
	wg.Done()
}

func RemoveImages() error {
	err:= os.RemoveAll("images/")
	if err != nil {
		return err
	}
	return nil
}

