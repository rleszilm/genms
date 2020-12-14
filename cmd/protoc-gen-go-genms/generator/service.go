package generator

import (
	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
)

type serviceRunner struct {
	fr   *fileRunner
	svc  *descriptorpb.ServiceDescriptorProto
	opts *options
}

func (r *serviceRunner) Generate(defaults map[string]interface{}) ([]*plugin_go.CodeGeneratorResponse_File, error) {
	resp := []*plugin_go.CodeGeneratorResponse_File{}
	if proto.HasExtension(r.svc.GetOptions(), annotations.E_MicroService) {
		ms := newMicroService(r.fr, r.svc, r.opts)

		chunks, err := generate(ms, defaults)
		if err != nil {
			return nil, err
		}
		resp = append(resp, chunks...)

	}

	return resp, nil
}

func newServiceRunner(fr *fileRunner, svc *descriptorpb.ServiceDescriptorProto, opts *options) *serviceRunner {
	return &serviceRunner{
		fr:   fr,
		svc:  svc,
		opts: opts,
	}
}
