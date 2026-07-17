### [2026-06-05] gcpurls.Template placeholder naming and struct field matching
- **Context**: Implementing direct types and IdentityV2 for `EventarcChannelConnection` (Issue #9294).
- **Problem**: In `eventarcchannelconnection_identity.go`, initializing the `gcpurls.Template` with the format `projects/{project}/locations/{location}/channelConnections/{channel_connection}` against the `EventarcChannelConnectionIdentity` struct caused an initialization panic:
  ```
  panic: field "channel_connection" not found in struct v1alpha1.EventarcChannelConnectionIdentity
  ```
  This happened because `gcpurls.Template` parses template parameters case-insensitively against the struct fields by lowercasing both. The struct field `ChannelConnection` lowercases to `channelconnection`, but the placeholder `{channel_connection}` has an underscore, resulting in a lookup key `channel_connection`. These two keys do not match.
- **Solution**: Renamed the placeholder parameter in the template format to `{channelconnection}` to align with the lowercase representation of the `ChannelConnection` field:
  ```go
  var EventarcChannelConnectionIdentityFormat = gcpurls.Template[EventarcChannelConnectionIdentity]("eventarc.googleapis.com", "projects/{project}/locations/{location}/channelConnections/{channelconnection}")
  ```
- **Impact**: Future implementers of KCC Direct resources using `gcpurls.Template` must ensure that all template placeholders match their corresponding struct field names case-insensitively (e.g. avoid underscores in the template placeholder unless the struct field itself contains underscores).

### [2026-06-15] Handling underscores in gcpurls.Template placeholders with struct fields
- **Context**: Implementing direct types and IdentityV2 for `EventarcGoogleApiSource` (Issue #10266).
- **Problem**: The GCP resource path segment name has underscores (`googleApiSources/{google_api_source}`). If the struct field is camel-cased to `GoogleApiSource`, `strings.ToLower(f.Name)` evaluates to `googleapisource`, which fails to match `{google_api_source}` casing-insensitive, causing initialization panics.
- **Solution**: Named the struct field with an explicit underscore: `Google_api_source string`. This allows the lowercase of the struct field `google_api_source` to perfectly match the placeholder `{google_api_source}`, resolving the mapping successfully.
- **Impact**: Future developers can safely keep underscores in the GCP template placeholders by declaring fields in Go with matching underscores (e.g., `Google_api_source`), allowing clean alignment with canonical GCP URL formats.

