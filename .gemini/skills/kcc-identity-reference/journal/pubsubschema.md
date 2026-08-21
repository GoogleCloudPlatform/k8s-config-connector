# PubSubSchema Identity and Reference Transition Journal

## Observations

- `PubSubSchema` is a new transition to the `IdentityV2` and `refs.Ref` patterns, meaning there were no existing `_identity.go` or `_reference.go` files for it.
- `PubSubSchema` is a global resource mapped to GCP projects (`projects/{project}/schemas/{schema}`) without a location (region/zone) property in GCP or in the CRD.
- The resource's spec contains `projectRef` of type `refs.ProjectRef`.
- Since the status does not contain an `externalRef` or `Name` field, no cross-check was performed in `GetIdentity(ctx, reader)`.

## Shortcomings in SKILL.md

- There are no specific shortcomings identified; the `SKILL.md` provided precise and accurate instructions for implementing identity and reference files for a resource that had neither.

## Learnings

- Implementing identity and reference types from scratch for a resource with `gcpurls.Template` is straightforward, clean, and produces exactly zero schema drift (as verified by `dev/tasks/diff-crds`).
