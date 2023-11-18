package macros_test

import (
	"testing"
	"webdoky3/revamp/preprocessors/src/preprocessor"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
	"webdoky3/revamp/preprocessors/src/run-macros/runner"
)

func TestDomxrefMacroWithPlain(t *testing.T) {
	env := environment.Environment{
		Locale: "uk",
	}
	registry := registry.NewRegistry(&preprocessor.Book{
		NonExhaustive: nil,
		Sections: []preprocessor.Section{
			{
				IsSeparator: false,
				Chapter: &preprocessor.Chapter{
					Path: "uk/docs/Web/API/Element/index.md",
				},
			},
		},
	})
	macrosRunner := runner.NewMacrosRunner(&env, registry)
	input := "{{domxref(\"Element\")}}"
	expected := "<a class=\"\" href=\"/uk/docs/Web/API/Element\" title=\"\"><code>Element</code></a>"
	actual := macrosRunner.Run(input)
	if actual != expected {
		if len(actual) != len(expected) {
			t.Errorf("Expected %d, got %d", len(expected), len(actual))
		}
		// Show difference between strings
		t.Log("Expected:")
		t.Log(expected)
		t.Log("Actual:")
		t.Log(actual)
		difference := ""
		for i := 0; i < len(actual); i++ {
			if i >= len(expected) || expected[i] != actual[i] {
				difference += "^"
			} else {
				difference += " "
			}
		}
		t.Log(difference)
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
