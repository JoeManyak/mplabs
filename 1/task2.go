package main

import (
	"bufio"
	"os"
	"strconv"
)

const div2 = 'a' - 'A'
const pageSize = 45

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	var rows []string
	var words []string
	var wordsPage []int
	var page = 1
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
		if row%pageSize == 0 {
			page++
		}
		goto RowChecker
	}

	if i == symbols-1 {
		words = append(words, rows[row][breaker:symbols])
		wordsPage = append(wordsPage, page)
		breaker = i + 1
		i++
		if i >= symbols {
			row++
			if row%pageSize == 0 {
				page++
			}
			goto RowChecker
		}
	}

	if rows[row][i] == ' ' {
		words = append(words, rows[row][breaker:i])
		wordsPage = append(wordsPage, page)
		breaker = i + 1
		i++
	}
	i++

	goto WordChecker
CountWords:
	var uniqueWords []string
	var uniquePages [][]int
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
		words[i] = words[i][:j] + string(words[i][j]+div2)
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
	var currentWord = 0
WordCountCycle:
	if currentWord >= wordLen {
		goto Sort
	}
	if j >= countedLen {
		goto ADD
	}
	if uniqueWords[j] == words[currentWord] {
		var z = 0
	uniqueWordChecker:
		if wordsPage[currentWord] == uniquePages[j][z] {
			goto Skip
		}
		z++
		if z != len(uniquePages[j]) {
			goto uniqueWordChecker
		}
		uniquePages[j] = append(uniquePages[j], wordsPage[currentWord])
	Skip:
		currentWord++
		j = 0
		goto WordCountCycle
	}
	j++
	goto WordCountCycle
ADD:
	//not count word lesser than 2 symbols
	if len(words[currentWord]) > 1 {
		uniqueWords = append(uniqueWords, words[currentWord])
		uniquePages = append(uniquePages, []int{wordsPage[currentWord]})
		countedLen++
	}
	currentWord++
	j = 0
	goto WordCountCycle
Sort:
	i = -1
	var z = 0
SortI:
	j = 0
	i++
	if i == len(uniqueWords)-1 {
		goto END
	}
SortJ:
	if j == len(uniqueWords)-i-1 {
		goto SortI
	}
SortZ:
	if z >= len(uniqueWords[j]) {
		z = 0
		j++
		goto SortJ
	}
	if z >= len(uniqueWords[j+1]) {
		uniquePages[j], uniquePages[j+1] = uniquePages[j+1], uniquePages[j]
		uniqueWords[j], uniqueWords[j+1] = uniqueWords[j+1], uniqueWords[j]
		j++
		z = 0
		goto SortJ
	}
	if uniqueWords[j][z] > uniqueWords[j+1][z] {
		uniquePages[j], uniquePages[j+1] = uniquePages[j+1], uniquePages[j]
		uniqueWords[j], uniqueWords[j+1] = uniqueWords[j+1], uniqueWords[j]
	}
	if uniqueWords[j][z] == uniqueWords[j+1][z] {
		z++
		goto SortZ
	}
	j++
	z = 0
	goto SortJ
END:
	output, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	i = 0
OUT:
	if i == countedLen {
		return
	}

	if len(uniquePages[i]) > 100 {
		i++
		goto OUT
	}
	_, err = output.Write([]byte(uniqueWords[i] + ": "))
	if err != nil {
		panic(err)
	}
	//print(uniqueWords[i] + ": ")
	var k = 0
FormatString:
	if k == len(uniquePages[i]) {
		k = 0
		i++
		goto OUT
	}
	//print(uniquePages[i][k])
	_, err = output.Write([]byte(strconv.Itoa(uniquePages[i][k])))
	if err != nil {
		panic(err)
	}
	k++
	if k < len(uniquePages[i]) {
		_, err = output.Write([]byte(", "))
		if err != nil {
			panic(err)
		}
		//print(", ")
	} else {
		_, err = output.Write([]byte(";\n"))
		if err != nil {
			panic(err)
		}
		//println(";")
	}
	goto FormatString

}
