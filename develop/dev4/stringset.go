package main

func getWordsSet(words []string) map[string][]string {
	result := make(map[string][]string, len(words))

	for _, word := range words {
		keyExists := false

		for key := range result {
			if equalWords(key, word) {
				result[key] = append(result[key], word)
				keyExists = true
				break
			}
		}

		if !keyExists {
			result[word] = []string{word}
		}
	}

	return result
}

func getLetters(word string) map[rune]int {
	letters := make(map[rune]int)

	for _, r := range word {
		if _, ok := letters[r]; !ok {
			letters[r] = 0
		}

		letters[r] += 1
	}

	return letters
}

func equalWords(word1, word2 string) bool {
	letters := getLetters(word1)

	for _, r := range word2 {
		if exist, ok := letters[r]; ok && exist > 0 {
			letters[r] -= 1
		} else {
			return false
		}
	}

	for _, value := range letters {
		if value != 0 {
			return false
		}
	}

	return true
}
