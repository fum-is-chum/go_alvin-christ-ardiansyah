package main

import "fmt"

type student struct {
	name string

	encodedName string

	score int
}

// use ROT-13 as substitution cipher
type Chiper interface {
	Encode() string

	Decode() string
}

func (s *student) Encode() string {
	encodedName := make([]byte, len(s.name))

	for i, char := range s.name {
		if char == 32 {
			encodedName[i] = 32
		} else {
			encodedName[i] = 'a' + byte((int(char-'a')+13)%26)
		}
	}
	// your code here
	s.encodedName = string(encodedName)
	return s.encodedName

}

func (s *student) Decode() string {

	decodedName := make([]byte, len(s.encodedName))

	for i, char := range s.encodedName {
		if char == 32 {
			decodedName[i] = 32
		} else {
			decodedName[i] = 'a' + byte((int(char-'a')+13)%26)
		}
	}
	// your code here
	s.name = string(decodedName)
	return s.name

}

func main() {

	var menu int

	var a student = student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student Name: ")

		fmt.Scan(&a.name)

		fmt.Println("\nEncode of student's name " + a.name + "is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Encoded Student Name: ")

		fmt.Scan(&a.encodedName)

		fmt.Println("\nDecode of student's name " + a.name + "is : " + c.Decode())

	}

}
