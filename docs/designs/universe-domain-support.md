# Technical Design: Universe Domain Support (Google Cloud Dedicated / Sovereign Clouds)

Status: Proposal
Tracking issue: [#5995](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5995)

## 1. Overview

Google Cloud Dedicated (GDC), Trusted Partner Cloud and sovereign offerings such as
S3NS ("Cloud de Confiance") run in a **universe** other than the public
`googleapis.com` one.

Config Connector has no notion of a universe today. `grep -ri universe` returns no
hits outside `third_party/`. The controller deploys successfully against a sovereign
cluster (it only talks to the Kubernetes API server to get that far) and then fails on
every reconcile, because every GCP client dials `*.googleapis.com`.

The design targets **any** universe. It introduces two explicitly configured, independent
values — a universe **domain** and a universe **prefix** — neither derived from the other or
from the credential (§5).

The work separates into two workstreams of very different size:

*   **Endpoints** (§6) — the domain flowing from the `ConfigConnector` CRD to all three GCP
    client stacks. Well-bounded, mechanical, high value.
*   **Project-ID prefixing** (§7) — in a universe, project IDs carry the prefix
    (`<prefix>:my-project`). This breaks parsing code in ways that are diffuse and much
    harder to enumerate.

Both are required for a working universe deployment. The proposal is to land **all of the
endpoint work plus the configuration surface for both values** near-term, and defer only the
project-ID *parsing* fixes — which then arrive as small self-contained PRs against a settled
configuration surface.

## 2. Goals

*   Let an operator declare the universe once, on the `ConfigConnector` object, and have
    every controller honour it.
*   Work for **any** universe, not a particular one. No implementation may hardcode a domain
    or prefix, or infer either from the other.
*   Cover the **direct controller** stack (215 kinds) completely. Per the migration tracker
    ([#10588](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/10588)) this
    is where essentially all resources are heading, so it is where universe support has the
    longest useful life.
*   Cover the **Terraform-based** stack (234 kinds — the largest today, 172 of them with no
    direct controller available).
*   Make the endpoint-override sites in direct controllers universe-aware, and make it hard
    to add a non-universe-aware one in the future.
*   Land the configuration surface for the universe prefix, so the project-ID parsing fixes
    that follow have an agreed source of truth.
*   Change nothing when no universe is configured. The public-universe code path must be
    byte-for-byte what it is today.

## 3. Non-Goals

*   Supporting more than one universe within a single KCC installation.
*   Any change to resource schemas or KRM semantics. A universe changes *where* the API is,
    not *what* it accepts.
*   Rewriting service names — see §5.
*   Fixing project-ID *parsing* in the first iterations. §7 scopes it; only the configuration
    input lands early.
*   Full DCL coverage in the first iteration. See §6.6.

## 4. Prior art

This design follows [GoogleCloudDataproc/hadoop-connectors#1752](https://github.com/GoogleCloudDataproc/hadoop-connectors/pull/1752)
(merged 2026-07-28), which added universe support to the GCS Hadoop connector and was
validated against a live S3NS universe. Three things carry over directly:

1.  **Precedence**: explicit configuration > `GOOGLE_CLOUD_UNIVERSE_DOMAIN` env var >
    `googleapis.com`.
2.  **An explicit endpoint override still wins** over the universe-derived endpoint.
3.  **A dedicated universe knob is worth having even though a raw endpoint override
    exists.** The reviewer there initially asked why the existing root-URL property was
    not sufficient. The answer that settled it applies verbatim to KCC and is recorded
    in §6.7.

## 5. Background: what actually differs in a universe

Three things differ. The design targets **any** universe; S3NS appears below only as a
worked example with concrete values, and nothing in the implementation may assume its
particular domain or prefix.

| | Public universe | Any universe | S3NS (example) |
| :--- | :--- | :--- | :--- |
| API **endpoints** | `<service>.googleapis.com` | `<service>.<domain>` | `<service>.s3nsapis.fr` |
| **Project IDs** | `my-project` | `<prefix>:my-project` | `s3ns:my-project` |
| **Service-agent emails** | `...@<agent>.iam.gserviceaccount.com` | `...@<agent>.<prefix>-system.iam.gserviceaccount.com` | — |

Two independent values are in play and **must not be conflated or derived from one another**:

*   the **universe domain** (`s3nsapis.fr`) — an API host suffix;
*   the **universe prefix** (`s3ns`) — a project-ID and service-agent qualifier.

There is no rule relating them. `s3nsapis.fr` and `s3ns` happen to share a stem; another
universe's pair may not. Both are therefore configured explicitly (§6.2), and neither is ever
inferred from the other or from the credential. Any code that computes one from the other is
a bug, and the two-universe test matrix in §9 exists specifically to catch it.

**Services absent from a universe.** Not every GCP service exists in every universe, so some
CRDs will have no backing API. KCC already has the mechanism for this and needs no new work:
`ConfigConnector.spec.experiments.resourceSettings` supports include/exclude by group and
kind. This should be documented for universe users rather than reimplemented.

**Service names do not change.** `<service>.googleapis.com` remains the service's
identifier in every universe — which is why enabling an API in a sovereign universe uses
exactly the same name as in public GCP. This is confirmed, and it is what makes §8 safe.

**Service-agent emails are not a KCC code problem.** A `grep` for `gserviceaccount.com`
across `pkg/`, `apis/` and `operator/` finds only documentation comments in generated CRD
field descriptions — KCC never constructs a service-agent email in Go, and performs no
validation on IAM member strings. Users supply those addresses in their own YAML
(`IAMPolicyMember.spec.member` and friends), so they are already free to write
universe-correct values. Worth documenting for users; no code change.

## 6. Workstream A: endpoints

### 6.1 Current state

KCC has three GCP client stacks, each configuring endpoints differently.

| Stack | Configured in | Kinds served today | Difficulty |
| :--- | :--- | :--- | :--- |
| Direct controllers | `pkg/config/controllerconfig.go` | **215** (173 files use `RESTClientOptions()`, 44 use `GRPCClientOptions()`) | Low — one chokepoint |
| Terraform-based | `pkg/tf/provider/provider.go` | **234** (172 with no direct controller available) | Low — the vendored provider already supports it |
| DCL-based | `pkg/dcl/clientconfig/config.go` | **65** (51 with no direct controller available) | High — see §6.6 |

Counts are from `pkg/controller/resourceconfig/static_config.go`, which is the authoritative
per-GVK `DefaultController` table (518 kinds total, plus 4 IAM special cases).

Note what this says about sequencing: **57% of kinds are not served by a direct controller
today.** Fixing only `ControllerConfig` would leave the larger half of the resource surface
broken in a universe. The "Terraform" stack here is not Terraform the tool — KCC vendors
`terraform-provider-google-beta`'s Go packages under `third_party/` and calls the provider's
CRUD functions in-process as a library. There is no `terraform` binary, no HCL, and no state
file anywhere in the reconcile path. It is simply the second of three code paths that build
GCP API URLs, and currently the biggest one.

#### Direct controllers

`ControllerConfig.RESTClientOptions()` and `GRPCClientOptions()` are the single funnel
through which almost every direct controller builds its client options. Both already
carry the exact TODO this design closes:

```go
// pkg/config/controllerconfig.go:146 and :169
// TODO: support endpoints?
// if m.config.Endpoint != "" {
//     opts = append(opts, option.WithEndpoint(m.config.Endpoint))
// }
```

`google.golang.org/api` (pinned at `v0.287.1`) exposes `option.WithUniverseDomain`, and
generated clients resolve their default endpoint from a universe-domain template
(`internal/cba.go:77`). `DialSettings.Validate()` declares no incompatibility between
`WithUniverseDomain` and `WithHTTPClient`, which matters because mockgcp intercepts
traffic with a `RoundTripper` on `HTTPClient`
(`config/tests/samples/create/harness.go:498`) rather than by overriding endpoints. The
existing test suite is therefore unaffected.

The Go library also handles two things the Java connector had to handle by hand: mTLS is
skipped for non-GDU universes (`internal/cba.go:109`, `:167`), and credential handling
adjusts because token exchange is unavailable outside the GDU (`internal/creds.go:251`).
No KCC work is needed for either. DirectPath is the exception — see §6.8.

#### Terraform-based

`pkg/tf/provider/provider.go:108` builds the provider config map with exactly four keys —
`access_token`, `scopes`, `user_project_override`, `billing_project`. The vendored
provider *does* have a `universe_domain` schema attribute, added by
[#1382](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/1382) in March
2024, but KCC never sets it.

**The vendored provider is already fully universe-capable.** #1382 added the attribute, and
[#1407](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/1407) (backporting
magic-modules#9463) added the base-path rewriting. `providerConfigure` does this immediately
before `SetEndpointDefaults` reads the map:

```go
// third_party/.../google-beta/provider/provider.go
// Replace hostname by the universe_domain field.
if config.UniverseDomain != "" && config.UniverseDomain != "googleapis.com" {
    for key, basePath := range transport_tpg.DefaultBasePaths {
        transport_tpg.DefaultBasePaths[key] = strings.ReplaceAll(basePath, "googleapis.com", config.UniverseDomain)
    }
}
```

The ordering is correct — the substitution runs before the `*_custom_endpoint` defaults are
resolved from `DefaultBasePaths`. So **no vendored change is needed**: KCC simply never set
the attribute. This is one line in `pkg/tf/provider/provider.go`.

This is also the same mechanism upstream settled on. `terraform-provider-google-beta` now
resolves every product's base URL through `transport.BaseUrl()`
(`google-beta/transport/base_url.go`), whose universe handling is the identical
`strings.ReplaceAll`. Host substitution is upstream-sanctioned, not a workaround.

**One caveat worth knowing.** That backport mutates the package-level
`transport_tpg.DefaultBasePaths` map in place, so it is process-global and not reversible:
once a provider is configured with a universe, every later provider built in the same process
inherits the rewritten paths, including one configured for the public universe. In production
this is harmless — `tfprovider.New` runs once per manager process — but it makes tests that
configure a universe order-dependent. This is why `buildProviderConfig` is factored out of
`New`: it lets the attribute mapping be tested without calling `Provider.Configure`. Fixing
the global mutation belongs upstream, not here.

#### Endpoint overrides inside direct controllers

`pkg/controller/direct/` contains **27 `option.WithEndpoint(...)` call sites across 27
files**. These override whatever universe resolution produced, so they win, and each must
be made universe-aware:

*   **15 inline literals**, e.g. `option.WithEndpoint("networksecurity.googleapis.com:443")`
    (13 are NetworkSecurity; the others are Batch ResourceAllowance and Grafeas). One
    special case: `recaptchaenterprise/client.go:42` uses a
    `public-preview-recaptchaenterprise.googleapis.com` prefix.
*   **12 location-computed**, e.g. `fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)`
    across AIPlatform/VertexAI (6), Colab (2), GKEMulticloud, ParameterManager (2, using
    the `parametermanager.<loc>.rep.googleapis.com` regional form), and Tags
    (`location + "-cloudresourcemanager.googleapis.com"`).

### 6.2 API surface

Add an optional grouped field to `ConfigConnectorSpec`
(`operator/pkg/apis/core/v1beta1/configconnector_types.go`):

```go
// Universe configures Config Connector to target a Google Cloud universe other
// than the public one — for example Google Cloud Dedicated, or a sovereign
// cloud such as S3NS. Leave unset for public Google Cloud.
//+kubebuilder:validation:Optional
Universe *UniverseSpec `json:"universe,omitempty"`

type UniverseSpec struct {
    // Domain is the API host suffix of the universe. API endpoints are resolved
    // as <service>.<domain> instead of <service>.googleapis.com.
    // Example: "s3nsapis.fr".
    // If unset, the GOOGLE_CLOUD_UNIVERSE_DOMAIN environment variable is used;
    // if that is also unset, "googleapis.com" applies.
    //+kubebuilder:validation:Optional
    Domain string `json:"domain,omitempty"`

    // Prefix is the universe qualifier applied to project IDs
    // (<prefix>:my-project) and to service-agent email domains. Example: "s3ns".
    // It is NOT derived from Domain and must be set explicitly.
    // If unset, the GOOGLE_CLOUD_UNIVERSE_PREFIX environment variable is used.
    //+kubebuilder:validation:Optional
    Prefix string `json:"prefix,omitempty"`
}
```

A grouped object rather than two flat fields, for three reasons: the two values are always
set together and meaningless apart; it leaves room for the further universe knobs that
Cloud Foundation Fabric's `modules/project` found necessary (`unavailable_services`,
`unavailable_service_identities`) without adding more top-level spec fields; and it keeps
`ConfigConnectorSpec` from accumulating a second unrelated-looking `universe*` entry. Flat
`universeDomain` / `universePrefix` is a viable alternative if maintainers prefer it — see
§10, question 2.

Resolution precedence for each value independently, matching hadoop-connectors#1752 and the
Google SDK convention:

```
spec.universe.domain  >  GOOGLE_CLOUD_UNIVERSE_DOMAIN  >  "googleapis.com"
spec.universe.prefix  >  GOOGLE_CLOUD_UNIVERSE_PREFIX  >  ""
```

`GOOGLE_CLOUD_UNIVERSE_DOMAIN` is a cross-SDK Google standard, honoured by
`google.golang.org/api` itself. **`GOOGLE_CLOUD_UNIVERSE_PREFIX` is not** — there is no
established Google environment variable for the project prefix, so this one is KCC-defined.
It is proposed for symmetry and for parity with the domain's configuration story; if Google
later standardises a name, KCC should adopt it and alias this one.

For the domain, `""` and `"googleapis.com"` must normalize to the same thing everywhere, so
that the two spellings of "public universe" cannot produce divergent behaviour. (The
equivalent normalization was a review finding on hadoop-connectors#1752, where it caused
client-cache fragmentation.)

**Validation.** Setting `prefix` without `domain`, or vice versa, is almost certainly a
misconfiguration and should be rejected by the operator with a clear message rather than
half-applied. A universe needs both.

The field lives on `ConfigConnector` only, not on `ConfigConnectorContext`. A universe is a
property of the whole Google Cloud deployment the cluster talks to, not of a namespace or a
project, so a per-namespace value would be meaningless in practice and would let an operator
create an unserviceable configuration. In namespaced mode the operator propagates the
CC-level value into every per-namespace StatefulSet — there is existing precedent for
exactly this in `applyConfigConnectorExperiments`
(`operator/pkg/controllers/configconnectorcontext/namespaced.go`).

The name matches the Terraform provider's `universe_domain` and the
`GOOGLE_CLOUD_UNIVERSE_DOMAIN` environment variable, so users moving between the two do not
have to learn a second vocabulary.

### 6.3 Plumbing

The path mirrors `--billing-project` exactly:

```
ConfigConnector.spec.universe.{domain,prefix}
  └─ operator/pkg/k8s/constants.go            UniverseDomainFlag = "--universe-domain"
     │                                        UniversePrefixFlag = "--universe-prefix"
     └─ setFlagForManagerContainer(...)       (Deployment in cluster mode,
                                               StatefulSet in namespaced mode)
        └─ cmd/manager/main.go                flag.StringVar(&universeDomain, ...)
           │                                  flag.StringVar(&universePrefix, ...)
           └─ kccmanager.Config.{UniverseDomain,UniversePrefix}
              ├─ config.ControllerConfig.{UniverseDomain,UniversePrefix}  (direct)
              ├─ tfprovider.Config.UniverseDomain                        (Terraform)
              └─ clientconfig.Options                     (DCL, via ControllerConfig)
```

The prefix is carried through the same wiring as the domain in Phase 1, even though nothing
consumes it yet. That is deliberate: the plumbing is the cheap half and is identical for both
values, whereas the parsing fixes that *use* the prefix (§7) are diffuse. Landing the
plumbing early means each subsequent parsing fix is a small self-contained PR against an
already-agreed configuration surface, rather than each one having to re-litigate where the
prefix comes from.

The Terraform provider has no notion of a project prefix — prefixed project IDs are just
opaque strings to it, exactly as legacy domain-scoped IDs (`google.com:my-project`) always
have been — so only the domain flows there.

### 6.4 Direct controllers

Add the field and two helpers to `ControllerConfig`:

```go
// UniverseDomain is the Google Cloud universe to target. Empty means the public
// universe ("googleapis.com").
UniverseDomain string

// GetUniverseDomain returns the configured universe domain, defaulting to
// "googleapis.com". Empty and "googleapis.com" are equivalent.
func (c *ControllerConfig) GetUniverseDomain() string

// Endpoint rewrites a default endpoint into the configured universe.
//   Endpoint("networksecurity.googleapis.com:443")
//     -> "networksecurity.s3nsapis.fr:443"
// It returns the argument unchanged in the public universe.
func (c *ControllerConfig) Endpoint(defaultEndpoint string) string
```

`RESTClientOptions()` and `GRPCClientOptions()` each gain three lines, replacing the TODO:

```go
if c.UniverseDomain != "" {
    opts = append(opts, option.WithUniverseDomain(c.UniverseDomain))
}
```

That alone covers the 173 + 44 controllers that build options through the funnel.

The 27 override sites then become a mechanical one-line change each:

```go
-opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))
+opts = append(opts, option.WithEndpoint(m.config.Endpoint("networksecurity.googleapis.com:443")))
```

The same helper handles the location-computed forms, since it substitutes the host suffix
rather than pattern-matching a service name:

```go
opts = append(opts, option.WithEndpoint(
    m.config.Endpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location))))
```

To stop the pattern regressing, add a unit test in `pkg/controller/direct` that walks the
package AST and fails on any `option.WithEndpoint` argument reaching a `googleapis.com`
literal without passing through `ControllerConfig.Endpoint`. This is cheap, and it is the
only thing that keeps 27 fixed sites from becoming 40 broken ones as #10588 lands more
controllers.

### 6.5 Terraform-based

One part, not two. As established in §6.1, the vendored provider already performs the
base-path substitution; KCC only has to pass the attribute.

`pkg/tf/provider/provider.go` gains a `UniverseDomain` field on `Config` and one map entry:

```go
if config.UniverseDomain != "" {
    cfgMap["universe_domain"] = config.UniverseDomain
}
```

Guarded on non-empty rather than set unconditionally: passing `"googleapis.com"` would be a
no-op substitution, but it would still route the public universe through the in-place
`DefaultBasePaths` mutation, which it does not go through today. The no-op guarantee is worth
more than the uniformity.

The config-map construction is factored into `buildProviderConfig` so it can be tested
without invoking `Provider.Configure`, which mutates provider package state (§6.1).

### 6.6 DCL-based

Proposed for a later phase — but on cost grounds, not because the surface is small. DCL is
the **default controller for 65 kinds, 51 of which have no direct alternative today.** Those
51 kinds are unreachable in a universe until this is solved, and #10588 does not have a dated
plan for retiring them. Anyone reading this as "DCL doesn't matter" is misreading it; the
honest statement is that it is the most expensive third of the work and the least mechanical,
so it goes last.

DCL base paths are hardcoded per *resource* inside the vendored library:

```go
// third_party/.../services/google/compute/firewall_policy_association_internal.go:41
return dcl.Nprintf("https://www.googleapis.com/compute/v1/", params)
```

`dcl.WithBasePath` exists but is a single **global** override, useless when one client talks
to many services. The only remaining seam is a host-rewriting `http.RoundTripper` installed
on the DCL HTTP client in `pkg/dcl/clientconfig/config.go`. The `www.googleapis.com/compute/v1/`
alias needs an explicit special case; a blind `strings.ReplaceAll` would produce
`www.s3nsapis.fr`, which does not exist.

### 6.7 Why a dedicated field, and not just an endpoint override

Worth stating up front, because it is the first question a reviewer asks — it was asked on
hadoop-connectors#1752 and the answer is what settled it.

1.  **Credential/universe validation.** `option.WithUniverseDomain` makes the client
    validate that the credential's universe matches the target before sending a token
    (`internal/settings.go:270`, `internal/creds.go:251`). A raw endpoint override carries no
    such check, so a misconfiguration would **silently send credentials to the wrong
    universe**. That is a security property, not a convenience.
2.  **gRPC.** An endpoint override is per-client and does not compose. `WithUniverseDomain`
    is resolved by the library for both REST and gRPC transports, including mTLS suppression
    (§6.1).
3.  **Convention.** `GOOGLE_CLOUD_UNIVERSE_DOMAIN` is the cross-SDK standard. Users arriving
    from Terraform, gcloud or the Java connectors already know it.

The two are not in conflict: an explicit endpoint override still takes precedence over the
universe-derived endpoint, exactly as in hadoop-connectors#1752.

### 6.8 DirectPath

Unlike the Java client, Go's `checkDirectPathEndPoint` (`transport/grpc/dial.go:488`) does
**not** gate DirectPath on the endpoint being `googleapis.com`. DirectPath is opt-in per
client and additionally requires `metadata.OnGCE()`, so it is unlikely to trigger for KCC's
control-plane clients — but "unlikely" is not "verified".

Action: audit whether any client KCC constructs enables DirectPath; if so, set
`GOOGLE_CLOUD_DISABLE_DIRECT_PATH=true` when a non-default universe is configured. This
mirrors what hadoop-connectors#1752 had to do explicitly on the gRPC path.

## 7. Workstream B: project-ID prefixing

**Confirmed:** in a universe, project IDs carry the universe prefix — `<prefix>:my-project`,
e.g. `s3ns:my-project`. This is the same shape as legacy domain-scoped project IDs
(`google.com:my-project`), so GCP APIs handle it; the risk is entirely in KCC's own parsing.

Split this workstream in two, because the halves have very different costs:

*   **The configured input** — `spec.universe.prefix` and `--universe-prefix`, per §6.2/§6.3.
    Cheap, identical wiring to the domain, **included in Phase 1**. Nothing consumes it yet.
*   **The parsing fixes** — every site that mis-parses a prefixed project ID. Diffuse, no
    reliable static signature, **deferred**.

Landing the input first means each parsing fix afterwards is a small self-contained PR
against a settled configuration surface. Without it, the first such PR has to invent the
configuration too, and every subsequent one inherits that argument.

A sovereign deployment is not actually working until the parsing half is done — it is only
reaching the right endpoints. That is a real limitation and should be stated plainly in the
release notes for Phase 1, not glossed.

### 7.1 Confirmed breakage

`BigQueryDataset` fails outright. BigQuery returns `FullID` as `projectID:datasetID`, and
three call sites split it on `:` and require exactly two tokens:

```go
// pkg/controller/direct/bigquerydataset/bigquerydataset_controller.go:183
tokens := strings.Split(createdMetadata.FullID, ":")
if len(tokens) == 2 {
    resourceID := tokens[1]
    ...
} else {
    return fmt.Errorf("Error getting resourceID: %s. The full ID of the created "+
        "BigQueryDataset is expected to be in the format of projectID:datasetID", ...)
}
```

With a prefixed project the FullID is `s3ns:my-project:my_dataset` — three tokens — and
creation fails with that error. Same pattern at `bigquerydataset_mappings.go:84` and `:99`.

The correct fix is `strings.LastIndex(FullID, ":")` rather than `Split`, which is right in
both universes.

### 7.2 Known-adjacent, needs audit

*   `pkg/util/identity/identity.go:28` — `ProjectIDRegexp = "[a-z][a-z0-9-]{4,28}[a-z0-9]"`
    rejects a colon. Currently referenced only by its own test, so harmless today, but it is
    the canonical definition and will be reused.
*   `apis/common/projects/mapper.go` — project ID ↔ number conversion. Splits links on `/`,
    so a colon inside a segment survives; needs confirming against a real universe response
    (does CRM return `s3ns:my-project` or `my-project` as `projectId`?).
*   Any other `strings.Split(x, ":")` on a value that may contain a project ID.

### 7.3 Why the parsing half is deferred

The endpoint work is bounded and enumerable: 3 config sites plus 27 override sites, all
findable by grep. Project-ID parsing is diffuse — it hides in every resource that handles a
composite identifier, and there is no reliable static signature for it. Attempting both at
once would produce an unreviewable PR and would stall the part that is ready.

Endpoints first is also the right order operationally: with endpoints fixed, a universe user
can determine empirically which resources break on prefixing, which is far better data than
any audit could produce.

### 7.4 Guidance for the parsing fixes

Recorded here so the eventual PRs are consistent:

*   Prefer parsing that is correct in **both** universes over branching on whether a prefix is
    configured. `strings.LastIndex(s, ":")` beats `strings.Split(s, ":")` regardless of
    universe, and needs no configuration at all. Reach for `UniversePrefix` only where the
    ambiguity is genuinely unresolvable without knowing the prefix.
*   Never strip a prefix that was not configured, and never add one the user did not ask for.
*   Per the Cloud Foundation Fabric rules, identifiers built from a **project number** are
    unaffected by prefixing; only **project-ID**-derived strings need care. That distinction
    is a useful triage filter when auditing a controller.

## 8. What deliberately does not change

`pkg/controller/direct/` contains 163 `googleapis.com` occurrences across 85 files. Only the
27 in §6.1 are endpoints. The rest must be left alone, and this section exists so a future
reviewer does not "fix" them:

*   **Cloud Asset Inventory resource names** — `//compute.googleapis.com/projects/...`.
    These are *service names*: logical registry identifiers, not hosts, and confirmed
    unchanged across universes (§5). KCC formalises this as `IdentityV2.Host()` in
    `apis/common/identity/identity.go:34`. Roughly 119 of the 163 occurrences are this.
*   **OAuth scope URLs** — `https://www.googleapis.com/auth/...` in `pkg/gcp/scopes.go`.
    Identifiers, not endpoints. Universe-independent.
*   **Self-links written back to status** — e.g.
    `https://www.googleapis.com/compute/v1/projects/%s/global/networks/%s` in
    `compute/computenetwork_controller.go:190`. These echo what the API returns; in a
    universe the API returns the universe form and KCC should echo *that*, not rewrite it.
    Driven by what the server sends, not by client configuration.

## 9. Testing strategy

*   **Unit.** Table tests for `ControllerConfig.Endpoint` covering the public universe
    (identity), a sovereign universe, the `:443` suffix form, the regional
    `<loc>-aiplatform` form, the `parametermanager.<loc>.rep` form, and the
    `public-preview-` prefix form. Table tests for `GetUniverseDomain` precedence
    (flag > env > default) and for `""` ≡ `"googleapis.com"`. Unit test on the TF config map
    asserting `universe_domain` is absent when unset and present when set.
*   **Operator.** Extend the golden manifest tests in `operator/pkg/test/controller/` to
    assert `--universe-domain` reaches the manager container in both cluster and namespaced
    mode, and is absent when the field is unset.
*   **Integration.** Because mockgcp intercepts at the `RoundTripper` layer, an existing mock
    test can be run with a universe configured and assert the outbound `Request.Host`. This
    gives real end-to-end coverage of endpoint resolution in CI, with no sovereign access
    required, and should gate the change.
*   **Live validation, on two universes.** The author of this proposal has access to both a
    live S3NS universe and a Google Cloud Dedicated environment, and validated
    hadoop-connectors#1752 against S3NS by an equivalent method (containerised client,
    `GOOGLE_CLOUD_UNIVERSE_DOMAIN=s3nsapis.fr`, real bucket operations).

    Testing both universes rather than one is deliberate and worth the extra run: it is the
    only way to catch an implementation that has accidentally hardcoded something
    S3NS-specific. In particular the universe *domain* and the universe *prefix* are
    independent values (§5), and a single-universe test cannot distinguish a correct
    implementation from one that assumes a relationship between them.

    This will be run and reported before the PR is marked ready, since the first question
    asked on #1752 — before review even started — was whether the change had been tested in
    a real TPC universe.
*   **No-op guarantee.** The public-universe path must produce identical client options.
    Assert this directly rather than inferring it from a green suite.

## 10. Open questions for maintainers

Two earlier questions are now **resolved**, confirmed by an operator running both S3NS and a
Google Cloud Dedicated environment:

*   *Do service names change in a universe?* **No.** `<service>.googleapis.com` remains the
    service identifier everywhere — which is why enabling an API in a universe uses the same
    name as public GCP. This is what makes §8 safe and keeps ~119 occurrences out of scope.
*   *Are project IDs prefixed?* **Yes** — `<prefix>:my-project`. Hence §7.

Remaining:

1.  ~~Is a vendored TF backport acceptable?~~ **Moot.** #1407 already backported the base-path
    substitution; the vendored provider is universe-capable and KCC only had to pass the
    attribute. No `third_party/` change is proposed.
2.  **Is `spec.universe.{domain,prefix}` the shape you want**, or flat `universeDomain` /
    `universePrefix`? §6.2 argues for the grouped object; the flat form is a defensible
    alternative and trivial to switch to before anything ships.
3.  **CC-only, or CCC too?** §6.2 argues a universe is a cluster-wide property and a
    per-namespace value would be meaningless, but `billingProject` sets a contrary precedent.
4.  **Is `GOOGLE_CLOUD_UNIVERSE_PREFIX` an acceptable KCC-defined env var**, given there is no
    Google-standard name for the project prefix? If one exists internally, KCC should use it
    instead.
5.  **Is Phase 3 (DCL) wanted?** 65 kinds depend on it, 51 with no direct alternative. If
    there is a dated plan to retire DCL that makes the work moot, that changes the calculus —
    but from outside, those 51 kinds look permanently stranded in a universe otherwise.

## 11. Interim workaround for users

Until this lands, users on a sovereign universe can set `GOOGLE_CLOUD_UNIVERSE_DOMAIN` on
the `cnrm-controller-manager` pod. `google.golang.org/api` honours it
(`internal/settings.go:245-258`), so direct controllers that build clients through
`RESTClientOptions()` / `GRPCClientOptions()` and do not override the endpoint will resolve
correctly.

The limits should be stated plainly: it does nothing for Terraform-based or DCL-based
resources, nothing for the 27 controllers that override the endpoint explicitly, and nothing
for project-ID prefixing (§7). It is worth trying as a diagnostic — if resources backed by
direct controllers start reconciling and TF-backed ones do not, that confirms the analysis
above.

## 12. Implementation phases

1.  **Phase 1 — configuration surface and chokepoint.** `spec.universe.{domain,prefix}` CC
    field with validation, `--universe-domain` / `--universe-prefix` flags, operator
    propagation, `ControllerConfig.{UniverseDomain,UniversePrefix}` + option wiring, TF
    attribute + vendored base-path backport, unit and operator tests, mockgcp Host assertion.
    Unblocks the 215 direct-controller kinds and the 234 TF-backed kinds for endpoints, and
    settles where the prefix comes from for everything that follows.
2.  **Phase 2 — endpoint overrides.** The 27 sites in §6.1 converted to
    `ControllerConfig.Endpoint`, plus the AST regression test and the DirectPath audit (§6.8).
3.  **Phase 3 — DCL.** Host-rewriting RoundTripper. Not optional in the long run: 65 kinds
    depend on it, 51 with no direct alternative (§6.6). Last only because it is the most
    expensive.
4.  **Phase 4 — project-ID parsing.** Consumes the prefix landed in Phase 1. Separate design;
    §7.4 is its guidance. Start with the confirmed `BigQueryDataset` breakage, which is a
    self-contained fix worth landing on its own merit (`LastIndex` is more correct than
    `Split` in the public universe too).

Phases 1 and 2 are independent of each other and can be reviewed separately. Phase 4 depends
only on Phase 1's configuration surface, not on Phases 2 or 3.

**Scope honesty:** at the end of Phase 2, KCC reaches the right endpoints for ~87% of kinds
but is not yet fully functional in a universe — prefixed project IDs will still break an
unknown subset of resources. Release notes for any phase before 4 must say so.

## 13. References

*   [#5995](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5995) — original question
*   [#10588](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/10588) — direct controller migration tracker
*   [#1382](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/1382) — vendored backport of the `universe_domain` provider attribute
*   [hadoop-connectors#1752](https://github.com/GoogleCloudDataproc/hadoop-connectors/pull/1752) — prior art; universe support in the GCS Hadoop connector, validated on S3NS
*   `terraform-provider-google-beta`, `google-beta/transport/base_url.go` — upstream universe substitution
*   `google.golang.org/api/option.WithUniverseDomain`, `google.golang.org/api/internal/{settings,cba,creds}.go`
