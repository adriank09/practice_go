package main

import (
	"regexp"
	"testing"
)

// TestHelloName calls main.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`¥b` + name + `¥b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestAdd(t *testing.T) {
	expected := 3
	result := Add(1, 2)
	if result != expected {
		t.Fatalf("Add(1,2) was %v, expected = %v", result, expected)
	}
}
