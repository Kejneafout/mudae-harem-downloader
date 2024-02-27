package main

import (
	"fmt"

	"mudae-harem-downloader/src"
)

func main() {
	dataJson, err := src.FetchSeriesValues("inputs/1_series_values.txt")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	err = src.FetchNotesImages("inputs/2_notes_images.txt", dataJson)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	err = src.SaveJsonToFile(dataJson, "data.json")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	err = src.DownloadImages(dataJson.Characters, "images/")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	err = src.ReplaceRemotePathsWithLocal(&dataJson, "images/")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	err = src.CreateZip("exports/", "data.json", "images/")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}
