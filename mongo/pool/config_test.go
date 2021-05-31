package pool_test

import (
	"os"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/rleszilm/genms/mongo/pool"
)

func TestNewFromEnv(t *testing.T) {
	testcases := []struct {
		desc   string
		input  map[string]string
		expect *pool.Config
		err    bool
	}{
		{
			desc:  "defaults only",
			input: map[string]string{},
			expect: &pool.Config{
				URI:             "mongodb://localhost:27017",
				AppName:         "",
				MaxPoolSize:     25,
				MaxConnIdleTime: 30 * time.Second,
				Database:        "vvv-repl",
				Timeout:         5 * time.Second,
				ReadPref:        "primarypreferred",
			},
		},
		{
			desc: "all values",
			input: map[string]string{
				"TEST_URI":            "mongodb://test-uri",
				"TEST_APPNAME":        "test-app",
				"TEST_POOL_SIZE":      "100",
				"TEST_CONN_IDLE_TIME": "100s",
				"TEST_DATABASE":       "vvv-test-repl",
				"TEST_TIMEOUT":        "50s",
				"TEST_READ_PREF":      "primary",
			},
			expect: &pool.Config{
				URI:             "mongodb://test-uri",
				AppName:         "test-app",
				MaxPoolSize:     100,
				MaxConnIdleTime: 100 * time.Second,
				Database:        "vvv-test-repl",
				Timeout:         50 * time.Second,
				ReadPref:        "primary",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tc.input {
				os.Setenv(k, v)
			}

			actual, err := pool.NewFromEnv("test")
			if err != nil {
				if tc.err {
					return
				}
				t.Error("Unexpected error returned:", err)
				return
			}

			if diff := deep.Equal(actual, tc.expect); diff != nil {
				t.Error("Parsed values are not as expected:", diff)
			}
		})
	}
}
