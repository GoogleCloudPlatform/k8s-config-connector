---
name: update-go-version
description: Guides through updating the Go version (go and toolchain) across all modules, workspaces, Dockerfiles, and scripts in Config Connector.
---

# Upgrading Go Version in KCC

This skill outlines the process for updating the Go version in the Config Connector (KCC) codebase.

## Workflow

### 1. Update Go Modules and Workspaces

Update each `go.mod` file and the root `go.work` file.
Specifically, set the `go` directive to the minor version (e.g., `1.26`) and the `toolchain` directive to the precise patch version (e.g., `go1.26.4`).

The affected files include:
*   `go.mod` (root)
*   `go.work` (root)
*   `mockgcp/go.mod`
*   `mockgcp/tools/patch-proto/go.mod`
*   `dev/tools/controllerbuilder/go.mod`
*   `dev/tools/crd-mcp-server/go.mod`
*   `experiments/tools/licensescan/go.mod`
*   `experiments/kubectl-plan/go.mod`
*   `experiments/kompanion/go.mod`

`dev/codebots/update-golang-version` automates steps 1–3 with `go mod edit -go=... -toolchain=...` and `go work edit`, but it also creates its own branch and commits.

If a dependency requires a newer patch release (for example `go 1.27.2`), `go get` raises that module's `go` line to match, and `go.work` must be at least as high as every module it uses.

### 2. Update Dockerfiles

Update base and builder images in Dockerfiles to target the precise patch version (e.g., `golang:1.26.4`).

Key files to check/update:
*   `operator/Dockerfile`
*   `build/tooling/Dockerfile`
*   `build/builder/Dockerfile`
*   `.devcontainer/Dockerfile`

### 3. Update Setup Scripts

Update Go version variables in environment setup scripts.
*   `scripts/environment-setup/golang-setup.sh` (e.g., update `GO_VERSION="1.26.1"` to `"1.26.4"`)

### 4. Update golangci-lint and the CI Workflows (minor upgrades)

golangci-lint refuses to lint a module that targets a newer Go than the one it was built with. When moving to a new minor version (e.g., `1.26` to `1.27`):
*   Bump `version:` in `dev/tasks/generate-github-actions` and `GOLANGCI_LINT_VERSION` in `Makefile` to a golangci-lint release built with the new Go (`golangci-lint version` prints it).
*   Update the hard-coded `go-version:` in `dev/tasks/generate-github-actions`.
*   Run `dev/tasks/generate-github-actions` to regenerate `.github/workflows/`.

A newer golangci-lint can add govet analyzers through `enable-all` in `.golangci.yaml`. Run it with `--max-same-issues=0 --max-issues-per-linter=0` to see every finding, then fix the findings or disable the analyzer.

A new Go minor version can also change the compiler's export data format. The generators that load packages with `golang.org/x/tools/go/packages` (`dev/tools/openapi-to-krm/cmd/generate-mapper`, controllerbuilder's `prunetypes`) then log `export data version N is greater than maximum supported version M` and emit `invalid type` into generated mappers. Bump `golang.org/x/tools` to a release that reads the new format (`go get golang.org/x/tools@<version>` with `GOWORK=off` in the root module, `dev/tools/controllerbuilder` and `mockgcp`), then run `make generate`. `dev/tools/goimports.sh` takes the goimports version from the root `go.mod`, so this also moves `make fmt` to the new goimports.

`go run <tool>@<version>` builds the tool with its own, often older, `golang.org/x/tools`, so tools that load packages are built from `go.mod` instead: `dev/tasks/generate-crds`, the `Makefile` and `operator/Makefile` run `go run sigs.k8s.io/controller-tools/cmd/controller-gen` without `@version`. Go 1.27 also removed `GODEBUG=gotypesalias=0`, which controller-tools v0.16.5 relied on for KCC's type aliases (it panicked in `localNamedToSchema`); v0.17.3 and later resolve aliases with `Alias.Rhs()`. controller-tools v0.18.0 generates identical CRDs; its deepcopy output names an alias's target type instead of the alias.

### 5. Tidy and Format

Once the files are updated, run `go mod tidy -diff` on all modified modules to check that the dependency graph is fully resolved and compliant with the updated toolchain. Previous Go upgrades changed only the `go` and `toolchain` lines; if tidy reports unrelated drift (the root `go.mod` is not kept tidy), leave it out of the upgrade PR.

Run `make fmt` and keep its changes: gofmt output can change between Go versions (Go 1.27 re-aligns trailing comments in some generated `.pb.go` files).

Alternatively, run:
```bash
make ready-pr
```
which runs presubmits, formatting, and tidying.

### 6. Validate

Validate that the codebase builds and tests pass under the updated version configuration.
```bash
go vet ./...
go test ./...
```

New Go versions can add vet checks. `go test` runs a subset of vet (including `printf`), so a new vet finding can fail a package's test build.
