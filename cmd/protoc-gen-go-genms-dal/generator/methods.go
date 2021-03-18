package generator

import (
	"fmt"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
)

// QueryStructField returns the field name that holds the given query.
func QueryStructField(query *annotations.DalOptions_Query) string {
	switch query.Mode {
	case annotations.DalOptions_Query_QueryMode_Auto, annotations.DalOptions_Query_QueryMode_ProviderStub:
		return fmt.Sprintf("query%s string", protocgenlib.ToTitleCase(query.Name))
	default:
		return ""
	}
}
