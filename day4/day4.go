package day4

import (
	"advent22/helpers"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
)

func Day4() {
	input := helpers.GetPuzzleInput("4")
	rows := strings.Split(input, "\n")
	ch := make(chan int32)
	var count int32 = 0

	for _, row := range rows {
		row := row

		go func() {
			assigned_sections := strings.Split(row, ",")
			first_assignment := assigned_sections[0]
			second_assignment := assigned_sections[1]
			first_boundaries := strings.Split(first_assignment, "-")
			second_boundaries := strings.Split(second_assignment, "-")
			first_upper, _ := strconv.Atoi(first_boundaries[1])
			first_lower, _ := strconv.Atoi(first_boundaries[0])
			second_upper, _ := strconv.Atoi(second_boundaries[1])
			second_lower, _ := strconv.Atoi(second_boundaries[0])
			var res int32

			// pairs_fully_overlap := (first_lower <= second_lower && first_upper >= second_upper) || (second_lower <= first_lower && second_upper >= first_upper)
			pairs_partially_overlap := (first_lower <= second_upper && first_upper >= second_upper) || (first_lower <= second_lower && first_upper >= second_lower) || (second_lower <= first_upper && second_upper >= first_upper) || (second_lower <= first_lower && second_upper >= first_lower)

			if pairs_partially_overlap {
				res = 1
			} else {
				res = 0
			}
			ch <- res
		}()
	}

	for range rows {
		atomic.AddInt32(&count, <-ch)
	}
	fmt.Println(count)
}
