package main

import (
	"testing"
)

func TestSetupFlags(t *testing.T) {
	var timeLimit *int
	var fileName *string

	fileName, timeLimit = setupFlags(fileName, timeLimit)
	fileNameWant := "problems.csv"
	timeLimitWant := 30

	if *fileName != fileNameWant {
		t.Errorf("got %s want %s", *fileName, fileNameWant)
	}
	if *timeLimit != timeLimitWant {
		t.Errorf("got %d want %d", *timeLimit, timeLimitWant)
	}
}

func TestReadCsvFileAndReturnFileLines(t *testing.T) {
	fileName := "./1.csv"
	linesWant := 2
	lines := readCsvFileAndReturnFileLines(&fileName)
	if len(lines) != linesWant {
		t.Errorf("got %d want %d", len(lines), linesWant)
	}
}
func TestParseLines(t *testing.T) {
	question := "1+1"
	answer := "2"
	lines := [][]string{{question, answer}}

	problems := parseLines(lines)
	if len(problems) != len(lines) {
		t.Errorf("got %d want %d", len(problems), len(lines))
	}
	if problems[0].question != question {
		t.Errorf("got %s want %s", problems[0].question, question)
	}
	if problems[0].answer != answer {
		t.Errorf("got %s want %s", problems[0].answer, answer)
	}
}

func TestQuizGame(t *testing.T) {
	question1 := "1+1"
	answer1 := "2"
	lines := [][]string{{question1, answer1}}
	var timeLimit *int
	*timeLimit = 30

	quizGame(lines, timeLimit)

}
