package css_l10n

import (
	"encoding/json"
	"log"
	"os"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

var plainL10nData map[string]map[string]string

var l10nData = map[string]map[string]func(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error){
	"absoluteLength": {
		"uk": absoluteLength,
	},
	"absoluteLength0ForNone": {
		"uk": absoluteLength0ForNone,
	},
	"absoluteLength0IfColumnRuleStyleNoneOrHidden": {
		"uk": absoluteLength0IfColumnRuleStyleNoneOrHidden,
	},
	"absoluteLengthOr0IfBorderBottomStyleNoneOrHidden": {
		"uk": absoluteLengthOr0IfBorderBottomStyleNoneOrHidden,
	},
	"absoluteLengthOr0IfBorderLeftStyleNoneOrHidden": {
		"uk": absoluteLengthOr0IfBorderLeftStyleNoneOrHidden,
	},
	"absoluteLengthOr0IfBorderRightStyleNoneOrHidden": {
		"uk": absoluteLengthOr0IfBorderRightStyleNoneOrHidden,
	},
	"absoluteLengthOr0IfBorderTopStyleNoneOrHidden": {
		"uk": absoluteLengthOr0IfBorderTopStyleNoneOrHidden,
	},
	"absoluteLengthOrNone": {
		"uk": absoluteLengthOrNone,
	},
	"absoluteLengthOrNormal": {
		"uk": absoluteLengthOrNormal,
	},
	"absoluteLengthOrPercentage": {
		"uk": absoluteLengthOrPercentage,
	},
	"allElementsAndPseudos": {
		"uk": allElementsAndPseudos,
	},
	"allElementsExceptTableDisplayTypes": {
		"uk": allElementsExceptTableDisplayTypes,
	},
	"allElementsExceptTableElementsWhenCollapse": {
		"uk": allElementsExceptTableElementsWhenCollapse,
	},
	"allElementsNoEffectIfDisplayNone": {
		"uk": allElementsNoEffectIfDisplayNone,
	},
	"allElementsSVGContainerElements": {
		"uk": allElementsSVGContainerElements,
	},
	"allElementsThatCanReferenceImages": {
		"uk": allElementsThatCanReferenceImages,
	},
	"allElementsUAsNotRequiredWhenCollapse": {
		"uk": allElementsUAsNotRequiredWhenCollapse,
	},
	"angleBasicShapeOrPath": {
		"uk": angleBasicShapeOrPath,
	},
	"angleRoundedToNextQuarter": {
		"uk": angleRoundedToNextQuarter,
	},
	"anyElementEffectOnProgressAndMeter": {
		"uk": anyElementEffectOnProgressAndMeter,
	},
	"applyingToMultiple": {
		"uk": applyingToMultiple,
	},
	"asAutoOrColor": {
		"uk": asAutoOrColor,
	},
	"asDefinedForBasicShapeWithAbsoluteURIOtherwiseAsSpecified": {
		"uk": asDefinedForBasicShapeWithAbsoluteURIOtherwiseAsSpecified,
	},
	"asLength": {
		"uk": asLength,
	},
	"asSpecifiedButVisibleOrClipReplacedToAutoOrHiddenIfOtherValueDifferent": {
		"uk": asSpecifiedButVisibleOrClipReplacedToAutoOrHiddenIfOtherValueDifferent,
	},
	"asSpecifiedURLsAbsolute": {
		"uk": asSpecifiedURLsAbsolute,
	},
	"autoOnAbsolutelyPositionedElementsValueOfAlignItemsOnParent": {
		"uk": autoOnAbsolutelyPositionedElementsValueOfAlignItemsOnParent,
	},
	"basicShapeOtherwiseNo": {
		"uk": basicShapeOtherwiseNo,
	},
	"beforeAndAfterPseudos": {
		"uk": beforeAndAfterPseudos,
	},
	"directChildrenOfElementsWithDisplayMozBoxMozInlineBox": {
		"uk": directChildrenOfElementsWithDisplayMozBoxMozInlineBox,
	},
	"eachOfShorthandPropertiesExceptUnicodeBiDiAndDirection": {
		"uk": eachOfShorthandPropertiesExceptUnicodeBiDiAndDirection,
	},
	"elementsWithDisplayBoxOrInlineBox": {
		"uk": elementsWithDisplayBoxOrInlineBox,
	},
	"elementsWithDisplayMarker": {
		"uk": elementsWithDisplayMarker,
	},
	"elementsWithDisplayMozBoxMozInlineBox": {
		"uk": elementsWithDisplayMozBoxMozInlineBox,
	},
	"elementsWithOverflowNotVisibleAndReplacedElements": {
		"uk": elementsWithOverflowNotVisibleAndReplacedElements,
	},
	"firstLetterPseudoElementsAndInlineLevelFirstChildren": {
		"uk": firstLetterPseudoElementsAndInlineLevelFirstChildren,
	},
	"lpc": {
		"uk": lpc,
	},
	"maskElements": {
		"uk": maskElements,
	},
	"normalOnElementsForPseudosNoneAbsoluteURIStringOrAsSpecified": {
		"uk": normalOnElementsForPseudosNoneAbsoluteURIStringOrAsSpecified,
	},
	"referToLineHeight": {
		"uk": referToLineHeight,
	},
	"referToSizeOfMaskPaintingArea": {
		"uk": referToSizeOfMaskPaintingArea,
	},
	"sameAsBoxOffsets": {
		"uk": sameAsBoxOffsets,
	},
	"sameAsMargin": {
		"uk": sameAsMargin,
	},
	"sameAsMaxWidthAndMaxHeight": {
		"uk": sameAsMaxWidthAndMaxHeight,
	},
	"sameAsMinWidthAndMinHeight": {
		"uk": sameAsMinWidthAndMinHeight,
	},
	"sameAsWidthAndHeight": {
		"uk": sameAsWidthAndHeight,
	},
	"specifiedValueNumberClipped0To1": {
		"uk": specifiedValueNumberClipped0To1,
	},
	"startOrNamelessValueIfLTRRightIfRTL": {
		"uk": startOrNamelessValueIfLTRRightIfRTL,
	},
	"twoAbsoluteLengthOrPercentages": {
		"uk": twoAbsoluteLengthOrPercentages,
	},
	"xulImageElements": {
		"uk": xulImageElements,
	},
}

func HasKey(key string) bool {
	_, ok := plainL10nData[key]
	if ok {
		return true
	}
	_, ok = l10nData[key]
	return ok
}

func Localize(env *environment.Environment, reg *registry.Registry, key string, arg1 string, arg2 string) (string, error) {
	locales, ok := plainL10nData[key]
	if ok {
		translation, ok := locales[env.Locale]
		if ok {
			return translation, nil
		} else {
			log.Fatal("Locale not found")
		}
	}
	complexLocales, ok := l10nData[key]
	if ok {
		translationFunc, ok := complexLocales[env.Locale]
		if ok {
			return translationFunc(env, reg, arg1, arg2)
		} else {
			log.Fatal("Locale not found")
		}
	}
	return key, nil
}

func init() {
	rawJson, err := os.ReadFile("revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n/plain.json")
	if err != nil {
		log.Fatal(err)
	}
	plainL10nData = make(map[string]map[string]string)
	err = json.Unmarshal(rawJson, &plainL10nData)
	if err != nil {
		log.Fatal(err)
	}
}
