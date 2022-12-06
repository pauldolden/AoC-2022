package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
    "os"
)

func GetPuzzleInput(day string) string {
	client := &http.Client{}

	r := strings.NewReader("")

    path := fmt.Sprint("https://adventofcode.com/2022/day/", day, "/input")

	req, err := http.NewRequest("GET", path, r)

	if err != nil {
		panic(err)
	}
    aoc_key := os.Getenv("AOC_SESSION_KEY")

    cookie := fmt.Sprint("session=", aoc_key)


	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(bs)
}
