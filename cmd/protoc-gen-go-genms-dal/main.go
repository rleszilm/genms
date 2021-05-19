package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator/cache"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator/keyvalue"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator/rest"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator/sql/postgres"
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
	if !proto.HasExtension(msgOpts, annotations.E_MessageOptions) {
		return nil
	}
	ext := proto.GetExtension(msgOpts, annotations.E_MessageOptions)
	dalOpts := ext.(*annotations.DalOptions)

	// write interfaces
	if err := generator.GenerateInterface(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	// write cache
	if err := cache.GenerateCache(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	// write keyvalue
	if err := keyvalue.GenerateKeyValue(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	// write map cache
	if err := cache.GenerateMap(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	// write lru cache
	if err := cache.GenerateLRU(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	// write updater
	if err := cache.GenerateUpdater(plugin, file, msg, dalOpts); err != nil {
		return err
	}

	for _, be := range dalOpts.Backends {
		switch be {
		case annotations.DalOptions_Postgres:
			if err := postgres.GenerateCollection(plugin, file, msg, dalOpts); err != nil {
				return err
			}
		case annotations.DalOptions_Rest:
			if err := rest.GenerateCollection(plugin, file, msg, dalOpts); err != nil {
				return err
			}
		}
	}
	return nil
}
