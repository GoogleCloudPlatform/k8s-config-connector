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

`--manageability` builds the list of every resource KCC could support, with a reason for every resource it rules out. Run it from the repo root:

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

**Buckets:**

| Bucket | Rule |
|---|---|
| Fully Manageable | CREATE + READ + DELETE. UPDATE is optional. |
| Partially Manageable: no DELETE | CREATE + READ, no DELETE. |
| Partially Manageable: no CREATE | READ + UPDATE, no CREATE (acquire-only). |
| Needs review | READ can't be resolved: only `List` with a user-specified ID, or no `Get`/`List`. |
| Brownfield (no googleapis resource) | v1alpha1 TF/DCL kinds that no googleapis resource maps to. |
| Dropped | Matched a drop rule (see below). |

**Drop reasons**, in precedence order. Every dropped row has `drop_reason`, `drop_detail` and a `drop_link` to the rule or data:

| Reason | Rule |
|---|---|
| `skip-list` | Matches `coverage_skip.json`. |
| `example-api` | `google.example.*` sample APIs. |
| `deprecated` | Deprecated proto package, service or RPCs, or listed in `deprecated_services.json`. |
| `non-ga-only` | Only defined in alpha/beta packages. |
| `v1beta1-brownfield` | Already served by a KCC kind with a `v1beta1` version. |
| `no-crud` | No CRUD or List RPCs. |
| `read-only` | Neither CREATE nor UPDATE. |
| `nondeclarative:pattern` | Ephemeral, data-plane, snapshot or backup collection. |
| `nondeclarative:bulk-rpc` | Collection has bulk or purge RPCs. |
| `policy:<reason>` | Listed in `policy_exclusions.json`: Datapath API, Private Cloud, Drip risk. |

**KCC status** (`kcc_status`):
- `net-new`: no KCC kind yet.
- `implemented (direct)`: the kind has a direct controller.
- `in progress (types only)`: types and CRD exist, but no controller yet.
- `brownfield (v1alpha1 TF/DCL)`: only a TF or DCL controller.
- `v1beta1`: already served by a v1beta1 kind (dropped as `v1beta1-brownfield`).

Kinds are matched to protos through their `+kcc:proto` annotations. TF/DCL kinds have no annotations, so they're matched by name, with a few manual aliases in the script (`MANUAL_KIND_ALIASES`, `MANUAL_KIND_OVERRIDES`).

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

`first_googleapis_sha` and `first_googleapis_link` point at the earliest KCC googleapis pin whose upstream tree has the message: the global pin first, then the per-service `PROTO_SHA`s in `apis/*/generate.sh`. Each pin is compiled from the full googleapis tree at that commit. This differs from the curated subset that `generate-proto.sh` builds into `.build/`, so a message can be "pinned" here and still need a new glob in `generate-proto.sh` before `generate-types` can see it. Missing pin commits are fetched by SHA (network); pass `--no-fetch-pins` to skip them, in which case they're listed in `meta.skipped_googleapis_pins`.

**API groups** (`api_group`, `api_family`):
- `Google Cloud`: `google.cloud.*`, plus the Cloud APIs whose protos live elsewhere: Spanner, Bigtable, IAM, Cloud Storage, Storage Transfer, Logging, Monitoring, Pub/Sub, Firestore, Datastore, App Engine, GKE, DLP, Access Context Manager, `google.api.*` (API Keys, Cloud Quotas, Service Usage), `google.devtools.*` and Grafeas.
- `Other Google APIs`: everything else, with `api_family` set to `Workspace`, `Ads & Marketing`, `Shopping`, `Maps` or `Misc`.
- The summary and top-K list put Google Cloud first.

### Output (`manageability_list.json`)

```json
{
  "meta": {
    "schema_version": 1,
    "googleapis_sha": "<resolved googleapis commit>",
    "kcc_ref": "<KCC commit used for GitHub links>",
    "googleapis_pins": [{"sha": "...", "used_by": ["global (apis/git.versions)"]}],
    "skipped_googleapis_pins": [],
    "config_fallbacks": [],
    "counts": {"Google Cloud": {"Fully Manageable": {"net-new": 110}}}
  },
  "resources": [{"id": "...", "bucket": "...", "kcc_status": "...", "api_group": "...", "...": "..."}]
}
```

`resources` uses the same columns as `manageability_list.csv`.

**Reproducibility.** For the same googleapis and KCC SHAs, both files are byte-for-byte identical:
- There are no timestamps or local paths.
- Rows have a total order (API group, bucket, KCC status, id) and every row uses the same fixed key order.
- The rule data (`coverage_skip.json`, `policy_exclusions.json`, `deprecated_services.json`, `services_acronyms.json`, `cp_resources_list.json`) and the controllerbuilder rules are read from the analyzed KCC tree, not from wherever the script runs.

To get identical output on any machine:
- Pass full SHAs (or check `meta.googleapis_sha` / `meta.kcc_ref`); `LOCAL` uses the working tree as-is.
- Let the script fetch pins: `meta.skipped_googleapis_pins` must be empty.
- `meta.config_fallbacks` lists any rule file the analyzed tree didn't have, which then came from the script's directory.

## Unified Resource Policy

When implementing missing resources identified by this tool, follow the **Unified Direct Pattern**:

1.  **Single logical Kind**: If a GCP resource supports multiple hierarchies (e.g., Global and Regional) but shares the same logical name, do **not** create separate KRM Kinds.
2.  **Extended Existing Kinds**: If a resource identified as "missing" is actually a hierarchical variant of an existing KCC Kind (e.g., Regional Secret Manager Secret), the correct action is to **extend the existing Kind** with a `location` field rather than creating a new one.
3.  **Location-Aware Identity**: Use `gcpurls.Template` to implement identity logic that branches based on the `location` field.

## Implementation Workflow

This tool is integrated into the project's agentic workflows. The chore file `.agents/greenfield-direct-new-resource-types.md` uses this script to automatically identify and create implementation tasks for missing resources.
