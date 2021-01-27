package generator

import (
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"
	"unicode"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	tokenRegex = regexp.MustCompile("[\\s_-]")
)

// PackageName returns the name of the package.
func PackageName(file *protogen.File) string {
	return string(file.GoPackageName)
}

// PackagePath returns the import path of the package.
func PackagePath(file *protogen.File) string {
	return string(file.GoImportPath)
}

// DalPackageName returns the dal package name.
func DalPackageName(file *protogen.File) string {
	return string("dal_" + file.GoPackageName)
}

// DalPackagePath returns the import path of the package.
func DalPackagePath(file *protogen.File) string {
	toks := append([]string{strings.ReplaceAll(file.GoImportPath.String(), "\"", "")}, "dal")
	return path.Join(toks...)
}

// QualifiedPackageName returns the qualified package name.
func QualifiedPackageName(outfile *protogen.GeneratedFile, path string) string {
	ident := protogen.GoIdent{GoImportPath: protogen.GoImportPath(path)}
	return strings.Split(outfile.QualifiedGoIdent(ident), ".")[0]
}

// MessageName returns the name of the message for which code is being generated.
func MessageName(msg *protogen.Message) string {
	return msg.GoIdent.GoName
}

// ServiceName returns the name of the service for which code is being generated.
func ServiceName(svc *protogen.Service) string {
	return svc.GoName
}

// GoFieldName returns the field type.
func GoFieldName(field *protogen.Field) string {
	if field == nil {
		log.Println("nil field")
		return "nil"

	}
	return field.GoName
}

// GoFieldType returns the field type.
func GoFieldType(outfile *protogen.GeneratedFile, field *protogen.Field) (string, error) {
	if field == nil {
		return "", fmt.Errorf("field is nil")

	}

	if field.Message != nil {
		return "*" + outfile.QualifiedGoIdent(field.Message.GoIdent), nil
	}
	if field.Enum != nil {
		return outfile.QualifiedGoIdent(field.Enum.GoIdent), nil
	}
	return ToGoType(field.Desc.Kind()), nil
}

// QueryStructField returns the field name that holds the given query.
func QueryStructField(query *annotations.DalOptions_Query) string {
	switch query.Mode {
	case annotations.DalOptions_Query_QueryMode_Auto, annotations.DalOptions_Query_QueryMode_ProviderStub:
		return fmt.Sprintf("query%s string", ToTitleCase(query.Name))
	default:
		return ""
	}
}

// QualifiedType returns the qualified type of the message.
func QualifiedType(outfile *protogen.GeneratedFile, msg *protogen.Message) string {
	return outfile.QualifiedGoIdent(msg.GoIdent)
}

// QualifiedDalType returns the qualified type of the message.
func QualifiedDalType(outfile *protogen.GeneratedFile, msg *protogen.Message) string {
	m := protogen.GoIdent{
		GoName:       msg.GoIdent.GoName,
		GoImportPath: msg.GoIdent.GoImportPath + "/dal",
	}
	return outfile.QualifiedGoIdent(m)
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
