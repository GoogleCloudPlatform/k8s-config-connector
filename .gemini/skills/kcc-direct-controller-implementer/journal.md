### [2026-05-19] Direct Controllers for APIs without GAPIC or google-api-go-client
- **Context**: Implementing `AIStreamsCluster` which only has a gRPC client in `genproto` (no GAPIC, no google-api-go-client).
- **Problem**: Need to make API calls without an existing high-level SDK or REST client, while ensuring KCC's authentication, routing (mockgcp), and LRO waiting work.
- **Solution**: 
  1. Use `config.GRPCClientOptions()` to get the base `grpc.DialOption`s.
  2. Append `option.WithEndpoint("API_ENDPOINT:443")`.
  3. Call `grpc.Dial(ctx, opts...)`. The connection handles auth and mockgcp routing natively.
  4. Use this connection to instantiate the generated proto client (e.g. `pb.NewAIStreamsClient(conn)`) and the long-running operations client (`lro.NewOperationsClient(conn)`).
- **Impact**: Enables KCC direct controllers to support purely gRPC-based proto APIs or new partner APIs that have not yet published standard client libraries.

ApiHubDeployment: Dropped 'Annotations' from spec since proto uses map of string to AttributeValues (wrapper message). Used IdentityV2 standard.
