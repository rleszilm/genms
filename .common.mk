## Build config
BUILD_ARGS ?= CGO_ENABLED=0
BUILD_FLAGS ?= -a -ldflags "-s"

BUILD_MODULES ?= `go list ./... | grep -v tools`

## Testing config
ifndef TESTS
TESTS=$(BUILD_MODULES)
endif

ifndef TEST_OPTS
TEST_OPTS=-race
else
TEST_OPTS += -race
endif
