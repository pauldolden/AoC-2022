package day3

import (
	"advent22/helpers"
	"fmt"
	"strings"
	"unicode"
)

var row_segments = [][]string{}
var duplicates = [][]string{}

func Day3Part1() {
	input := helpers.GetPuzzleInput("3")
	priorities := setPriorities()
	rows := strings.Split(strings.TrimSpace(input), "\n")

	for _, row := range rows {
		row_length := len(row)
		row_left := row[:row_length/2]
		row_right := row[row_length/2:]
		row_halves := []string{row_left, row_right}
		row_segments = append(row_segments, row_halves)
	}

	for _, segment := range row_segments {
		segment := segment
		chars_to_compare := strings.Split(segment[0], "")
		local_duplicates := []string{}
		for _, c := range chars_to_compare {
			if strings.Contains(segment[1], c) {
				if !strings.Contains(strings.Join(local_duplicates, ""), c) {
					local_duplicates = append(local_duplicates, c)
				}
			}
		}
		duplicates = append(duplicates, local_duplicates)
	}
	count := 0
	for _, i := range duplicates {
		for _, j := range i {
			count = count + priorities[j]
		}
	}
	fmt.Println(count)
}

func Day3Part2() {
	input := helpers.GetPuzzleInput("3")
	priorities := setPriorities()
	badges := []string{}
	ss := strings.Split(input, "\n")

	chunks := chunkSlice(ss, 3)

	for _, chunk := range chunks {
		cs := strings.Split(chunk[0], "")
		var badge string
		for _, c := range cs {
			if strings.Contains(chunk[1], c) && strings.Contains(chunk[2], c) {
				badge = c
				break
			}
		}
		badges = append(badges, badge)
	}
	count := 0
	for _, badge := range badges {
		count = count + priorities[badge]
	}
	fmt.Println(count)
}

func setPriorities() map[string]int {
	m := make(map[string]int)
	for letter := 'a'; letter <= 'z'; letter++ {
		upper_letter := unicode.ToUpper(letter)
		// Seems fairly OK to hardcode these as it refers to ASCII value
		idx := int(letter) - 96
		m[string(letter)] = idx
		m[string(upper_letter)] = idx + 26 // e.g. a = 1, A = 27
	}

	return m
}

func chunkSlice(s []string, chunkSize int) [][]string {
	chunks := [][]string{}

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}

		chunks = append(chunks, s[i:end])
	}
	return chunks
}
