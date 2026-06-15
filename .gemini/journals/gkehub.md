# GKE Hub Journal

### [2026-06-15] Types Generation and Protos Dependency for GKEHubFleet
- **Context**: Implementing direct types for `GKEHubFleet`.
- **Problem**: GKE Hub Fleets API definitions (`fleet.proto`) were added in Feb 2026, which is newer than the baseline pinned googleapis SHA in `apis/git.versions`. Running `generate-types` under the baseline SHA failed with `proto: not found`.
- **Solution**: Updated `apis/git.versions` to the latest googleapis commit to generate types for `GKEHubFleet` cleanly. Once types and CRD files were generated, reverted `apis/git.versions` to the baseline pin. Since the generated types are static Go code, they compile perfectly fine without needing the newer googleapis protos compiled in the global `.pb` cache.
- **Impact**: Enables implementing greenfield types for resources introduced in newer protos without upgrading the global googleapis pin (which would otherwise break generated mappers for other services like `firestore` and `sql`).
