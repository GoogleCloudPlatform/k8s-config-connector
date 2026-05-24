### 2026-05-24 Implementing Direct Controller and E2E Tests for AppHubDiscoveredService
- **Context**: Implementing the Phase 2 (Controller and E2E Tests) for `AppHubDiscoveredService` resource.
- **Problem**: 
  1. `AppHubDiscoveredService` is a read-only ("discovered") resource in Google Cloud, meaning there is no manual GCP API to `Create` or `Delete` it.
  2. The Google Cloud REST Go library sends request paths using `/discoveredservices/` (all lowercase), whereas MockGCP's gRPC-gateway was strictly expecting the canonical path `/discoveredServices/` (capital `S`), causing a 404 router mismatch.
- **Solution**:
  1. Implemented on-demand "auto-discovery" of services in the mock `GetDiscoveredService` method: if a requested discovered service is not found, MockGCP automatically creates and returns a mock discovered service.
  2. Added a path rewrite in MockGCP HTTP RoundTripper to dynamically redirect any incoming request containing `/discoveredservices/` to the canonical `/discoveredServices/` path before routing.
- **Impact**: All read-only AppHub discovered services are now automatically and correctly mocked, routed, and fully testable in the KCC E2E suite hermetically.
