# Memcache Tribal Knowledge

### 2026-05-16 MemcacheInstance Implementation
- **Context**: Implementing `mockgcp` for `MemcacheInstance` (PR #8204).
- **Problem**: KRM `memcacheParameters` maps to GCP `parameters` in the underlying API, and some fields like `authorized_network` expect full resource URLs.
- **Solution**: Used `httptogrpc` to bridge the HTTP/JSON traffic to a Go implementation using `cloud.google.com/go/memcache/apiv1beta2/memcachepb`. Populated `MemcacheNodes` dynamically based on `nodeCount` and `zones` to provide plausible discovery endpoints.
- **Impact**: Future agents working on Memcache will know how the KRM fields map to the mock implementation and what to expect in the `status` field (e.g., discovery endpoints).
