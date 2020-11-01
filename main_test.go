package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	name     string
	dividend int
	divisor  int
	expected int
	isErr    bool
} {
	{ "valid-data", 100, 10, 10, false },
	{ "invalid-data", 100, 0, 0, true },
	{ "expect 5", 50, 10, 5, false },
	{ "expect 5", 70, 7, 10, false },
}

// TestDivision uses table driven tests to run large amounts of data against the same test
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.dividend, tt.divisor)
			if tt.isErr {
				if err == nil {
					t.Log(fmt.Sprintf("Passed %d and %d and got %d", tt.dividend, tt.divisor, got))
					t.Error("expected error but did not get one")
				}
			} else {
				t.Log(fmt.Sprintf("Passed %d and %d and got %d", tt.dividend, tt.divisor, got))
				if err != nil {
					t.Error("Did not expect an error, but got error of", err.Error())
				}
			}
			if got != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, got)
			}
		})
	}
}

// TestDivide tests valid data
func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

// TestBadDivide tests invalid data (division by zero)
func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
