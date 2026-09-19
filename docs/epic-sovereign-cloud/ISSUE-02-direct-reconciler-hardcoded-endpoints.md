# Issue 2: Eliminate Hardcoded googleapis.com Endpoints Across Direct Reconcilers

**Status:** Open  
**Epic:** KCC Support in Sovereign Cloud  
**Components:** `pkg/controller/direct/*`, `pkg/gcp`

---

## 1. Problem Description

Multiple direct reconcilers implement client constructors by appending literal string endpoints:
- `pkg/controller/direct/grafeas/grafeasnote_controller.go`:
  `opts = append(opts, option.WithEndpoint("containeranalysis.googleapis.com:443"))`
- `pkg/controller/direct/networksecurity/*`:
  `opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))`
- `pkg/controller/direct/vectorsearch/vectorsearchcollection_controller.go`:
  `opts = append(opts, option.WithEndpoint("https://vectorsearch.googleapis.com"))`
- `pkg/controller/direct/assuredworkloads/assuredworkloadsworkload_controller.go`:
  `opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-assuredworkloads.googleapis.com:443", location)))`
- `pkg/controller/direct/aiplatform/*`:
  `opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))`

When running in sovereign or isolated universes, requests to these endpoints fail to resolve or reject tokens with authentication errors.

## 2. Root Cause Analysis

Controllers construct `option.WithEndpoint(...)` with explicit `.googleapis.com` strings rather than consulting the universe domain context.

## 3. Required Implementation

1. **Create Universe Endpoint Resolver Utility in `pkg/gcp/universe.go`:**
   ```go
   func GetUniverseDomain() string {
       if val := os.Getenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN"); val != "" {
           return val
       }
       return "googleapis.com"
   }

   func FormatEndpoint(service string, location string) string {
       domain := GetUniverseDomain()
       if location != "" {
           return fmt.Sprintf("%s-%s.%s:443", location, service, domain)
       }
       return fmt.Sprintf("%s.%s:443", service, domain)
   }
   ```
2. **Refactor Direct Controllers:**
   Replace literal strings with `gcp.FormatEndpoint("containeranalysis", "")`, `gcp.FormatEndpoint("networksecurity", "")`, etc.
3. **Respect `option.WithUniverseDomain`:**
   Pass `option.WithUniverseDomain(gcp.GetUniverseDomain())` into all client option arrays where supported by Google Cloud Go client libraries.

## 4. Acceptance Criteria

- [ ] Zero literal `*.googleapis.com` references in direct controller client initialization logic.
- [ ] Direct controllers correctly target `<service>.<universe_domain>` when `GOOGLE_CLOUD_UNIVERSE_DOMAIN` is set.
