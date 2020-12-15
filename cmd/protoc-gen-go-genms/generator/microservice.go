package generator

import (
	"text/template"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	microserviceTemplateSource = `// {{ .ms.name }}ServerService implements {{ .ms.name }}Service
type {{ .ms.name }}ServerService struct {
	service.Deps
	{{ .ms.name }}Server

	grpcServer *grpc_service.Server
	{{ if .ms.enableRest -}} restServer *rest_service.Server {{- end }}
	{{ if .ms.enableGraphQL -}} graphqlServer *graphql_service.Server {{- end }}
}

// Initialize implements service.Service.Initialize
func (s *{{ .ms.name }}ServerService) Initialize(ctx context.Context) error {
	s.grpcServer.WithService(func(server *grpc.Server) {
		Register{{ .ms.name }}Server(server, s)
	})

	{{ if .ms.enableRest -}}
	if err := s.restServer.WithGrpcProxy(ctx, "{{ .ms.name }}", Register{{ .ms.name }}HandlerFromEndpoint); err != nil {
		return err
	}
	{{- end }}
	{{ if .ms.enableGraphQL -}}
	if err := s.graphqlServer.WithGrpcProxy(ctx, "{{ .ms.name }}", Register{{ .ms.name }}GraphqlWithOptions); err != nil {
		return err
	}
	{{- end }}
	return nil
}

// Shutdown implements service.Service.Shutdown
func (s *{{ .ms.name }}ServerService) Shutdown(_ context.Context) error {
	return nil
}

func (s *{{ .ms.name }}ServerService) NameOf() string {
	return "{{ .ms.name }}"
}

func (s *{{ .ms.name }}ServerService) String() string {
	return s.NameOf()
}

// New{{ .ms.name }}ServerService returns a new {{ .ms.name }}ServerService
func New{{ .ms.name }}ServerService(impl {{ .ms.name }}Server, grpcServer *grpc_service.Server {{- if .ms.enableRest }}, restServer *rest_service.Server{{ end }} {{- if .ms.enableGraphQL }}, graphqlServer *graphql_service.Server{{ end }}) *{{ .ms.name }}ServerService {
	server := &{{ .ms.name }}ServerService{
		{{ .ms.name }}Server: impl,
		grpcServer: grpcServer,
		{{ if .ms.enableRest -}}
			restServer: restServer,
		{{- end }}
		{{ if .ms.enableGraphQL -}}
			graphqlServer: graphqlServer,
		{{- end }}
	}

	grpcServer.WithDependencies(server)
	{{ if .ms.enableRest -}}
		restServer.WithDependencies(server)
	{{- end }}
	{{ if .ms.enableGraphQL -}}
		graphqlServer.WithDependencies(server)
	{{- end }}

	return server
}
`
)

var (
	microServiceTemplate = template.Must(
		template.New("microserviceTemplate").
			Parse(microserviceTemplateSource))
)

type microService struct {
	fr      *fileRunner
	reqOpts *options
	opts    *annotations.MicroServiceOptions
	svc     *descriptorpb.ServiceDescriptorProto
}

func (s *microService) FileRunner() *fileRunner {
	return s.fr
}

func (s *microService) FileName() string {
	return filename(s.fr.file, s.reqOpts)
}

func (s *microService) Replacements(defaults map[string]interface{}) map[string]interface{} {
	vals := cloneValues(defaults)

	vals["ms"] = map[string]interface{}{
		"name":          s.svc.GetName(),
		"enableRest":    s.opts.GetRest(),
		"enableGraphQL": s.opts.GetGraphql(),
	}

	return vals
}

func (s *microService) Imports() []string {
	imports := []string{
		`"context"`,
		`"github.com/rleszilm/gen_microservice/service"`,
		`grpc_service "github.com/rleszilm/gen_microservice/service/grpc"`,
		`"google.golang.org/grpc"`,
	}

	if s.opts.GetRest() {
		imports = append(imports, `rest_service "github.com/rleszilm/gen_microservice/service/rest"`)
	}

	if s.opts.GetGraphql() {
		imports = append(imports, `graphql_service "github.com/rleszilm/gen_microservice/service/graphql"`)
	}
	return imports
}

func (s *microService) Constants() map[string]interface{} {
	return nil
}

func (s *microService) Variables() map[string]string {
	return nil
}

func (s *microService) Logic() *template.Template {
	return microServiceTemplate
}

func (s *microService) Outline() *template.Template {
	return nil
}

func newMicroService(fr *fileRunner, svc *descriptorpb.ServiceDescriptorProto, reqOpts *options) *microService {
	ms := &microService{
		fr:      fr,
		svc:     svc,
		reqOpts: reqOpts,
	}

	ext := proto.GetExtension(svc.GetOptions(), annotations.E_MicroService)
	opts, ok := ext.(*annotations.MicroServiceOptions)

	if ok {
		ms.opts = opts
	}

	return ms
}
