# 2026-05-18 - spanner migration
- Moved mockgcp spanner to httptogrpc.
- Spanner admin has two services: DatabaseAdmin and InstanceAdmin. Both were registered using mux.AddService.
- Replaced mux.RewriteHeaders with mux.OverrideHeaders to remove "X-Content-Type-Options".
- Updated imports in service.go, database.go and instance.go to use cloud.google.com/go/spanner/admin/...
