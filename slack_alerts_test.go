package slackAlerts

import (
	"regexp"
	"testing"
)

func TestMessage(t *testing.T) {
	name := "ok"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Message("A message *with some bold statements* and _some italicized text_.")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Message("hi") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
