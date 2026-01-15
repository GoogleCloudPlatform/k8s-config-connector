A good mockgcp service:

* Has normalization in `normalize.go` and follows the service registration pattern (does not appear in `mockgcp/mock_http_roundtrip.go`)
* Has tests that use `gcloud` (so can be tested "early" without KCC)
* Uses the official google protos (with grpc-gateway for the HTTP bridge)
* Implements the same version as is used by the controller (in particular for Terraform and DCL resources)