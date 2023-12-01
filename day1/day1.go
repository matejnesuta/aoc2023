package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func task(taskNum int) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	regForNotNumbers := regexp.MustCompile("[^1234567890]+")

	scanner := bufio.NewScanner(file)
	var sum int = 0
	for scanner.Scan() {
		str := scanner.Text()
		if taskNum == 2 {
			str = regexp.MustCompile("one").ReplaceAllString(str, "o1e")
			str = regexp.MustCompile("two").ReplaceAllString(str, "t2o")
			str = regexp.MustCompile("three").ReplaceAllString(str, "t3e")
			str = regexp.MustCompile("four").ReplaceAllString(str, "f4r")
			str = regexp.MustCompile("five").ReplaceAllString(str, "f5e")
			str = regexp.MustCompile("six").ReplaceAllString(str, "s6x")
			str = regexp.MustCompile("seven").ReplaceAllString(str, "s7n")
			str = regexp.MustCompile("eight").ReplaceAllString(str, "e8t")
			str = regexp.MustCompile("nine").ReplaceAllString(str, "n9e")
		}
		replaceStr := regForNotNumbers.ReplaceAllString(str, "")
		number, _ := strconv.Atoi(string(replaceStr[0]) + string(replaceStr[len(replaceStr)-1]))
		sum += number
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
