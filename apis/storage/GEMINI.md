# Directory: apis/storage

This directory contains the Go type definitions for the Config Connector resources related to Google Cloud Storage.

## Structure

This directory follows the standard structure for `apis`, with versioned subdirectories (e.g., `v1beta1`) containing the `*_types.go` files for each Cloud Storage resource (e.g., `storagebucket_types.go`, `storagebucketaccesscontrol_types.go`).

These Go structs define the schema for the KCC resources that manage Cloud Storage resources.

When you are working with Cloud Storage resources in KCC, you will need to import the types from this directory.

The structure of this directory is representative of all the service directories under `apis/`.

See also the root `GEMINI.md` and `apis/GEMINI.md`.
