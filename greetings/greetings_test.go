package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName_ContainsName(t *testing.T) {
	name := "Jim"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Jim") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty_ReturnsError(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
