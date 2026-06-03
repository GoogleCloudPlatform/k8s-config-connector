# DataflowFlexTemplateJob Learning Journal

## Identity and Reference Pattern Implementation

When moving `DataflowFlexTemplateJob` to the identity and refs pattern, I observed the following:

1. **Server-generated IDs**: Dataflow jobs have server-generated IDs (`jobID`) that are different from the user-provided job names used during launch. To handle this in `IdentityV2`:
    - `getIdentityFromSpec` checks for `status.jobId`. If empty, it returns `nil, nil`, indicating the identity is not yet known.
    - This pattern is similar to `CloudBuildTrigger` and handles the transition from "desiring a job with name X" to "managing a job with ID Y".

2. **Region vs Location**: Dataflow uses `region` in its spec instead of the more common `location`. The `getIdentityFromSpec` implementation was tailored to handle this by reading from `spec.region`.

3. **Normalization of ExternalRef**: Since `externalRef` contains the full GCP resource name including the server-generated `jobID`, it must be normalized in E2E tests.
    - I updated `tests/e2e/normalize.go`'s `findLinksInKRMObject` to include `.status.externalRef`.
    - This allows the existing `Replacements` logic to extract the `jobID` from the link and replace it with `${jobID}` in the golden object files.
    - This is a general improvement that should likely be applied to all resources moving to the `externalRef` pattern.

4. **Controller Integration**: The direct controller was updated to use `GetIdentity`. This simplifies the logic for resolving `projectID` and `location`, and provides a standard way to handle the resource's identity throughout its lifecycle.
