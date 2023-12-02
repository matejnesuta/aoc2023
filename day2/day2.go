package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func helper(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

func task(taskNum int) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int = 0
	var red int = 12
	var green int = 13
	var blue int = 14
	for scanner.Scan() {
		var redMax int = 0
		var greenMax int = 0
		var blueMax int = 0
		id, _ := strconv.Atoi(strings.Fields(strings.Split(scanner.Text(), ":")[0])[1])
		games := strings.Split(strings.Split(scanner.Text(), ":")[1], ";")
		for _, game := range games {
			v := strings.Split(game, ",")
			for i := 0; i < len(v); i++ {
				field := strings.Fields(v[i])
				num := helper(field[0])
				if taskNum == 1 {
					if (string(field[1]) == "red" && num > red) ||
						(string(field[1]) == "green" && num > green) ||
						(string(field[1]) == "blue" && num > blue) {
						goto innerEnd
					}
				} else {
					if string(field[1]) == "red" && num > redMax {
						redMax = num
					} else if string(field[1]) == "green" && num > greenMax {
						greenMax = num
					} else if string(field[1]) == "blue" && num > blueMax {
						blueMax = num
					}
				}
			}
		}
		if taskNum == 1 {
			sum += id
		} else {
			sum += redMax * greenMax * blueMax
		}
	innerEnd:
	}

	fmt.Printf("Task %d result: ", taskNum)
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	task(1)
	task(2)
}
