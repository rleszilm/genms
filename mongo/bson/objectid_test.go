package bson_test

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/rleszilm/genms/mongo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestToObjectID(t *testing.T) {
	testcases := []struct {
		desc   string
		input  interface{}
		expect bson.ObjectID
		err    bool
	}{
		{
			desc:   "from ObjectID",
			input:  bson.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			expect: bson.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
		},
		{
			desc:   "from primitive.ObjectID",
			input:  primitive.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			expect: bson.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
		},
		{
			desc:   "from [12]byte",
			input:  [12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expect: bson.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
		},
		{
			desc:   "from []byte",
			input:  []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expect: bson.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
		},
		{
			desc:  "from []byte - too short",
			input: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			err:   true,
		},
		{
			desc:  "from []byte - too long",
			input: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 13},
			err:   true,
		},
		{
			desc:   "from string",
			input:  "deadbeefdeadbeefdeadbeef",
			expect: bson.ObjectID([12]byte{222, 173, 190, 239, 222, 173, 190, 239, 222, 173, 190, 239}),
		},
		{
			desc:  "from string - too short",
			input: "deadbeefdeadbeefdeadbee",
			err:   true,
		},
		{
			desc:  "from string - too long",
			input: "deadbeefdeadbeefdeadbeefd",
			err:   true,
		},
		{
			desc:  "from string - non-hexstring",
			input: "xxxxxxxxxxxxxxxxxxxxxxxx",
			err:   true,
		},
		{
			desc:  "from invalid type",
			input: 0,
			err:   true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err := bson.ToObjectID(tc.input)
			if err != nil && tc.err {
				return
			} else if err != nil {
				t.Fatal("error returned when not expected:", err)
			} else if tc.err {
				t.Fatal("error not returned when expected")
			}

			if diff := deep.Equal(actual, tc.expect); diff != nil {
				t.Error("objectid is not as expected:", diff)
			}
		})
	}
}
