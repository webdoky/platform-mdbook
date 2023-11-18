package helpers_test

import (
	"regexp"
	"testing"
	"webdoky3/revamp/preprocessors/src/helpers"
)

func TestReplaceAllStringSubmatchFunc(t *testing.T) {
	var tests = []struct {
		re   *regexp.Regexp
		str  string
		repl func([]string) string
		want string
	}{
		{
			re:   regexp.MustCompile(`\$(\w+)`),
			str:  "Hello $name, you have $count new messages.",
			repl: func(groups []string) string { return groups[1] },
			want: "Hello name, you have count new messages.",
		},
		{
			re:   regexp.MustCompile(`\$(\w+)`),
			str:  "Hello $name, you have $count new messages.",
			repl: func(groups []string) string { return groups[0] },
			want: "Hello $name, you have $count new messages.",
		},
		{
			re:   regexp.MustCompile(`\$(\w+)`),
			str:  "Hello $name, you have $count new messages.",
			repl: func(groups []string) string { return groups[1] + "!" },
			want: "Hello name!, you have count! new messages.",
		},
		{
			re:   regexp.MustCompile(`\$(\w+)`),
			str:  "Hello $name, you have $count new messages.",
			repl: func(groups []string) string { return groups[0] + "!" },
			want: "Hello $name!, you have $count! new messages.",
		},
		{
			// Works with unmatched regexp
			re:   regexp.MustCompile(`\$(\w+)`),
			str:  "Hello name, you have count new messages.",
			repl: func(groups []string) string { return groups[0] + "!" },
			want: "Hello name, you have count new messages.",
		},
	}

	for _, test := range tests {
		got := helpers.ReplaceAllStringSubmatchFunc(test.re, test.str, test.repl)
		if got != test.want {
			t.Errorf("ReplaceAllStringSubmatchFunc(%q, %q, f) = %q, want %q", test.re, test.str, got, test.want)
		}
	}
}
