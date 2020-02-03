package http

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
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

func DownloadFile(filepath string, url string) error {

	// Creating file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Getting data
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	//Writting data to a file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		return err
	}

	return nil

}
