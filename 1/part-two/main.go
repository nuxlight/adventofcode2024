package main

import (
	"log"
	"os"
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
	var total = 0
	for _, a := range first_col {
		var finded = 0
		for _, b := range sec_col {
			if a == b {
				finded = finded + 1
			}
		}
		total = total + (a * finded)
	}
	log.Println(total)
}
