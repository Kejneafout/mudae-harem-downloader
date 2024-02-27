package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Character struct {
	Rank    string `json:"rank"`
	Name    string `json:"name"`
	Series  string `json:"series"`
	Value   string `json:"value"`
	Note    string `json:"note"`
	Image   string `json:"image"`
}

type Metadata struct {
	Title string `json:"title"`
	Total string `json:"total"`
}

type Data struct {
	Metadata   Metadata    `json:"metadata"`
	Characters []Character `json:"characters"`
}

func FetchSeriesValues(inputFile string) (data Data, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var head []string
	var body []string

	// Read the first 3 lines for metadata
	for i := 0; i < 3 && scanner.Scan(); i++ {
		head = append(head, scanner.Text())
	}

	// Read the rest of the lines for character data
	for scanner.Scan() {
		body = append(body, scanner.Text())
	}

	title := strings.TrimSpace(head[0])
	total := regexp.MustCompile(`\d+`).FindString(head[2])

	data.Metadata.Title = title
	data.Metadata.Total = total

	var characters []Character

	regex := regexp.MustCompile(`(#\d+) - (.+?) - (.+?) (\d+ ka)`)

	for _, line := range body {
		if strings.TrimSpace(line) == "" {
			continue
		}

		elements := regex.FindStringSubmatch(line)
		if elements != nil {
			character := Character{
				Rank:   elements[1],
				Name:   elements[2],
				Series: elements[3],
				Value:  elements[4],
			}
			characters = append(characters, character)
		} else {
			fmt.Printf("Error parsing SV line: %s\n", line)
		}
	}

	data.Characters = characters

	return data, scanner.Err()
}
