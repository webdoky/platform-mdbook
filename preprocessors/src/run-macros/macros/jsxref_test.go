package macros_test

import (
	"testing"
	"webdoky3/preprocessors/src/preprocessor"
	"webdoky3/preprocessors/src/run-macros/macros"
	"webdoky3/preprocessors/src/run-macros/registry"
	"webdoky3/preprocessors/src/run-macros/runner"
)

func TestJsxrefMacroWithPlain(t *testing.T) {
	env := macros.Environment{
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
	expected := "<a href=\"/uk/docs/Web/JavaScript/Reference/Global_Objects/Object\"><code>Object</code></a>"
	actual := macrosRunner.Run(input)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestJsxrefMacroWithDisplayName(t *testing.T) {
	env := macros.Environment{
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
	input := "{{jsxref(\"Object\", \"Об'єкт\")}}"
	expected := "<a href=\"/uk/docs/Web/JavaScript/Reference/Global_Objects/Object\"><code>Об'єкт</code></a>"
	actual := macrosRunner.Run(input)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
