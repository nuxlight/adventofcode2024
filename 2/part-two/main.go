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
		if len(numbers) == 1 {
			continue
		}
		if check_line(numbers) {
			total = total + 1
			continue
		} else {
			var sub_total = 0
			for i := 0; i < len(numbers); i++ {
				mod := delete_element_from_array(numbers, i)

				if check_line(mod) {
					sub_total = sub_total + 1
				}
			}
			if sub_total > 0 {
				total = total + 1
				continue
			}
		}

	}
	log.Println(total)
}

func delete_element_from_array(array []string, index int) []string {
	var final_array = []string{}
	for i, el := range array {
		if i != index {
			final_array = append(final_array, el)
		}

	}
	return final_array
}

func check_line(array []string) bool {
	if check_equality(array) {
		if check_decreasing(array) {
			if check_separation(array) {
				return true
			}
		}
		if check_increase(array) {

			if check_separation(array) {
				return true
			}
		}
	}
	return false
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
