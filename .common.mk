## Build config
BUILD_ARGS ?= CGO_ENABLED=0
BUILD_FLAGS ?= -a -ldflags "-s"

## Testing config
TEST_MODE ?= unit
TESTS_OPTS ?= -race
TEST_EXCLUDE = tools
TEST_INTEGRATION = mongo

ifeq ($(TEST_MODE),integration)
	TESTS ?= `go list ./... | egrep -v $(TEST_EXCLUDE) | egrep $(TEST_INTEGRATION)`
else
	TESTS ?= `go list ./... | egrep -v $(TEST_EXCLUDE) | egrep -v $(TEST_INTEGRATION)`
endif
