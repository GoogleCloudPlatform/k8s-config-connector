## 2026-05-18 - speech migration
- Moved mockgcp speech to httptogrpc.
- Updated imports to use "cloud.google.com/go/speech/apiv2/speechpb".
- Removed mockgcp-specific speech v2 proto generation from Makefile.
- Deleted generated speech proto files in mockgcp/generated/mockgcp/cloud/speech/v2/.
- Removed RewriteError from NewHTTPMux as it is not supported by httptogrpc.
