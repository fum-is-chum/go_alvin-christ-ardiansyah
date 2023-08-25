package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name string
	score int
}

type StudentList struct {
	students []Student
	_totalScore int
}

func (s StudentList) GetStudentWithMinimumScore() []Student {
	minScore := s.students[0].score
	students := []Student{s.students[0]}

	for i := 1; i < len(s.students); i++ {
		if s.students[i].score < minScore {
			minScore = s.students[i].score
			students = []Student{s.students[i]}
		} else if s.students[i].score == minScore {
			students = append(students, s.students[i])
		}
	}

	return students
}

func (s StudentList) GetStudentWithMaximumScore() []Student {
	maxScore := s.students[0].score
	students := []Student{s.students[0]}

	for i := 1; i < len(s.students); i++ {
		if s.students[i].score > maxScore {
			maxScore = s.students[i].score
			students = []Student{s.students[i]}
		} else if s.students[i].score == maxScore {
			students = append(students, s.students[i])
		}
	}

	return students
}

func (s StudentList) GetAverageScore() float64 {
	return float64(s._totalScore) / float64(len(s.students))
}

func (s *StudentList) addNewStudent(newStudent Student) {
	s.students = append(s.students, newStudent)
	s._totalScore += newStudent.score
}

func main() {
	studentList := StudentList{}

	for i:=1; i<=5; i++ {
		newStudent := Student{}
		fmt.Print(fmt.Sprintf("Input %d Student's Name: ", i))
		fmt.Scanln(&newStudent.name)

		fmt.Print(fmt.Sprintf("Input %d Student's Score: ", i))
		fmt.Scanln(&newStudent.score)

		studentList.addNewStudent(newStudent)
	}

	// print average
	fmt.Println(fmt.Sprintf("Average Score: %v", studentList.GetAverageScore()))

	// print student(s) with minimum score
	var minimumResult strings.Builder
	minimumStudentList := studentList.GetStudentWithMinimumScore();
	for idx,student := range minimumStudentList {
		minimumResult.WriteString(fmt.Sprintf("%s (%d)", student.name, student.score))
		if idx < len(minimumStudentList) - 1 {
			minimumResult.WriteRune(',')
		}
	}
	fmt.Println(fmt.Sprintf("Min Score of Students: %s", minimumResult.String()))
	
	// print student(s) with maximum score
	var maximumResult strings.Builder
	maximumStudentList := studentList.GetStudentWithMaximumScore();
	for idx,student := range maximumStudentList {
		maximumResult.WriteString(fmt.Sprintf("%s (%d)", student.name, student.score))
		if idx < len(maximumStudentList) - 1 {
			maximumResult.WriteRune(',')
		}
	}
	fmt.Println(fmt.Sprintf("Max Score of Students: %s", minimumResult.String()))
}