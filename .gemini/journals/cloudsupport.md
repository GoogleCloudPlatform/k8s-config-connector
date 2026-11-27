# CloudSupport Journal

### [2026-11-15] Greenfield Direct Controller, Fixtures, and Fuzzer for CloudSupportSupportEventSubscription
- **Context**: Greenfield direct controller, E2E fixtures, and fuzzer for `CloudSupportSupportEventSubscription` (#15738).
- **Action**:
  1. Generated direct controller (`pkg/controller/direct/cloudsupport/cloudsupportsupporteventsubscription_controller.go`) and round-trip fuzzer (`pkg/controller/direct/cloudsupport/cloudsupportsupporteventsubscription_fuzzer.generated.go`) using `controllerbuilder`.
  2. Implemented `identity.ServerGeneratedIdentity` for `CloudSupportSupportEventSubscriptionIdentity`:
     - Added `HasIdentitySpecified() bool` method.
     - Updated `getIdentityFromCloudSupportSupportEventSubscriptionSpec` to read directly from `Spec.ResourceID` instead of defaulting to `metadata.name`.
     - In `GetIdentity`, allowed adoption of server-generated `status.ExternalRef` when `spec.resourceID` is empty.
     - Added `ExternalIdentifier()` to implement `identity.ExternalIdentifier`.
     - In controller `Find()` and `Delete()`, added `!a.id.HasIdentitySpecified()` check to skip invalid queries before creation.
     - In controller `Create()`, omitted setting `resource.Name` so the service assigns the ID, and parsed `created.GetName()` to set `status.ExternalRef`.
     - In controller `AdapterForObject`, ensured `common.NormalizeGreenfieldReferences` is called prior to `GetIdentity`.
  3. Registered `CloudSupportSupportEventSubscription` in `pkg/controller/resourceconfig/static_config.go` using `python3 dev/tasks/generate_static_config.py`.
  4. Added `CloudSupportSupportEventSubscription` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  5. Created minimal and maximal E2E fixtures under `pkg/test/resourcefixture/testdata/basic/cloudsupport/v1alpha1/cloudsupportsupporteventsubscription/` with dependencies (`PubSubTopic` and `IAMPolicyMember` for `cloud-support-apievents@system.gserviceaccount.com`).
  6. Updated alpha missing field exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
  7. Verified 200-seed fuzzer round-trip (`FOCUS=CloudSupportSupportEventSubscription go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`), apichecks tests (`TestGenerateScripts`, `TestDirectResourceFileNaming`, `TestCRDFieldPresenceInTestsForAlpha`), and CAI tests (`TestRegisteredTemplatesMatchCAI`, `TestVerifyLabelsTestCasesExist`).
- **Observation**:
  - Live execution against real GCP (`./hack/record-gcp`) failed with HTTP 403 `Permission 'cloudsupport.supportSubscriptions.create' denied on resource '//cloudresourcemanager.googleapis.com/organizations/128653134652'` because the test service account (`overseer-kcc-tester@cnrm-barni-4.iam.gserviceaccount.com`) lacks organization-level permissions (`roles/cloudsupport.admin`). Golden logs and fixtures are prepared for MockGCP simulation in Step 3.
