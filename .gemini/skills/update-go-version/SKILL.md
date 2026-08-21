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

### 4. Tidy and Format

Once the files are updated, run `go mod tidy` on all modified modules to ensure the dependency graph is fully resolved and compliant with the updated toolchain.
Alternatively, run:
```bash
make ready-pr
```
which runs presubmits, formatting, and tidying.

### 5. Validate

Validate that the codebase builds and tests pass under the updated version configuration.
```bash
go test ./...
```
