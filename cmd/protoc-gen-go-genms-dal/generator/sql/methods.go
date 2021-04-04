package generator_sql

import (
	"fmt"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// QueryMethod returns the function code for a query.
func QueryMethod(outfile *protogen.GeneratedFile, msg *protogen.Message, fields *generator.Fields, query *annotations.DalOptions_Query) (string, error) {
	/*	tmplSrc := `// {{ ToTitleCase .Query.Name }} implements {{ QualifiedDalType .Outfile .Msg }}Collection.{{ ToTitleCase .Query.Name }}
		func (x *{{ MessageName .Msg }}Collection){{ ToTitleCase .Query.Name }}({{ .FuncArgs }}) ([]*{{ QualifiedType .Outfile .Msg }}, error) {
			filter := &{{ QualifiedDalType .Outfile .Msg }}Filter{
				{{ range .QueryArgs -}}
					{{ . }},
				{{ end -}}
			}
			return x.find(ctx, "{{ ToSnakeCase .Query.Name }}", x.query{{ ToTitleCase .Query.Name }}, filter)
		}`

		tmpl, err := template.New("queryMethod").
			Funcs(template.FuncMap{
				"ToSnakeCase": protocgenlib.ToSnakeCase,
				"ToTitleCase": protocgenlib.ToTitleCase,
				"ToLower":     strings.ToLower,
			}).
			Parse(tmplSrc)

		if err != nil {
			return "", err
		}

		if query.Mode == annotations.DalOptions_Query_QueryMode_InterfaceStub {
			return "", nil
		}

		ctx := generator.QualifiedPackageName(outfile, "context")
		funcArgs := []string{fmt.Sprintf("ctx %s.Context", ctx)}
		queryArgs := []string{}
		for _, f := range query.Args {
			field := fields.ByName(f)

			fieldType, err := generator.GoFieldType(outfile, field)
			if err != nil {
				return "", err
			}
			funcArgs = append(funcArgs, fmt.Sprintf("%s %s", f, fieldType))

			pointer := '&'
			if field.Desc.Kind().String() == "message" {
				pointer = ' '
			}
			queryArgs = append(queryArgs, fmt.Sprintf("%s: %c%s", generator.GoFieldName(field), pointer, f))
		}

		values := map[string]interface{}{
			"Outfile":   outfile,
			"Msg":       msg,
			"Query":     query,
			"FuncArgs":  strings.Join(funcArgs, ", "),
			"QueryArgs": queryArgs,
		}

		buf := &bytes.Buffer{}
		if err := tmpl.Execute(buf, values); err != nil {
			return "", err
		}

		return buf.String(), nil
	*/
	return "", nil
}

// QueryImplementation returns the function code for a query.
func QueryImplementation(outfile *protogen.GeneratedFile, msg *protogen.Message, query *annotations.DalOptions_Query) (string, error) {
	/*
		tmplSrc := `// generate {{ ToTitleCase .Query.Name }} query
		coll.query{{ ToTitleCase .Query.Name }} = {{ .P.GenerateSQL }}.MustGenerateQuery("{{ QualifiedDalType .Outfile .Message }}-Query-{{ ToTitleCase .Query.Name }}", queries.{{ ToTitleCase .Query.Name }}(), queryReplacements)`

		tmpl, err :=
			template.New("queryImplementation").
				Funcs(template.FuncMap{
					"QualifiedDalType": generator.QualifiedKind,
					"ToTitleCase":      protocgenlib.ToTitleCase,
				}).
				Parse(tmplSrc)

		if err != nil {
			return "", err
		}

		if query.Mode == annotations.DalOptions_Query_QueryMode_InterfaceStub {
			return "", nil
		}

		p := map[string]string{
			"GenerateSQL": generator.QualifiedPackageName(outfile, "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator/sql"),
		}

		values := map[string]interface{}{
			"P":       p,
			"Outfile": outfile,
			"Message": msg,
			"Query":   query,
		}

		buf := &bytes.Buffer{}
		if err := tmpl.Execute(buf, values); err != nil {
			return "", err
		}

		return buf.String(), nil
	*/
	return "", nil
}

// QueryTemplate returns the generated query template.
func QueryTemplate(msg *protogen.Message, fields *generator.Fields, query *annotations.DalOptions_Query) (string, error) {
	/*tmplSrc := `// {{ ToTitleCase .query.Name }} implements {{ MessageName .msg }}QueryTemplateProvider.{{ ToTitleCase .query.Name }}.
	func (x *{{ MessageName .msg }}Queries) {{ ToTitleCase .query.Name }}() string {
		return ` + "`" + `SELECT {{ "{{ .fields }}" }} FROM {{ "{{ .table }}" }}
		{{- if .clauses }}
			WHERE
				{{ .clauses }};` + "`" + `
		{{ end -}}
	}
	`
		tmpl, err :=
			template.New("query").
				Funcs(template.FuncMap{
					"MessageName": generator.MessageName,
					"ToTitleCase": protocgenlib.ToTitleCase,
				}).
				Parse(tmplSrc)
		if err != nil {
			return "", err
		}

		switch query.Mode {
		case annotations.DalOptions_Query_QueryMode_InterfaceStub, annotations.DalOptions_Query_QueryMode_ProviderStub:
			return "", nil
		}

		queryArgs := []string{}
		for _, arg := range query.Args {
			qName := fields.QueryNameByFieldName(arg)
			queryArgs = append(queryArgs, fmt.Sprintf("%s = :%s", qName, qName))
		}

		values := map[string]interface{}{
			"msg":     msg,
			"query":   query,
			"clauses": strings.Join(queryArgs, " AND\n\t\t\t"),
		}

		buf := &bytes.Buffer{}
		if err := tmpl.Execute(buf, values); err != nil {
			return "", err
		}

		return buf.String(), nil
	*/
	return "", nil
}

// NullTypeToGoType returns a statement that gives the value of the sql nulltype as the
// required go type.
func NullTypeToGoType(outfile *protogen.GeneratedFile, obj string, name string, field *protogen.Field) (string, error) {
	/*kind, err := generator.GoFieldType(outfile, field)
	if err != nil {
		return "", nil
	}

	switch kind {
	case "bool":
		return fmt.Sprintf("%s%s.Bool", obj, name), nil
	case "float64":
		return fmt.Sprintf("%s%s.Float64", obj, name), nil
	case "float32":
		return fmt.Sprintf("float32(%s%s.Float64)", obj, name), nil
	case "int32":
		return fmt.Sprintf("%s%s.Int32", obj, name), nil
	case "int64":
		return fmt.Sprintf("%s%s.Int64", obj, name), nil
	case "string":
		return fmt.Sprintf("%s%s.String", obj, name), nil
	default:
		if field.Enum != nil {
			eType, err := generator.GoFieldType(outfile, field)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%s(%s%s.Int32)", eType, obj, name), nil
		}
		return fmt.Sprintf("%s%s", obj, name), nil
	}
	*/
	return "", nil
}

// ProtoToNullType returns the sql null type for the given proto type
func ProtoToNullType(outfile *protogen.GeneratedFile, field *protogen.Field) (string, error) {
	/*
		pkg := generator.QualifiedPackageName(outfile, "database/sql")
		kind, err := generator.GoFieldType(outfile, field)
		if err != nil {
			return "", nil
		}

		switch kind {
		case "bool":
			return pkg + ".NullBool", nil
		case "float64":
			return pkg + ".NullFloat64", nil
		case "float32":
			return pkg + ".NullFloat64", nil
		case "int32":
			return pkg + ".NullInt32", nil
		case "int64":
			return pkg + ".NullInt64", nil
		case "string":
			return pkg + ".NullString", nil
		default:
			if field.Enum != nil {
				return pkg + ".NullInt32", nil
			}
			return kind, nil
		}
	*/
	return "", nil
}

// QueryTemplateProviderMethod returns the method that provides a query template
func QueryTemplateProviderMethod(query *annotations.DalOptions_Query) string {
	switch query.Mode {
	case annotations.DalOptions_Query_QueryMode_Auto, annotations.DalOptions_Query_QueryMode_ProviderStub:
		return fmt.Sprintf("%s() string", protocgenlib.ToTitleCase(query.Name))
	default:
		return ""
	}
}
