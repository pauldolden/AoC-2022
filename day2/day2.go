package day2

import (
	"advent22/helpers"
	"fmt"
	"strings"
)

type Option struct {
	score int
	beats *Option
	loses *Option
}

var rock = Option{}
var paper = Option{}
var scissors = Option{}

func Day2() {
    rock.loses = &paper
    rock.beats = &scissors
    rock.score = 1
    paper.loses = &scissors
    paper.beats = &rock
    paper.score = 2
    scissors.loses = &rock
    scissors.beats = &paper
    scissors.score = 3

    input := helpers.GetPuzzleInput("2")

	s := strings.TrimSpace(input)

	games := strings.Split(s, "\n")

	score := 0

	for _, game := range games {
		values := strings.Split(game, " ")
		play(values[0], values[1], &score)
	}

	fmt.Println(score)
}

func play(c string, p string, s *int) {
	var compOption *Option
    score := 0
	switch c {
	// Lose
	case "A":
		compOption = &rock
		// Draw
	case "B":
		compOption = &paper
		// Win
	case "C":
		compOption = &scissors
	default:
		break
	}

    switch p {
    case "X":
        score += compOption.beats.score
    case "Y":
        score += compOption.score
        score += 3
    case "Z":
        score += compOption.loses.score
        score += 6
    }

	*s += score
}
