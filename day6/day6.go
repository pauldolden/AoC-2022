package day6

import (
	"advent22/helpers"
	"fmt"
	"io"
	"strings"
)

const MARKER_SIZE = 14

func Day6() {
	input := helpers.GetPuzzleInput("6")
	r := strings.NewReader(input)
	buffer := make([]byte, 1)
	recent_marker_check := make([]byte, 0)
	count := 0

	for {
		n, err := r.Read(buffer)
		count++

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(recent_marker_check) < MARKER_SIZE {
			recent_marker_check = append(recent_marker_check, buffer[:n]...)
		} else {
			recent_marker_check = recent_marker_check[1:]
			recent_marker_check = append(recent_marker_check, buffer[:n]...)
		}
		temp_map := map[byte]int{}
		if len(recent_marker_check) == MARKER_SIZE {
			for _, i := range recent_marker_check {
				temp_map[i]++
			}

			if len(temp_map) == MARKER_SIZE {
				break
			}
		}
	}
	fmt.Println(count)
}
