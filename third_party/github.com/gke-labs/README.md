# Forks of the multi-cluster leader election components

Copies of the two Go modules that Config Connector's multi-cluster leader election uses:

| Directory | Module | Version | Upstream source |
|---|---|---|---|
| `multicluster-leader-election/` | `github.com/gke-labs/multicluster-leader-election` | `v0.0.0-20260219060047-1ae5cb16f27e` | [gke-labs/multicluster-leader-election@1ae5cb16f27e](https://github.com/gke-labs/multicluster-leader-election/tree/1ae5cb16f27e) |
| `kube-etl/syncer/` | `github.com/gke-labs/kube-etl/syncer` | `v0.0.0-20261027010125-063e4a30cf84` | [gke-labs/kube-etl@063e4a30cf84, `syncer/`](https://github.com/gke-labs/kube-etl/tree/063e4a30cf84/syncer) |

- **Wiring.** `go.mod` and `mockgcp/go.mod` replace both modules with these directories. Their `require` lines keep the upstream version that each copy was taken from.
- **Source.** Each directory is the module exactly as `go mod download` fetches it. When copied, each directory's `h1:` hash matched the module's `go.sum` entry. Module downloads don't record file modes, so no file is executable.
- **Local changes.** None yet. Keep them small, and list each one here so that the next update can reapply it.

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
4. Reapply the local changes listed above.
