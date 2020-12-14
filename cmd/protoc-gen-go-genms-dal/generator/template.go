package generator

import (
	"errors"
	"path"
	"regexp"
	"strings"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	tokenRegex = regexp.MustCompile("[\\s_-]")
)

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

func ToFieldType(s string, msgIn interface{}) (string, error) {
	msg, ok := msgIn.(*protogen.Message)
	if !ok {
		return "", errors.New("msgIn is not a *protogen.Message")
	}

	field := msg.Desc.Fields().ByJSONName(ToCamelCase(s))
	if field == nil {
		return "", errors.New(s + " is not a valid field (" + ToCamelCase(s) + ")")
	}

	if field.Enum() != nil {
		return "", errors.New("enum query fields are not supported for generation")
	}

	if field.Message() != nil {
		return "", errors.New("message query fields are not supported for generation")
	}

	return field.Kind().String(), nil
}

func ToGoType(t protoreflect.Kind) string {
	switch t.String() {
	case "bool":
		return "bool"
	case "double":
		return "float64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "string":
		return "string"
	default:
		return t.GoString()
	}
}

func PackageName(prefix string, dirs ...string) string {
	toks := append([]string{strings.ReplaceAll(prefix, "\"", "")}, dirs...)
	return "\"" + path.Join(toks...) + "\""
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
