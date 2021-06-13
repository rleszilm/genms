// Package rest_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package rest_dal_multi

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	ioutil "io/ioutil"
	http "net/http"
	url "net/url"
	strconv "strconv"
	template "text/template"
	time "time"

	copier "github.com/jinzhu/copier"
	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	dal "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal"
	rest "github.com/rleszilm/genms/rest"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// TypeTwoCollection is an autogenerated implementation of dal.TypeTwoCollection.
type TypeTwoCollection struct {
	dal.UnimplementedTypeTwoCollection

	name string

	client *http.Client
	config *TypeTwoConfig

	url                     *url.URL
	urlAll                  string
	urlTmplOneParam         *template.Template
	urlTmplMultipleParam    *template.Template
	urlTmplMessageParam     *template.Template
	urlTmplWithComparator   *template.Template
	urlTmplWithRest         *template.Template
	urlTmplProviderStubOnly *template.Template
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *TypeTwoCollection) Initialize(_ context.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *TypeTwoCollection) Shutdown(_ context.Context) error {
	return nil
}

// String returns the name of the Collection.
func (x *TypeTwoCollection) String() string {
	if x.name != "" {
		return "rest-dal-multi-type-two-" + x.name
	}
	return "rest-dal-multi-type-two"
}

// NameOf returns the name of the Collection.
func (x *TypeTwoCollection) NameOf() string {
	return x.String()
}

// DoReq executes the given http request.
func (x *TypeTwoCollection) DoReq(ctx context.Context, label string, req *http.Request) ([]*multi.TypeTwo, error) {
	var err error
	var resp *http.Response

	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(rest.TagCollection, "type_two"),
		tag.Upsert(rest.TagInstance, x.name),
		tag.Upsert(rest.TagMethod, label),
		tag.Upsert(rest.TagRestMethod, req.Method),
	)
	stats.Record(ctx, rest.MeasureInflight.M(1))
	defer func(ctx context.Context) {
		if resp != nil {
			ctx, _ = tag.New(ctx,
				tag.Upsert(rest.TagResponseCode, strconv.Itoa(resp.StatusCode)),
			)
		}

		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, rest.MeasureLatency.M(dur), rest.MeasureInflight.M(-1))
	}(ctx)

	ctx, cancel := context.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	resp, err = x.client.Do(req.WithContext(ctx))
	if err != nil {
		rest.Logs().Error("could not execute rest request:", err)
		stats.Record(ctx, rest.MeasureError.M(1))
		return nil, err
	}

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rest.Logs().Error("could not read rest response:", err)
		stats.Record(ctx, rest.MeasureError.M(1))
		return nil, err
	}

	TypeTwoScanners := []*TypeTwoScanner{}
	if err := json.Unmarshal(buff, &TypeTwoScanners); err != nil {
		rest.Logs().Error("could not unmarshal rest response:", err)
		stats.Record(ctx, rest.MeasureError.M(1))
		return nil, err
	}

	TypeTwos := []*multi.TypeTwo{}
	for _, c := range TypeTwoScanners {
		TypeTwos = append(TypeTwos, c.TypeTwo())
	}
	return TypeTwos, nil
}

// All implements dal.TypeTwoCollection.All
func (x *TypeTwoCollection) All(ctx context.Context) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)
	u.Path = x.urlAll

	req := &http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    u,
	}

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	return x.DoReq(ctx, "all", req)
}

// OneParam implements dal.TypeTwoCollection.OneParam
func (x *TypeTwoCollection) OneParam(ctx context.Context, scalar_int32 int32) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)

	req := &http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    u,
	}

	queryValues := url.Values{}
	queryValues.Add("scalar_int32", fmt.Sprintf("%v", scalar_int32))

	req.URL.RawQuery = queryValues.Encode()

	pathValues := map[string]interface{}{}
	pathBuf := &bytes.Buffer{}
	if err := x.urlTmplOneParam.Execute(pathBuf, pathValues); err != nil {
		return nil, err
	}
	req.URL.Path = pathBuf.String()

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	return x.DoReq(ctx, "one_param", req)
}

