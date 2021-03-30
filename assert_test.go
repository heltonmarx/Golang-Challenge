package sample1

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func assertBool(t *testing.T, expected bool, actual bool, msg string) {
	if expected != actual {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
	}
}

func assertInt(t *testing.T, expected int, actual int, msg string) {
	if expected != actual {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
	}
}

func assertFloat(t *testing.T, expected float64, actual float64, msg string) {
	if expected != actual {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
	}
}

func assertFloats(t *testing.T, expected []float64, actual []float64, msg string) {
	if len(expected) != len(actual) {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
		return
	}
	sort.Float64s(expected)
	sort.Float64s(actual)
	for i, expectedValue := range expected {
		if expectedValue != actual[i] {
			t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
			return
		}
	}
}

func assertErrors(t *testing.T, expected []error, actual error, msg string) {
	if actual == nil {
		t.Error(msg, fmt.Sprintf("expected : %v, got : nil", expected))
		return
	}
	// expected error message should be present in the error message
	for _, err := range expected {
		if !strings.Contains(actual.Error(), err.Error()) {
			t.Error(msg, fmt.Sprintf("expected : %v, got : %v", err.Error(), actual))
			return
		}
	}
}
