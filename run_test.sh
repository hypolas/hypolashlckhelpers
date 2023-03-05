#!/usr/bin/env bash

# Test simple mode
echo "# Test simple mode"
export HYPOLAS_LOGS_FILE=test/log.log
export HYPOLAS_HEALTHCHECK_HTTP_URL=OK
go test -v .


# Test passing variable in environment variable value
echo "# Test passing variable in environment variable value"
export THISVAR=OK
export HYPOLAS_HEALTHCHECK_HTTP_URL="\$THISVAR"
go test -v .

# Test passing custom ID in environment variable
echo "# Test passing custom ID in environment variable"
export HYPOLAS_HEALTHCHECK_ID="MYID"
export HYPOLAS_HEALTHCHECK_MYID_HTTP_URL=OK
go test -v .
