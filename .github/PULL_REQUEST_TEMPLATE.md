### BRIEF Change description

<!--
Describe what this pull request does.

PR Linkage & Multi-PR Staging Protocol:
- Intermediate PR (Phase 1 or 2): Use "Part of #<issue_id>". Do NOT use "Fixes #" or "Closes #".
- Terminal PR (Phase 3 or single-PR field addition): Use "Fixes #<issue_id>".
-->
Fixes #
<!-- or: Part of # -->

#### WHY do we need this change?

#### Special notes for your reviewer:

#### Confidentiality & Cleanliness Verification

- [ ] No Buganizer IDs (`b/<id>`), internal URLs, or customer names in the title, description, commit messages, code comments, or release note.

#### Does this PR add something which needs to be 'release noted'?
<!--
Release Note Protocol:
- Intermediate PR (Phase 1 or 2): In the release note block, enter "NONE".
- Terminal PR (Phase 3 or single-PR field addition): Provide the customer-facing release note.
If the PR requires additional action from users switching to the new release, include the string "action required".
-->
```release-note

```

- [ ] Reviewer reviewed release note.

#### Additional documentation e.g., references, usage docs, etc.:

<!--
This section can be blank if this pull request does not require any additional documentation.

When adding links which point to resources within git repositories, like
usage documentation, please reference a specific commit and avoid
linking directly to the master branch. This ensures that links reference a
specific point in time, rather than a document that may change over time.

See here for guidance on getting permanent links to files: https://help.github.com/en/articles/getting-permanent-links-to-files

Please use the following format for linking documentation:
- [Usage]: <link>
- [Other doc]: <link>
-->
```docs

```

#### Intended Milestone

Please indicate the intended milestone. 
- [ ] Reviewer tagged PR with the actual milestone.

### Archetype-Specific Pre-Flight Checklists

Complete the checklist corresponding to the controller archetype modified in this PR:

#### Direct Controller Archetype
- [ ] `make manifests`
- [ ] `make resource-docs`
- [ ] `TestFocusedMappers`
- [ ] MockGCP alignment
- [ ] `make ready-pr`

#### Terraform-Based Controller Archetype
- [ ] TPG compiles (`make vet`)
- [ ] `config/servicemappings/` updated
- [ ] `make manifests`
- [ ] `make resource-docs`
- [ ] `make ready-pr`

#### DCL-Based Controller Archetype
- [ ] `make upgrade-dcl`
- [ ] `make ready-pr`
- [ ] CRD diff verified

### Tests you have done

<!--

Make sure you have run "make ready-pr" to run required tests and ensure this PR is ready to review. 

Also if possible, share a bit more on the tests you have done. 

For example if you have updated the pubsubtopic sample, you can share the test logs from running the test case locally.

go test -v -tags=integration ./config/tests/samples/create -test.run TestAll -run-tests pubsubtopic

-->

- [ ] Run `make ready-pr` to ensure this PR is ready for review.
- [ ] Perform necessary E2E testing for changed resources.
