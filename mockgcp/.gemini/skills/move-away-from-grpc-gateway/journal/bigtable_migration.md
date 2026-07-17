# 2026-05-17 - bigtable migration
- Moved mockgcp bigtable to httptogrpc.
- Bigtable admin protos were under `google/bigtable/admin/v2` in `third_party/googleapis`.
- Registered both `BigtableInstanceAdminClient` and `BigtableTableAdminClient`.
- Removed `RewriteError` as instructed by the skill.