// MultipleParam implements dal.TypeTwoCollection.MultipleParam
func (x *TypeTwoCollection) MultipleParam(ctx context.Context, scalar_int32 int32, scalar_int64 int64, scalar_float32 float32) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)

	req := &http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    u,
	}

	queryValues := url.Values{}
	queryValues.Add("scalar_int32", fmt.Sprintf("%v", scalar_int32))
	queryValues.Add("scalar_int64", fmt.Sprintf("%v", scalar_int64))
	queryValues.Add("scalar_float32", fmt.Sprintf("%v", scalar_float32))

	req.URL.RawQuery = queryValues.Encode()

	pathValues := map[string]interface{}{}
	pathBuf := &bytes.Buffer{}
	if err := x.urlTmplMultipleParam.Execute(pathBuf, pathValues); err != nil {
		return nil, err
	}
	req.URL.Path = pathBuf.String()

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	return x.DoReq(ctx, "multiple_param", req)
}

// MessageParam implements dal.TypeTwoCollection.MessageParam
func (x *TypeTwoCollection) MessageParam(ctx context.Context, obj_message *multi.TypeTwo_Message) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)

	req := &http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    u,
	}

	queryValues := url.Values{}
	queryValues.Add("obj_message", fmt.Sprintf("%v", obj_message))

	req.URL.RawQuery = queryValues.Encode()

	pathValues := map[string]interface{}{}
	pathBuf := &bytes.Buffer{}
	if err := x.urlTmplMessageParam.Execute(pathBuf, pathValues); err != nil {
		return nil, err
	}
	req.URL.Path = pathBuf.String()

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	return x.DoReq(ctx, "message_param", req)
}

// WithComparator implements dal.TypeTwoCollection.WithComparator
func (x *TypeTwoCollection) WithComparator(ctx context.Context, scalar_int32 int32) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)

	req := &http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    u,
	}

	queryValues := url.Values{}
	queryValues.Add("scalar_int32", fmt.Sprintf("%v", scalar_int32))

	req.URL.RawQuery = queryValues.Encode()

	pathValues := map[string]interface{}{}
	pathBuf := &bytes.Buffer{}
	if err := x.urlTmplWithComparator.Execute(pathBuf, pathValues); err != nil {
		return nil, err
	}
	req.URL.Path = pathBuf.String()

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	return x.DoReq(ctx, "with_comparator", req)
}

// WithRest implements dal.TypeTwoCollection.WithRest
func (x *TypeTwoCollection) WithRest(ctx context.Context, scalar_int32 int32, scalar_int64 int64, scalar_float32 float32, scalar_float64 float64) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)

	req := &http.Request{
		Method: "POST",
		Header: http.Header{},
		URL:    u,
	}

	queryValues := url.Values{}
	queryValues.Add("query_rest_scalar_int32", fmt.Sprintf("%v", scalar_int32))

	req.URL.RawQuery = queryValues.Encode()

	pathValues := map[string]interface{}{"scalar_int64": &scalar_int64}
	pathBuf := &bytes.Buffer{}
	if err := x.urlTmplWithRest.Execute(pathBuf, pathValues); err != nil {
		return nil, err
	}
	req.URL.Path = pathBuf.String()

	bodyValues := map[string]interface{}{"scalar_float32": &scalar_float32}
	bodyBytes, err := json.Marshal(bodyValues)
	if err != nil {
		return nil, err
	}
	bodyRC := ioutil.NopCloser(bytes.NewReader(bodyBytes))
	req.Body = bodyRC

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	req.Header.Add("scalar_float64", fmt.Sprintf("%v", scalar_float64))

	return x.DoReq(ctx, "with_rest", req)
}

// ProviderStubOnly implements dal.TypeTwoCollection.ProviderStubOnly
func (x *TypeTwoCollection) ProviderStubOnly(ctx context.Context) ([]*multi.TypeTwo, error) {
	u := &url.URL{}
	copier.Copy(u, x.url)

	req := &http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    u,
	}

	queryValues := url.Values{}

	req.URL.RawQuery = queryValues.Encode()

	pathValues := map[string]interface{}{}
	pathBuf := &bytes.Buffer{}
	if err := x.urlTmplProviderStubOnly.Execute(pathBuf, pathValues); err != nil {
		return nil, err
	}
	req.URL.Path = pathBuf.String()

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}

	return x.DoReq(ctx, "provider_stub_only", req)
}

