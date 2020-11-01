package main

import (
	"fmt"
	"testing"
)

// tests sets up a slice of structs that we will use for our table driven test.
var tests = []struct {
	name     string  // the name of the test
	dividend float32 // the number that will be divided
	divisor  float32 // what we will divide by
	expected float32 // the expected result (will be zero if an error is returned)
	isErr    bool    // whether or not the data should generate an error
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-7", 70.0, 7.0, 10.0, false},
	{"expect-0", 0, 1000, 0, false},
}

// TestDivision uses table driven tests to run large amounts of data against the same test
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.dividend, tt.divisor)
			if tt.isErr {
				if err == nil {
					t.Log(fmt.Sprintf("Passed %f and %f and got %f", tt.dividend, tt.divisor, got))
					t.Error("expected error but did not get one")
				}
			} else {
				t.Log(fmt.Sprintf("Passed %f and %f and got %f", tt.dividend, tt.divisor, got))
				if err != nil {
					t.Error("Did not expect an error, but got error of", err.Error())
				}
			}
			if got != tt.expected {
				t.Errorf("expected %f but got %f", tt.expected, got)
			}
		})
	}
}

// NOTE: with our table driven test, the next two tests are redundant

// TestDivide tests valid data
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

// TestBadDivide tests invalid data (division by zero)
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
