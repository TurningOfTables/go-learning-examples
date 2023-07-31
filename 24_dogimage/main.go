// Makes a request to the dog.ceo website for a random image and saves it locally

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://dog.ceo/api/breeds/image/random"

type DogImageInfo struct {
	URL    string `json:"message"`
	Status string
}

func main() {
	dogImageInfo := imageInfo()
	fmt.Println(dogImageInfo.URL)
	fileName := determineFilename(dogImageInfo.URL)
	imageResponse := retrieveImage(dogImageInfo.URL)
	saveImageToFile(imageResponse, fileName)
}

func imageInfo() DogImageInfo {
	fmt.Printf("Requesting Dog Image Info...")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error requesting Dog Image Info: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("...success\n")
	} else {
		fmt.Printf("...failed\n")
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
	}

	var result DogImageInfo
	json.Unmarshal(body, &result)

	return result
}

func retrieveImage(url string) *http.Response {
	fmt.Printf("Retrieving Dog Image...")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
		resp.Body.Close()
	}

	fmt.Printf("...success")
	return resp
}

func saveImageToFile(resp *http.Response, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error saving file contents: %v", err)
	}
	defer resp.Body.Close()
}

func determineFilename(url string) string {
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[i+1:]
			break
		}
	}

	return url
}
