package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*	Task: Description 1

	Fundamentals of Go Tasks
	Task: Student Grade Calculator
	Create a Go console application that allows students to calculate their
	average grade based on different subjects. The application should prompt
	the student to enter their name and the number of subjects they have taken.
	For each subject, the student should enter the subject name and the grade
	obtained (numeric value). After entering all subjects and grades,
	the application should display the student's name, individual subject
		grades, and the calculated average grade.
	Requirements:

	Use variables and data types to store student data.✅
	Use conditional statements to validate input (e.g., ensure grade values are within a valid range).✅
	Implement loops to handle multiple subjects and grades.✅
	Utilize collections (e.g., List, Dictionary) to store subject names and corresponding grades.✅
	Define a method to calculate the average grade based on the entered grades.✅
	Use string interpolation to display the results in a user-friendly format.✅
	Write test for your code [Optional]✅

*/


// subject key with grade value
var profile = map[string]int{}


func averageGradeCalculator() float32 {
	total := 0
	for _, value := range profile {
		total += value
	}
	return float32(total) / float32(len(profile))
}

func main() {
	fmt.Print("Input Name and Number of subjects you take: ")

	reader := bufio.NewReader(os.Stdin)

	// accept name and # subjects
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fields := strings.Fields(input)
	if len(fields) < 2 {
		log.Fatal("Invalid input. Please enter Your name and number of subjects.")
	}
	studentName := fields[0]
	subjects, err := strconv.Atoi(fields[1])
	if err != nil || subjects <= 0 {
		log.Fatal("Invalid number of subjects.")
	}

	// input loop for each subject
	for i := 0; i < subjects; i++ {
		for {
			fmt.Printf("Enter subject name and grade for subject #%d: ", i+1)
			line, _ := reader.ReadString('\n')
			inputFields := strings.Fields(line)

			if len(inputFields) < 2 {
				fmt.Println("Invalid input. Format: <subject> <grade>. Try again!")
				continue
			}

			subject := inputFields[0]
			grade, err := strconv.Atoi(inputFields[1])
			if err != nil {
				fmt.Println("Invalid grade. Must be a number.")
				continue
			}
			if grade < 0 || grade > 100 {
				fmt.Println("Grade must be between 0 and 100 inclusive. Try again.")
				continue
			}

			profile[subject] = grade
			break // valid input
		}
	}

	// handle output
	fmt.Println("Student name: ", studentName)

	for key, value := range profile {
		fmt.Printf("Subject: %s, Grade: %d\n", key, value)
	}

	average := averageGradeCalculator()
	fmt.Println("Average result: ", average)
}