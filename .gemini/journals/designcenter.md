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

### [2026-11-24] Greenfield Step 2: Implement direct controller, fuzzer, and E2E fixtures for DesignCenterCatalogTemplateRevision
- **Context**: Greenfield implementation of direct controller, round-trip fuzzer, and real-GCP recorded E2E fixtures for `DesignCenterCatalogTemplateRevision` (`designcenter.cnrm.cloud.google.com/v1alpha1`, Issue #16348).
- **Resolutions & Findings**:
  - **Controller and Fuzzer**: Generated with `controllerbuilder generate-controller` and `generate-fuzzer --prepopulate-resource`. The GCP API does not support `UpdateCatalogTemplateRevision` (revisions are immutable), and controllerbuilder generated an update handler rejecting non-empty spec field updates.
  - **GCP API Artifact Requirements**: CatalogTemplateRevision creation requires an `artifact_location` oneof (e.g. `gitSource`) pointing to a valid Terraform module repository containing `metadata.yaml` with valid `spec.info.actuationTool` configurations (such as `GoogleCloudPlatform/terraform-google-cloud-run` at ref tag `v0.34.2`).
  - **Fixtures & Golden Logs**: Recorded E2E fixtures against real GCP (`cnrm-barni-4`) using `./hack/record-gcp "fixtures/designcentercatalogtemplaterevision-minimal"` and `"fixtures/designcentercatalogtemplaterevision-maximal"`. Audit probe markers and Cloud Audit Logs confirm successful actuation and deletion.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterCatalogTemplateRevision`), CRD field presence checks, and presubmit unit tests all passed cleanly.

### [2026-11-19] Prepopulating DesignCenterApplicationTemplate
- **Context**: Prepopulating direct KRM types, identity, reference, mapper, and judgement queue for `DesignCenterApplicationTemplate` (Issue #16161).
- **Generation & Mapper Setup**:
  - Added `--resource DesignCenterApplicationTemplate:ApplicationTemplate` to `generate-types` and `generate-identity` in `apis/designcenter/generate.sh`.
  - Appended `google.cloud.apphub.v1` to `--service` on `generate-mapper` in `apis/designcenter/generate.sh` to automatically emit mappers for AppHub types referenced in `AppHubApplicationParameters` (`Scope`, `Attributes`, `Criticality`, `Environment`, `ContactInfo`).
  - Added `DesignCenterApplicationTemplate` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` with golden output for `DesignCenterApplicationTemplate`.

### [2026-11-25] Triaging judgement queue and finalizing KRM types for DesignCenterApplicationTemplate
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterApplicationTemplate` (Issue #16442).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: `DesignCenterApplicationTemplate` is a child resource under `DesignCenterSpace` (`projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `SpaceRef` with required, immutable `ParentRef *DesignCenterSpaceRef` (`json:"parentRef"`).
  - `possible-reference-by-description` on `.spec.saasRuntimeContext.saasNames`: Field format is `projects/{project}/locations/{location}/saas/{saas}`. To avoid breaking CRD backward compatibility on existing kinds sharing `SaaSRuntimeContext` in status/observedState (`DesignCenterCatalogTemplateRevision`), forked `ApplicationTemplate_SaaSRuntimeContext` into `designcenterapplicationtemplate_types.go` and converted `saasNames []string` to `SaasRefs []saasservicemgmtv1alpha1.SaaSServiceMgmtSaaSRef` (`json:"saasRefs,omitempty"`). Implemented handwritten bidirectional mappers and roundtrip unit tests in `pkg/controller/direct/designcenter/designcenterapplicationtemplate_mapper.go`.
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applicationTemplates`.
  - `well-known-proto-type`: Confirmed `google.protobuf.Value` representation and resolved as accepted.
  - `message-map-derived`: Confirmed map types and mappings for Terraform UI input/output display types and resolved as accepted.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.
  - Updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-25] Greenfield Step 2: Implement direct controller, fuzzer, and E2E fixtures for DesignCenterApplicationTemplate
- **Context**: Greenfield implementation of direct controller, round-trip fuzzer, and real-GCP recorded E2E fixtures for `DesignCenterApplicationTemplate` (`designcenter.cnrm.cloud.google.com/v1alpha1`, Issue #16450).
- **Resolutions & Findings**:
  - **Controller and Fuzzer**: Generated using `controllerbuilder generate-controller` and `generate-fuzzer --prepopulate-resource`. The controller registers with `static_config.go` and uses the standard REST GAPIC client with `NormalizeGreenfieldReferences` and `CompareGreenfieldSpec`.
  - **Fixtures & Golden Logs**: Created `designcenterapplicationtemplate-minimal` and `designcenterapplicationtemplate-maximal` fixtures under parent `DesignCenterSpace` (`projects/${projectId}/locations/us-central1/spaces/dc-space-${uniqueId}`). Recorded against real GCP (`cnrm-barni-4`) using `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`.
  - **Lifecycle & Updates**: Verified Create, Update (`displayName`, `description`), and Delete phases against real GCP, passing all 7/7 HTTP report validation checks and re-reconciliation testing.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterApplicationTemplate`), CRD field presence checks, and presubmit unit tests all passed cleanly.

