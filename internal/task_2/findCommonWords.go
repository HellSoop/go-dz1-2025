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
			for _, w := range words_slice {
				common_words[w] = true
			}

		} else {  // other - check overlap
			// create set for current file
			words_set := make(map[string]bool)
			for _, w := range words_slice {
				words_set[w] = true
			}

			// check overlap with commons
			for word := range common_words {
				if !words_set[word] {  // default value for bool is false so it will be returned if word isn't in set
					delete(common_words, word)
				}
			}
		}
	}

	// write output
	file, _ := os.Create(outputFilename)
	defer file.Close()

	result_line := ""
	for word := range common_words {
		result_line += word + " "
	}
	
	file.Write([]byte(result_line))

	return nil
}
