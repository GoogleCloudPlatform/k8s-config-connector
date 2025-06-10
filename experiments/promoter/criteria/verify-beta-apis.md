Run `go test ./tests/apichecks/...` to verify the Beta CRDs
    * If hit `[missing_field]` error, comment out that field from "apis/${SERVICE}/v1beta1/", "pkg/controller/direct/{SERVICE}/", and "pkg/test/resourcefixture/testdata/basic/${SERVICE}/" using "/* NOTYET " commenter.
