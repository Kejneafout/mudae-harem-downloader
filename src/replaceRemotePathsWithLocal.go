package src

import (
	"fmt"
	"path"
	"strings"
)

func ReplaceRemotePathsWithLocal(data *Data, imagesDirectory string) (err error) {
	for index := range data.Characters {
		rank := strings.TrimLeft(data.Characters[index].Rank, "#")
		name := strings.ReplaceAll(data.Characters[index].Name, " ", "_")
		extension := path.Ext(data.Characters[index].Image)
		localImagePath := fmt.Sprintf("images/%d_%s_%s%s", index+1, rank, name, extension)
		data.Characters[index].Image = localImagePath
	}

	// Save modified data to JSON file
	err = SaveJsonToFile(*data, "data.json")
	if err != nil {
		fmt.Printf("Error saving JSON data to file: %s\n", err)
		return err
	}

	return nil
}
