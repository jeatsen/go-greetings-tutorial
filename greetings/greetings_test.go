package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
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

func TestHellos(t *testing.T) {
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hellos(names)

	if (err != nil) {
		t.Fatalf(`Hellos should not return an error, %v`, err)
	}

	if len(names) != len(messages) {
		t.Fatalf(`Number of messages should match the number of names`)
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"Gladys", ""}
	messages, err := Hellos(names)
	if (len(messages) != 0 || err == nil) {
		t.Fatalf(`Hellos should return an error if contains an empty name`)
	}
}
