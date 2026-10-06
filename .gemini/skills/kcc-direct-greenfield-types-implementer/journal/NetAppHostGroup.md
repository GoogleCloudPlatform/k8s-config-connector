# NetAppHostGroup Journal

## Observations & Patterns

### 1. Proto Source Pinning
- `HostGroup` was added in a newer commit of googleapis (`e09e85d32`) than the repository's default pinned version in `apis/git.versions`.
- `apis/netapp/generate.sh` was updated to compile proto descriptors using commit `e09e85d32` to `.build/googleapis-netapp.pb` via `./generate-proto.sh "e09e85d32" "${REPO_ROOT}/.build/googleapis-netapp.pb"` and pass `--proto-source-path` to `controllerbuilder`.

### 2. Acronym Naming Convention
- Proto field `os_type` translates to Go struct field `OSType` (`json:"osType,omitempty"`) rather than `OsType` to align with KCC acronym conventions and automated mapper generation.

### 3. Identity and URL Template
- The GCP URL format for `HostGroup` is `projects/{project}/locations/{location}/hostGroups/{hostGroup}`.
- Because `netapp.googleapis.com/HostGroup` is not present in `cloudassetinventory_names.jsonl`, an exception was added to `pkg/gcpurls/registry_test.go` (`//netapp.googleapis.com/projects/{}/locations/{}/hostGroups/{}`).