### [2026-11-25] Prepopulating DesignCenterPolicy
- **Context**: Prepopulating direct KRM types, identity, reference, mapper, and judgement queue for `DesignCenterPolicy` (Issue #16469).
- **Generation & Ordering**:
  - Added `--resource DesignCenterPolicy:Policy` to `generate-types` and `generate-identity` in `apis/designcenter/generate.sh`.
  - Ordered `DesignCenterPolicy:Policy` first among resources in `generate-types` and `generate-identity` so that downstream resources (`DesignCenterCatalogTemplateRevision` and `DesignCenterApplicationTemplate`) that traverse shared types `Parameter` and `CloudControlConfig` with `forceAll == true` are visited after and preserve all output fields during `deduplicateAndSortOutputMessages`, preventing CRD backward compatibility failures on existing kinds.
  - Added `DesignCenterPolicy` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` with golden output for `DesignCenterPolicy`.

### [2026-11-26] Greenfield Step 1b: Triaging judgement queue and finalizing KRM types for DesignCenterPolicy
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterPolicy` (Issue #16490).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: `DesignCenterPolicy` is a child resource under `DesignCenterApplicationTemplate` (`projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}/policies/{policy}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `ApplicationTemplateRef` with required, immutable `ParentRef *DesignCenterApplicationTemplateRef` (`json:"parentRef"`).
  - `authoritative-proto-reference` on `.spec.cloudControlConfigs[].cloudControl`: Replaced with `CloudControlRef *cloudsecuritycompliancev1alpha1.CloudSecurityComplianceCloudControlRef` (`json:"cloudControlRef,omitempty"`) in `CloudControlConfig`, moving `CloudControlConfig` to `designcenterpolicy_types.go`.
  - `INPUT_ONLY` field placed in status: Removed `CloudControlConfigs` from `DesignCenterPolicyObservedState` per `REFERENCES_AND_NESTED_STRUCTS.md` because `cloud_control_configs` is marked `INPUT_ONLY` in proto and not returned on GET.
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}/policies/{policy}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applicationTemplates.policies`.
  - `well-known-proto-type`: Confirmed `google.protobuf.Value` representation and resolved as accepted.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.

### [2026-11-25] Prepopulating DesignCenterApplication
- **Context**: Prepopulating direct KRM types, identity, reference, mapper, and judgement queue for `DesignCenterApplication` (Issue #16468).
- **Generation & Mapper Setup**:
  - Added `--resource DesignCenterApplication:Application` to `generate-types` and `generate-identity` in `apis/designcenter/generate.sh`.
  - Appended `google.cloud.config.v1` to `--service` on `generate-mapper` in `apis/designcenter/generate.sh` to emit mappers for `google.cloud.config.v1.TerraformError` referenced in `DeploymentErrorObservedState`.
  - Handwrote `DeploymentMetadataObservedState` and `DeploymentGroupMetadataObservedState` mappers in `pkg/controller/direct/designcenter/designcenterapplication_mapper.go` to bridge output-only types pruned from spec/input mappings.
  - Added `DesignCenterApplication` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Added `designcenterapplications.designcenter.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` due to empty object schema on `status.observedState.compositeApplicationParameters.projectParametersMap` (from proto `ProjectParametersList` where all fields are output-only).
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` with golden output for `DesignCenterApplication`.

### [2026-11-26] Greenfield Step 1b: Triaging judgement queue and finalizing KRM types for DesignCenterApplication
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, reference, and mappers for `DesignCenterApplication` (Issue #16508).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: `DesignCenterApplication` is a child resource under `DesignCenterSpace` (`projects/{project}/locations/{location}/spaces/{space}/applications/{application}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `SpaceRef` with required, immutable `ParentRef *DesignCenterSpaceRef` (`json:"parentRef"`).
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/applications/{application}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applications`.
  - `serviceAccount`: Converted `ServiceAccount` to `ServiceAccountRef *refsv1beta1.IAMServiceAccountRef` (`json:"serviceAccountRef,omitempty"`), verified auto-generated proto mapping in `mapper.generated.go`.
  - `message-map-derived` & empty schema fix: Defined `ProjectParametersList` and `DeploymentMetadata` in `designcenterapplication_types.go` with their output-only fields (`project_parameters` and output-only metadata fields) to avoid empty object schemas in OpenAPI/CRD generation. Implemented `DeploymentMetadata_FromProto`/`ToProto` in `designcenterapplication_mapper.go`. This resolved the empty object schema and allowed removing `designcenterapplications` from `knownInvalidCRDs` in `tests/apichecks/crds_test.go`.
  - `possible-reference-by-description` on `serviceAccountMap` and `kubernetesServiceAccount`: Deferred to `tests/apichecks/testdata/exceptions/refs_deferred.txt`: `serviceAccountMap` is a `map[string]string` mapping component URIs to SA emails, and `kubernetesServiceAccount` is a Kubernetes SA within the target GKE namespace.
  - `possible-reference-by-name` on `rolesAPIsConfig[].projectID` and `authoritative-proto-reference` on `roles[].role`: Resolved as `not-applicable` because `rolesAPIsConfig[].projectID` specifies the target project ID string, and `role` is an IAM role string name (`roles/editor`).
  - `well-known-proto-type`: Confirmed `google.protobuf.Value` representation across 18 entries and marked resolved/accepted.
  - `untriaged-bulk-generation`: Resolved with resolution `accepted`.

### [2026-11-26] Greenfield Step 2: Implement direct controller, fuzzer, and E2E fixtures for DesignCenterPolicy
- **Context**: Greenfield implementation of direct controller, round-trip fuzzer, and real-GCP recorded E2E fixtures for `DesignCenterPolicy` (`designcenter.cnrm.cloud.google.com/v1alpha1`, Issue #16505).
- **Resolutions & Findings**:
  - **Controller and Fuzzer**: Generated using `controllerbuilder generate-controller` and `generate-fuzzer --prepopulate-resource`. The controller registers with `static_config.go` and uses the standard REST GAPIC client with `NormalizeGreenfieldReferences` and `CompareGreenfieldSpec`.
  - **API Behavior on Create**: Real GCP requires a non-empty URI (`policyURI` or `policyRevisionURI`) even though protobuf marked it optional; creating a policy without URI fails with HTTP 400 Bad Request (`The request was invalid: invalid policy: uri cannot be empty: invalid argument`). Used `organizations/12345/locations/global/frameworks/test-framework` in test fixtures.
  - **Fixtures & Golden Logs**: Created `designcenterpolicy-minimal` and `designcenterpolicy-maximal` fixtures under parent `DesignCenterApplicationTemplate` and grandparent `DesignCenterSpace` (`projects/${projectId}/locations/us-central1/spaces/dc-space-${uniqueId}`). Recorded against real GCP (`cnrm-barni-4`) using `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`.
  - **Lifecycle & Updates**: Verified Create, Update (`displayName`, `description`, `applyConditions`), and Delete phases against real GCP, passing all 7/7 HTTP report validation checks and re-reconciliation testing. Both fixtures produced rich `_audit_probe.log` entries with authentic service audit probe data.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterPolicy`), `TestCRDFieldPresenceInTestsForAlpha`, and presubmit unit tests all passed cleanly.

### [2026-11-26] Greenfield Step 2: Implement direct controller, fuzzer, and E2E fixtures for DesignCenterApplication
- **Context**: Greenfield implementation of direct controller, round-trip fuzzer, and real-GCP recorded E2E fixtures for `DesignCenterApplication` (`designcenter.cnrm.cloud.google.com/v1alpha1`, Issue #16521).
- **Resolutions & Findings**:
  - **Controller and Fuzzer**: Generated using `controllerbuilder generate-controller` and `generate-fuzzer --prepopulate-resource`. The controller registers with `static_config.go` and uses the standard REST GAPIC client with `NormalizeGreenfieldReferences` and `CompareGreenfieldSpec`.
  - **GCP API Nuances**:
    - Creation of `Application` requires `scope.type` and `deploymentRegion` (required for regional-scoped applications).
    - Source was configured using a shared template revision URI from the Google Catalog (`googlespace`), e.g. `projects/gcpdesigncenter/locations/us-central1/spaces/googlespace/sharedTemplates/googlespace-simple-agent-runtime-1kqzt4m/revisions/r-2`, under a parent `DesignCenterSpace` with `spec.enableGCPSharedTemplates: true`.
    - Delete operation returns an LRO, which the controller waits on via `op.Wait(ctx)`.
    - GCP returns `projectParameters[].roles` and `rolesApisConfig[].roles` in non-deterministic order. Added sort normalizers for `.projectParameters[].roles` and `.rolesApisConfig[].roles` (and in KRM status `.status.observedState.projectParameters[].roles` / `rolesAPIsConfig[].roles`) to ensure deterministic golden files.
  - **Fixtures & Golden Logs**: Created `designcenterapplication-minimal` and `designcenterapplication-maximal` fixtures under parent `DesignCenterSpace`. Recorded against real GCP (`cnrm-barni-4`) using `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`. Both minimal and maximal passed all 7/7 HTTP report validation checks and re-reconciliation testing.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterApplication`), CRD field presence checks, and presubmit unit tests all passed cleanly.

### [2026-11-26] Prepopulating DesignCenterComponent
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DesignCenterComponent (Issue #16541).
- **Findings & Solution**:
  - Configured `apis/designcenter/generate.sh` with dedicated `--resource DesignCenterComponent:Component` under `generate-types` and `generate-identity`.
  - Mechanically generated `designcentercomponent_types.go`, `designcentercomponent_identity.generated.go`, `designcentercomponent_reference.generated.go`, `designcentercomponent_identity_generated_test.go`, CRD manifest, deepcopy, and updated `pkg/controller/direct/designcenter/mapper.generated.go`.
  - Added `DesignCenterComponent` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` in strict alphabetical order.
  - Initialized open entries in `apis/designcenter/judgement_queue.yaml` for `DesignCenterComponent` (including `untriaged-bulk-generation` with `status: open`).
  - Recorded untriaged missing fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-26] Greenfield Step 1b: Triaging judgement queue and finalizing KRM types for DesignCenterComponent
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterComponent` (Issue #16544).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: `DesignCenterComponent` is a child resource under `DesignCenterApplicationTemplate` (`projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}/components/{component}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `ApplicationTemplateRef` with required, immutable `ParentRef *DesignCenterApplicationTemplateRef` (`json:"parentRef"`).
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}/components/{component}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applicationTemplates.components`.
  - `possible-reference-by-name` on `.spec.applicationInfo.apphubApplicationID`: Resolved as `not-applicable` because it is a user-defined ID string for the created App Hub application, not a reference to an existing resource.
  - `well-known-proto-type`: Confirmed `google.protobuf.Value` representation across parameters and observedState schema entries and marked resolved/accepted.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.
  - Updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-27] Greenfield Step 2: Implement direct controller, E2E fixtures, and fuzzer for DesignCenterComponent
- **Context**: Autogenerated direct controller and round-trip fuzzer, created minimal and maximal E2E fixtures, and recorded golden HTTP logs against real GCP for `DesignCenterComponent` (Issue #16552).
- **Resolutions & Findings**:
  - **Controller and Fuzzer**: Generated using `controllerbuilder generate-controller` and `generate-fuzzer --prepopulate-resource`. The controller registers with `static_config.go` and uses the standard REST GAPIC client with `NormalizeGreenfieldReferences` and `CompareGreenfieldSpec`.
  - **GCP API Nuances**:
    - Creation of `Component` under `ApplicationTemplate` requires the parent `DesignCenterApplicationTemplate` to have `compositionType: COMPOSITE`. Standard application templates reject components with `400 Bad Request: Application template ... has a STANDARD composition type and cannot contain a component of type STANDARD_APPLICATION_TEMPLATE`.
    - Ancestor `DesignCenterSpace` must have `enableGCPSharedTemplates: true` when referencing shared templates from the Google catalog (e.g. `projects/gcpdesigncenter/locations/us-central1/spaces/googlespace/sharedTemplates/googlespace-simple-agent-runtime-1kqzt4m/revisions/r-2`).
    - Synchronous operations: `CreateComponent`, `UpdateComponent`, and `DeleteComponent` are synchronous RPCs (non-LRO).
  - **Fixtures & Golden Logs**: Created `designcentercomponent-minimal` and `designcentercomponent-maximal` fixtures with parent `DesignCenterApplicationTemplate` (COMPOSITE) and ancestor `DesignCenterSpace`. Recorded against real GCP (`cnrm-barni-4`) using `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`. Both minimal and maximal passed all 7/7 HTTP report validation checks and re-reconciliation testing.
  - **Validation**: 200-seed focused round-trip fuzzer (`TestFocusedMappers/DesignCenterComponent`), CRD field presence checks, and presubmit unit tests all passed cleanly.

### [2026-11-27] Prepopulating DesignCenterConnection
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DesignCenterConnection (Issue #16566).
- **Findings & Solution**:
  - Configured `apis/designcenter/generate.sh` with dedicated `--resource DesignCenterConnection:Connection` under `generate-types` and `generate-identity`.
  - Mechanically generated `designcenterconnection_types.go`, `designcenterconnection_identity.generated.go`, `designcenterconnection_reference.generated.go`, `designcenterconnection_identity_generated_test.go`, CRD manifest, deepcopy, and updated `pkg/controller/direct/designcenter/mapper.generated.go`.
  - To prevent CRD backward-compatibility breakages for existing DesignCenter CRDs (`DesignCenterApplication`, `DesignCenterApplicationTemplate`, `DesignCenterCatalogTemplateRevision`, `DesignCenterComponent`), declared `ParameterObservedState` in `designcenterpolicy_types.go` with `Key`, `Value`, and `ReadOnly` fields so `controllerbuilder` preserves those fields in `types.generated.go`.
  - Added `DesignCenterConnection` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` in strict alphabetical order.
  - Initialized open entries in `apis/designcenter/judgement_queue.yaml` for `DesignCenterConnection` (`untriaged-bulk-generation` with `status: open`).
  - Recorded untriaged missing fields in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-27] Greenfield Step 1b: Triaging judgement queue and finalizing KRM types for DesignCenterConnection
