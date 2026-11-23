### [2026-11-11] Prepopulating DesignCenterSpace
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DesignCenterSpace (Issue #15404).
- **Problem**: The proto for `google.cloud.designcenter.v1` is not present in the default pinned `googleapis` commit in `apis/git.versions`. In addition, adding `cloud.google.com/go/designcenter v0.1.0` directly from upstream pulls in newer `genproto` and `google.golang.org/api` dependencies that cause transitive dependency upgrades across dozens of client libraries, dirtying `mockgcp/go.mod` and breaking roundtrip fuzzers and golden logs.
- **Solution**: Pinned `PROTO_SHA="547f8fdb4b91dbac70d540b3757520f0608ae356"` in `apis/designcenter/generate.sh` and vendored `cloud.google.com/go/designcenter` under `third_party/cloud.google.com/go/designcenter` with a `replace` directive in `go.mod` to isolate its module requirements and avoid cascading dependency upgrades.
- **Impact**: Ensures `DesignCenter` proto generation and client dependencies compile cleanly without breaking other controllers, fuzzers, or golden tests across the repository.

### [2026-11-12] Triaging judgement queue and finalizing KRM types for DesignCenterSpace
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterSpace` (Issue #15454).
- **Resolutions & Findings**:
  - `location-or-parent-ref`: Space is a regional project-parented resource (`projects/{project}/locations/{location}/spaces/{space}`). Inlined `*parent.ProjectAndLocationRef` in `DesignCenterSpaceSpec`.
  - `possible-reference-by-description` on `.spec.gcsBucket`: Converted `gcsBucket` to `gcsBucketRef` (`*storagev1beta1.StorageBucketRef`) referencing Cloud Storage bucket.
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces`.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.

### [2026-11-18] Greenfield: Implement direct controller, E2E fixtures, and fuzzer for DesignCenterSpace
- **Context**: Implemented direct controller, fuzzer, and minimal/maximal E2E fixtures for `DesignCenterSpace` (Issue #15981).
- **Findings & Solutions**:
  - **API Preconditions**: Creating a `Space` in GCP Design Center requires the GCP project to be designated as an App Hub Management Project (`gcloud apphub boundary update --project=PROJECT_ID --location=global --crm-node=projects/PROJECT_ID`). Without this boundary setup, creation fails with `FAILED_PRECONDITION: Invalid resource state: Current project is not a management project`.
  - **Field Mapping (Near Miss)**: `EnableGCPSharedTemplates` in KRM versus `EnableGcpSharedTemplates` in proto caused a near-miss during auto-mapping generation. Implemented custom mapping in `pkg/controller/direct/designcenter/designcenterspace_mapper.go` to cleanly map `EnableGCPSharedTemplates` and enable full roundtrip fuzzer coverage.
  - **StorageBucket Reference Formatting**: GCP Design Center API expects bare bucket name `{$bucket_name}` in `Space.gcs_bucket` instead of the full resource identity URI `projects/{project}/buckets/{bucket}`. Added normalization in mapper to extract the bucket name for proto and prepend the canonical identity URI when reading from proto.
  - **CRUD Operations**: The Space API uses synchronous CRUD methods (`CreateSpace`, `UpdateSpace`, `DeleteSpace`) rather than LROs. Controller handles synchronous responses directly.
  - **Fixtures & Golden Output**: Created minimal and maximal fixtures (`designcenterspace-minimal`, `designcenterspace-maximal`) and successfully recorded them against real GCP (`cnrm-barni-4`). Re-reconciliation, Server-Side Apply, and roundtrip fuzzers passed cleanly.

### [2026-11-18] Prepopulating DesignCenterCatalog
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DesignCenterCatalog (Issue #16049).
- **Findings & Solution**:
  - Configured `apis/designcenter/generate.sh` with dedicated `--resource DesignCenterCatalog:Catalog` under `generate-types` and `generate-identity`.
  - Mechanically generated `designcentercatalog_types.go`, `designcentercatalog_identity.generated.go`, `designcentercatalog_reference.generated.go`, `designcentercatalog_identity_generated_test.go`, CRD manifest, deepcopy, and `pkg/controller/direct/designcenter/mapper.generated.go`.
  - Added `DesignCenterCatalog` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Initialized open entries in `apis/designcenter/judgement_queue.yaml` for `DesignCenterCatalog` (including `untriaged-bulk-generation` with `status: open`).
  - Recorded untriaged missing fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-19] Triaging judgement queue and finalizing KRM types for DesignCenterCatalog
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterCatalog` (Issue #16066).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: DesignCenterCatalog is a child resource under DesignCenterSpace (`projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `SpaceRef` with required, immutable `ParentRef *DesignCenterSpaceRef` (`json:"parentRef"`).
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.catalogs`.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.
  - Updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-19] Greenfield: Implement direct controller, E2E fixtures, and fuzzer for DesignCenterCatalog
- **Context**: Implemented direct controller, round-trip fuzzer, and minimal/maximal E2E fixtures for `DesignCenterCatalog` (Issue #16093).
- **Findings & Solutions**:
  - **Generation**: Generated direct controller (`pkg/controller/direct/designcenter/designcentercatalog_controller.go`) and fuzzer (`pkg/controller/direct/designcenter/designcentercatalog_fuzzer.generated.go`) with one-shot `controllerbuilder` commands without adding them to `apis/designcenter/generate.sh`.
  - **Operations**: `CreateCatalog`, `GetCatalog`, and `UpdateCatalog` are synchronous operations, while `DeleteCatalog` returns an LRO. The generated controller cleanly handles the delete LRO waiting.
  - **E2E Dependencies & Recording**: `DesignCenterCatalog` requires a parent `DesignCenterSpace` in `dependencies.yaml`. Recorded authentic golden logs against real GCP (`cnrm-barni-4`) with `RECORD_AUDIT_PROBE=1 ./hack/record-gcp` for both `designcentercatalog-minimal` and `designcentercatalog-maximal`.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterCatalog`), CRD field presence checks, and re-reconciliation passed cleanly.

### [2026-11-19] Prepopulating DesignCenterCatalogTemplate
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DesignCenterCatalogTemplate (Issue #16143).
- **Findings & Solution**:
  - Configured `apis/designcenter/generate.sh` with dedicated `--resource DesignCenterCatalogTemplate:CatalogTemplate` under `generate-types` and `generate-identity`.
  - Mechanically generated `designcentercatalogtemplate_types.go`, `designcentercatalogtemplate_identity.generated.go`, `designcentercatalogtemplate_reference.generated.go`, `designcentercatalogtemplate_identity_generated_test.go`, CRD manifest, deepcopy, and updated `pkg/controller/direct/designcenter/mapper.generated.go`.
  - Added `DesignCenterCatalogTemplate` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Initialized open entries in `apis/designcenter/judgement_queue.yaml` for `DesignCenterCatalogTemplate` (including `untriaged-bulk-generation` with `status: open`).
  - Recorded untriaged missing fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-19] Triaging judgement queue and finalizing KRM types for DesignCenterCatalogTemplate
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterCatalogTemplate` (Issue #16155).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: DesignCenterCatalogTemplate is a child resource under DesignCenterCatalog (`projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}/templates/{catalog_template}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `CatalogRef` with required, immutable `ParentRef *DesignCenterCatalogRef` (`json:"parentRef"`).
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}/templates/{catalog_template}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.catalogs.templates`.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.
  - Updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-19] Greenfield: Implement direct controller, E2E fixtures, and fuzzer for DesignCenterCatalogTemplate
- **Context**: Implemented direct controller, round-trip fuzzer, and minimal/maximal E2E fixtures for `DesignCenterCatalogTemplate` (Issue #16169).
- **Findings & Solutions**:
  - **Generation**: Generated direct controller (`pkg/controller/direct/designcenter/designcentercatalogtemplate_controller.go`) and fuzzer (`pkg/controller/direct/designcenter/designcentercatalogtemplate_fuzzer.generated.go`) with one-shot `controllerbuilder` commands without modifying `apis/designcenter/generate.sh`.
  - **Operations**: `CreateCatalogTemplate`, `GetCatalogTemplate`, and `UpdateCatalogTemplate` are synchronous operations, while `DeleteCatalogTemplate` returns an LRO. The generated controller handles waiting for delete LRO completion.
  - **E2E Dependencies & Recording**: `DesignCenterCatalogTemplate` requires an ancestor chain of `DesignCenterSpace` and `DesignCenterCatalog` in `dependencies.yaml`. Recorded authentic golden logs against real GCP (`cnrm-barni-4`) with `RECORD_AUDIT_PROBE=1 ./hack/record-gcp` for both `designcentercatalogtemplate-minimal` and `designcentercatalogtemplate-maximal`.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterCatalogTemplate`), CRD field presence checks, and re-reconciliation passed cleanly.

