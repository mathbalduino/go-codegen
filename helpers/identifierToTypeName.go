package helpers

import "strings"

const invalidCharacters = " .[]{}(),;<->"
const pointerReplacement = "Ptr_"
const sliceReplacement = "Slc_"
const arrReplacement = "Arr_"
const variadicReplacement = "Variadic_"

// IdentifierToAsciiTypeName will transform a type identifier into a
// ASCII-only GO type name
//
// Example: "*int" turns into "Ptr_int"
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

// IdentifierToTypeName will transform a type identifier into
// a valid GO type name. Note that it will use non-ASCII symbols
//
// Example: "*int" turns into "ᕽint"
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

// identifierToTypeNameDictionary maps from invalid
// GO type name symbols to valid alternatives
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

// GetCharReplacement will return the valid equivalent of an
// invalid GO type name symbol
func GetCharReplacement(a rune) (rune, bool) {
	b, exists := identifierToTypeNameDictionary[a]
	return b, exists
}
