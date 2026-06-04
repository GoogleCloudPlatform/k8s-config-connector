# Journal: AccessContextManagerGCPUserAccessBinding Fuzzer Implementation

## Observations and Learnings

1. **Google Identity Protos Missing from Default Compilation**
   The protoc command in `dev/tools/controllerbuilder/generate-proto.sh` did not compile `google/identity` protos, only `google/cloud` and specific other paths. Since `GcpUserAccessBinding` resides under `google/identity/accesscontextmanager/v1/`, the proto path `${THIRD_PARTY}/googleapis/google/identity/*/*/*.proto` had to be added, and cached `.build/googleapis*.pb` files deleted to force reconstruction.

2. **Pointer Types Required for Automatic Mapper Generation**
   In the scaffolded `_types.go` file, scalar string fields should use `*string` (pointers) rather than `string`. If a non-pointer `string` is used in the spec, the KCC `mappergenerator` tool automatically expects a custom map function like `AccessContextManagerGCPUserAccessBindingSpec_GroupKey_ToProto` instead of using direct assignment or `direct.ValueOf`. Declaring fields as `*string` allows the mapper to be generated automatically.

3. **Organization and Folder Parent References**
   The templated `_types.go` scaffolds `projectRef *refsv1beta1.ProjectRef` and `location string`. Since `GcpUserAccessBinding` is parented by an Organization and has no location, we had to replace those with `organizationRef *refsv1beta1.OrganizationRef` and remove the location field.
