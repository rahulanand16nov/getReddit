package http

import (
	"io"
	"net/http"
	"os"
	"strings"
	"encoding/json"
)

func getJSON(url string, target interface{}) error {
	// Getting data
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	return json.NewDecoder(r.Body).Decode(target)
}

func downloadFile(filepath string, url string) error {
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
