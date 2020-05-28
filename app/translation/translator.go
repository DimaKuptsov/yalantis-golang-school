package translation

import (
	str "strings"
)

type Translator interface {
	Translate(stringForTranslate string) string
}

type PigLatinTranslator struct {
	vowelLatters  map[string]struct{}
	punctualMarks map[string]struct{}
}

func GetTranslator() Translator {
	englishVowelLatters := map[string]struct{}{
		"a": struct{}{},
		"e": struct{}{},
		"i": struct{}{},
		"o": struct{}{},
		"u": struct{}{},
	}
	punctualMarks := map[string]struct{}{
		",":  struct{}{},
		".":  struct{}{},
		"?":  struct{}{},
		"!":  struct{}{},
		":":  struct{}{},
		";":  struct{}{},
		"\"": struct{}{},
	}
	return PigLatinTranslator{vowelLatters: englishVowelLatters, punctualMarks: punctualMarks}
}

func (translator PigLatinTranslator) Translate(stringForTranslate string) string {
	const pigLatinSuffix = "ay"
	var translatedString string
	var letterForMove string
	var needMoveLatter bool
	var isFirstWordLetter bool
	for i, letter := range str.ToLower(stringForTranslate) {
		stringLetter := string(letter)
		if i == 0 {
			isFirstWordLetter = true
		}
		if isFirstWordLetter {
			_, isPunctualMark := translator.punctualMarks[stringLetter]
			if stringLetter == " " || isPunctualMark {
				translatedString += stringLetter
				continue
			}
			_, isVowel := translator.vowelLatters[stringLetter]
			if !isVowel {
				letterForMove = stringLetter
				needMoveLatter = true
			}
		}
		if isFirstWordLetter && needMoveLatter {
			isFirstWordLetter = false
			needMoveLatter = false
			continue
		}
		_, isPunctualMark := translator.punctualMarks[stringLetter]
		if stringLetter == " " || isPunctualMark {
			translatedString += letterForMove
			translatedString += pigLatinSuffix
			letterForMove = ""
			isFirstWordLetter = true
			needMoveLatter = false
		}
		translatedString += stringLetter
	}
	return translatedString
}