// NewTypeTwoCollection returns a new TypeTwoCollection.
func NewTypeTwoCollection(instance string, client *http.Client, urls TypeTwoUrlTemplateProvider, config *TypeTwoConfig) (*TypeTwoCollection, error) {
	coll := &TypeTwoCollection{
		name:   instance,
		client: client,
		config: config,
	}

	u, err := url.Parse(config.URL)
	if err != nil {
		return nil, err
	}
	coll.url = u

	coll.urlAll = urls.All()
	if urls.OneParam() != "" {
		urlTmplOneParam, err := template.New("urlTmplOneParam").
			Funcs(template.FuncMap{}).
			Parse(urls.OneParam())
		if err != nil {
			return nil, err
		}
		coll.urlTmplOneParam = urlTmplOneParam
	}

	if urls.MultipleParam() != "" {
		urlTmplMultipleParam, err := template.New("urlTmplMultipleParam").
			Funcs(template.FuncMap{}).
			Parse(urls.MultipleParam())
		if err != nil {
			return nil, err
		}
		coll.urlTmplMultipleParam = urlTmplMultipleParam
	}

	if urls.MessageParam() != "" {
		urlTmplMessageParam, err := template.New("urlTmplMessageParam").
			Funcs(template.FuncMap{}).
			Parse(urls.MessageParam())
		if err != nil {
			return nil, err
		}
		coll.urlTmplMessageParam = urlTmplMessageParam
	}

	if urls.WithComparator() != "" {
		urlTmplWithComparator, err := template.New("urlTmplWithComparator").
			Funcs(template.FuncMap{}).
			Parse(urls.WithComparator())
		if err != nil {
			return nil, err
		}
		coll.urlTmplWithComparator = urlTmplWithComparator
	}

	if urls.WithRest() != "" {
		urlTmplWithRest, err := template.New("urlTmplWithRest").
			Funcs(template.FuncMap{}).
			Parse(urls.WithRest())
		if err != nil {
			return nil, err
		}
		coll.urlTmplWithRest = urlTmplWithRest
	}

	if urls.ProviderStubOnly() != "" {
		urlTmplProviderStubOnly, err := template.New("urlTmplProviderStubOnly").
			Funcs(template.FuncMap{}).
			Parse(urls.ProviderStubOnly())
		if err != nil {
			return nil, err
		}
		coll.urlTmplProviderStubOnly = urlTmplProviderStubOnly
	}

	return coll, nil
}

// TypeTwoScanner is an autogenerated struct that
// is used to parse query results.
type TypeTwoScanner struct {
	ScalarInt32   int32                  `json:"scalar_int32"`
	ScalarInt64   int64                  `json:"scalar_int64"`
	ScalarFloat32 float32                `json:"scalar_float32"`
	ScalarFloat64 float64                `json:"scalar_float64"`
	ScalarString  string                 `json:"scalar_string"`
	ScalarBool    bool                   `json:"scalar_bool"`
	ScalarEnum    multi.TypeTwo_Enum     `json:"scalar_enum"`
	ObjMessage    *multi.TypeTwo_Message `json:"obj_message"`

	Renamed         string `json:"aliased"`
	IgnoredPostgres string `json:"ignored_postgres"`
	RenamedPostgres string `json:"renamed_postgres"`

	RenamedRest string `json:"aliased_rest"`
}

// TypeTwo returns a new multi.TypeTwo populated with scanned values.
func (x *TypeTwoScanner) TypeTwo() *multi.TypeTwo {
	y := &multi.TypeTwo{}

	y.ScalarInt32 = x.ScalarInt32
	y.ScalarInt64 = x.ScalarInt64
	y.ScalarFloat32 = x.ScalarFloat32
	y.ScalarFloat64 = x.ScalarFloat64
	y.ScalarString = x.ScalarString
	y.ScalarBool = x.ScalarBool
	y.ScalarEnum = x.ScalarEnum
	y.ObjMessage = x.ObjMessage

	y.Renamed = x.Renamed
	y.IgnoredPostgres = x.IgnoredPostgres
	y.RenamedPostgres = x.RenamedPostgres

	y.RenamedRest = x.RenamedRest
	return y
}

// TypeTwoConfig is a struct that can be used to configure a TypeTwoCollection
type TypeTwoConfig struct {
	URL     string            `envconfig:"url"`
	Name    string            `envconfig:"name"`
	Timeout time.Duration     `envconfig:"timeout" default:"5s"`
	Headers map[string]string `envconfig:"headers"`
}

// TypeTwoUrlTemplateProvider is an interface that returns the query templated that should be executed
// to generate the queries that the collection will use.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeTwoUrlTemplateProvider
type TypeTwoUrlTemplateProvider interface {
	All() string
	OneParam() string
	MultipleParam() string
	MessageParam() string
	WithComparator() string
	WithRest() string
	ProviderStubOnly() string
}
