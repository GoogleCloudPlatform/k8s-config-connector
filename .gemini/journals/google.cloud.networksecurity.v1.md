### [2026-05-27] Scaffolded NetworkSecurityMirroringDeploymentGroup
- **Context**: Scaffolding Greenfield NetworkSecurityMirroringDeploymentGroup (Issue #8731)
- **Problem**: The provided GCP APIs SHA (`731d7f2ab6`) did not contain the proto for `MirroringDeploymentGroup`. Running the scaffolder directly failed with `proto: not found`.
- **Solution**: The `update-gcp-dependencies` task fetched a newer SHA from master (`7496288011d66f2b34be84377500d114dc74e006`) that successfully included `MirroringDeploymentGroup`. I updated `apis/git.versions` to point to the commit where it was available, and re-ran the scaffolder, which correctly pulled and parsed the proto files.
- **Impact**: When scaffolding Greenfield resources and encountering `proto: not found` for the provided APIs SHA, agents should bump `apis/git.versions` to include the target API definition or pull the latest from `googleapis`.

### [2026-05-27] Using identity.IdentityV2 pattern
- **Context**: Implementing IdentityV2 for NetworkSecurityMirroringDeploymentGroup
- **Problem**: KCC Direct identities must use the new `identity.IdentityV2` and `refs.Ref` interface instead of the older unstructured methods.
- **Solution**: Followed the `workstationcluster_identity.go` pattern, which uses `gcpurls.Template` tied to `identity.IdentityV2`. I also created a modern `refs.Ref` implementation in `networksecuritymirroringdeploymentgroup_reference.go` leveraging `refs.NormalizeWithFallback`.
- **Impact**: Provides a template for direct controller identity implementation.
