package macros

import (
	"log"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func csssyntax(env *environment.Environment, reg *registry.Registry, _ string) (string, error) {
	log.Println("csssyntax, " + env.Path)
	return "Помилка: не вдалося знайти синтаксис", nil
}
