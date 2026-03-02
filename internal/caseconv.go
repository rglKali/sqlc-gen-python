package internal

import (
	"strings"
	"unicode"
)

// toWords splits any of camelCase, PascalCase, snake_case, or UPPER_CASE into words.
func toWords(s string) []string {
	// Replace underscores with spaces to handle snake_case / UPPER_CASE
	s = strings.ReplaceAll(s, "_", " ")

	var words []string
	var current strings.Builder

	runes := []rune(s)
	for i, r := range runes {
		if r == ' ' {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
			continue
		}

		// Detect transition from lower/digit to upper (camel/pascal boundary)
		if unicode.IsUpper(r) && current.Len() > 0 {
			// Peek ahead: if next char is lower, we're starting a new word
			// e.g. "XMLParser" → "XML", "Parser"
			prevIsUpper := unicode.IsUpper(runes[i-1])
			nextIsLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if !prevIsUpper || nextIsLower {
				words = append(words, current.String())
				current.Reset()
			}
		}

		current.WriteRune(r)
	}
	if current.Len() > 0 {
		words = append(words, current.String())
	}

	return words
}

// toSnakeCase converts any supported format to snake_case.
// e.g. "camelCase" → "camel_case", "PascalCase" → "pascal_case", "UPPER_CASE" → "upper_case"
func toSnakeCase(s string) string {
	words := toWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "_")
}

// toUpperCase converts any supported format to UPPER_CASE.
// e.g. "camelCase" → "CAMEL_CASE", "snake_case" → "SNAKE_CASE"
func toUpperCase(s string) string {
	words := toWords(s)
	for i, w := range words {
		words[i] = strings.ToUpper(w)
	}
	return strings.Join(words, "_")
}

// toCamelCase converts any supported format to camelCase.
// e.g. "snake_case" → "snakeCase", "PascalCase" → "pascalCase"
func toCamelCase(s string) string {
	words := toWords(s)
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
		} else {
			words[i] = capitalize(w)
		}
	}
	return strings.Join(words, "")
}

// toPascalCase converts any supported format to PascalCase.
// e.g. "snake_case" → "SnakeCase", "camelCase" → "CamelCase"
func toPascalCase(s string) string {
	words := toWords(s)
	for i, w := range words {
		words[i] = capitalize(w)
	}
	return strings.Join(words, "")
}

// capitalize lowercases a word then uppercases its first rune.
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ToLower(s)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
