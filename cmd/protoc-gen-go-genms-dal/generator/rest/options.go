package rest

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
)

// Options is a wrapper for the rest annotations
type Options struct {
	annotations.DalOptions_Query_Rest
}
