### 2026-06-15 MapManagement MapConfig Greenfield Types Implementation
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `MapManagementMapConfig` (Issue #10284).
- **Problem**: 
  1. The proto definitions are at `google.maps.mapmanagement.v2beta`, which is not compiled by default in KCC's `generate-proto.sh`.
  2. Compiling googleapis from a newer commit (like master or `731d7f2ab6`) globally via `git.versions` introduces breaking changes to other services during `validate-prereqs.sh` (`controller-gen` schema mismatches or parsing failures).
  3. `gcpurls.Template` parses placeholder keys case-insensitively. A snake_case placeholder like `{map_config}` matches a Go struct field containing an underscore (`Map_Config`), but we want to use the cleaner standard `MapConfig`.
- **Solution**: 
  1. Modified `dev/tools/controllerbuilder/generate-proto.sh` to include maps/mapmanagement.
  2. Created an isolated, self-contained `generate.sh` that compiles the specific googleapis commit `d23047f4bae1f796aff726e6e3399292d29804ef` directly to a local, private `mapmanagement.pb` file under `apis/mapmanagement/v1alpha1/`.
  3. Passed this private `.pb` path via `--proto-source-path` to the type/mapper generator.
  4. Used camelCase placeholder `{mapConfig}` in the `gcpurls.Template` definition so it matches the Go field name `MapConfig` case-insensitively.
- **Impact**: Keeps the global `git.versions` completely stable (preventing cross-service contamination), keeps the workspace 100% clean, and guarantees that mapmanagement types can always be generated reliably in any environment.
