package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

func DownloadImages(characters []Character, imagesDirectory string) (err error) {
	if len(characters) == 0 {
		fmt.Println("No data to download images.")
		return nil
	}

	if _, err := os.Stat(imagesDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(imagesDirectory, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating images directory: %s\n", err)
			return err
		}
	}

	// Create a custom HTTP client with a custom User-Agent header
	client := &http.Client{}

	for index, character := range characters {
		imageUrl := character.Image
		rank := regexp.MustCompile(`\d+`).FindString(character.Rank)
		name := strings.ReplaceAll(character.Name, " ", "_")
		extension := path.Ext(imageUrl)
		localImagePath := path.Join(imagesDirectory, fmt.Sprintf("%d_%s_%s%s", index+1, rank, name, extension))

		// Create a GET request with the custom client
		req, err := http.NewRequest("GET", imageUrl, nil)
		if err != nil {
			fmt.Printf("Error creating request for %s: %s\n", character.Name, err)
			continue
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

		// Perform the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error downloading image for %s: %s\n", character.Name, err)
			continue
		}
		defer resp.Body.Close()

		// Read the response body
		imageData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading image data for %s: %s\n", character.Name, err)
			continue
		}

		// Write image data to local file
		err = ioutil.WriteFile(localImagePath, imageData, 0644)
		if err != nil {
			fmt.Printf("Error saving image for %s: %s\n", character.Name, err)
			continue
		}

		fmt.Printf("Image for %s downloaded and saved to %s\n", character.Name, localImagePath)
	}

	return nil
}
