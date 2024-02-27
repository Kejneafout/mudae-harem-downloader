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

	regex := regexp.MustCompile(`^(#\d+).{3}([^|]+)(?: \| (.+)).{2} (.+)`)

	for _, line := range body {
		if strings.TrimSpace(line) == "" {
			continue
		}

		elements := regex.FindStringSubmatch(line)
		if elements != nil {
			for i, character := range data.Characters {
				if character.Name == elements[2] {
					data.Characters[i].Note = elements[3]
					data.Characters[i].Image = elements[4]
					break
				}
			}
		} else {
			fmt.Printf("Error parsing NI line: %s\n", line)
		}
	}

	return scanner.Err()
}
