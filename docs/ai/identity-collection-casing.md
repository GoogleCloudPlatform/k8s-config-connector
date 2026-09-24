# Identity Collection Segments GCP Has Never Seen

**Status:** Found by a proto-based audit of existing resources. `TestIdentityCollectionCasing` ratchets the subset verifiable against recorded HTTP traffic. None of the nine identified mismatches is currently modified in-place, because altering a collection segment changes a published `status.externalRef`, requiring careful backwards-compatibility analysis.

## Problem Description

In `template/apis/identity.go`, collection segments were historically constructed by lowercasing the proto message name and appending `s`:

```go
return i.parent.String() + "/{{.ProtoMessageName | ToLower}}s/" + i.id
```

In contrast, canonical GCP collection segments typically use camelCase. An audit of the 90 `_identity.go` files that correspond to declared `google.api.resource` patterns revealed **9 casing mismatches**:

- `apphubdiscoveredservice` (expected: `discoveredServices`)
- `apphubdiscoveredworkload` (expected: `discoveredWorkloads`)
- `bigquerydatapolicy` (expected: `dataPolicies`)
- `clouddmsconversionworkspace` (expected: `conversionWorkspaces`)
- `dataprocnodegroup` (expected: `nodeGroups`)
- `discoveryenginedatastore` (expected: `dataStores`)
- `managedkafkaconsumergroup` (expected: `consumerGroups`)
- `netappbackupvault` (expected: `backupVaults`)
- `storagemanagedfolder` (expected: `managedFolders`)

The collection segment forms part of the GCP resource name (`status.externalRef`, and the format `ParseXExternal` accepts in `spec...Ref.external`), so it must match the API format.

**Note on KRM naming:** This issue is distinct from KRM schema field naming:

| Scope | Example | Configured By | Test Enforcement |
|---|---|---|---|
| KRM field name | `spec.forwardingRules` | `GetJSONForKRM` | `TestCRDsAcronyms` |
| GCP collection segment | `projects/p/locations/l/lbTrafficExtensions/x` | Identity template | `TestIdentityCollectionCasing` |

## Evidence from Recorded Traffic and Protos

When checked against recorded GCP traffic under `pkg/test/resourcefixture/testdata/`, camelCase is what GCP actually exchanges (e.g. `backupVaults` 27 files to 0, `nodeGroups` 6 to 0, `conversionWorkspaces` 5 to 0).

`StorageManagedFolder` illustrates the discrepancy clearly. GCP's response body in `_http.log`:

```json
{
  "createTime": "2024-04-01T12:34:56.123456Z",
  "metageneration": "1",
  "name": "projects/_/buckets/bucket-${uniqueId}/managedFolders/managedfolder-${uniqueId}/",
  "updateTime": "2024-04-01T12:34:56.123456Z"
}
```

The protobuf declaration agrees: `+kcc:spec:proto=google.storage.control.v2.ManagedFolder`, whose `google.api.resource` pattern specifies `managedFolders`.

In contrast, three internal KCC locations exhibited discrepancies:

| Source | Format |
|---|---|
| GCP (response above and proto pattern) | `projects/_/buckets/{b}/managedFolders/{f}/` |
| `ParseManagedFolderExternal` (`_identity.go:107`) | Requires `buckets` and `managedfolders` (exactly 6 tokens) |
| `ManagedFolderRef` doc comment (`_reference.go:34`) | Mentions `projects/{p}/locations/{l}/managedfolders/{id}` |

The doc comment mentioned `locations` where the parser demands `buckets`, and both diverged from GCP on casing. Furthermore, GCP's trailing slash causes `strings.Split` to yield seven tokens, causing a canonical GCP name to fail the token length check before casing validation is even reached.

## Why Discrepancies Remained Undetected

Reconciliation succeeds because the underlying client API calls construct the request paths independently using correct casing. What diverged is the identifier string populated in `status.externalRef` and parsed by reference resolvers. This path is user-facing via `NormalizedExternal` (`_reference.go`), which parses user-supplied `external:` references.

Because both the serializer (`_identity.go`) and parser shared the same constant, KCC round-tripped its own values internally without encountering errors. Existing presubmits did not compare this pair: `TestCRDsAcronyms` validates field names, `shortname_pluralization.txt` validates CRD `shortNames`, and `naming_violations.txt` validates file naming conventions.

## Test Coverage and Ratchet Mechanism

`TestIdentityCollectionCasing` extracts each Identity's collection segment and compares it against segments observed in recorded HTTP traffic, ratcheting exceptions through `testdata/exceptions/identity_collection_casing.txt`.

Recorded traffic was chosen over proto parsing because `.build/googleapis.pb` is gitignored and not present in standard CI environments, which would cause proto checks to be skipped. In contrast, `_http.log` fixtures are committed and provide fast test execution across hundreds of fixture files.

Trade-offs of this approach:
- Because the fixtures record both KCC requests and GCP responses, an incorrect casing is only flagged when another part of the system or response uses the correct casing.
- Resources without existing test fixtures cannot be validated against recorded logs. Consequently, the baseline exception list tracks the 4 resources with fixtures (`netappbackupvault`, `discoveryenginedatastore`, `storagemanagedfolder`, `dataprocnodegroup`) rather than all 9 detected via proto audit.

Additionally, `TestIdentityCollectionRegex` ensures the regular expression anchors properly on `parent.String()`. Without this anchor, matching can inadvertently bind to intermediate segments (such as `/locations/` within the parent's `String()`), masking actual mismatches.

## Next Steps

- **Unratcheted Resources:** Resources lacking recorded fixtures require either test fixtures or a hermetic proto inspection step.
- **Backwards Compatibility:** Correcting existing identifiers in `status.externalRef` requires a migration plan (e.g. accepting both lowercase and camelCase in parsers while emitting canonical camelCase).
- **Scaffolding Guardrails:** Greenfield generator enhancements read declared `google.api.resource` patterns directly rather than applying heuristic lowercasing to proto message names.
