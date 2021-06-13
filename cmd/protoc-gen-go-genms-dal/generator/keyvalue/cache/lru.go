package cache

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/go-test/deep"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
)

// LRU is a struct that generates a lru cache file.
type LRU struct {
	File    *File
	Message *generator.Message
	Fields  *generator.Fields
	Queries *generator.Queries
	Opts    *annotations.DalOptions

	plugin   *protogen.Plugin
	filename string
}

// NewLRU returns a new updater renderer.
func NewLRU(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *LRU {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/keyvalue/cache/%s.genms.cache.lru.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	cfile := NewFile(outfile, file)
	cmsg := generator.NewMessage(cfile.Generator(), msg)
	cfields := generator.NewFields(cmsg)
	cqueries := generator.NewQueries(cfile.Generator(), cfields, opts)

	return &LRU{
		File:     cfile,
		Message:  cmsg,
		Fields:   cfields,
		Queries:  cqueries,
		Opts:     opts,
		plugin:   plugin,
		filename: filename,
	}
}

// GenerateLRU generates the updater for the collection.
func GenerateLRU(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewLRU(plugin, file, msg, opts)
	return c.render()
}

func (c *LRU) render() error {
	steps := []func() error{
		c.definePackage,
		c.defineCollection,
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	outfile := c.File.Outfile()
	original, err := outfile.Content()
	if err != nil {
		return err
	}
	formatted, err := imports.Process(c.filename, original, nil)

	if diff := deep.Equal(original, formatted); diff != nil {
		formattedOutfile := c.plugin.NewGeneratedFile(c.filename, ".")
		if _, err := formattedOutfile.Write(formatted); err != nil {
			return err
		}
		outfile.Skip()
	}

	return nil
}

func (c *LRU) definePackage() error {
	tmplSrc := `// Package {{ .File.CachePackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.CachePackageName }}
`

	tmpl, err := template.New("defineLRUPackage").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, c); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *LRU) defineCollection() error {
	tmplSrc := `// {{ .C.Message.Name }}LRU defines a LRU cache implementing {{ .P.KeyValue }}.{{ .C.Message.Name }}ReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type {{ .C.Message.Name }}LRU struct {
	{{ .P.Service }}.Dependencies
	Nil{{ .C.Message.Name }}Cache

	name string
	reader {{ .P.KeyValue }}.{{ .C.Message.Name }}Reader
	writer {{ .P.KeyValue }}.{{ .C.Message.Name }}Writer
	lru *{{ .P.LRU }}.ARCCache
	all []*{{ .C.Message.QualifiedKind }}
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *{{ .C.Message.Name }}LRU) Initialize(ctx {{ .P.Context }}.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *{{ .C.Message.Name }}LRU) Shutdown(_ {{ .P.Context }}.Context) error {
	return nil
}

// String returns the name of the LRU.
func (x *{{ .C.Message.Name }}LRU) String() string {
	{{- $pkg := .C.File.CachePackageName -}}
	if x.name != "" {
		return "{{ ToDashCase $pkg }}-{{ ToDashCase .C.Message.Name }}-lru-" + x.name
	}
	return "{{ ToDashCase $pkg }}-{{ ToDashCase .C.Message.Name }}-lru"
}

// NameOf returns the name of the LRU.
func (x *{{ .C.Message.Name }}LRU) NameOf() string {
	return x.String()
}

// All implements implements {{ .P.KeyValue }}.{{ .C.Message.Name }}ReadAller.
func (x *{{ .C.Message.Name }}LRU) All(ctx {{ .P.Context }}.Context) ([]*{{ .C.Message.QualifiedKind }}, error) {
	start := {{ .P.Time }}.Now()
	ctx, _ = {{ .P.Tag }}.New(ctx,
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagCollection, "{{ ToSnakeCase .C.Message.Name }}"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagInstance, x.name),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagMethod, "all"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagType, "lru"),
	)
	{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)
		{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureLatency.M(dur), {{ .P.Cache }}.MeasureInflight.M(-1))
	}()

	return x.all, nil
}

// GetByKey implements {{ .P.KeyValue }}.{{ .C.Message.Name }}Reader.
func (x *{{ .C.Message.Name }}LRU) GetByKey(ctx {{ .P.Context }}.Context, key {{ .P.KeyValue }}.{{ .C.Message.Name }}Key) (*{{ .C.Message.QualifiedKind }}, error) {
	start := {{ .P.Time }}.Now()
	ctx, _ = {{ .P.Tag }}.New(ctx,
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagCollection, "{{ ToSnakeCase .C.Message.Name }}"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagInstance, x.name),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagMethod, "get"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagType, "lru"),
	)
	{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)
		{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureLatency.M(dur), {{ .P.Cache }}.MeasureInflight.M(-1))
	}()
	
	if val, ok := x.lru.Get(key); ok {
		{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureHit.M(1))
		return val.(*{{ .C.Message.QualifiedKind }}), nil
	}
	{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureMiss.M(1))

	if x.reader != nil {
		val, err := x.reader.GetByKey(ctx, key)
		if err != nil {
			return nil, {{ .P.Fmt }}.Errorf("lru: {{ .C.Message.Name }}.GetByKey - %w", err)
		}
		x.lru.Add(key, val)
		return val, nil
	}

	{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureError.M(1))
	return nil, {{ .P.Fmt }}.Errorf("lru: {{ .C.Message.Name }}.GetByKey - %w", {{ .P.Cache }}.ErrGetValue)
}

// SetByKey implements {{ .P.KeyValue }}.{{ .C.Message.Name }}Writer.
func (x *{{ .C.Message.Name }}LRU) SetByKey(ctx {{ .P.Context }}.Context, key {{ .P.KeyValue }}.{{ .C.Message.Name }}Key, val *{{ .C.Message.QualifiedKind }}) (*{{ .C.Message.QualifiedKind }}, error) {
	start := {{ .P.Time }}.Now()
	ctx, _ = {{ .P.Tag }}.New(ctx,
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagCollection, "{{ ToSnakeCase .C.Message.Name }}"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagInstance, x.name),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagMethod, "set"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagType, "lru"),
	)
	{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)
		{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureLatency.M(dur), {{ .P.Cache }}.MeasureInflight.M(-1))
	}()

	if x.writer != nil {
		upd, err := x.writer.SetByKey(ctx, key, val)
		if err != nil {
			{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureError.M(1))
			return nil, {{ .P.Fmt }}.Errorf("lru: {{ .C.Message.Name }}.SetBykey - %w", err)
		}
		val = upd
	}

	x.lru.Add(key, val)

	all := []*{{ .C.Message.QualifiedKind }}{}
	for _, k := range x.lru.Keys() {
		y, _ := x.lru.Get(k)
		all = append(all, y.(*{{ .C.Message.QualifiedKind }}))
	}
	x.all = all

	return val, nil
}

// WithReader tells the {{ .C.Message.Name }}LRU where to source values from if they don't exist in cache.
func (x *{{ .C.Message.Name }}LRU) WithReader(r {{ .P.KeyValue }}.{{ .C.Message.Name }}Reader) {
	x.reader = r
}

// WithWriter tells the {{ .C.Message.Name }}LRU where to source values from if they don't exist in cache.
func (x *{{ .C.Message.Name }}LRU) WithWriter(w {{ .P.KeyValue }}.{{ .C.Message.Name }}Writer) {
	x.writer = w
}

// New{{ .C.Message.Name }}LRU returns a new {{ .C.Message.Name }}LRU cache.
func New{{ .C.Message.Name }}LRU(name string, i int) (*{{ .C.Message.Name }}LRU, error) {
	arc, err := {{ .P.LRU }}.NewARC(i)
	if err != nil {
		return nil, {{ .P.Fmt }}.Errorf("lru: arc init - %w", err)
	}

	return &{{ .C.Message.Name }}LRU{
		name: name,
		lru: arc,
	}, nil
}
`

	tmpl, err := template.New("defineLRU").
		Funcs(template.FuncMap{
			"ToDashCase":  protocgenlib.ToDashCase,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Cache":    c.File.QualifiedPackageName("github.com/rleszilm/genms/cache"),
		"Context":  c.File.QualifiedPackageName("context"),
		"Fmt":      c.File.QualifiedPackageName("fmt"),
		"KeyValue": c.File.QualifiedPackageName(path.Join(c.File.DalPackagePath(), "keyvalue")),
		"LRU":      c.File.QualifiedPackageName("github.com/hashicorp/golang-lru"),
		"Log":      c.File.QualifiedPackageName("github.com/rleszilm/genms/log"),
		"Service":  c.File.QualifiedPackageName("github.com/rleszilm/genms/service"),
		"Stats":    c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"Tag":      c.File.QualifiedPackageName("go.opencensus.io/tag"),
		"Time":     c.File.QualifiedPackageName("time"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
