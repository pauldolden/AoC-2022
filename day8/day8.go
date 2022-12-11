package day8

import (
	"advent22/helpers"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Position struct {
	x int
	y int
}

func Day8() {
	input := helpers.GetPuzzleInput("8")
	grid, height, width := buildGrid(input)
	heighest_scenic_score := 0
	// For each row
	for y := 0; y < height; y++ {
		// For each column
		for x := 0; x < width; x++ {
			// So for each tree we do the below
			tree_score := []int{}
			tree := grid[y][x]
			// Is an edge or rather, is already visible just add a visible tree and move on
			if y == 0 || x == 0 || y == height-1 || x == width-1 {
				continue
			}
			// It is not an edge so we want to check each direction
			for _, dir := range directions {
				x_move, y_move := dir[0], dir[1]
				dir_score := 0
				// if the tree is smaller than the one next to it (in that direction) just move on
				// It is not visible in that direction
				// The tree is bigger than its neighbour so MIGHT be visible, so we need to keep checking this direction
				// Until we reach an edge or a bigger tree_score
				number_of_moves := 1
				for k := 1; k <= width; k++ {
					x_loc := x + x_move*k
					y_loc := y + y_move*k
					// Feels gross but just avoids index out of range errors
					if x_loc >= width {
						x_loc = width - 1
					}
					if x_loc <= 0 {
						x_loc = 0
					}
					if y_loc >= height {
						y_loc = height - 1
					}
					if y_loc <= 0 {
						y_loc = 0
					}
					if grid[y_loc][x_loc] >= tree || x_loc == 0 || x_loc == width-1 || y_loc == 0 || y_loc == height-1 {
						break
					}
					number_of_moves++
				}
				dir_score = number_of_moves
				tree_score = append(tree_score, dir_score)
			}
			if tree_score[0]*tree_score[1]*tree_score[2]*tree_score[3] > heighest_scenic_score {
				heighest_scenic_score = tree_score[0] * tree_score[1] * tree_score[2] * tree_score[3]
			}
		}
	}
	fmt.Println(heighest_scenic_score)
}

func buildGrid(input string) (grid map[int][]int, height int, width int) {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

	grid = make(map[int][]int)
	row := 0

	for scanner.Scan() {
		trees := strings.Split(scanner.Text(), "")
		grid[row] = mapStrToInt(trees)
		row++
	}

	height = row
	width = len(grid[0])
	return
}

func mapStrToInt(ss []string) []int {
	is := []int{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		is = append(is, i)
	}

	return is
}
