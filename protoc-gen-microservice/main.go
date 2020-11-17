package main

import (
	"log"
	"os"

	"io/ioutil"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/rleszilm/gen_microservice/protoc-gen-microservice/generator"
)

func main() {
	var genError error

	resp := &plugin.CodeGeneratorResponse{}
	defer func() {
		// If some error has been occurred in generate process,
		// add error message to plugin response
		if genError != nil {
			message := genError.Error()
			resp.Error = &message
		}
		buf, err := proto.Marshal(resp)
		if err != nil {
			log.Fatalln(err)
		}

		os.Stdout.Write(buf)
	}()

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		genError = err
		return
	}

	req := &plugin.CodeGeneratorRequest{}
	if err = proto.Unmarshal(buf, req); err != nil {
		genError = err
		return
	}

	run := generator.NewRunner()
	chunks, err := run.Generate(req)

	if err != nil {
		genError = err
	}

	resp.File = append(resp.File, chunks...)
}
