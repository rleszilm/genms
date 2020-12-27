package rest

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

// QueryProviderMethod returns the method that provides a query template
func QueryProviderMethod(query *annotations.DalOptions_Query) string {
	switch query.Mode {
	case annotations.DalOptions_Query_QueryMode_Auto, annotations.DalOptions_Query_QueryMode_ProviderStub:
		return fmt.Sprintf("%s() (scheme string, method string, host string, path string, headers map[string]string, query []string, body []string)", generator.ToTitleCase(query.Name))
	default:
		return ""
	}
}

// QueryMethod returns the function code for a query.
func QueryMethod(outfile *protogen.GeneratedFile, msg *protogen.Message, fields *generator.Fields, query *annotations.DalOptions_Query) (string, error) {
	tmplSrc := `{{ with $state := . }}
// {{ ToTitleCase .Query.Name }} implements {{ QualifiedDalType .Outfile .Msg }}Collection.{{ ToTitleCase .Query.Name }}
func (x *{{ MessageName .Msg }}Collection){{ ToTitleCase .Query.Name }}({{ .FuncArgs }}) ([]*{{ QualifiedType .Outfile .Msg }}, error) {
	scheme, method, host, path, headers, args, body := x.queries.{{ ToTitleCase .Query.Name }}()

	values := {{ .P.URL }}.Values{}
	for _, arg := range args {
		switch arg {
		{{ range .Query.Args -}}	
		case "{{ . }}":
			values.Add("{{ . }}", {{ $state.P.Fmt }}.Sprintf("%v", {{ . }}))
		{{ end -}}
		}
	}
	
	bodyArgs := map[string]interface{}{}
	for _, arg := range body {
		switch arg {
		{{ range .Query.Args -}}	
		case "{{ . }}":
			bodyArgs["{{ . }}"] = {{ . }}
		{{ end -}}
		}
	}
	
	bodyBuf, err := {{ .P.JSON }}.Marshal(bodyArgs)
	if err != nil {
		return nil, err
	}

	bodyBytes := {{ .P.Bytes }}.NewBuffer(bodyBuf)
	
	url := {{ .P.URL }}.URL{
		Scheme: scheme,
		Host: host,
		Path: path,
		RawQuery: values.Encode(),
	}

	req, err := http.NewRequest(method, {{ .P.URL }}.String(), bodyBytes)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return x.find(ctx, "{{ ToSnakeCase .Query.Name }}", req)
}
{{ end }}
`

	tmpl, err := template.New("queryMethod").
		Funcs(template.FuncMap{
			"MessageName":      generator.MessageName,
			"QualifiedType":    generator.QualifiedType,
			"QualifiedDalType": generator.QualifiedDalType,
			"ToSnakeCase":      generator.ToSnakeCase,
			"ToTitleCase":      generator.ToTitleCase,
			"ToLower":          strings.ToLower,
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

	p := map[string]string{
		"Context": generator.QualifiedPackageName(outfile, "context"),
		"URL":     generator.QualifiedPackageName(outfile, "net/url"),
		"Fmt":     generator.QualifiedPackageName(outfile, "fmt"),
		"JSON":    generator.QualifiedPackageName(outfile, "encoding/json"),
		"IOUtil":  generator.QualifiedPackageName(outfile, "io/ioutil"),
		"Bytes":   generator.QualifiedPackageName(outfile, "bytes"),
	}

	values := map[string]interface{}{
		"Outfile":   outfile,
		"Msg":       msg,
		"Query":     query,
		"FuncArgs":  strings.Join(funcArgs, ", "),
		"QueryArgs": queryArgs,
		"Fields":    fields,
		"P":         p,
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, values); err != nil {
		return "", err
	}

	return buf.String(), nil
}
