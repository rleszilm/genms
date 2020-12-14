package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator/sql/postgres"
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
					fmt.Println("could not write output:", err)
				}
			}
		}

		if genError != nil {
			fmt.Println("generror", genError)
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

	for _, file := range plugin.Files {
		for _, msg := range file.Messages {
			if err := generate(plugin, file, msg); err != nil {
				genError = err
				return
			}
		}
	}
}

func generate(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message) error {
	msgOpts := msg.Desc.Options()
	if !proto.HasExtension(msgOpts, annotations.E_GenmsDal) {
		return nil
	}
	ext := proto.GetExtension(msgOpts, annotations.E_GenmsDal)
	dalOpts := ext.(*annotations.DalOptions)

	// if no database type is enabled return
	if !dalOpts.Sql.Enabled {
		return nil
	}

	// write interfaces
	if err := generator.GenerateInterface(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	if dalOpts.Sql.Enabled {
		switch dalOpts.Sql.Variant {
		case annotations.DalOptions_SQL_SQLVariant_Postgres:
			if err := postgres.GenerateCollection(plugin, file, msg, dalOpts); err != nil {
				fmt.Println("count not generate collection", err)
				return err
			}
		default:
			return fmt.Errorf("%v is not a supported sql variant", dalOpts.Sql.Variant)
		}
	}

	return nil
}
