---
name: Issue for Identity and Reference
description: For TF/DCL resources to be migrated to Direct, if generate.sh and types.go PR is closed, create an issue for creating identity and reference files.
schedule: "@daily"
skipPR: true
---
<!--
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Role
You are a software development assistant for the Kubernetes Config Connector project.
You have access to the following tools:
- GitHub CLI (`gh`)
- git
Use `gh` to perform your duties.

# Filter Criteria
This is the criteria to identify the resource Group and Kind that need to be migrated.

Find resources that meet the following criteria:
Look for issues labelled with `area/direct` and `step/gen-types` that are closed as completed within the last 7 days. These issues indicate that the generate.sh and types.go files have been created for a specific Group and Kind.
For each of these issues, identify the resource Group and Kind of the resource that has been migrated to Direct.

# Task
Use gh cli tool to create github issue.
In a single run create at most one issue to avoid overwhelming the team.
For a resource Group and Kind identified in the previous step, check if an issue already exists (open or closed) and create a new one if not.
If an issue already exists for that resource Group and Kind, use `gh issue edit <issue> --add-label` to add the labels if they do not exist.
If an issue already exists, skip to the next one that meets the criteria and repeat the process.
The issue will be linked to the main epic by the cross-reference in the issue body.
If more than 10 open issues labeled with 'overseer' and 'step/identity-reference' already exist, do not create new issues to avoid overwhelming the team. Instead, log a message indicating that there are already 10 pending issues and skip creating new ones until some of the existing issues are resolved.
Created issues should be clear and actionable, providing enough context for developers to understand what needs to be done.
IMPORTANT:
* Before creating an issue for a resource, check if an issue already exists (open or closed) to avoid duplicates.
* The issue title should be in the format: `Create Identity and Reference files for <resource_group> <resource_kind>`
* Use `gh` tool to create issue.
* Append a link to this chore file (`.agents/tf2d-identity-reference-issue.md`) at the end of the issue body for traceability.

## Issue Title

Title should be: `Create Identity and Reference files for <resource_group> <resource_kind>`,
where `<resource_group>` and `<resource_kind>` are replaced with the actual resource Group and Kind of the resource identified for migration.

## Issue Labels
The issue should be labeled with the following labels:
* `overseer` to indicate that the issue was created by Overseer.
* `area/direct` to indicate that the issue is related to Direct migration.
* `priority/medium` to indicate the priority level of the issue.
* `step/identity-reference` to indicate the step in the migration process.

Use gh tool to create the issue with the appropriate title, labels, and body content as described in the instructions.

## Issue Body
The issue body should contain this text template with the appropriate resource Group, Kind, and API Version filled in.
The body template is treated as markdown. So retain the quotes formatting and the code block formatting as is when filling in the Group and Kind.
Replace the <issue_number> placeholder with the actual issue number of the identified source issue that created the generate.sh and types.go files for this resource.
Dynamically determine the API version (e.g., v1beta1, v1alpha1) based on the directory structure created in the previous generation step.
Convert `<kind>` to lowercase for file paths (e.g. `lowercase_kind`).
Prefer the newer version and if not able to identify yse v1beta1.

------------ BEGIN ISSUE BODY TEMPLATE ------------
As part of moving resources from terraform and DCL controllers to direct controllers (Epic #5954), we need to create the Go identity and reference for `${resource_group}${resource_name}`.

Currently, `${resource_group}${resource_name}` is managed by the Terraform or DCL controller. The goal is to create the Go identity and reference in `apis/${resource_group}/v1beta1/` so that we can eventually migrate the controller implementation to the "direct" approach.

### Instructions

1.  **Add `apis/${resource_group}/v1beta1/${resource_name}_identity.go`**:
    Create a file `apis/${resource_group}/v1beta1/${resource_name}_identity.go`.
    The following are samples of similar identity files.
    - `artifactregistry/v1beta1/artifactregistryrepository_identity.go`
    - `apis/iam/v1beta1/serviceaccountkey_identity.go`
    The correct identity URL formats can be found in the page https://docs.cloud.google.com/asset-inventory/docs/asset-names. Please use that url format in place of "URLKey/{keyValue}". Please break out a separate field for each curly brace delineated '{field}' in the format.
    Please ensure the Identity class implements identity.IdentityV2.
    implement the GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) method on the ${resource_group}${resource_name} resource
