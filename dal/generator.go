package dal

import (
	"bytes"
	"text/template"
)

// RenderQuery panics if unable to render the query.
func RenderQuery(name string, source string, replacements interface{}) (string, error) {
	tmpl, err := template.New(name).Parse(source)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, replacements); err != nil {
		return "", err
	}

	return buf.String(), nil
}
