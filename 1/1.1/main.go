package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./1/1.1/input.txt")
	if err != nil {
		panic(err.Error())
	}

	//words that not counting
	var forbidden = []string{"in", "on", "out", "of", "the", "a", "an", "and"}
	var forbiddenCount = len(forbidden)

	var rows []string
	var words []string
	var row = 0
	var rowCount = 0
	scanner := bufio.NewScanner(file)

Scanner:
	if !scanner.Scan() {
		rowCount = len(rows)
		goto RowChecker
	}
	rows = append(rows, scanner.Text())
	goto Scanner
RowChecker:
	var symbols = 0
	var breaker = 0
	var i = 0
	if row == rowCount {
		goto CountWords
	}
	symbols = len(rows[row])
WordChecker:
	if i >= symbols {
		row++
		goto RowChecker
	}

	if i == symbols-1 {
		words = append(words, strings.ToLower(rows[row][breaker:symbols]))
		breaker = i + 1
		i++
		if i >= symbols {
			row++
			goto RowChecker
		}
	}
	if rows[row][i] == ' ' {
		words = append(words, strings.ToLower(rows[row][breaker:i]))
		breaker = i + 1
		i++
	}
	i++
	goto WordChecker
CountWords:
	var uniqueWords []string
	var uniqueCount []int
	var countedLen = 0
	var wordLen = len(words)
	var k = 0
	var currentWord = 0
	var j = 0
WordCountCycle:
	if currentWord >= wordLen {
		goto SORT
	}
	if j >= countedLen {
		goto ADD
	}
	if uniqueWords[j] == words[currentWord] {
		uniqueCount[j]++
		currentWord++
		j = 0
		goto WordCountCycle
	}
	j++
	goto WordCountCycle
ADD:
	k = 0
ForbiddenChecker:
	if words[currentWord] == forbidden[k] {
		j = 0
		currentWord++
		goto WordCountCycle
	}
	k++
	if k != forbiddenCount {
		goto ForbiddenChecker
	}
	uniqueWords = append(uniqueWords, words[currentWord])
	uniqueCount = append(uniqueCount, 1)
	currentWord++
	countedLen++
	j = 0
	goto WordCountCycle

SORT:
	i = -1
	j = 0
SortI:
	i++
	if i == countedLen-1 {
		goto END
	}
SortJ:
	if j == countedLen-i-1 {
		j = 0
		goto SortI
	}

	if uniqueCount[j] < uniqueCount[j+1] {
		uniqueCount[j], uniqueCount[j+1] = uniqueCount[j+1], uniqueCount[j]
		uniqueWords[j], uniqueWords[j+1] = uniqueWords[j+1], uniqueWords[j]
	}
	j++
	goto SortJ
END:
	i = 0
OUT:
	if i == countedLen {
		return
	}
	println(uniqueWords[i], ":", uniqueCount[i])
	i++
	goto OUT
}