The file likely includes something like the following
Example:
```go
const (
	// ${resource_group}${resource_name}IdentityURL is the format for the externalRef of a ${resource_group}${resource_name}.
	${resource_group}${resource_name}IdentityURL = "URLKey/{keyValue}"
)

var (
	_ identity.IdentityV2 = &${resource_group}${resource_name}Identity{}
	_ identity.Resource   = &${resource_group}${resource_name}{}
)

var ${resource_group}${resource_name}IdentityFormat = gcpurls.Template[${resource_group}${resource_name}Identity](
	"${resource_group}.googleapis.com",
	${resource_group}${resource_name}IdentityURL,
)

// ${resource_group}${resource_name}Identity represents the identity of a ${resource_group}${resource_name}.
// +k8s:deepcopy-gen=false
type ${resource_group}${resource_name}Identity struct {
	KeyValue string
}
```

2. **Add `apis/${resource_group}/v1beta1/${resource_name}_identity_test.go`**:
    Create a file `apis/${resource_group}/v1beta1/${resource_name}_identity_test.go`.
    It should unit test `apis/${resource_group}/v1beta1/${resource_name}_identity.go`.
    The following are samples of similar identity test files.
    - `apis/artifactregistry/v1beta1/artifactregistryrepository_identity_test.go`
    - `apis/iam/v1beta1/serviceaccountkey_identity_test.go`

3. **Add `apis/${resource_group}/v1beta1/${resource_name}_reference.go`**:
    Create a file `apis/${resource_group}/v1beta1/${resource_name}_reference.go`.
    The following are samples of similar reference files.
    - `apis/artifactregistry/v1beta1/artifactregistryrepository_reference.go`
    - `apis/iam/v1beta1/serviceaccountkey_reference.go`
Reference URL formats can be looked up from https://docs.cloud.google.com/asset-inventory/docs/asset-names
Please implement the full suite of methods like Normalize, ValidateExternal, and ParseExternalToIdentity.
The file likely includes something like the following
Example:
```go
var ${resource_group}${resource_name}GVK = schema.GroupVersionKind{
	Group:   "${resource_group}.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "${resource_group}${resource_name}",
}

var _ refsv1beta1.Ref = &${resource_group}${resource_name}Ref{}

// ${resource_group}${resource_name}Ref is a reference to a ${resource_group}${resource_name} resource.
type ${resource_group}${resource_name}Ref struct {
	// A reference to an externally managed ${resource_group}${resource_name} resource.
	// Should be in the format "URLKey/{keyValue}".
	External string `json:"external,omitempty"`

	// The name of a ${resource_group}${resource_name} resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ${resource_group}${resource_name} resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&${resource_group}${resource_name}Ref{})
}

func (r *${resource_group}${resource_name}Ref) GetGVK() schema.GroupVersionKind {
	return ${resource_group}${resource_name}GVK
}
```

4. **Add `apis/${resource_group}/v1beta1/${resource_name}_reference_test.go`**:
    Create a file `apis/${resource_group}/v1beta1/${resource_name}_reference_test.go`.
    It should unit test `apis/${resource_group}/v1beta1/${resource_name}_reference.go`.
    The following are samples of similar reference test files.
    - `apis/artifactregistry/v1beta1/artifactregistryrepository_reference_test.go`
    - `apis/bigquery/v1beta1/bigquerytable_reference_test.go`

5.  **Copyright Headers**:
    Ensure that new files have the correct copyright header with the current year (dynamically insert the current year):
    ```go
    // Copyright <current_year> Google LLC
    ```
    Please do *not* change the copyright on existing files.

6. **Validate changes**:
   - Running `make all-binary` and `make test` will ensure the new code compiles and the tests pass. Please fix any issue discovered by this compilation.

7. **Create PR**:
   - Create a Pull Request with your changes.
   - Apply the same labels to the PR as are on this issue.
   - Include a link to the chore file (`.agents/tf2d-identity-reference-issue.md`) in the PR description.
   - Include `Fixes #<issue-number>` in the PR description.

This issue is part of Epic #5954.

This issue is a continuation of the #<issue_number> which created the generate.sh and types.go files for this resource. Please refer to that issue for more context on the resource and the migration effort.
------------ END ISSUE BODY TEMPLATE ------------
