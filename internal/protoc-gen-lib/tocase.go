package protocgenlib

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	tokenRegex = regexp.MustCompile("[\\s_-]")
)

func AsPointer(s string) string {
	if s[0] == '*' {
		return s
	}
	return "*" + s
}

func ToTitleCase(s string) string {
	tokens := tokenize(s)
	for i, tok := range tokens {
		tokens[i] = strings.Title(tok)
	}
	return strings.Join(tokens, "")
}

func ToCamelCase(s string) string {
	tokens := tokenize(s)
	tokens[0] = strings.ToLower(tokens[0])
	for i := 1; i < len(tokens); i++ {
		tokens[i] = strings.Title(tokens[i])
	}
	return strings.Join(tokens, "")
}

func ToSnakeCase(s string) string {
	tokens := tokenize(s)
	for i, tok := range tokens {
		tokens[i] = strings.ToLower(tok)
	}
	return strings.Join(tokens, "_")
}

func ToDashCase(s string) string {
	tokens := tokenize(s)
	for i, tok := range tokens {
		tokens[i] = strings.ToLower(tok)
	}
	return strings.Join(tokens, "-")
}

func tokenize(s string) []string {
	strs := []string{}
	tokens := tokenRegex.Split(s, -1)
	for _, tok := range tokens {
		if tok == "" {
			continue
		}

		for len(tok) > 0 {
			str, rem := parseToken(tok)
			strs = append(strs, str)
			tok = rem
		}
	}
	return strs
}

func parseToken(s string) (string, string) {
	var lastCap int
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			lastCap = i
			continue
		}
		break
	}

	for i := lastCap + 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			return s[:i], s[i:]
		}
	}

	return s, ""
}
