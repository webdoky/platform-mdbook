package macros_test

import (
	"testing"
	"webdoky3/revamp/preprocessors/src/preprocessor"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
	"webdoky3/revamp/preprocessors/src/run-macros/runner"
)

func TestJsxrefMacroWithPlain(t *testing.T) {
	env := environment.Environment{
		Locale: "uk",
	}
	registry := registry.NewRegistry(&preprocessor.Book{
		NonExhaustive: nil,
		Sections: []preprocessor.Section{
			{
				IsSeparator: false,
				Chapter: &preprocessor.Chapter{
					Path: "uk/docs/Web/JavaScript/Reference/Global_Objects/Object/index.md",
				},
			},
		},
	})
	macrosRunner := runner.NewMacrosRunner(&env, registry)
	input := "{{jsxref(\"Object\")}}"
	expected := "<a class=\"\" href=\"/uk/docs/Web/JavaScript/Reference/Global_Objects/Object\" title=\"\"><code>Object</code></a>"
	actual := macrosRunner.Run(input)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestJsxrefMacroWithDisplayName(t *testing.T) {
	env := environment.Environment{
		Locale: "uk",
	}
	registry := registry.NewRegistry(&preprocessor.Book{
		NonExhaustive: nil,
		Sections: []preprocessor.Section{
			{
				IsSeparator: false,
				Chapter: &preprocessor.Chapter{
					Path: "uk/docs/Web/JavaScript/Reference/Global_Objects/Object/index.md",
				},
			},
		},
	})
	macrosRunner := runner.NewMacrosRunner(&env, registry)
	input := "{{jsxref(\"Object\", \"Об'єкт\", \"\", true)}}"
	expected := "<a class=\"\" href=\"/uk/docs/Web/JavaScript/Reference/Global_Objects/Object\" title=\"\">Об&#39;єкт</a>"
	actual := macrosRunner.Run(input)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
