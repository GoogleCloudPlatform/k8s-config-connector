# Issue 1: Propagate Universe Domain to Terraform Provider Shim & Controller Manager

**Status:** Open  
**Epic:** KCC Support in Sovereign Cloud  
**Components:** `pkg/tf/provider`, `pkg/controller/kccmanager`, `pkg/config`

---

## 1. Problem Description

When KCC runs in an isolated Google Cloud Universe (such as Sovereign Cloud Germany, Sovereign Cloud France, or private air-gapped environments), reconcilers utilizing the embedded Terraform provider fail with:
```
googleapi: Error 401: Request had invalid authentication credentials.
reason: "ACCESS_TOKEN_TYPE_UNSUPPORTED"
service: "pubsub.googleapis.com"
```
Even if `GOOGLE_CLOUD_UNIVERSE_DOMAIN` is set in the controller manager pod environment, the Terraform provider shim ignores it.

In `pkg/tf/provider/provider.go`:
```go
type Config struct {
    GCPAccessToken         string
    Scopes                 []string
    UserProjectOverride    bool
    BillingProject         string
    EnableMetricsTransport bool
    // UniverseDomain is missing!
}
```
And during provider configuration in `New(ctx context.Context, config Config)`:
```go
cfgMap["scopes"] = config.Scopes
cfgMap["user_project_override"] = config.UserProjectOverride
cfgMap["billing_project"] = config.BillingProject
// cfgMap["universe_domain"] is never set!
```

## 2. Root Cause Analysis

The underlying `terraform-provider-google-beta` explicitly supports `universe_domain` via schema attribute and environment variables, but KCC's shim does not expose this parameter. Additionally, `kccmanager.go` does not extract the active universe domain from environment variables or ConfigConnector CR spec.

## 3. Required Implementation

1. **Extend `tfprovider.Config`:**
   Add `UniverseDomain string` to `pkg/tf/provider/provider.go`.
2. **Populate `cfgMap["universe_domain"]`:**
   In `provider.go`, if `config.UniverseDomain != ""`, set:
   ```go
   cfgMap["universe_domain"] = config.UniverseDomain
   ```
   If empty, fallback to `os.Getenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN")`.
3. **Propagate from `kccmanager.go`:**
   In `pkg/controller/kccmanager/kccmanager.go`, extract `universeDomain := os.Getenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN")` and assign `tfCfg.UniverseDomain = universeDomain`.
4. **Extend `ConfigConnector` and `ConfigConnectorContext` Spec:**
   Add optional `universeDomain: string` field to CRD specs so administrators can declare the universe domain declaratively at the cluster or namespace level.

## 4. Acceptance Criteria

- [ ] Terraform provider resolves endpoints as `https://<service>.<universe_domain>/` when `GOOGLE_CLOUD_UNIVERSE_DOMAIN` is set.
- [ ] In the absence of an explicit universe domain, provider defaults to `googleapis.com` without regression.
- [ ] Unit tests verifying `cfgMap["universe_domain"]` serialization in `pkg/tf/provider/provider_test.go`.
