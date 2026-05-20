# Journal: CloudDMSMigrationJob Identity and Reference Migration

## Observations
- `CloudDMSMigrationJob` is project-scoped and located in `apis/clouddms/v1alpha1/`.
- The CAI resource type is `datamigration.googleapis.com/MigrationJob`.
- The URL format is `projects/{project}/locations/{location}/migrationJobs/{migrationJob}`.
- The resource already had an `ExternalRef` field in its status, which I used for cross-checking in `GetIdentity`.
- `CloudDMSMigrationJobSpec` uses `parent.ProjectAndLocationRef`, which is easily resolved using `refs` package helpers.

## Challenges
- Identifying where the controller was located was slightly tricky as it wasn't in the usual `pkg/controller/direct/clouddms` directory. However, implementing the identity and reference pattern is independent of the controller's immediate presence.

## Design Decisions
- Followed the `ArtifactRegistryRepository` pattern as instructed.
- Used `migrationjob_` prefix for filenames to match existing `connectionprofile_` files in the same directory.
- Verified template initialization with a temporary test to avoid runtime panics.
