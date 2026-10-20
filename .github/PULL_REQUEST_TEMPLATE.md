### BRIEF Change description

<!--
Describe what this pull request does.

PR Linkage & Multi-PR Staging Protocol:
* Intermediate PR (Phase 1 or 2 of a multi-phase resource delivery):
  Use "Part of #<issue_id>". Do NOT use "Fixes #" or "Closes #".
  In the release note block below, enter "NONE".
* Terminal PR (Phase 3 or a single-PR field addition / bugfix):
  Use "Fixes #<issue_id>" and provide a customer-facing release note.
-->
Fixes #
Part of #

#### WHY do we need this change?

#### Special notes for your reviewer:

#### Pre-Flight Checklist by Controller Archetype:
<!--
Check off the items that apply to your change.
-->
- [ ] **Direct Controller**:
  - [ ] `make manifests` run and CRD diffs verified
  - [ ] Mappers tested: `go test -v ./pkg/controller/direct/... -run TestFocusedMappers`
  - [ ] Hermetic MockGCP tests pass: `hack/compare-mock <service> <resource>`
  - [ ] Controller registered via anonymous import in `pkg/controller/direct/register/register.go`
- [ ] **Terraform-Based Controller**:
  - [ ] TPG compiles cleanly: `make vet`
  - [ ] `config/servicemappings/<service>.yaml` updated
  - [ ] `make manifests` and `make resource-docs` executed
  - [ ] Integration tests pass: `go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...`
- [ ] **DCL-Based Controller**:
  - [ ] `make upgrade-dcl` executed
  - [ ] CRD diffs verified in `config/crds/resources/`
- [ ] **General PR Readiness**:
  - [ ] `make ready-pr` passes cleanly
  - [ ] **SOUL.md Invariant 7**: No Buganizer IDs (`b/<id>`), internal URLs, or customer names in the title, description, commit messages, code comments, test fixtures, or release note.

#### Does this PR add something which needs to be 'release noted'?
<!--
If no (or if this is an intermediate Phase 1/2 PR), write "NONE" in the release-note block below.
If yes, a customer-facing release note is required:
Enter your extended release note in the block below. If the PR requires additional action from users switching to the new release, include the string "action required".
-->
```release-note

```

- [ ] Reviewer reviewed release note.

#### Additional documentation e.g., references, usage docs, etc.:

```docs

```

#### Intended Milestone

Please indicate the intended milestone. 
- [ ] Reviewer tagged PR with the actual milestone.

### Tests you have done

<!--
Make sure you have run "make ready-pr" to run required tests and ensure this PR is ready to review. 
-->
- [ ] Run `make ready-pr` to ensure this PR is ready for review.
- [ ] Perform necessary E2E testing for changed resources.
