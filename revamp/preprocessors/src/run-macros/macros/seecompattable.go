package macros

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func seecompattable(env *environment.Environment, reg *registry.Registry, _ string) (string, error) {
	return "<div class=\"notecard experimental\" id=\"sect1\"><p><strong>Експериментальне:</strong> <strong>Це <a href=\"/uk/docs/MDN/Writing_guidelines/Experimental_deprecated_obsolete#eksperymentalne\">експериментальна технологія</a></strong><br>Уважно звіртеся з <a href=\"#sumisnist-iz-brauzeramy\">таблицею Сумісності з браузерами</a>, перш ніж використовувати це в роботі.</p></div>", nil
}
