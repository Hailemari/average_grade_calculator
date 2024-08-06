package main

import (
	"bytes"
	"testing"
	"bufio"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		gradesTotal float64
		subjectsCount int
		expected float64
	}{
		{300, 3, 100},
		{270, 3, 90},
		{0, 1, 0},
		{100, 0, 0},
	}

	for _, tt := range tests {
		result := calculateAverage(tt.gradesTotal, tt.subjectsCount)
		if result != tt.expected {
			t.Errorf("calculateAverage(%v, %v) = %v; want %v", tt.gradesTotal, tt.subjectsCount, result, tt.expected)
		}
	}
}

func TestGetName(t *testing.T) {
	input := "John Doe\n"
	reader := bufio.NewReader(bytes.NewReader([]byte(input)))
	name, err := getName(reader)
	if err != nil {
		t.Fatal(err)
	}
	if name != "John Doe" {
		t.Errorf("getName() = %v; want %v", name, "John Doe")
	}
}

func TestGetSubjectsAndGrades(t *testing.T) {
	input := "Math\n85\nScience\n90\ndone\n"
	expectedSubjects := map[string]float64{
		"Math":    85,
		"Science": 90,
	}
	reader := bufio.NewReader(bytes.NewReader([]byte(input)))
	subjects, gradesTotal, subjectsCount, err := getSubjectsAndGrades(reader)
	if err != nil {
		t.Fatal(err)
	}
	if len(subjects) != len(expectedSubjects) {
		t.Errorf("getSubjectsAndGrades() = %v; want %v", subjects, expectedSubjects)
	}
	for k, v := range expectedSubjects {
		if subjects[k] != v {
			t.Errorf("getSubjectsAndGrades() subject %v = %v; want %v", k, subjects[k], v)
		}
	}
	if gradesTotal != 175 {
		t.Errorf("getSubjectsAndGrades() gradesTotal = %v; want %v", gradesTotal, 175)
	}
	if subjectsCount != 2 {
		t.Errorf("getSubjectsAndGrades() subjectsCount = %v; want %v", subjectsCount, 2)
	}
}
