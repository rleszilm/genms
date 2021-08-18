package protocgenlib

import (
	"log"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func ToGoKind(k protoreflect.Kind, prefix string) string {
	log.Println("togo:", prefix, k.String(), k.GoString())
	switch k.String() {
	case "bool":
		return prefix + "bool"
	case "double":
		return prefix + "float64"
	case "float":
		return prefix + "float32"
	case "int32":
		return prefix + "int32"
	case "int64":
		return prefix + "int64"
	case "string":
		return prefix + "string"
	case "bytes":
		return "[]byte"
	default:
		return prefix + k.GoString()
	}
}
