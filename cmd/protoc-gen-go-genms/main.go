package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var plugin *protogen.Plugin
	var genError error

	defer func() {
		if plugin != nil {
			plugin.Error(genError)

			out, err := proto.Marshal(plugin.Response())
			if err != nil {
				genError = err
			} else {
				if _, err := os.Stdout.Write(out); err != nil {
					log.Fatalln(err)
				}
			}
		}

		if genError != nil {
			log.Fatalln(genError)
		}
	}()

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		genError = err
		return
	}

	req := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(buf, req); err != nil {
		genError = err
		return
	}

	opts := protogen.Options{}
	plugin, err = opts.New(req)
	if err != nil {
		genError = err
		return
	}
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, file := range plugin.Files {
		for _, msg := range file.Services {
			if err := generate(plugin, file, msg); err != nil {
				genError = err
				return
			}
		}
	}
}

func generate(plugin *protogen.Plugin, file *protogen.File, svc *protogen.Service) error {
	svcOpts := svc.Desc.Options()
	if !proto.HasExtension(svcOpts, annotations.E_ServiceOptions) {
		return nil
	}
	ext := proto.GetExtension(svcOpts, annotations.E_ServiceOptions)
	opts := ext.(*annotations.ServiceOptions)

	return generator.GenerateMicroService(plugin, file, svc, opts)
}
