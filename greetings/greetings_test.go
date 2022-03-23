package greetings

import (
	"regexp"
	"testing"
)

func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHelloName(t *testing.T) {
	name := "lxp"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("lxp") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
