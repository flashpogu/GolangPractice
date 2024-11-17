package main

import (
	"fmt"
	"strings"
)

func strToMap(str string) map[string]int {
	lowerStr := strings.Fields(strings.ToLower(str))
	// fmt.Println(lowerStr)
	count := make(map[string]int)
	for _, word := range lowerStr {
		word = strings.Trim(word, `.,"-`)
		count[word]++
	}
	return count
}

func main() {
	str := `As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overwelmed the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, converted with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled accross his path, but otherwise there seemed to be no life striring in the world`

	frq := strToMap(str)
	for word, count := range frq {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}
