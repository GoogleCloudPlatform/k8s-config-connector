---
name: kcc-direct-identity-implementer
description: Implement the IdentityV2 and ExternalIdentifier interfaces for a direct KCC resource. Use this when implementing IdentityV2 and refs.Ref for a resource.
---

# KCC Direct Identity Implementer

This skill guides the implementation of the IdentityV2 and ExternalIdentifier interfaces for direct KCC resources, ensuring they follow the canonical `gcpurls.Template` pattern.

## Inputs
- `resource_kind`: The KCC Kind (e.g., `VertexAIExampleStore`).
- `template`: The GCP URL template (e.g., `projects/{project}/locations/{location}/exampleStores/{example_store}`).
- `api_version`: The KCC API version.

## Workflow

1.  **Implement Identity**:
    Create `apis/<service>/<api_version>/<resource_lower>_identity.go`.
    - Use `identity.IdentityV2`.
    - Use `gcpurls.Template` for URL parsing.
    - Implement `ExternalIdentifier()`.
    - Implement `ParentString() string` to return the GCP parent URI (e.g., `projects/{project}` or `projects/{project}/locations/{location}`).

2.  **Implement Identity Tests**:
    Create `apis/<service>/<api_version>/<resource_lower>_identity_test.go`.
    - Ensure every `<resource_lower>_identity.go` file has a corresponding test file.
    - The test file should verify that `FromExternal()` can parse valid external identities and `String()` can generate correct external identities.

3.  **Parent & Hierarchy Support**:
    - Ensure the `Spec` struct has a `ProjectRef` (and `Location` if applicable).
    - **Hierarchical Branching**: If the GCP resource supports multiple hierarchies (e.g., Global and Regional), implement logic in `String()` and `FromExternal()` to handle both. Use the presence of the `location` field in the Spec to determine which pattern to use.

## Journaling
Append any template mapping complexities (e.g., multi-parent or multi-hierarchy resources) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
