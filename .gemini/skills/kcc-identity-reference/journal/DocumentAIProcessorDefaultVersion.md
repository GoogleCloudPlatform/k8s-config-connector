# DocumentAIProcessorDefaultVersion Identity & Reference Migration Journal

## Overview
Successfully migrated `DocumentAIProcessorDefaultVersion` (`documentai.cnrm.cloud.google.com/v1alpha1`) to the modern `identity.IdentityV2` and `refs.Ref` patterns using `gcpurls.Template` and code generation via `controllerbuilder generate-identity`.

## Context & Observations

1. **RPC Request Mapping & Explicit Pattern Override**:
   - `DocumentAIProcessorDefaultVersion` is mapped from the proto request `google.cloud.documentai.v1.SetDefaultProcessorVersionRequest`.
   - Because `SetDefaultProcessorVersionRequest` is an RPC request message rather than a standard resource proto message, it lacks `(google.api.resource)` annotations.
   - Running `controllerbuilder generate-identity` required passing `--pattern DocumentAIProcessorDefaultVersion=projects/{project}/locations/{location}/processors/{processor}` to explicitly specify the canonical URL pattern.

2. **Declaration-by-Declaration Override ("Hand-written wins")**:
   - `DocumentAIProcessorDefaultVersion` is a brownfield Terraform-managed resource.
   - Its schema contains only `spec.resourceID` (the processor resource name) and `spec.version` without explicit `spec.location` or `spec.projectRef`.
   - In accordance with KCC guidelines, no schema changes were made.
   - Hand-wrote `getIdentityFromDocumentAIProcessorDefaultVersionSpec` in `documentaiprocessordefaultversion_identity.go` to handle:
     - Full external references in `spec.resourceID` (or `metadata.name`) via `identity.FromExternal(resourceID)`.
     - Standard fallback via `refs.ResolveProjectID` and `refs.GetLocation` if provided.
   - Hand-wrote `Normalize` in `documentaiprocessordefaultversion_reference.go` using `refs.NormalizeWithFallback` which checks for `Ready == True` in status conditions before returning the identity, correctly raising `k8s.NewReferenceNotReadyError` when not ready.
   - `controllerbuilder generate-identity` automatically recognized the hand-written declarations and omitted them from `*.generated.go` without collisions.

3. **Judgement Queue Resolution**:
   - Resolved the `identity-root-unknown` entry in `apis/documentai/judgement_queue.yaml` with status `resolved` and resolution `edited`.

4. **Schema Verification & CAIS Compatibility**:
   - `dev/tasks/diff-crds` confirmed zero changes to the CRD schema.
   - `TestRegisteredTemplatesMatchCAI` in `pkg/gcpurls` passed directly because the template `projects/{project}/locations/{location}/processors/{processor}` matches the CAIS entry for `documentai.googleapis.com/Processor`.
   - All unit tests and conformance tests pass cleanly.
