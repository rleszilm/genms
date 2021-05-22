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

// Updater is a struct that generates an updater file.
type Updater struct {
	File    *File
	Message *generator.Message
	Fields  *generator.Fields
	Queries *generator.Queries
	Opts    *annotations.DalOptions

	plugin   *protogen.Plugin
	filename string
}

// NewUpdater returns a new updater renderer.
func NewUpdater(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *Updater {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/keyvalue/cache/%s.genms.updater.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	cfile := NewFile(outfile, file)
	cmsg := generator.NewMessage(cfile.Generator(), msg)
	cfields := generator.NewFields(cmsg)
	cqueries := generator.NewQueries(cfile.Generator(), cfields, opts)

	return &Updater{
		File:     cfile,
		Message:  cmsg,
		Fields:   cfields,
		Queries:  cqueries,
		Opts:     opts,
		plugin:   plugin,
		filename: filename,
	}
}

// GenerateUpdater generates the updater for the collection.
func GenerateUpdater(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewUpdater(plugin, file, msg, opts)
	return c.render()
}

func (c *Updater) render() error {
	steps := []func() error{
		c.definePackage,
		c.defineUpdater,
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

func (c *Updater) definePackage() error {
	tmplSrc := `// Package {{ .File.CachePackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.CachePackageName }}
`

	tmpl, err := template.New("defineUpdaterPackage").
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

func (c *Updater) defineUpdater() error {
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}

// {{ .C.Message.Name }}Updater is an autogenerated implementation of {{ .C.Message.QualifiedDalKind }}Updater.
type {{ .C.Message.Name }}Updater struct {
	{{ .P.Service }}.Dependencies

	name string
	reader {{ .P.KeyValue }}.{{ .C.Message.Name }}ReadAller
	writer {{ .P.KeyValue }}.{{ .C.Message.Name }}Writer
	key {{ .P.KeyValue }}.{{ .C.Message.Name }}KeyFunc
	interval {{ .P.Time }}.Duration
	done chan struct{}
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *{{ .C.Message.Name }}Updater) Initialize(ctx {{ .P.Context }}.Context) error {
	go x.update(ctx)
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *{{ .C.Message.Name }}Updater) Shutdown(_ {{ .P.Context }}.Context) error {
	return nil
}

// NameOf returns the name of the updater.
func (x *{{ .C.Message.Name }}Updater) NameOf() string {
	return x.name
}

// String returns the name of the updater.
func (x *{{ .C.Message.Name }}Updater) String() string {
	return x.name
}


func (x *{{ .C.Message.Name }}Updater) update(ctx {{ .P.Context }}.Context) {
	ctx, _ = {{ .P.Tag }}.New(ctx,
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagCollection, "{{ ToSnakeCase .C.Message.Name }}"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagInstance, x.name),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagMethod, "update"),
		{{ .P.Tag }}.Upsert({{ .P.Cache }}.TagType, "updater"),
	)

	ticker := {{ .P.Time }}.NewTicker(1)
	for {
		select {
		case <- ctx.Done():
			return
		case <- ticker.C:
			{{ .P.Cache }}.Logs().Infof("starting update for %s", x.name)
			start := {{ .P.Time }}.Now()
			{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureInflight.M(1))
			
			vals, err := x.reader.All(ctx)
			if err != nil {
			}

			for _, val := range vals {
				{{ .P.Cache }}.Logs().Trace("updater {{ .C.Message.Name }} storing value:", x.key(val), val)
				if err = x.writer.SetByKey(ctx, x.key(val), val); err != nil {
					{{ .P.Cache }}.Logs().Error("updater {{ .C.Message.Name }} could not store value:", x.key(val), val, err)
					break
				}
			}
		
			{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureInflight.M(-1))

			if err != nil {
				{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureError.M(1))
			}

			stop := {{ .P.Time }}.Now()
			dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)
			{{ .P.Stats }}.Record(ctx, {{ .P.Cache }}.MeasureLatency.M(dur))

			if x.interval == 0 {
				{{ .P.Cache }}.Logs().Infof("updater %s is terminating", x.name)
				return
			}
			{{ .P.Cache }}.Logs().Infof("scheduling next update for %v", x.interval)
			ticker.Reset(x.interval)
		}
	}
}

// WithReadAller tells the {{ .C.Message.Name }}Map where to source values from if they don't exist in cache.
func (x *{{ .C.Message.Name }}Updater) WithReadAller(r {{ .P.KeyValue }}.{{ .C.Message.Name }}ReadAller) {
	x.reader = r
}

// WithWriter tells the {{ .C.Message.Name }}Map where to source values from if they don't exist in cache.
func (x *{{ .C.Message.Name }}Updater) WithWriter(w {{ .P.KeyValue }}.{{ .C.Message.Name }}Writer) {
	x.writer = w
}

// New{{ .C.Message.Name }}Updater returns a new {{ .C.Message.Name }}Updater.
func New{{ .C.Message.Name }}Updater(name string, k {{ .P.KeyValue }}.{{ .C.Message.Name }}KeyFunc, i {{ .P.Time }}.Duration) *{{ .C.Message.Name }}Updater {
	return &{{ .C.Message.Name }}Updater{
		name: name,
		key: k,
		interval: i,
		done: make(chan struct{}),
	}
}
`

	tmpl, err := template.New("defineUpdater").
		Funcs(template.FuncMap{
			"ToCamelCase": protocgenlib.ToCamelCase,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Cache":    c.File.QualifiedPackageName("github.com/rleszilm/genms/cache"),
		"Context":  c.File.QualifiedPackageName("context"),
		"KeyValue": c.File.QualifiedPackageName(path.Join(c.File.DalPackagePath(), "keyvalue")),
		"Service":  c.File.QualifiedPackageName("github.com/rleszilm/genms/service"),
		"Stats":    c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"Sync":     c.File.QualifiedPackageName("sync"),
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
