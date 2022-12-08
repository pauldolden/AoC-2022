package day5

import (
	"advent22/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day5() {
	input := helpers.GetPuzzleInput("5")
	pattern_and_moves := strings.Split(input, "\n\n")
	pattern := pattern_and_moves[0]
	moves := pattern_and_moves[1]
	column_map := map[int][]string{}
	pattern_rows := strings.Split(pattern, "\n")
	move_rows := strings.Split(strings.TrimSpace(moves), "\n")
	item_rgx := regexp.MustCompile(`\[.]`)
	for idx, row := range reverse(pattern_rows) {
		if idx == 0 {
			continue
		}
		row_items := (ChunkString(row, 4))

		for idx, item := range row_items {
			idx := idx + 1
			if item_rgx.Match([]byte(item)) {
				column_map[idx] = append(column_map[idx], string(item[1]))
			}
		}
	}

	rgx := regexp.MustCompile(`(\d+)`)

	for _, move_row := range move_rows {
		digits := rgx.FindAllString(string(move_row), 3)
		num_to_move, _ := strconv.Atoi(digits[0])
		from, _ := strconv.Atoi(digits[1])
		to, _ := strconv.Atoi(digits[2])

		slice_index := len(column_map[from]) - num_to_move
		items_to_move := column_map[from][slice_index:]
		column_map[from] = column_map[from][:slice_index]
		column_map[to] = append(column_map[to], items_to_move...)
	}
	fmt.Println(column_map)
}

func reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}
