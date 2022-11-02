package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var csvFilename *string
	csvFilename = setupFlags(csvFilename)

	lines := readCsvFileAndReturnFileLines(csvFilename)
	quizGame(lines)
}

func setupFlags(s *string) *string {
	s = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	return s
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

func quizGame(lines [][]string) {
	problems := parseLines(lines)
	var correct = 0
	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))
}
