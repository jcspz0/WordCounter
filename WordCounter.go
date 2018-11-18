package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func SortWords(wordsCountMap map[string]int) []string {
	var words []string
	for word := range wordsCountMap {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	if (len(os.Args) < 2){
		fmt.Printf("there is not a file to read!!")
		return
	}
	cmdArg := os.Args[1:2][0]
	file, err := ioutil.ReadFile(cmdArg)
	if err != nil {
		fmt.Print(err)
	}
	
	regex, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		fmt.Print(err)
	}

	//words := strings.Fields(regex.ReplaceAllString(strings.ToLower(string(file)), " "))
	words := strings.ToLower(string(file))
	
	onlyWordsWithoutTrash := regex.ReplaceAllString(words, " ")
	separetedWords := strings.Fields(onlyWordsWithoutTrash)

	wordsCountMap := make(map[string]int)
	for _, word := range separetedWords {
		wordsCountMap[word]++
	}

	sortedWords := SortWords(wordsCountMap)
	for _, word := range sortedWords {
		fmt.Println(word, wordsCountMap[word])
	}
}