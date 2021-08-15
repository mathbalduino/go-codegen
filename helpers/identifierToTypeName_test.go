package helpers

import (
	"fmt"
	"testing"
)

func TestIdentifierToAsciiTypeName(t *testing.T) {
	t.Run("Type identifiers without pointer, array, slice, variadic or invalid chars should not be modified", func(t *testing.T) {
		typeIdentifier := "abc"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != typeIdentifier {
			t.Fatalf("Not expected to change")
		}
	})
	t.Run("Every '*' in the type identifier should be replaced", func(t *testing.T) {
		typeIdentifier := "*abc"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != fmt.Sprintf("%s%s", pointerReplacement, typeIdentifier[1:]) {
			t.Fatalf("Expected to replace pointer '*' char")
		}
	})
	t.Run("Blank spaces should be removed", func(t *testing.T) {
		typeIdentifier := "   "
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "" {
			t.Fatalf("Blank spaces not removed")
		}
	})
	t.Run("Brackets should be removed", func(t *testing.T) {
		typeIdentifier := "{}{}}{}{"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "" {
			t.Fatalf("Brackets not removed")
		}
	})
	t.Run("Parenthesis should be removed", func(t *testing.T) {
		typeIdentifier := "())(())("
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "" {
			t.Fatalf("Parenthesis not removed")
		}
	})
	t.Run("Commas should be removed", func(t *testing.T) {
		typeIdentifier := ",,"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "" {
			t.Fatalf("Commas not removed")
		}
	})
	t.Run("Double-commas should be removed", func(t *testing.T) {
		typeIdentifier := ";;;"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "" {
			t.Fatalf("Double-commas not removed")
		}
	})
	t.Run("Chan '<', '-' and '>' should be removed", func(t *testing.T) {
		typeIdentifier := "<->"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "" {
			t.Fatalf("Chan symbols not removed")
		}
	})
	t.Run("Slices should be properly rewritten", func(t *testing.T) {
		typeIdentifier := "[]someTypeName"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != fmt.Sprintf("%s%s", sliceReplacement, typeIdentifier[2:]) {
			t.Fatalf("Slice typeIdentifier not properly replaced")
		}
	})
	t.Run("Arrays should be properly rewritten", func(t *testing.T) {
		typeIdentifier := "[5]someTypeName"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != fmt.Sprintf("%s5%s", arrReplacement, typeIdentifier[3:]) {
			t.Fatalf("Array typeIdentifier not properly replaced")
		}
	})
	t.Run("Variadics should be properly rewritten", func(t *testing.T) {
		typeIdentifier := "...someTypeName"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != fmt.Sprintf("%s%s", variadicReplacement, typeIdentifier[3:]) {
			t.Fatalf("Variadic typeIdentifier not properly replaced")
		}
	})
	t.Run("Maps should be properly rewritten", func(t *testing.T) {
		typeIdentifier := "map[string]int"
		n := IdentifierToAsciiTypeName(typeIdentifier)
		if n != "mapstringint" {
			t.Fatalf("Array typeIdentifier not properly replaced")
		}
	})
}

func TestIdentifierToTypeName(t *testing.T) {
	t.Run("Type identifiers without invalid characters, should not be modified", func(t *testing.T) {
		typeIdentifier := "simpleTypeIdentifier"
		n := IdentifierToTypeName(typeIdentifier)
		if n != typeIdentifier {
			t.Fatalf("Not expected to be modified")
		}
	})
	t.Run("Should replace only the invalid chars correctly", func(t *testing.T) {
		prefix, suffix := "abc", "def"
		for invalidChar, replacementChar := range identifierToTypeNameDictionary {
			input := fmt.Sprintf("%s%c%s", prefix, invalidChar, suffix)
			output := fmt.Sprintf("%s%c%s", prefix, replacementChar, suffix)
			if IdentifierToTypeName(input) != output {
				t.Fatalf("Expected to be replaced correctly")
			}
		}
	})
}

func TestGetCharReplacement(t *testing.T) {
	t.Run("Should return nil and false if the char doesn't have a replacement", func(t *testing.T) {
		r, ok := GetCharReplacement('a')
		if ok {
			t.Fatalf("Expected to be false")
		}
		if r != 0 {
			t.Fatalf("Expected to be zero")
		}
	})
	t.Run("Should return the rune and true if the char have a replacement", func(t *testing.T) {
		r, ok := GetCharReplacement('(')
		if !ok {
			t.Fatalf("Expected to be true")
		}
		if r != identifierToTypeNameDictionary['('] {
			t.Fatalf("Expected to be equal the replacement")
		}
	})
}
