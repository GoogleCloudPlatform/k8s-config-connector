# Forks of the multi-cluster leader election components

Copies of the two Go modules that Config Connector's multi-cluster leader election uses:

| Directory | Module | Version | Upstream source |
|---|---|---|---|
| `multicluster-leader-election/` | `github.com/gke-labs/multicluster-leader-election` | `v0.0.0-20260219060047-1ae5cb16f27e` | [gke-labs/multicluster-leader-election@1ae5cb16f27e](https://github.com/gke-labs/multicluster-leader-election/tree/1ae5cb16f27e) |
| `kube-etl/syncer/` | `github.com/gke-labs/kube-etl/syncer` | `v0.0.0-20261027010125-063e4a30cf84` | [gke-labs/kube-etl@063e4a30cf84, `syncer/`](https://github.com/gke-labs/kube-etl/tree/063e4a30cf84/syncer) |

- **Wiring.** `go.mod` and `mockgcp/go.mod` replace both modules with these directories. Their `require` lines keep the upstream version that each copy was taken from.
- **Two builds of the same code:**
  - KCC's manager imports the forks' API and client packages, so it builds them with KCC's `go.mod`.
  - The `multiclusterlease` and `krmsyncer` images are built from these directories as standalone modules, each with its own `go.mod` and `go.sum` (see `build/builder/Dockerfile`).
  - Dependency fixes for those images go into the fork's `go.mod`. KCC's `go.mod` doesn't change what they link.
- **Source.** Each directory is the module exactly as `go mod download` fetches it. When copied, each directory's `h1:` hash matched the module's `go.sum` entry. Module downloads don't record file modes, so no file is executable.
- **Local changes.** Keep them small, and list each one here so that the next update can reapply it:
  - **CRD embedding** (added files): `kube-etl/syncer/config/crd/embed.go` and `multicluster-leader-election/config/crd/bases/embed.go` embed each fork's CRD, so that `config-connector ha setup` installs it. `pkg/cli/cmd/ha` doesn't build without them.
  - **The syncer's Config Connector wildcard skips the system groups:**
    - Files: `kube-etl/syncer/controllers/kcc_glob.go` and `kcc_glob_test.go` (added), and the two calls to `kccGlobMatchesGroup` in `krmsyncer_controller.go` (in `expandRule` and `ruleMatchesGVK`).
    - Change: the `*.cnrm.cloud.google.com` rule no longer matches `core.cnrm.cloud.google.com` and its subgroups (ConfigConnector, ConfigConnectorContext, MultiClusterLease).
    - Why: `ha setup` lets another cluster's syncer read only what the `cnrm-viewer` role covers, which excludes these groups, so the syncer must not watch them.
    - Not in upstream yet.

## Building a fork on its own

From the repository root, `GOWORK=off` stops the repository's `go.work` from taking over:

```shell
GOWORK=off go build -C third_party/github.com/gke-labs/multicluster-leader-election ./...
GOWORK=off go build -C third_party/github.com/gke-labs/kube-etl/syncer ./...
```

## Updating a fork

1. Download the new version outside this module. `Dir` in the output is the module's source:

   ```shell
   (cd "$(mktemp -d)" && go mod download -json github.com/gke-labs/kube-etl/syncer@<version>)
   ```

2. Replace the directory with a copy of `Dir`, and make the copy writable:

   ```shell
   rm -rf third_party/github.com/gke-labs/kube-etl/syncer
   cp -r <Dir> third_party/github.com/gke-labs/kube-etl/syncer
   chmod -R u+w third_party/github.com/gke-labs/kube-etl/syncer
   ```

3. Set the new version in the `require` lines of `go.mod` and `mockgcp/go.mod`, update the table above, and run `make ensure`.
4. Reapply the local changes listed above, and check that the fork still builds on its own (see above).
