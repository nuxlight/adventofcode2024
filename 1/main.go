package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	log.Println("start script")
	data, err := os.ReadFile("input-data")
	if err != nil {
		log.Fatalln(err.Error())
	}
	lines := strings.Split(string(data), "\n")
	var first_col []int
	var sec_col []int
	for _, line := range lines {
		line := strings.Split(line, "   ")
		if len(line) > 1 {
			val, _ := strconv.Atoi(line[0])
			first_col = append(first_col, val)
			val, _ = strconv.Atoi(line[1])
			sec_col = append(sec_col, val)
		}
	}
	//var total []int
	slices.Sort(first_col)
	slices.Sort(sec_col)
	var total = 0
	for i, c := range first_col {
		var t = 0
		if c > sec_col[i] {
			t = c - sec_col[i]
		} else {
			t = sec_col[i] - c
		}

		total = total + t
	}
	log.Println(total)
}
