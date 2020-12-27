package rest

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator"
	generator_sql "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator/sql"
	"google.golang.org/protobuf/compiler/protogen"
)

// Collection is a struct that generates a colelction file.
type Collection struct {
	Outfile     *protogen.GeneratedFile
	File        *protogen.File
	Message     *protogen.Message
	Opts        *annotations.DalOptions
	Fields      *generator.Fields
	QueryFields string
}

func (c *Collection) render() error {
	steps := []func() error{
		c.definePackage,
		c.defineCollection,
		c.defineService,
		c.defineDefaultQueries,
		c.defineQueries,
		c.defineNewCollection,
		c.defineScanner,
		c.defineConfig,
		c.defineQueryProvider,
		c.defineMetrics,
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collection) definePackage() error {
	tmplSrc := `// Package rest_{{ DalPackageName .File }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package rest_{{ DalPackageName .File }}
`

	tmpl, err := template.New("definePackage").
		Funcs(template.FuncMap{
			"DalPackageName": generator.DalPackageName,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, c); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineCollection() error {
	tmplSrc := `{{ with $state := . }}
// {{ MessageName .C.Message }}Collection is an autogenerated implementation of {{ QualifiedDalType .C.Outfile .C.Message }}Collection.
type {{ MessageName .C.Message }}Collection struct {
	{{ .P.Collection }}.Unimplemented{{ MessageName .C.Message }}Collection

	client *{{ .P.HTTP }}.Client
	config *{{ MessageName .C.Message }}Config
	
	queries {{ MessageName .C.Message }}QueryProvider
}
{{ end }}
`

	tmpl, err := template.New("defineCollection").
		Funcs(template.FuncMap{
			"DalPackageName":   generator.DalPackageName,
			"MessageName":      generator.MessageName,
			"QualifiedDalType": generator.QualifiedDalType,
			"QueryStructField": generator.QueryStructField,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": generator.QualifiedPackageName(c.Outfile, generator.DalPackagePath(c.File)),
		"HTTP":       generator.QualifiedPackageName(c.Outfile, "net/http"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineService() error {
	tmplSrc := `// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *{{ MessageName .C.Message }}Collection) Initialize(_ {{ .P.Context }}.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *{{ MessageName .C.Message }}Collection) Shutdown(_ {{ .P.Context }}.Context) error {
	return nil
}

// NameOf returns the name of a service. This must be unique if there are multiple instances of the same
// service.
func (x *{{ MessageName .C.Message }}Collection) NameOf() string {
	return "rest_{{ DalPackageName .C.File }}_" + x.config.Name
}

// String returns a string identifier for the service.
func (x *{{ MessageName .C.Message }}Collection) String() string {
	return x.NameOf()
}
`

	tmpl, err := template.New("defineService").
		Funcs(template.FuncMap{
			"DalPackageName": generator.DalPackageName,
			"MessageName":    generator.MessageName,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Context": generator.QualifiedPackageName(c.Outfile, "context"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineDefaultQueries() error {
	tmplSrc := `// All implements {{ QualifiedDalType .C.Outfile .C.Message }}Collection.All
func (x *{{ MessageName .C.Message }}Collection) All(ctx {{ .P.Context }}.Context) ([]*{{ QualifiedType .C.Outfile .C.Message }}, error) {
	scheme, method, host, path, headers, _, _ := x.queries.All()

	url := url.URL{
		Scheme: scheme,
		Host: host,
		Path: path,
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return x.find(ctx, "all", req)
}

func (x *{{ MessageName .C.Message }}Collection) find(ctx {{ .P.Context }}.Context, label string, req *http.Request) ([]*{{ QualifiedType .C.Outfile .C.Message }}, error) {
	var err error
	start := {{ .P.Time }}.Now()
	{{ .P.Stats }}.Record(ctx, {{ ToCamelCase (MessageName .C.Message) }}Inflight.M(1))
	defer func() {
		stop := {{ .P.Time }}.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)

		if err != nil {
			ctx, err = {{ .P.Tag }}.New(ctx,
				{{ .P.Tag }}.Upsert({{ ToCamelCase (MessageName .C.Message) }}QueryError, label),
			)
		}

		ctx, err = {{ .P.Tag }}.New(ctx,
			{{ .P.Tag }}.Upsert({{ ToCamelCase (MessageName .C.Message) }}QueryName, label),
		)

		{{ .P.Stats }}.Record(ctx, {{ ToCamelCase (MessageName .C.Message) }}Latency.M(dur), {{ ToCamelCase (MessageName .C.Message) }}Inflight.M(-1))
	}()

	req.WithContext(ctx)
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}

	buff, err := {{ .P.IOUtil }}.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	{{ MessageName .C.Message }}s := []*{{ QualifiedType .C.Outfile .C.Message }}{}
	if err := {{ .P.JSON }}.Unmarshal(buff, &{{ MessageName .C.Message }}s); err != nil {
		return nil, err
	}
	return {{ MessageName .C.Message }}s, nil
}
`

	tmpl, err := template.New("defineDefaultQueries").
		Funcs(template.FuncMap{
			"DalPackageName":   generator.DalPackageName,
			"QualifiedDalType": generator.QualifiedDalType,
			"MessageName":      generator.MessageName,
			"QualifiedType":    generator.QualifiedType,
			"ToCamelCase":      generator.ToCamelCase,
			"ToLower":          strings.ToLower,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	fieldArgs := []string{}
	for _, name := range c.Fields.QueryNames() {
		field := c.Fields.ByQueryName(name)
		if field == nil {
			continue
		}

		fname := generator.GoFieldName(field)
		fieldArgs = append(fieldArgs, fmt.Sprintf(`if arg.%s != nil {
			fields = append(fields, "%s = :%s")
		}`, fname, name, name))
	}

	p := map[string]string{
		"Context": generator.QualifiedPackageName(c.Outfile, "context"),
		"Time":    generator.QualifiedPackageName(c.Outfile, "time"),
		"Tag":     generator.QualifiedPackageName(c.Outfile, "go.opencensus.io/tag"),
		"Stats":   generator.QualifiedPackageName(c.Outfile, "go.opencensus.io/stats"),
		"URL":     generator.QualifiedPackageName(c.Outfile, "net/url"),
		"JSON":    generator.QualifiedPackageName(c.Outfile, "encoding/json"),
		"IOUtil":  generator.QualifiedPackageName(c.Outfile, "io/ioutil"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C":         c,
		"P":         p,
		"FieldArgs": fieldArgs,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineQueries() error {
	tmplSrc := `{{ range .Queries }}
	{{ . }}
{{ end }}
`

	tmpl, err := template.New("defineQueries").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	queries := []string{}
	for _, q := range c.Opts.Queries {
		query, err := QueryMethod(c.Outfile, c.Message, c.Fields, q)
		if err != nil {
			return err
		}
		queries = append(queries, query)
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{"Queries": queries}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineNewCollection() error {
	tmplSrc := `// New{{ MessageName .C.Message }}Collection returns a new {{ MessageName .C.Message }}Collection.
func New{{ MessageName .C.Message }}Collection(queries {{ MessageName .C.Message }}QueryProvider, config *{{ MessageName .C.Message }}Config) (*{{ MessageName .C.Message }}Collection, error) {
	register{{ ToTitleCase (MessageName .C.Message) }}MetricsOnce.Do(register{{ ToTitleCase (MessageName .C.Message) }}Metrics)

	coll := &{{ MessageName .C.Message }}Collection{
		Unimplemented{{ MessageName .C.Message }}Collection: {{ .P.Collection }}.Unimplemented{{ MessageName .C.Message }}Collection{},
		config: config,
		queries: queries,
		client: &{{ .P.HTTP }}.Client{
			Timeout: config.Timeout,
		},
	}
	
	return coll, nil
}
`

	tmpl, err := template.New("defineNewCollection").
		Funcs(template.FuncMap{
			"MessageName":      generator.MessageName,
			"ToTitleCase":      generator.ToTitleCase,
			"DalPackageName":   generator.DalPackageName,
			"QualifiedDalType": generator.QualifiedDalType,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": generator.QualifiedPackageName(c.Outfile, generator.DalPackagePath(c.File)),
		"HTTP":       generator.QualifiedPackageName(c.Outfile, "net/http"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineScanner() error {
	tmplSrc := `// {{ MessageName .C.Message }}Scanner is an autogenerated struct that
// is used to parse query results.
type {{ MessageName .C.Message }}Scanner struct {
	{{ range .ScanFields -}}
		{{ . }}
	{{ end }}
}
		
// {{ MessageName .C.Message }} returns a new {{ QualifiedType .C.Outfile .C.Message }} populated with scanned values.
func (x *{{ MessageName .C.Message }}Scanner) {{ MessageName .C.Message }}() *{{ QualifiedType .C.Outfile .C.Message }} {
	return &{{ QualifiedType .C.Outfile .C.Message }}{
		{{ range .ScanValues -}}
			{{ . }},
		{{ end }}
	}
}
`

	tmpl, err := template.New("defineScanner").
		Funcs(template.FuncMap{
			"MessageName":   generator.MessageName,
			"QualifiedType": generator.QualifiedType,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	scannerFields := []string{}
	scannerValues := []string{}
	for _, q := range c.Fields.QueryNames() {
		field := c.Fields.ByQueryName(q)
		fname := generator.GoFieldName(field)

		pn, err := generator_sql.ProtoToNullType(c.Outfile, field)
		if err != nil {
			return err
		}
		scannerFields = append(scannerFields, fmt.Sprintf("%s %s `json:\"%s\"`", fname, pn, q))

		nt, err := generator_sql.NullTypeToGoType(c.Outfile, "x.", fname, field)
		if err != nil {
			return err
		}
		scannerValues = append(scannerValues, fmt.Sprintf("%s: %s", fname, nt))
	}

	p := map[string]string{}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C":          c,
		"P":          p,
		"ScanFields": scannerFields,
		"ScanValues": scannerValues,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineConfig() error {
	tmplSrc := `// {{ MessageName .C.Message }}Config is a struct that can be used to configure a {{ MessageName .C.Message }}Collection
	type {{ MessageName .C.Message }}Config struct {
		Name string ` + "`" + `envconfig:"name"` + "`" + `
		Timeout {{ .P.Time }}.Duration ` + "`" + `envconfig:"timeout" default:"5s"` + "`" + `
		Headers map[string]string ` + "`" + `envconfig:"headers"` + "`" + `
	}
`

	tmpl, err := template.New("defineConfig").
		Funcs(template.FuncMap{
			"MessageName": generator.MessageName,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Time": generator.QualifiedPackageName(c.Outfile, "time"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineQueryProvider() error {
	tmplSrc := `// {{ MessageName .C.Message }}QueryProvider is an interface that returns the query templated that should be executed
// to generate the queries that the collection will use.
type {{ MessageName .C.Message }}QueryProvider interface {
	All() (scheme string, method string, host string, path string, headers map[string]string, query []string, body []string)
	{{ range .C.Opts.Queries -}}
		{{ QueryProviderMethod . }}
	{{ end -}}
}

// {{ MessageName .C.Message }}QueryDetails is a function that provides the data needed to perform a query.
type {{ MessageName .C.Message }}QueryDetails func() (scheme string, method string, host string, path string, headers map[string]string, query []string, body {{ .P.IO }}.Reader)
`

	tmpl, err := template.New("defineQueryProvider").
		Funcs(template.FuncMap{
			"QueryProviderMethod": QueryProviderMethod,
			"MessageName":         generator.MessageName,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"IO": generator.QualifiedPackageName(c.Outfile, "io"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineMetrics() error {
	tmplSrc := `// define metrics
	var (
		{{ ToCamelCase (MessageName .C.Message) }}QueryName = {{ .P.Tag }}.MustNewKey("dal_rest_{{ ToSnakeCase (MessageName .C.Message) }}")
		{{ ToCamelCase (MessageName .C.Message) }}QueryCode = {{ .P.Tag }}.MustNewKey("dal_rest_{{ ToSnakeCase (MessageName .C.Message) }}_code")
		{{ ToCamelCase (MessageName .C.Message) }}QueryError = {{ .P.Tag }}.MustNewKey("dal_rest_{{ ToSnakeCase (MessageName .C.Message) }}_error")
	
		{{ ToCamelCase (MessageName .C.Message) }}Latency = stats.Float64("{{ ToSnakeCase (MessageName .C.Message) }}_latency", "Latency of {{ MessageName .C.Message }} queries", stats.UnitMilliseconds)
		{{ ToCamelCase (MessageName .C.Message) }}Inflight = stats.Int64("{{ ToSnakeCase (MessageName .C.Message) }}_inflight", "Count of {{ MessageName .C.Message }} queries in flight", stats.UnitDimensionless)
	
		register{{ ToTitleCase (MessageName .C.Message) }}MetricsOnce {{ .P.Sync }}.Once
	)
	
	func register{{ ToTitleCase (MessageName .C.Message) }}Metrics() {
		views := []*{{ .P.View }}.View{
			{
				Name:        "dal_rest_{{ ToSnakeCase (MessageName .C.Message) }}_latency",
				Measure:     {{ ToCamelCase (MessageName .C.Message) }}Latency,
				Description: "The distribution of the query latencies",
				TagKeys:     []{{ .P.Tag }}.Key{ {{ ToCamelCase (MessageName .C.Message) }}QueryName, {{ ToCamelCase (MessageName .C.Message) }}QueryCode, {{ ToCamelCase (MessageName .C.Message) }}QueryError},
				Aggregation: {{ .P.View }}.Distribution(0, 25, 100, 200, 400, 800, 10000),
			},
			{
				Name:        "dal_rest_{{ ToSnakeCase (MessageName .C.Message) }}_inflight",
				Measure:     {{ ToCamelCase (MessageName .C.Message) }}Inflight,
				Description: "The number of queries being processed",
				TagKeys:     []{{ .P.Tag }}.Key{ {{ ToCamelCase (MessageName .C.Message) }}QueryName, {{ ToCamelCase (MessageName .C.Message) }}QueryCode},
				Aggregation: {{ .P.View }}.Sum(),
			},
		}
	
		if err := {{ .P.View }}.Register(views...); err != nil {
			{{ .P.Log }}.Fatal("Cannot register metrics:", err)
		}
	}
`

	tmpl, err := template.New("defineMetrics").
		Funcs(template.FuncMap{
			"ToCamelCase":   generator.ToCamelCase,
			"ToSnakeCase":   generator.ToSnakeCase,
			"ToTitleCase":   generator.ToTitleCase,
			"MessageName":   generator.MessageName,
			"QueryTemplate": generator_sql.QueryTemplate,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Log":  generator.QualifiedPackageName(c.Outfile, "log"),
		"Sync": generator.QualifiedPackageName(c.Outfile, "sync"),
		"Tag":  generator.QualifiedPackageName(c.Outfile, "go.opencensus.io/tag"),
		"View": generator.QualifiedPackageName(c.Outfile, "go.opencensus.io/stats/view"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.Outfile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

// NewCollection returns a new collection renderer.
func NewCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *Collection {
	filename := fmt.Sprintf("dal/rest/%s.genms.dal.%s.go", file.GeneratedFilenamePrefix, strings.ToLower(msg.GoIdent.GoName))
	outFile := plugin.NewGeneratedFile(filename, ".")
	fields := generator.NewFields(msg, annotations.DalOptions_BackEnd_Rest)

	return &Collection{
		Outfile:     outFile,
		File:        file,
		Message:     msg,
		Opts:        opts,
		Fields:      fields,
		QueryFields: strings.Join(fields.QueryNames(), ", "),
	}
}

// GenerateCollection generates the dal interface for the collection
func GenerateCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewCollection(plugin, file, msg, opts)
	return c.render()
}

// NullTypeToGoType returns a statement that gives the value of the sql nulltype as the
// required go type.
func NullTypeToGoType(outFile *protogen.GeneratedFile, obj string, name string, field *protogen.Field) (string, error) {
	return generator_sql.NullTypeToGoType(outFile, obj, name, field)
}

// ProtoToNullType returns the sql null type for the given proto type
func ProtoToNullType(outFile *protogen.GeneratedFile, field *protogen.Field) (string, error) {
	return generator_sql.ProtoToNullType(outFile, field)
}
