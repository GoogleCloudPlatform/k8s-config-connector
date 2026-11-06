---
name: Support new field(s)
about: Request field(s) that exist in the GCP API but is missing from an existing
  KCC resource.
title: 'ai:feat: Add field <field_path> to <resource_kind_name>'
labels: enhancement, overseer
assignees: ''
type: Feature

---

<!--
Use this template to request missing fields on an EXISTING KCC resource.
Overseer picks up issues labeled `overseer` and implements the field by following
the skill referenced at the bottom of this issue.

To keep the generated PR small and easy to review:
- One Kind per issue.
- Only fill in the sections below. Please do NOT paste implementation plans,
  code snippets, or unrelated context. Extra content can lead Overseer to make
  unnecessary changes.
- For a brand-new resource, use the "Greenfield Resource Workflow" template.
  For bugs, use the "Bug" template.
-->

## Checklist

- [ ] I did not find a related open enhancement request.
- [ ] I understand that enhancement requests filed in the GitHub repository are by default low priority.
- [ ] If this request is time-sensitive, I have opened an issue with [GCP support](https://cloud.google.com/support-hub). **NOTE:** Only KCC on GKE is supported by GCP Support.

## KCC Kind

<!-- The CamelCase KCC Kind, e.g. SpannerDatabase -->


## Controller type

<!-- Check one. Terraform-based CRDs carry the label `cnrm.cloud.google.com/tf2crd: "true"`. -->

- [ ] Direct
- [ ] Terraform

## Missing field(s)

<!--
One per line: the GCP API field path, and whether it belongs in spec (user-settable)
or status.observedState (output-only).
Example:
- `encryptionConfig.kmsKeyName`: spec
- `earliestVersionTime`: status.observedState
-->

- 

## GCP API reference

<!-- Link to the REST API documentation for the field. -->


## Use case (optional)

<!-- Example configuration yaml. -->

---

## Instructions for Overseer (please do not edit)

Implement the field(s) listed in this issue on the given Kind.

Select the skill based on "Controller type":
- **Direct**: follow `.gemini/skills/add-missing-field/SKILL.md`.
- **Terraform**: follow `.gemini/skills/update-terraform-fields/SKILL.md`, then `.gemini/skills/test-terraform-fields/SKILL.md` for test coverage and golden files. Use .gemini/skills/reviewgen-legacy-feature/SKILL.md to conduct the review.

Please update the skill where you find it can be improved (and put observations that aren't necessarily as general in the journal.md file in the same folder)

When creating the PR, include the release note block:
```release-note
<resource_kind_name>: Added support for <field_path>.
```
For example:
```release-note
BigtableTable: Added support for spec.columnFamily[].type.
```


