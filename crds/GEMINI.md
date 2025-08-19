# Directory: crds

This directory contains the raw YAML definitions for the Config Connector Custom Resource Definitions (CRDs).

These files are generated from the Go type definitions in the `apis/` directory using `controller-gen`. They should not be edited manually.

These CRDs are applied to a Kubernetes cluster to install the KCC resource types. They are bundled into the installation manifests in the `config/installbundle` directory.

When you need to see the exact schema for a KCC resource, you can look at the corresponding CRD file in this directory.

See also the root `GEMINI.md` and the `apis/GEMINI.md` for more context.
