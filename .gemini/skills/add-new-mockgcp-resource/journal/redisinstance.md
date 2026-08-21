# Journal: Match MockGCP with Real GCP for RedisInstance

During the matching process for `redis.Instance` (`RedisInstance`), several key insights and practices were identified:

## Volatile Field Normalization in MockGCP (`normalize.go`)
1. **Maintenance Fields**: Real GCP returns dynamic, time-based maintenance fields such as `availableMaintenanceVersions` and `maintenanceVersion`. These fields are volatile and are not currently simulated in MockGCP. They were removed in `ConfigureVisitor` (`replacements.RemovePath`) to ensure consistent test logs.
2. **Operation Done Status**: Real GCP LROs return `"done": false` during operation startup. However, MockGCP LRO JSON formatting omits the default `false` value. Removing the `.done` path in `ConfigureVisitor` successfully aligned the formats.
3. **Location / Zone Selection**: Real GCP can provision resources in dynamic zones depending on active region availability (e.g., `us-central1-f`), whereas MockGCP default behavior hardcodes `-a` suffix (e.g., `us-central1-a`). We normalized these dynamically generated zones to `us-central1-a` in `ConfigureVisitor`.

## Status Normalization in KRM (`tests/e2e/normalize.go`)
To prevent KRM status differences in `_generated_object_*.golden.yaml` due to dynamic runtime IP addresses or zones, we registered custom normalization rules in `tests/e2e/normalize.go`:
- `.status.host` normalized to `"10.20.30.40"`
- `.status.currentLocationId` normalized to `"us-central1-a"`
- `.status.nodes[].zone` normalized to `"us-central1-a"`

These normalizations successfully aligned both real and mock GCP targets, allowing both E2E test runs to pass with zero golden-file discrepancies.
