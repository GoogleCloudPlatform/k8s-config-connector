---
name: Support new field(s)
about: Request field(s) that exist in the GCP API but is missing from an existing
  KCC resource.
title: 'ai:feat: Add field <field_path> to <resource_kind_name>'
labels: enhancement, overseer, overseer/review
assignees: ''
type: Feature

---

<!--
Use this template to request missing fields on an EXISTING KCC resource.
Overseer picks up issues labeled `overseer` and implements the field by following
the skill referenced at the bottom of this issue.

At the bottom of this issue, keep ONLY the instructions block for your Kind's
controller type (Direct or Terraform) and delete the other one.

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

<!--
CHOOSE ONE: keep EITHER the Direct block OR the Terraform block below, and delete the other.
Do not edit the text inside the block you keep.
Terraform-based CRDs carry the label `cnrm.cloud.google.com/tf2crd: "true"`; otherwise the Kind uses a Direct controller.
-->

<!-- ===== Direct controller: delete this block if the Kind is Terraform-based ===== -->

## Instructions for Overseer

Implement the field(s) listed in this issue on the given Kind by following `.gemini/skills/add-missing-field/SKILL.md` end-to-end.

## Review Instructions

Verify the PR follows `.gemini/skills/add-missing-field/SKILL.md`.

<!-- ===== End of Direct controller block ===== -->

<!-- ===== Terraform controller: delete this block if the Kind uses a Direct controller ===== -->

## Instructions for Overseer

Implement the field(s) listed in this issue on the given Kind by following `.gemini/skills/update-terraform-fields/SKILL.md` end-to-end.

## Review Instructions

Use `.gemini/skills/reviewgen-legacy-feature/SKILL.md` to conduct the review.

<!-- ===== End of Terraform controller block ===== -->