- **Context**: Triaged `apis/designcenter/judgement_queue.yaml` and finalized KRM types, identity, and reference types for `DesignCenterConnection` (Issue #16571).
- **Resolutions & Findings**:
  - `location-or-parent-ref`, `parent-ref-guessed`, `identity-root-unknown`: `DesignCenterConnection` is a child resource under `DesignCenterComponent` (`projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}/components/{component}/connections/{connection}`). Per `judgement/references/PARENT_AND_IDENTITY.md` (Case B), replaced `ProjectRef`, `Location`, and `ComponentRef` with required, immutable `ParentRef *DesignCenterComponentRef` (`json:"parentRef"`).
  - `authoritative-proto-reference`: Converted `.spec.destinationComponentURI` to `DestinationComponentRef *DesignCenterComponentRef` (`json:"destinationComponentRef"`). Added custom mapper functions `DesignCenterConnectionSpec_FromProto` and `DesignCenterConnectionSpec_ToProto` in `pkg/controller/direct/designcenter/designcenterconnection_mapper.go` to cleanly map the `destination_component_uri` field.
  - `identity-not-in-cai`: Verified template `//designcenter.googleapis.com/projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template}/components/{component}/connections/{connection}` against proto `google.api.resource` pattern and marked accepted.
  - `verify-service-docs-link` & `verify-resource-docs-link`: Added service docs `https://cloud.google.com/application-design-center/docs` and REST reference `https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applicationTemplates.components.connections`.
  - `well-known-proto-type`: Confirmed `google.protobuf.Value` representation across parameters and marked resolved/accepted.
  - `untriaged-bulk-generation`: Resolved with resolution `edited`.
  - Updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.















