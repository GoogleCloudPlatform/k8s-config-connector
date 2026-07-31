---
name: kcc-direct-base-types-implementer
description: Base skill containing shared standards for all KCC direct resource types (both greenfield and brownfield).
---

# KCC Direct Base Types Implementer

This skill provides the mandatory baseline standards that apply to *all* new KRM types (`_types.go`) for direct resources in Config Connector, regardless of whether they are greenfield or brownfield migrations.

## Shared Standards for <kind>_types.go

After running the generator (via `generate.sh`), you must verify and enforce the following baseline requirements on the resulting `_types.go` file:

- **Copyright**: The file must start with `// Copyright 2026 Google LLC`.
- **CRD Labels**: Include at least these two labels in the type definition:
  ```go
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
  ```
  *(Note: See greenfield/brownfield skills for the correct `stability-level` label to append.)*
- **Status Fields**: `status.observedGeneration` must be exactly `*int64`.
- **Reference Fields**: Ensure that fields referencing other GCP/KCC resources are implemented as proper KCC reference fields (e.g., using `pubsubv1beta1.PubSubTopicRef` or `refsv1beta1.ProjectRef`), following the `Ref` suffix naming convention. You **MUST NOT** add new exceptions to `tests/apichecks/testdata/exceptions/missingrefs.txt`. All reference-like fields must be implemented as proper references.
  * Standard references include:
    * Service accounts should be references to `IAMServiceAccount`.
    * Cloud Storage buckets should be references to `StorageBucket`.
    * Service Directory configurations should be references to `ServiceDirectoryService`.
- **Implement a New Reference Types**:

  * When a field is a reference to another resource, and that reference resource does not have full KCC support yet, we should support it as external-only reference.
  The `<newkind>Ref` must **always** be defined and implemented in its own separate file named `<newkind>_reference.go` (e.g., `networksecurityinterceptdeploymentgroup_reference.go`) under `apis/<service>/v1alpha1`(e.g., `apis/networksecurity/v1alpha1/`) directory. This keeps the main type definitions clean and isolated from reference resolution boilerplate.
  Example external-only reference:
  ```
  var _ refsv1beta1.Ref = &NetworkSecurityInterceptDeploymentGroupRef{}
  var NetworkSecurityInterceptDeploymentGroupGVK = GroupVersion.WithKind("NetworkSecurityInterceptDeploymentGroup")
  
  // NetworkSecurityInterceptDeploymentGroupRef is a reference to a NetworkSecurityInterceptDeploymentGroup.
  
  type NetworkSecurityInterceptDeploymentGroupRef struct {
      /* A reference to an externally managed NetworkSecurityInterceptDeploymentGroup resource.
      Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptDeploymentGroups/{{interceptDeploymentGroupID}}". */
      External string `json:"external,omitempty"`
  
      /* NOTYET
      // The name of a NetworkSecurityInterceptDeploymentGroup resource.
      Name string `json:"name,omitempty"`
  
      // The namespace of a NetworkSecurityInterceptDeploymentGroup resource.
      Namespace string `json:"namespace,omitempty"`
      */
  }
  
  func (r *NetworkSecurityInterceptDeploymentGroupRef) GetGVK() schema.GroupVersionKind {
  return NetworkSecurityInterceptDeploymentGroupGVK
  }
  
  func (r *NetworkSecurityInterceptDeploymentGroupRef) GetNamespacedName() types.NamespacedName {
  return types.NamespacedName{}
  }
  
  func (r *NetworkSecurityInterceptDeploymentGroupRef) GetExternal() string {
  return r.External
  }
  
  func (r *NetworkSecurityInterceptDeploymentGroupRef) SetExternal(ref string) {
  r.External = ref
  }
  
  func (r *NetworkSecurityInterceptDeploymentGroupRef) ValidateExternal(ref string) error {
  id := &NetworkSecurityInterceptDeploymentGroupIdentity{}
  if err := id.FromExternal(ref); err != nil {
  return err
  }
  return nil
  }
  
  func (r *NetworkSecurityInterceptDeploymentGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
  if r.External == "" {
  return fmt.Errorf("external reference must be specified for %s", NetworkSecurityInterceptDeploymentGroupGVK.Kind)
  }
  return r.ValidateExternal(r.External)
  }
  ```
  * Also implement `<kind>_identity.go` of the new reference resource(e.g., `networksecurityinterceptdeploymentgroup_identity.go`), as the `FromExternal` function is used in `<kind>_reference.go`. 
  Example initial identity:
  ```
  var (
  _ identity.IdentityV2 = &NetworkSecurityInterceptDeploymentGroupIdentity{}
  )
  
  var (
  NetworkSecurityInterceptDeploymentGroupIdentityFormat = gcpurls.Template[NetworkSecurityInterceptDeploymentGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/interceptDeploymentGroups/{interceptdeploymentgroup}")
  )
  
  // NetworkSecurityInterceptDeploymentGroupIdentity is the identity of a GCP NetworkSecurityInterceptDeploymentGroup resource.
  // +k8s:deepcopy-gen=false
  type NetworkSecurityInterceptDeploymentGroupIdentity struct {
  Project                  string
  Location                 string
  InterceptDeploymentGroup string
  }
  
  func (i *NetworkSecurityInterceptDeploymentGroupIdentity) String() string {
  return NetworkSecurityInterceptDeploymentGroupIdentityFormat.ToString(*i)
  }
  
  func (i *NetworkSecurityInterceptDeploymentGroupIdentity) Host() string {
  return NetworkSecurityInterceptDeploymentGroupIdentityFormat.Host()
  }
  
  func (i *NetworkSecurityInterceptDeploymentGroupIdentity) FromExternal(ref string) error {
  parsed, match, err := NetworkSecurityInterceptDeploymentGroupIdentityFormat.Parse(ref)
  if err != nil {
  return fmt.Errorf("format of NetworkSecurityInterceptDeploymentGroup external=%q was not known (use %s): %w", ref, NetworkSecurityInterceptDeploymentGroupIdentityFormat.CanonicalForm(), err)
  }
  if !match {
  return fmt.Errorf("format of NetworkSecurityInterceptDeploymentGroup external=%q was not known (use %s)", ref, NetworkSecurityInterceptDeploymentGroupIdentityFormat.CanonicalForm())
  }
  
      *i = *parsed
      return nil
  }
  ```
  * Run `TestDirectResourceFileNaming` unit test and update `testdata/exceptions/naming_violations.txt`. We expect new entries like
  ```
  [naming_violation] file=apis/networksecurity/v1alpha1/networksecurityinterceptdeploymentgroup_identity.go prefix=networksecurityinterceptdeploymentgroup (expected a valid resource kind prefix)
  ```
  Since we do not yet support this resource, just add it as an external-only reference to unblock the development of other resources that depend on it.
