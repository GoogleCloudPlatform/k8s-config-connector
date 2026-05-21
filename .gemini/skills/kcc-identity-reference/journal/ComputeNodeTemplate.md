# ComputeNodeTemplate Identity and Reference Journal

## Observation: Region vs Location
- `ComputeNodeTemplate` uses `spec.region` instead of `spec.location`.
- The `gcpurls.Template` used `{Region}` and the struct field was `Region`.
- In `getIdentityFromComputeNodeTemplateSpec`, I had to extract `region` from the spec. Since `refs.GetLocation` only looks for `location`, I manually extracted it from the unstructured object or the typed object.

## Observation: CAI Metadata
- CAI metadata for `NodeTemplate` uses plural `nodeTemplates` in the path, which matched the `servicemapping`.
- CAI: `//compute.googleapis.com/projects/{{PROJECT_ID}}/regions/{{REGION}}/nodeTemplates/{{NODE_TEMPLATE}}`
- Template: `projects/{Project}/regions/{Region}/nodeTemplates/{NodeTemplate}`

## Observation: Existing GVK
- `ComputeNodeTemplateGVK` was already defined in `nodetemplate_types.go`, so I didn't need to define it in the reference file.

## Learning: IdentityV2 Implementation
- Implementing `GetIdentity` in a separate `_identity.go` file for a resource defined in `_types.go` is a clean way to add functionality without bloating the main types file.
- Using `common.ToUnstructured` was a mistake as it was undefined; the correct way is to check the type and then use `unstructured.NestedString` on `obj.(*unstructured.Unstructured).Object`. Actually I used `common.ValueOf` and `unstructured.NestedString` correctly in the final version.
