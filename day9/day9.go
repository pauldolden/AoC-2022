package day9

import (
	"bufio"
	"fmt"
	"strings"
)

var Directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func Day9() {
	// input := helpers.GetPuzzleInput("9")
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	fmt.Println(input)
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		fmt.Println(row)
	}
}
