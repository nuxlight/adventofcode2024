package main

import (
	"log"
	"os"
	"strings"
)

type position struct {
	X int
	Y int
}

func main() {
	log.Println("start script")
	data, err := os.ReadFile("input-data")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var total = 0
	lines := strings.Split(string(data), "\n")
	for y, line := range lines {
		letters := strings.Split(line, "")
		for x, letter := range letters {
			var pos = position{X: x, Y: y}

			if letter == "M" {
				if check_upper_right(lines, pos) {
					total = total + 1
					log.Println("check_upper_right")
					log.Println(pos)
				}
				if check_upper_left(lines, pos) {
					total = total + 1
					log.Println("check_upper_left")
					log.Println(pos)
				}
				if check_down_left(lines, pos) {
					total = total + 1
					log.Println("check_down_left")
					log.Println(pos)
				}
				if check_down_right(lines, pos) {
					total = total + 1
					log.Println("check_down_right")
					log.Println(pos)
				}
			}
		}
	}
	log.Println(total)
}

func check_upper_right(lines []string, pos position) bool {
	var word = []string{}
	read_pos := pos.X
	for i := pos.Y; i > pos.Y-3; i-- {
		if i >= 0 {
			letters := strings.Split(lines[i], "")
			if read_pos < len(letters) {

				word = append(word, letters[read_pos])
				read_pos = read_pos + 1
			}
		}
	}
	if check_word(word) {
		return true
	}
	return false
}

func check_upper_left(lines []string, pos position) bool {
	var word = []string{}
	read_pos := pos.X
	for i := pos.Y; i > pos.Y-3; i-- {
		if i >= 0 {
			letters := strings.Split(lines[i], "")
			if read_pos >= 0 {
				word = append(word, letters[read_pos])
				read_pos = read_pos - 1
			}
		}
	}
	if check_word(word) {
		return true
	}
	return false
}

func check_down_right(lines []string, pos position) bool {
	var word = []string{}
	read_pos := pos.X
	for i := pos.Y; i < pos.Y+3; i++ {

		if i < len(lines) {
			letters := strings.Split(lines[i], "")
			if read_pos < len(letters) {
				if len(letters) > 0 {
					word = append(word, letters[read_pos])
					read_pos = read_pos + 1
				}
			}
		}
	}
	if check_word(word) {
		return true
	}
	return false
}

func check_down_left(lines []string, pos position) bool {
	var word = []string{}
	read_pos := pos.X
	for i := pos.Y; i < pos.Y+3; i++ {
		if i < len(lines) {
			letters := strings.Split(lines[i], "")
			if read_pos >= 0 {
				if len(letters) > 0 {
					word = append(word, letters[read_pos])
					read_pos = read_pos - 1
				}
			}
		}
	}
	if check_word(word) {
		return true
	}
	return false
}

func check_word(word []string) bool {
	if strings.Join(word, "") == "MAS" {
		log.Println(word)
		return true
	}
	return false
}
