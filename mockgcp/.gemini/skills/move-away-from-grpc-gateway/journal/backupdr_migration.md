# backupdr migration

Migrated `mockbackupdr` to use `httptogrpc`. Removed `mux.RewriteError` which was used for 404 error manipulation, consistent with the `datastream` migration.
