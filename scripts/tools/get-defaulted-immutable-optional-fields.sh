#!/bin/bash

go test ./config/tests/servicemapping/... -v -run TestImmutableOptionalFieldsWithDefault > ./defaulted_immutable_optional_fields.txt