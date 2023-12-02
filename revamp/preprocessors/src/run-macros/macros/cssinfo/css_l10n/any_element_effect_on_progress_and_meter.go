package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/htmlelement"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func anyElementEffectOnProgressAndMeter(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	progressRef, err := htmlelement.Htmlelement(env, reg, "progress")
	if err != nil {
		return "", err
	}
	meterRef, err := htmlelement.Htmlelement(env, reg, "meter")
	if err != nil {
		return "", err
	}
	return "будь-який елемент; діє на " + progressRef + " і " + meterRef + ", але не на &lt;input type=\"range\"&gt; чи інші елементи", nil
}
