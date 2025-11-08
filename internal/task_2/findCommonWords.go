package task2

import (
	"os"
	"strings"
)

func FindCommonWords(outputFilename string, inputFilenames ...string) error {
	common_words := make(map[string]bool)
	
	for i, filename := range inputFilenames {
		data, err := os.ReadFile(filename)
		if err != nil {
			return ErrOpenFile
		}

		words_slice := strings.Fields(string(data))
		
		if i == 0 {  // first file - create a set
			common_words = createWordsSet(words_slice)

		} else {  // other - check overlap
			words_set := createWordsSet(words_slice)
			deleteNonCommonWords(common_words, words_set)
		}
	}

	err := writeWordsSetToFile(common_words, outputFilename)
	return err
}


func createWordsSet(words []string) map[string]bool {
	words_set := make(map[string]bool)
	for _, w := range words {
		words_set[w] = true
	}

	return  words_set
}


func deleteNonCommonWords(set map[string]bool, filter map[string]bool) {
	for word := range set {
		if !filter[word] {  // default value for bool is false so it will be returned if word isn't in set
			delete(set, word)
		}
	}
}


func writeWordsSetToFile(words_set map[string]bool, outputFilename string) error {
	file, err := os.Create(outputFilename)
	if err != nil {
		return ErrOpenFile
	}
	defer file.Close()

	result_line := ""
	for word := range words_set {
		result_line += word + " "
	}
	
	file.Write([]byte(result_line))
	
	return nil
}
