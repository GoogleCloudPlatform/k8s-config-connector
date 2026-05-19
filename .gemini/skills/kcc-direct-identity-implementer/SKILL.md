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

2.  **Parent References**:
    Ensure the `Spec` struct has a `ProjectRef` (and `Location` if applicable) for the parent. Use the `refs.` package.

## Journaling
Append any template mapping complexities (e.g., multi-parent resources) to `.gemini/skills/kcc-direct-identity-implementer/journal.md`.
