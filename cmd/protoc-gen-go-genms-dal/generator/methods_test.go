package generator_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator"
)

func TestTokenize(t *testing.T) {
	testcases := []struct {
		desc   string
		input  string
		expect []string
	}{
		{
			desc:   "non-char seperators",
			input:  "a b  c\td \te-f_g-_h_-i--j__k",
			expect: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"},
		},
		{
			desc:   "capitol seperators",
			input:  "AppleCherryDVD",
			expect: []string{"Apple", "Cherry", "DVD"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := generator.Tokenize(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("Tokens not as expected.\nDiff:\n    %s\n", strings.Join(deep.Equal(tc.expect, actual), "\n    "))
			}
		})
	}
}
