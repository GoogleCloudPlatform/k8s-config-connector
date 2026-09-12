# Bulk Regenerated Greenfield CRDs (Post-Processed Sandbox)

This directory contains the post-processed CustomResourceDefinition (CRD) OpenAPI v3 YAML manifests generated during the empirical Greenfield parity evaluation.

## Baseline Reference

* **Evaluation Baseline Date:** 2026-09-02
* **Upstream Master Baseline Commit SHA:** `25aedf2f10ef974e0820eb74db95286338ed66bd` (`origin/master` fork at `5b6602f14bdc913bcf67e359e44069591d3e509a`)
* **Evaluation Scope:** 122 GCP services evaluated via sandboxed package isolation.

## Post-Processing Pipeline

Each CRD YAML in this directory was generated via isolated `controller-gen` and post-processed with Config Connector's canonical CRD tooling pipeline:

1. `go run ./scripts/add-validation-to-crds --dir .build/sandbox-crds/<service>`
2. `go run ./scripts/crd-tools set-field spec.preserveUnknownFields=false --dir .build/sandbox-crds/<service>`
3. `go run ./scripts/crd-tools delete-field status --dir .build/sandbox-crds/<service>`
4. `go run ./scripts/crd-tools set-annotation cnrm.cloud.google.com/version=0.0.0-dev --dir .build/sandbox-crds/<service>`
5. `go run ./scripts/crd-tools delete-annotation controller-gen.kubebuilder.io/version --dir .build/sandbox-crds/<service>`
6. `go run ./scripts/crd-tools reflow-descriptions --dir .build/sandbox-crds/<service>`
7. `go run ./scripts/crd-tools backport-alpha --dir .build/sandbox-crds/<service>`

## Summary Statistics

* **Total Services in Sandbox:** 99
* **Total Generated CRD YAMLs:** 411 files (one file per Kind)
* **Exact Baseline Properties Analyzed (Passing subset):** 10,982 properties
* **Deterministic Exact Matches:** 9,068 properties (82.57%)
* **Deterministic Structural Parity (+ `*Ref` & `SecretRef`):** 9,304 properties (84.72%)
* **Reference Overrides (`*Ref`):** 235
* **Secret Overrides (`SecretRef`):** 1
* **Missing in Generated (Gaps / Domain Overrides):** 1,643
* **Added Upstream Proto Properties:** 2,925
