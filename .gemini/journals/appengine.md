### [2026-11-15] AppEngineAuthorizedCertificate Prepopulation
- **Context**: Prepopulated direct KRM types, identity, and judgement queue for `AppEngineAuthorizedCertificate` (`appengine.cnrm.cloud.google.com/v1alpha1`, proto `google.appengine.v1.AuthorizedCertificate`).
- **Problem**: App Engine resources use the `apps/{app}/...` resource URL hierarchy in GCP URLs instead of `projects/{project}/...`. The proto defines no `(google.api.resource)` annotation, resulting in `identity-root-unknown` in `judgement_queue.yaml`.
- **Solution**: Scaffolded types, identity, and mapper via `controllerbuilder generate-types --prepopulate-resource`, `generate-identity`, and `generate-mapper --prepopulate-resource` in `apis/appengine/generate.sh`. Registered `AppEngineAuthorizedCertificate` in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Impact**: Step 1b (`judgement/SKILL.md`) should triage `identity-root-unknown` by mapping the root `apps/{app}` to `spec.projectRef`.
