# KCC Resource Coverage Analysis Tool

This tool calculates the coverage of Google Cloud Platform (GCP) resources in Kubernetes Config Connector (KCC) relative to the published definitions in the `googleapis/googleapis` repository.

It is used to identify gaps in resource coverage and prioritize the "easiest" next resources to implement using the direct reconciliation model.

## Features

- **SHA-based Analysis**: Compare specific versions of both `googleapis` and `k8s-config-connector`.
- **Capability Detection**: Identifies which GCP resources support standard CRUD operations (Create, Read, Update, Delete) based on gRPC service definitions.
- **Hierarchy Analysis**: Detects "leaf" resources whose parents are Projects, Folders, Organizations, or Locations.
- **Heuristic Matching**: Links GCP proto types to KCC CRDs by accounting for service aliases and naming conventions (e.g., stripping "Compute" prefix from `ComputeInstance`).
- **Prioritization**: Recommends the Top K missing manageable resources to implement, sorted by ease of implementation (Leaf, Next Layer, etc.), displaying their available and missing CRUD operations.

## Usage

```bash
python3 hack/tools/greenfield/calculate_coverage.py <googleapis_sha> <kcc_sha> [k] [--update-gap] [--resources-list] \
    [--manageability] [--googleapis-repo PATH] [--output-dir DIR] [--no-fetch-pins]
```

### Parameters

