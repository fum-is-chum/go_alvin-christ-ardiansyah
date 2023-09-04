package main

import (
	"fmt"
	"sync"
	"unicode"
)


func countChar(s *string) {
	var m sync.Mutex
	var wg sync.WaitGroup

	charCount := make(map[rune]int)

	for _, c := range *s {
		if unicode.IsLetter(c) {
			wg.Add(1)
			go func(char rune) {
				defer wg.Done()

				m.Lock()
				defer m.Unlock()

				_ , exist := charCount[char]
				if exist {
					charCount[char]++
				} else {
					charCount[char] = 1
				}
			}(unicode.ToLower(c))
		}
	}

	wg.Wait()
	for char,count := range charCount {
		fmt.Println(fmt.Sprintf("%c: %d", char, count))
	}
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec turpis leo, varius id accumsan vel, luctus sed lacus."
	countChar(&text)
}