package middleware_grpc

import (
	"strings"
)

func parseMethod(method string) (string, string, string) {
	// /package.service/method
	methodTokens := strings.Split(method, "/")
	packageTokens := strings.Split(methodTokens[1], ".")
	return packageTokens[0], packageTokens[1], methodTokens[2]
}
