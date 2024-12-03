package main

import (
	"log"
	"math"
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
	var total = 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		//log.Println(numbers)
		if len(numbers) == 1 {
			continue
		}
		if check_equality(numbers) {
			if check_decreasing(numbers) {
				if check_separation(numbers) {
					log.Println(numbers)
					total = total + 1
					continue
				}
			}
			if check_increase(numbers) {

				if check_separation(numbers) {
					log.Println(numbers)
					total = total + 1
					continue
				}
			}
		}

	}
	log.Println(total)
}

func check_increase(array []string) bool {
	for i, el := range array {
		if i != 0 {
			cur, _ := strconv.Atoi(el)
			before, _ := strconv.Atoi(array[i-1])
			if cur > before {
				return false
			}
		}
	}
	return true
}

func check_decreasing(array []string) bool {
	for i, el := range array {
		if i != 0 {
			cur, _ := strconv.Atoi(el)
			before, _ := strconv.Atoi(array[i-1])
			if cur < before {
				return false
			}
		}
	}
	return true
}

func check_equality(array []string) bool {
	for i, el := range array {
		if i != 0 && i < len(array) {
			cur, _ := strconv.Atoi(el)
			before, _ := strconv.Atoi(array[i-1])
			if cur == before {
				return false
			}
		}
	}
	return true
}

func check_separation(numbers []string) bool {
	var fail = 0
	for i, el := range numbers {

		if i != 0 {

			cur, _ := strconv.Atoi(el)
			before, _ := strconv.Atoi(numbers[i-1])
			op := cur - before
			opp := math.Abs(float64(op))

			if opp > 3 {
				fail = fail + 1
			}
		}
	}
	if fail > 0 {
		return false
	}
	return true
}
