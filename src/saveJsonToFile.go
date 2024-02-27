package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func SaveJsonToFile(data Data, filename string) (err error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("JSON data saved to %s\n", filename)
	return nil
}
