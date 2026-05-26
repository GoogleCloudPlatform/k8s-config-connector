#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -x
cd /workspaces/k8s-config-connector/mockgcp
cd tools/patch-proto

APIGEE_API_PROTO=../../apis/mockgcp/cloud/apigee/v1/service.proto
if ! grep -q "{name=projects/\*}:provisionOrganization" ${APIGEE_API_PROTO}; then
  sed -i 's/{project=projects\/\*}:provisionOrganization/{name=projects\/\*}:provisionOrganization/g' ${APIGEE_API_PROTO}
fi

API_PROTO=../../apis/mockgcp/cloud/resourcemanager/v1/api.proto
sed -i 's/^syntax = "proto2";/syntax = "proto3";/' ${API_PROTO}

go run . --file ${API_PROTO} --message "TestIamPermissionsRequest" --mode "replace" <<'INNER_EOF'
  // REQUIRED: The resource for which the permission checking is to be performed.
  optional string resource = 1;

  // The set of permissions to check for the resource.
  repeated string permissions = 2 [json_name="permissions"];
INNER_EOF

if ! grep -q "rpc TestIamPermissions" ${API_PROTO}; then
  go run . --file ${API_PROTO} --service "FoldersServer" --mode "append" <<'INNER_EOF'
  // Returns permissions that a caller has on the specified project.
  rpc TestIamPermissions(TestIamPermissionsRequest) returns (TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v1/{resource=folders/*}:testIamPermissions"
      body: "*"
    };
  };
INNER_EOF

  go run . --file ${API_PROTO} --service "OrganizationsServer" --mode "append" <<'INNER_EOF'
  // Returns permissions that a caller has on the specified project.
  rpc TestIamPermissions(TestIamPermissionsRequest) returns (TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v1/{resource=organizations/*}:testIamPermissions"
      body: "*"
    };
  };
INNER_EOF

  go run . --file ${API_PROTO} --service "ProjectsServer" --mode "append" <<'INNER_EOF'
  // Returns permissions that a caller has on the specified project.
  rpc TestIamPermissions(TestIamPermissionsRequest) returns (TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v1/{resource=projects/*}:testIamPermissions"
      body: "*"
    };
  };
INNER_EOF
fi

sed -i 's/^syntax = "proto3";/syntax = "proto2";/' ${API_PROTO}
