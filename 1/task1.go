package main

import (
	"bufio"
	"os"
	"strconv"
)

const div1 = 'a' - 'A'
const N = 25

func main() {
	file, err := os.Open("input.txt")
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

	//as far as I know, in golang no other way to read line by line, so I used scanner
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
		words = append(words, rows[row][breaker:symbols])
		breaker = i + 1
		i++
		if i >= symbols {
			row++
			goto RowChecker
		}
	}
	if rows[row][i] == ' ' {
		words = append(words, rows[row][breaker:i])
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
	var thisWord string
	var j = 0
	i = 0
FixWords:
	if i == wordLen {
		goto StartCounting
	}
	thisWord = words[i]
FixWord:
	if j >= len(words[i]) {
		i++
		j = 0
		goto FixWords
	}
	if thisWord[j] >= 'A' && thisWord[j] <= 'Z' {
		words[i] = words[i][:j] + string(words[i][j]+div1)
		if j != len(words[i]) {
			words[i] += thisWord[j+1:]
		}
		thisWord = words[i]
	}
	//getting symbols
	if (thisWord[j] < 'a' || thisWord[j] > 'z') && !(thisWord[j] == '\'' || thisWord[j] == '-' ||
		(thisWord[j] >= '0' && thisWord[j] <= '9')) {
		words[i] = thisWord[:j]
		if j != len(thisWord) {
			words[i] += thisWord[j+1:]
		}
		j--
	}
	thisWord = words[i]
	j++
	goto FixWord

StartCounting:
	var k = 0
	var currentWord = 0
WordCountCycle:
	if currentWord >= wordLen {
		goto Sort
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
	//not count word lesser than 2 symbols
	if len(words[currentWord]) > 1 {
		uniqueWords = append(uniqueWords, words[currentWord])
		uniqueCount = append(uniqueCount, 1)
		countedLen++
	}
	currentWord++
	j = 0
	goto WordCountCycle

Sort:
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
	output, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	i = 0
OUT:
	if i == countedLen {
		err = output.Close()
		if err != nil {
			panic(err)
		}
		return
	}
	_, err = output.Write([]byte(uniqueWords[i] + ": " + strconv.Itoa(uniqueCount[i]) + "\n"))
	if err != nil {
		panic(err)
	}
	if i < N {
		println(uniqueWords[i], ":", uniqueCount[i])
	}
	i++
	goto OUT
}