### [2026-11-20] Prepopulating DesignCenterCatalogTemplateRevision
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DesignCenterCatalogTemplateRevision (Issue #16238).
- **Findings & Solution**:
  - Configured `apis/designcenter/generate.sh` with dedicated `--resource DesignCenterCatalogTemplateRevision:CatalogTemplateRevision` under `generate-types` and `generate-identity`.
  - Added handwritten mapper helper functions in `pkg/controller/direct/designcenter/designcentercatalogtemplaterevision_mapper.go` for `*Metadata` messages skipped by `generate-mapper` (`TfBlueprintMetadataObservedState`, `HelmChartMetadataObservedState`, `InferredMetadataObservedState`, `TerraformBlueprintUiMetadataObservedState`) and `Value_ToProto` for `structpb.Value`.
  - Mechanically generated `designcentercatalogtemplaterevision_types.go`, `designcentercatalogtemplaterevision_identity.generated.go`, `designcentercatalogtemplaterevision_reference.generated.go`, `designcentercatalogtemplaterevision_identity_generated_test.go`, CRD manifest, deepcopy, and updated `pkg/controller/direct/designcenter/mapper.generated.go`.
  - Added `DesignCenterCatalogTemplateRevision` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Initialized open entries in `apis/designcenter/judgement_queue.yaml` for `DesignCenterCatalogTemplateRevision` (including `untriaged-bulk-generation` with `status: open`).
  - Added `designcentercatalogtemplaterevisions.designcenter.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` due to empty object schema from `TerraformUiInputs` map value type.
  - Defined recursive types `ListValue` and `DisplayVariable` in `designcentercatalogtemplaterevision_types.go` with kubebuilder unknown fields annotations (`+kubebuilder:validation:items:XPreserveUnknownFields`, `+kubebuilder:validation:items:Type=object` on `ListValue.Values`, and `+kubebuilder:pruning:PreserveUnknownFields`, `+kubebuilder:validation:Schemaless`, `+kubebuilder:validation:XPreserveUnknownFields`, `+kubebuilder:validation:Type=object` on `DisplayVariable.Properties`) so `controller-gen` outputs valid OpenAPI schemas that satisfy Kubernetes OpenAPI v3 validation in envtest (`TestPreview`).
  - Recorded untriaged missing fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`, recursive types in `recursivetypes.txt`, and not representable refs in `refs_not_representable.txt`.

### [2026-11-20] Triaging judgement queue and finalizing KRM types for DesignCenterCatalogTemplateRevision
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterCatalogTemplateRevision` (Issue #16250).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: `DesignCenterCatalogTemplateRevision` is a child resource under `DesignCenterCatalogTemplate` (`projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}/templates/{catalog_template}/revisions/{revision}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `TemplateRef` with required, immutable `ParentRef *DesignCenterCatalogTemplateRef` (`json:"parentRef"`).
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}/templates/{catalog_template}/revisions/{revision}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.catalogs.templates.revisions`.
  - `authoritative-proto-reference`: Resolved `.spec.applicationTemplateRevisionSource` as `not-applicable` because `ApplicationTemplateRevision` is an unmanaged resource in KCC. Moved `DeveloperConnectSourceConfig` into `designcentercatalogtemplaterevision_types.go` and converted `.spec.developerConnectSourceConfig.developerConnectRepoURI` to `DeveloperConnectRepoRef *developerconnectv1alpha1.DeveloperConnectGitRepositoryLinkRef` (`resolution: edited`).
  - Added handwritten mappers `DeveloperConnectSourceConfig_FromProto` and `DeveloperConnectSourceConfig_ToProto` in `pkg/controller/direct/designcenter/designcentercatalogtemplaterevision_mapper.go`.
  - `deprecated-field`: Resolved `.spec.gcsSourceURI` (omitted from spec) and `.spec.gitSource.developerConnectConfig` (in nested `GitSource`) as `accepted`.
  - `message-map-derived`: Verified Go map types and mappings for `DisplayVariable`, `TemplateUiOutputDisplayVariable`, `HelmChartInput`, `TerraformUiInputDisplayVariable`, and `TerraformUiOutputDisplay`, and marked `accepted`.
  - `reference-not-representable`: Resolved `.spec.ociRepo.uri` as `accepted` (container image URI not representable as KCC reference today).
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.
  - Updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-23] Finalize Step 1b compliance for DesignCenterCatalogTemplateRevision
- **Context**: Addressed Step 1b review feedback on `DesignCenterCatalogTemplateRevision` (Issue #16241).
- **Resolutions & Findings**:
  - `deprecated-field` on `.spec.gitSource.developerConnectConfig`: Handwrote `GitSource` struct in `designcentercatalogtemplaterevision_types.go` without the deprecated `DeveloperConnectConfig` field, as `GitSource` is exclusive to `DesignCenterCatalogTemplateRevision` and superseded by `developerConnectSourceConfig`. Updated judgement queue resolution to `edited`.
  - `well-known-proto-type`: Resolved all 18 open `google.protobuf.Value` entries in `apis/designcenter/judgement_queue.yaml` as `accepted` with explanatory notes confirming generated `struct Value` KRM representation and `Value_ToProto` mapper.
  - Regenerated types, deepcopy, mappers, and CRD manifests via `./apis/designcenter/generate.sh`.
  - Removed stale `.spec.gitSource.developerConnectConfig.gitRepositoryLink` exception from `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.