- `<googleapis_sha>`: The git SHA, branch, or tag in the `googleapis/googleapis` repo. The pinned commit for standard gap analysis is [3f9c9d7](https://github.com/googleapis/googleapis/commit/3f9c9d72cb20768ca4ac9f12030faaf43b13c231) (or use `master` for the bleeding edge).
- `<kcc_sha>`: The git SHA in the `k8s-config-connector` repo. Use `LOCAL` to compare against your current working directory.
- `[k]`: (Optional) The number of top manageable resources to list. Defaults to 10.
- `--update-gap`: (Optional) If provided, updates the `gap_analysis.txt` file with the coverage snapshot.
- `--resources-list`: (Optional) If provided, outputs a full structured list of all resources with their layer and CRUD status to `resources_list.json`.
- `--manageability`: (Optional) Classifies every googleapis resource with the manageability rules below and writes `manageability_list.json` and `manageability_list.csv`. Requires `protoc`. When it's the only output flag, the legacy analysis is skipped.
- `--googleapis-repo PATH`: (Optional) Use an existing local googleapis clone (for example `.build/third_party/googleapis`) instead of cloning into `$TMPDIR/kcc_coverage`. The protos are exported with `git archive`, so the clone's checkout isn't modified. The commit must already be fetched.
- `--output-dir DIR`: (Optional) Where to write the manageability files. Defaults to this directory.
- `--no-fetch-pins`: (Optional) Don't fetch missing googleapis pin commits. The output is then incomplete for the pin columns, and the skipped pins are listed in the JSON `meta`.

### Example

Compare the pinned googleapis protos against your local KCC state with CRUD analysis and save both reports:
```bash
python3 hack/tools/greenfield/calculate_coverage.py 3f9c9d7 LOCAL 20 --update-gap --resources-list
```

## Metric Definitions

- **Total GCP Resources**: Every unique `google.api.resource` or `google.api.resource_definition` found in the protos.
  - **Manageable**: Resources that have at least one CRUD function (`CREATE`, `READ`, `UPDATE`, or `DELETE`) but are not in KCC.
    - **Adoptable**: Resources that have no `CREATE` function, but have at least one other CRUD function.
    - **Partially Manageable(Creatable)**: Resources that have the `CREATE` function but lack the `DELETE` function.
    - **Fully Manageable**: Resources that have both `CREATE` and `DELETE` functions.
      - **Leaf (Easy)**: Resources that are fully manageable and have flat parentage (Project, Folder, Org, or Location). These are generally the most straightforward to implement as direct resources.
      - **Next Layer**: Resources that are fully manageable and have one dependency.
  - **Unmanageable**: Resources that have no CRUD function.

## Manageability Mode

`--manageability` builds the list of every resource KCC could support. It gives a reason for every resource it rules out, and says what KCC has for each resource today. Run it from the repo root:

```bash
git -C .build/third_party/googleapis fetch --depth 1 origin master
python3 hack/tools/greenfield/calculate_coverage.py origin/master LOCAL 20 --manageability \
    --googleapis-repo .build/third_party/googleapis --output-dir /tmp/manageability
```

Unlike the legacy analysis, it compiles all googleapis protos with `protoc` and reads the descriptors, so RPC names, HTTP bindings, request fields and `deprecated` options are exact. It also reuses the rules that `controllerbuilder` enforces, read straight from this checkout: the `generate-types` non-GA, deprecated and non-declarative gates (`nondeclarative.go`, `deprecated.go`), and the `generate-identity` server-generated ID check.

**Operations.** CRUD RPCs are counted only from the resource's own package; for GA resources that's the newest GA major.
- Canonical verbs: `Create`/`Insert`, `Get`, `Update`/`Patch`, `Delete`.
- `Destroy`/`Disable`/`Deactivate`/`Remove` also count as DELETE.
- Aliases are counted but flagged: `Cancel`/`Abort`/`Finish`/`Terminate` (one-shot), `BatchCreate`/`BatchUpdate`/`Upload`/`Upsert`/`Set`, and `BatchGet`/`Describe`/`Read`.
- A `List` RPC stands in for `Get` only when the ID is server-generated: the `Create` request has a `parent` and no `*_id` field.

**Buckets.** The bucket says what the resource's API can do. It doesn't depend on what KCC has for the resource: that is in `kcc_status` and `kcc_stability`.

| Bucket | Rule |
|---|---|
| Fully Manageable | CREATE + READ + DELETE. UPDATE is optional. |
| Partially Manageable: no DELETE | CREATE + READ, no DELETE. |
| Partially Manageable: no CREATE | READ + UPDATE, no CREATE (acquire-only). |
| Needs review | READ can't be resolved: only `List` with a user-specified ID, or no `Get`/`List`. |
| No googleapis resource | A KCC kind that no googleapis resource on the list maps to (see below). |
| Dropped | Matched a drop rule (see below). |

**Drop reasons**, in precedence order. Every dropped row has `drop_reason`, `drop_detail` and a `drop_link` to the rule or data:

| Reason | Rule |
|---|---|
| `skip-list` | Matches `coverage_skip.json`. |
| `example-api` | `google.example.*` sample APIs. |
| `deprecated` | Deprecated proto package, service or RPCs, or listed in `deprecated_services.json`. |
| `non-ga-only` | Only defined in alpha/beta packages. |
| `no-crud` | No CRUD or List RPCs. |
| `read-only` | Neither CREATE nor UPDATE. |
| `nondeclarative:pattern` | Ephemeral, data-plane, snapshot or backup collection. |
| `nondeclarative:bulk-rpc` | Collection has bulk or purge RPCs. |
| `policy:<reason>` | Listed in `policy_exclusions.json`: Datapath API, Private Cloud, Drip risk. |

A resource is never dropped because KCC already supports it. A supported resource that matches one of these rules is dropped like any other and keeps its `kcc_status`: `serviceusage.googleapis.com/Service` is `read-only` here, and KCC serves it as a stable kind.

**KCC status** (`kcc_status`) says what KCC has for the resource:
- `net-new`: no KCC kind yet.
- `in progress (types only)`: types and CRD exist, but no controller yet.
- `brownfield (TF/DCL)`: only a TF or DCL controller.
- `implemented (direct)`: the kind has a direct controller.
- `implemented (other controller)`: the kind has a controller of its own that is neither direct nor TF/DCL. These are the IAM policy kinds.

**Stability** (`kcc_stability`) says which API version KCC serves the kind at:
- `stable`: the kind serves `v1beta1`.
- `alpha`: the kind only serves `v1alpha1`.
- empty: the resource is `net-new`.

These are the words of the `cnrm.cloud.google.com/stability-level` CRD label. The script reads the versions the CRD serves instead, because not every CRD has the label.

A resource can have several KCC kinds, which `kcc_kinds` lists with their versions and controller. When some of them are stable, only those decide `kcc_status`: a stable TF kind next to an alpha direct kind is `brownfield (TF/DCL)` and `stable`. `kind` and `kcc_link` name the first of the kinds that decided.

Kinds are matched to protos through their `+kcc:proto` annotations. TF/DCL kinds have no annotations, so they're matched by name, with a few manual aliases in the script (`MANUAL_KIND_ALIASES`, `MANUAL_KIND_OVERRIDES`).

**Kinds with no googleapis resource.** Every KCC kind is on the list. The kinds come from the CRDs in `config/crds/resources`, and their controllers from `pkg/controller/resourceconfig/static_config.go` and the direct controllers. A kind that no googleapis resource maps to gets a row of its own in the `No googleapis resource` bucket, with the id `(no googleapis resource) <Kind>`. Its `kcc_status` and `kcc_stability` are set like on any other row; it has no ops, because there is no proto to read them from. These kinds are:
- sub-objects of another resource, such as `ComputeRouterNAT` or `StorageBucketAccessControl`;
- kinds of APIs that publish no protos in googleapis, such as Cloud DNS or Apigee;
- kinds that only exist in KCC, such as `IAMPolicy` and `IAMPolicyMember` (`KCC_ONLY_KINDS` in the script keeps them from being matched by name);
- resources that the protos don't declare as one, such as `DataprocJob`.

`meta.kcc_kinds` counts them: `total` is every kind, `without_googleapis_resource` the kinds in this bucket.

**Kind names** (`kind`, `predicted_kind`). For a resource that isn't in KCC yet, `kind` is the name a new Kind must have under [`dev/tools/controllerbuilder/serviceprefixes`](../../../dev/tools/controllerbuilder/serviceprefixes): `<service prefix><proto message>`, with the acronyms in `acronyms.yaml` applied to the message words. `sqladmin` + `BackupRun` gives `SQLAdminBackupRun`. This is the name that `TestKindNamesFollowServicePrefixes` in `tests/apichecks` expects. For a resource already in KCC, `kind` is its existing Kind and `predicted_kind` is what the rule gives.

The API's service prefix is found the way that test finds it, from the proto package in `protoPackages`. A package that no `generate.sh` uses yet is matched without its version, then by API name.

**Flags** (`flags`) never change the bucket:

| Flag | Meaning |
|---|---|
| `nondeclarative:one-shot` | CREATE without UPDATE, and no DELETE or a `Cancel` RPC. `generate-types` needs `--allow-non-declarative`. |
| `alias:*` | The op only comes from a flagged alias verb. |
| `delete-via:*` | DELETE comes from `Destroy`, `Disable`, `Deactivate` or `Remove`. |
| `read-via-list` | READ is inferred from `List`. |
| `review:*` | Why the row needs review. |
| `deprecated-rpc:*` | Some managing RPCs are deprecated and weren't counted. |
| `not-in-kcc-pinned-googleapis` | The message isn't in the global googleapis pin (`apis/git.versions`). |
| `kind-prefix:pending` | The suggested Kind uses a service prefix that's still in `service-prefix-pending.yaml`. The test rejects the Kind until the prefix is approved. |
| `kind-prefix:unlisted` | The API is in neither service-prefix file, so the prefix is a guess from the API name. Propose one in `service-prefix-pending.yaml`. |
| `kind-name:repeats-prefix(<Kind>)` | The message starts with the words the prefix ends with (`CloudBuild` + `BuildTrigger`). `<Kind>` says them once (`CloudBuildTrigger`); using it needs a `knownKindNames` entry in the test. |

`first_googleapis_sha` and `first_googleapis_link` point at the earliest KCC googleapis pin whose upstream tree has the message: the global pin first, then the per-service `PROTO_SHA`s in `apis/*/generate.sh`. Each pin is compiled from the full googleapis tree at that commit. This differs from the curated subset that `generate-proto.sh` builds into `.build/`, so a message can be "pinned" here and still need a new glob in `generate-proto.sh` before `generate-types` can see it. Missing pin commits are fetched by SHA (network); pass `--no-fetch-pins` to skip them, in which case they're listed in `meta.skipped_googleapis_pins`.

**API groups** (`api_group`, `api_family`):
- `Google Cloud`: `google.cloud.*`, plus the Cloud APIs whose protos live elsewhere: Spanner, Bigtable, IAM, Cloud Storage, Storage Transfer, Logging, Monitoring, Pub/Sub, Firestore, Datastore, App Engine, GKE, DLP, Access Context Manager, `google.api.*` (API Keys, Cloud Quotas, Service Usage), `google.devtools.*` and Grafeas.
- `Other Google APIs`: everything else, with `api_family` set to `Workspace`, `Ads & Marketing`, `Shopping`, `Maps` or `Misc`.
- The summary and top-K list put Google Cloud first.

### Output (`manageability_list.json`)

```json
{
  "meta": {
    "schema_version": 2,
    "googleapis_sha": "<resolved googleapis commit>",
    "kcc_ref": "<KCC commit used for GitHub links: the newest one that changed an input of the list>",
    "kcc_uncommitted_inputs": [],
    "googleapis_pins": [{"sha": "...", "used_by": ["global (apis/git.versions)"]}],
    "skipped_googleapis_pins": [],
    "config_fallbacks": [],
    "kind_naming": {"source": "dev/tools/controllerbuilder/serviceprefixes", "approved_prefixes": 159, "pending_prefixes": 48, "acronyms": 40},
    "kcc_kinds": {"total": 809, "with_googleapis_resource": 701, "without_googleapis_resource": 108},
    "counts": {"Google Cloud": {"Fully Manageable": {"net-new": 102, "brownfield (TF/DCL)": 96}}},
    "stable_counts": {"Google Cloud": {"Fully Manageable": {"brownfield (TF/DCL)": 80}}}
  },
  "resources": [{"id": "...", "bucket": "...", "kcc_status": "...", "kcc_stability": "...", "api_group": "...", "...": "..."}]
}
```

`resources` uses the same columns as `manageability_list.csv`. `counts` is the number of rows per API group, bucket and `kcc_status`. `stable_counts` has the same shape and only counts the rows whose `kcc_stability` is `stable`.

Some queries:

```bash
# Resources KCC could support that have no controller yet.
jq '[.resources[] | select(.bucket != "Dropped" and (.kcc_status == "net-new" or .kcc_status == "in progress (types only)"))] | length' \
    hack/tools/greenfield/manageability_list.json
# Stable kinds that still run on TF or DCL.
jq -r '.resources[] | select(.kcc_stability == "stable" and .kcc_status == "brownfield (TF/DCL)") | .kind' \
    hack/tools/greenfield/manageability_list.json
# Resources KCC serves as stable that the rules drop, with the reason.
jq -r '.resources[] | select(.bucket == "Dropped" and .kcc_stability == "stable") | [.kind, .drop_reason] | @tsv' \
    hack/tools/greenfield/manageability_list.json
```

**Reproducibility.** The list depends only on the googleapis commit and on the KCC files it is built from. The same inputs give byte-for-byte identical files:
- There are no timestamps or local paths.
- Rows have a total order (API group, bucket, KCC status, id) and every row uses the same fixed key order.
- The rule data (`coverage_skip.json`, `policy_exclusions.json`, `deprecated_services.json`, `cp_resources_list.json`, the `serviceprefixes` files) and the controllerbuilder rules are read from the analyzed KCC tree, not from wherever the script runs.
- `meta.kcc_ref` is not `HEAD`. It is the newest commit that changed one of the KCC files the list is built from: the CRDs, `static_config.go`, the direct controllers, `apis/`, the rule data and the controllerbuilder rules (`KCC_INPUT_PATHS` in the script). The GitHub links point at it. A commit that only touches other files doesn't move it, and that includes the commit that checks the list in. So running the script again on the commit that contains the list writes the same file.

To get identical output on any machine:
- Pass full SHAs (or check `meta.googleapis_sha` / `meta.kcc_ref`); `LOCAL` uses the working tree as-is.
- Commit changes to the KCC inputs before running: `meta.kcc_uncommitted_inputs` must be empty. It lists the inputs with local changes. No commit has them, so links into them can show other lines, and the list changes once they are committed.
- Let the script fetch pins: `meta.skipped_googleapis_pins` must be empty.
- `meta.config_fallbacks` lists any rule file the analyzed tree didn't have, which then came from the script's own checkout.

### Checked-in snapshot

[`manageability_list.json`](manageability_list.json) is a checked-in run of this mode, so the list can be read without running the script. Its `meta` block names the googleapis and KCC commits it was generated from.

To update it to the newest googleapis, run the mode from the repo root without `--output-dir`, which writes next to the script:

```bash
git -C .build/third_party/googleapis fetch --depth 1 origin master
python3 hack/tools/greenfield/calculate_coverage.py origin/master LOCAL 20 --manageability \
    --googleapis-repo .build/third_party/googleapis
```

To regenerate it from the googleapis commit it already uses, pass that commit:

```bash
sha=$(jq -r .meta.googleapis_sha hack/tools/greenfield/manageability_list.json)
git -C .build/third_party/googleapis fetch --depth 1 origin "$sha"
python3 hack/tools/greenfield/calculate_coverage.py "$sha" LOCAL 20 --manageability \
    --googleapis-repo .build/third_party/googleapis
git status --short hack/tools/greenfield
```

The run is idempotent: `git status` prints nothing when the list is up to date with the checkout, however often it runs. A change to the script or to a KCC input shows up as a diff of the list.

A change to the script and the list it regenerates can go in one commit. A change to a KCC input, such as `policy_exclusions.json`, has to be committed before the list is generated (`meta.kcc_uncommitted_inputs`). Only the JSON is checked in: the `manageability_list.csv` written next to it has the same rows and is ignored by git.

## Unified Resource Policy

When implementing missing resources identified by this tool, follow the **Unified Direct Pattern**:

1.  **Single logical Kind**: If a GCP resource supports multiple hierarchies (e.g., Global and Regional) but shares the same logical name, do **not** create separate KRM Kinds.
2.  **Extended Existing Kinds**: If a resource identified as "missing" is actually a hierarchical variant of an existing KCC Kind (e.g., Regional Secret Manager Secret), the correct action is to **extend the existing Kind** with a `location` field rather than creating a new one.
3.  **Location-Aware Identity**: Use `gcpurls.Template` to implement identity logic that branches based on the `location` field.

## Implementation Workflow

This tool is integrated into the project's agentic workflows. The chore file `.agents/greenfield-direct-new-resource-types.md` uses this script to automatically identify and create implementation tasks for missing resources.
