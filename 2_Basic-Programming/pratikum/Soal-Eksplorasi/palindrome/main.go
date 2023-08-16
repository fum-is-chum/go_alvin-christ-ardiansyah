package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Masukkan kalimat: ")
 	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	result:= isPalindrome(text);
	if result {
		fmt.Print("Palindrome")
	} else {
		fmt.Print("Bukan palindrome")
	}
}

func isPalindrome(text string) bool {
	fmt.Println("Captured: " + text)
	textLength := len(text) - 1;
	for i:= 0; i < textLength / 2; i++ {
		if(text[i] != text[textLength - i]) {
			return false
		}
	}

	return true
}