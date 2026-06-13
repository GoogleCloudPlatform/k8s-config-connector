# ComputeReservation Direct Types & Greenfield Skill Journal

During the investigation of the `ComputeReservation` types and `generate.sh` configuration, the following findings and observations were made:

## 1. Multi-version and Greenfield/Brownfield Distinctions
- `ComputeReservation` was previously a legacy (Terraform-based) resource in KCC under `v1beta1`.
- Because of this, its direct controller migration occurred under `v1beta1` (as seen in `apis/compute/v1beta1/reservation_types.go` and `pkg/controller/direct/compute/reservation_controller.go`), rather than starting as a new greenfield `v1alpha1` resource.
- Therefore, the stability-level label in `reservation_types.go` is `stable`, corresponding to its `v1beta1` status, rather than `alpha` as standard greenfield resources.

## 2. Type Skipping Behavior of Controller Builder
- The KCC generator tool (`tooling/main.go generate-types`) automatically detects existing Go types with `// +kcc:proto` or `// +kcc:spec:proto` annotations (such as those in `reservation_types.go`).
- When such handwritten types are present, the generator skips outputting them in `types.generated.go`, printing messages like:
  `/* found existing non-generated go type with proto tag "google.cloud.compute.v1.Reservation", skipping`
- This is extremely powerful because it allows developers to hand-write and custom-validate fields (or map them to references like `ExtendedProjectRef`) while still generating helper and subsidiary types in `types.generated.go`.

## 3. Reference Mapping
- Fields such as `ShareSettings.ProjectMap.KeyRef` and `ShareSettingsProjectConfig.ProjectIDRef` are correctly mapped to proper KCC reference fields (`ExtendedProjectRef` and `ProjectRef` respectively).
- No new entries were added to the `missingrefs.txt` exceptions list, adhering to the high-quality engineering standard that all reference-like fields must use proper KCC references.
