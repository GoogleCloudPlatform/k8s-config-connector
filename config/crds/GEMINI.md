# Directory: config/crds

This directory contains the Custom Resource Definition (CRD) manifests for all the Config Connector resource types.

These CRDs are the authoritative source for the Kubernetes API schema of the KCC resources. They are generated from the Go type definitions in the `apis/` directory.

These manifests are bundled into the main installation YAMLs in `config/installbundle`.

When you need to understand the full schema of a KCC resource, including validation rules and documentation, the CRD files in this directory are the best place to look.

See also the root `GEMINI.md`, `apis/GEMINI.md`, and `config/GEMINI.md`.
