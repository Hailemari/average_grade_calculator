package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculateAverage calculates the average grade.
func calculateAverage(gradesTotal float64, subjectsCount int) float64 {
	if subjectsCount == 0 {
		return 0
	}
	return gradesTotal / float64(subjectsCount)
}

// getName prompts the user for their name until a non-empty string is provided.
func getName(reader *bufio.Reader) (string, error) {
	for {
		fmt.Print("What is your name? ")
		name, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		name = strings.TrimSpace(name)
		if name != "" {
			return name, nil
		}

		fmt.Println("Name cannot be empty. Please enter your name.")
	}
}

// getSubjectsAndGrades collects subject names and their corresponding grades.
func getSubjectsAndGrades(reader *bufio.Reader) (map[string]float64, float64, int, error) {
	subjects := make(map[string]float64)
	var gradesTotal float64
	var subjectsCount int

	for {
		fmt.Print("Enter the name of the subject (or type 'done' to finish): ")
		subjectName, err := reader.ReadString('\n')
		if err != nil {
			return nil, 0, 0, err
		}
		subjectName = strings.TrimSpace(subjectName)

		if strings.ToLower(subjectName) == "done" {
			break
		}

		if subjectName == "" {
			fmt.Println("Please enter a subject name.")
			continue
		}

		if _, exists := subjects[subjectName]; exists {
			fmt.Printf("Subject '%s' already exists. Please enter a different subject name.\n", subjectName)
			continue
		}

		for {
			fmt.Printf("Enter the grade for %s (0 - 100) -> ", subjectName)
			gradeStr, err := reader.ReadString('\n')
			if err != nil {
				return nil, 0, 0, err
			}
			gradeStr = strings.TrimSpace(gradeStr)

			grade, err := strconv.ParseFloat(gradeStr, 64)
			if err != nil || grade < 0 || grade > 100 {
				fmt.Println("Invalid grade. Please enter a numeric value between 0 and 100.")
				continue
			}

			subjects[subjectName] = grade
			gradesTotal += grade
			subjectsCount++
			break
		}
	}

	return subjects, gradesTotal, subjectsCount, nil
}

// printResults displays the student name, subjects, and average grade.
func printResults(name string, subjects map[string]float64, averageGrade float64) {
	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("\nSubjects and Grades:")
	for subject, grade := range subjects {
		fmt.Printf("\t%s: %.2f\n", subject, grade)
	}
	fmt.Printf("\nAverage Grade: %.2f\n", averageGrade)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name, err := getName(reader)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	subjects, gradesTotal, subjectsCount, err := getSubjectsAndGrades(reader)
	if err != nil {
		fmt.Println("Error reading subjects and grades:", err)
		return
	}

	if subjectsCount == 0 {
		fmt.Println("No subjects entered. Exiting...")
		return
	}

	averageGrade := calculateAverage(gradesTotal, subjectsCount)
	printResults(name, subjects, averageGrade)
}
