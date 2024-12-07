package main

import (
	"log"
	"os"
	"regexp"
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
	var total = 0
	var do = true
	for _, line := range lines {
		r, _ := regexp.Compile(`(\bmul\(+([1-9][0-9]*),([1-9][0-9]*)+\))|(\bdon't\(\))|(\bdo\(\))`)
		finded := r.FindAllString(line, -1)
		log.Println(finded)
		for _, op := range finded {
			log.Println(total)
			if op == "don't()" {
				do = false
				continue
			}
			if op == "do()" {
				do = true
				continue
			}
			if do {
				rr, _ := regexp.Compile(`([1-9][0-9]*),([1-9][0-9]*)`)
				numbers := rr.FindAllString(op, -1)
				slice := strings.Split(numbers[0], ",")
				nb_first, _ := strconv.Atoi(slice[0])
				nb_sec, _ := strconv.Atoi(slice[1])
				sub_total := nb_first * nb_sec
				total = total + sub_total

			}
		}
	}
	log.Println(total)
}
