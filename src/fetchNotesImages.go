package src

import (
        "bufio"
        "fmt"
        "os"
        "regexp"
        "strings"
)

func FetchNotesImages(inputFile string, data Data) (err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var body []string

	// Skip the first two lines
	for i := 0; i < 2 && scanner.Scan(); i++ {
		_ = scanner.Text()
	}
	// Read the rest of the lines for note and image data
	for scanner.Scan() {
		body = append(body, scanner.Text())
	}

	regex := regexp.MustCompile(`^(.+?)\s*(?:\|\s*(.+?))?\s*-\s*(https?:\/\/\S+)`)

	for _, line := range body {
		if strings.TrimSpace(line) == "" {
			continue
		}

		elements := regex.FindStringSubmatch(line)
		if elements != nil {
			for i, character := range data.Characters {
				if character.Name == elements[1] {
					data.Characters[i].Note = elements[2]
					data.Characters[i].Image = elements[3]
					break
				}
			}
		} else {
			fmt.Printf("Error parsing notes_images line: %s\n", line)
		}
	}

	return scanner.Err()
}
