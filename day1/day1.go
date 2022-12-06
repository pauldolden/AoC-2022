package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
    "advent22/helpers"
)

func Day1() {
    input := helpers.GetPuzzleInput("1")
	ss := strings.Split(input, "\n\n")
	em := make(map[int]int)
	for idx, se := range ss {
		count := 0
		for _, ae := range strings.Split(se, "\n") {
            if(ae == "") {
                continue
            }
			i, err := strconv.Atoi(ae)

			if err != nil {
				panic(err)
			}
			count += i
		}

		em[count] = idx
	}
    vm := []int{}

    for k := range em {
        vm = append(vm, k);
    }
    
    sort.Ints(vm);

    l := len(vm) - 3

    tt := vm[l:]
    ttc := 0
    for _, v := range tt {
        ttc += v
    }

    fmt.Println(ttc)
}
