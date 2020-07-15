package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func SaveJSON(data interface{}) error {
	dataJSON, _ := json.Marshal(data)
	err := ioutil.WriteFile("data.json", dataJSON,0644)
	if err != nil {
		return nil
	}
	return nil
}

func DownloadFile(filename string, url string) error {
	// Create the directory
	_ = os.Mkdir("Images", os.ModeDir)
	// Create the file
	out, err := os.Create("Images/"+filename + ".jpg")
	if err != nil {
		return err
	}
	defer out.Close()
	
	// Get the data
	resp, err := http.Get(url)
    if err != nil {
        return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
	}
	
	return nil
}

func RemoveImages() error {
	err:= os.RemoveAll("Images/")
	if err != nil {
		return err
	}
	return nil
}