- **Reference Types Location**: Whenever a reference type (e.g. `<Kind>Ref` implementing `refsv1beta1.Ref`) is needed, it must **always** be defined and implemented in its own separate file named `<kind>_reference.go` (e.g., `filestorebackup_reference.go`) rather than inside `_types.go`. This keeps the main type definitions clean and isolated from reference resolution boilerplate.
- **Service-Generated Fields**: Fields that are service-generated values (such as `etag`) **MUST NOT** be under `spec`. They should be put under status (specifically under `observedState`).
- **Acronym Capitalization**: Acronyms within field names must be either all in lowercase or all capitalized (e.g., use `vertexAISearchRuntimeConfig` instead of `vertexAiSearchRuntimeConfig`).
- **Manual Types Isolation**: Always move any manually defined/handled types (including complex structs, recursive types, custom schemas, union/oneof configs) from `types.generated.go` to `<kind>_types.go` before making any further modifications to them.
- **CRD Recursion and Commenting**: If a schema/type definition has deeply nested or recursive field types (such as `parameters` and `response` in function/schema declarations) that trigger OpenAPI validation errors or panics during conversion, comment them out or represent them as simple custom types to preserve CRD stability.
- **Preserve Unknown Fields**: Avoid using `x-kubernetes-preserve-unknown-fields: true` except on standard fields of type `apiextensionsv1.JSON` representing raw/dynamic unstructured user-provided payload structures (such as dynamic parameters/request payloads) where keeping unknown fields is strictly required to prevent data loss. Verify any such usage to ensure it is necessary.

- **KCC Proto Annotations**:
  To enable auto-generation of mappers, you must add the correct "kcc:proto" annotations to Go structs in `_types.go`:
  * The Spec struct must be annotated with `// +kcc:spec:proto=<proto_type>` (e.g. `// +kcc:spec:proto=google.cloud.compute.v1.ServiceAttachment`).
  * The ObservedState struct (if present) must be annotated with `// +kcc:observedstate:proto=<proto_type>`.
  * The Status struct (if there is no separate ObservedState struct) must be annotated with `// +kcc:status:proto=<proto_type>`.
  * Nested/referenced helper structs (both in Spec and Status) must be annotated with `// +kcc:proto=<proto_sub_type>` (e.g. `// +kcc:proto=google.cloud.compute.v1.ServiceAttachmentConnectedEndpoint`).
