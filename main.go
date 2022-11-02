package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var csvFilename *string
	var timeLimit *int
	csvFilename, timeLimit = setupFlags(csvFilename, timeLimit)
	lines := readCsvFileAndReturnFileLines(csvFilename)
	quizGame(lines, timeLimit)
}

func setupFlags(s *string, i *int) (*string, *int) {
	s = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	i = flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	return s, i
}
func readCsvFileAndReturnFileLines(fileName *string) [][]string {
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to read the CSV file " + *fileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}
	return lines
}
func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func quizGame(lines [][]string, timeLimit *int) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problems := parseLines(lines)
	var correct = 0
	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problem.question)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == problem.answer {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}
