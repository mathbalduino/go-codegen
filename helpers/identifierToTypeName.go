package helpers

import "strings"

const invalidCharacters = " .[]{}(),;<->"
const pointerReplacement = "Ptr_"
const sliceReplacement = "Slc_"
const arrReplacement = "Arr_"
const variadicReplacement = "Variadic_"

func IdentifierToAsciiTypeName(typeIdentifier string) string {
	typename := ""
	for i, char := range typeIdentifier {
		if strings.ContainsRune(invalidCharacters, char) {
			if char == '.' && i >= 2 {
				lastThreeChars := string(typeIdentifier[i-2]) + string(typeIdentifier[i-1]) + string(char)
				if lastThreeChars == "..." {
					typename += variadicReplacement
				}
				continue
			}
			if char == '[' {
				if i >= 2 {
					lastThreeChars := string(typeIdentifier[i-3]) + string(typeIdentifier[i-2]) + string(typeIdentifier[i-1])
					if lastThreeChars == "map" {
						continue
					}
				}

				if typeIdentifier[i+1] == ']' {
					typename += sliceReplacement
					continue
				}
				typename += arrReplacement
			}
			continue
		}

		if char == '*' {
			typename += pointerReplacement
			continue
		}

		typename += string(char)
	}
	return typename
}

func IdentifierToTypeName(typeIdentifier string) string {
	newStr := ""
	for _, currChar := range typeIdentifier {
		newUnicodeChar, hasReplacement := identifierToTypeNameDictionary[currChar]
		if !hasReplacement {
			newStr += string(currChar)
			continue
		}
		newStr += string(newUnicodeChar)
	}
	return newStr
}

var identifierToTypeNameDictionary = map[rune]rune{
	'(': 'ᑕ',
	')': 'ᑐ',
	'[': 'ⵎ',
	']': 'コ',
	'<': 'ᐸ',
	'>': 'ᐳ',
	'{': 'ᓬ',
	'}': 'ᕒ',
	' ': 'ᆢ',
	'*': 'ᕽ',
	'.': 'ꓸ',
	'-': 'ｰ',
	';': 'ꓼ',
	',': 'ꓹ',
}

func GetCharReplacement(a rune) (rune, bool) {
	b, exists := identifierToTypeNameDictionary[a]
	return b, exists
}
