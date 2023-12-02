package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func autoOnAbsolutelyPositionedElementsValueOfAlignItemsOnParent(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	alignItemsRef, err := cssxref.Cssxref(env, reg, "align-items")
	if err != nil {
		return "", err
	}
	justifySelfRef, err := cssxref.Cssxref(env, reg, "justify-self")
	if err != nil {
		return "", err
	}
	return "<code>auto</code> обчислюється в самого себе на абсолютно позиціонованих елементах, а на всіх інших рамках – в обчислене значення " + alignItemsRef + " на батьківському елементі (мінус усілякі історичні ключові слова), або в <code>start</code>, якщо рамка не має батьківського елемента. Його поведінка залежить від моделі компонування, як описано для " + justifySelfRef + ". Інакше – задане значення.", nil
}
