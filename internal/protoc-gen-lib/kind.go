package protocgenlib

import "google.golang.org/protobuf/reflect/protoreflect"

func ToGoKind(k protoreflect.Kind) string {
	switch k.String() {
	case "bool":
		return "bool"
	case "double":
		return "float64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "string":
		return "string"
	default:
		return k.GoString()
	}
}
