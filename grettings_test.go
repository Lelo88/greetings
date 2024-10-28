package greetings

import (
	"regexp"
	"testing"
)

// TestHelloEmpty calls greetings.Hello with an empty string, checking for an error.
// testing is a package that provides support for automated testing of Go packages.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	// want := "Hi, Gladys. Welcome!"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Gladys")
	if !want.MatchString(message) || err != nil {
		t.Fail()
		// Utilizamos t.Fatalf en lugar de t.Fail() para registrar información sobre el error y fallar el test.
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fail()
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}