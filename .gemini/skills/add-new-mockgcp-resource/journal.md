# Journal: Add New MockGCP Resource

## 2026-05-16

*   Created the skill based on `mockgcp/mockfirestore` pattern.
*   Initial focus is on `MemcacheInstance` (issue #8197).
*   Noted that `httptogrpc` can be used to avoid proto compilation.
*   Emphasized scoping `Previsit` in `normalize.go` to avoid global log corruption.
