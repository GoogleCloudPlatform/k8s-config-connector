#!/bin/bash

PROTO_FILES=`find .build -name '*.proto'`
FILTER="package google\..*\.${resource_group}\.(v[0-9]+)"
PACKAGE=`for file in ${PROTO_FILES}; do egrep "package google\..*\.${resource_group,,}\.(v[0-9]+)" $file ; done | cut -d';' -f1 | uniq | awk '{ print length($0), $0; }' | sort -t' ' -k1nr -k3d | tail -1 | cut -d' ' -f3`
echo "Using package $PACKAGE for resource group ${resource_group}"

CRD_FILE=$(ls config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_${resource_group,,}${resource_name,,}*.${resource_group,,}.cnrm.cloud.google.com.yaml 2>/dev/null | head -n 1)
CONTROLLER_TYPE="Terraform"
LABEL_NAME="tf2crd"
IS_MIGRATED=false

if [ -n "$CRD_FILE" ]; then
    if grep -q "dcl2crd: \"true\"" "$CRD_FILE" || grep -q "dcl2crd: true" "$CRD_FILE"; then
        CONTROLLER_TYPE="DCL"
        LABEL_NAME="dcl2crd"
    elif ! grep -q "tf2crd: \"true\"" "$CRD_FILE" && ! grep -q "tf2crd: true" "$CRD_FILE"; then
        IS_MIGRATED=true
    fi
fi

if [ "$IS_MIGRATED" = true ]; then
    echo "WARNING: The CRD for ${resource_group}${resource_name} ($CRD_FILE) does not have tf2crd or dcl2crd labels."
    echo "It is likely already migrated to the direct approach."
    exit 0
fi

if [ -f "apis/${resource_group,,}/v1beta1/generate.sh" ]; then
	echo "As part of moving resources from ${CONTROLLER_TYPE} controllers to direct controllers (Epic #5954), we need to create the Go types for \`${resource_group}${resource_name}\`.

Currently, \`${resource_group}${resource_name}\` is managed by the ${CONTROLLER_TYPE} controller (marked with \`${LABEL_NAME}=true\`). The goal is to create the Go types in \`apis/${resource_group,,}/v1beta1/\` so that we can eventually migrate the controller implementation to the \"direct\" approach.

### Instructions

1.  **Add to generate.sh**:
    Modify \`apis/${resource_group,,}/v1beta1/generate.sh\` to include \`${resource_group}${resource_name}\`.
    It likely maps to something like \`${PACKAGE}\`.
    Example:
    \`\`\`bash
    go run . generate-types \\
      --service ${PACKAGE} \\
      --api-version ${resource_group,,}.cnrm.cloud.google.com/v1beta1 \\
      --resource ${resource_group}SomethingElse:SomethingElse \\
      --resource ${resource_group}${resource_name}:${resource_name}
    \`\`\`

2.  **Generate Scaffolding**:
    Run \`apis/${resource_group,,}/v1beta1/generate.sh\`. This should create \`apis/${resource_group,,}/v1beta1/${resource_name,,}_types.go\`.

3.  **Iterate on Types**:
    Compare the generated CRD with the existing one using \`dev/tasks/diff-crds\`.
    Modify \`apis/${resource_group,,}/v1beta1/${resource_name,,}_types.go\` until the CRD matches the existing one at \`config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_${resource_group,,}${resource_name,,}s.${resource_group,,}.cnrm.cloud.google.com.yaml\`.

    **Acceptance Criteria:**
    - Running \`dev/tasks/diff-crds\` should not show differences (or minimal acceptable ones like descriptions).
    - Ensure that running the check_crd_equivalence MCP on the CRD should return EQUIVALENT.
    - Changes to the schema (fields added/removed) are NOT acceptable.

4.  **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    \`\`\`go
    // Copyright 2026 Google LLC
    \`\`\`
    Please do not change the copyright on existing files.

5.  **Labels**:
    Ensure the controller-runtime annotations match the existing CRD labels, including:
    \`\`\`go
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/managed-by-kcc=true\"
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/system=true\"
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/stability-level=stable\"
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/${LABEL_NAME}=true\"
    \`\`\`
    The goal is to maintain these annotations, not add an annotation if it is missing.

6.  **Status**:
    \`status.observedGeneration\` should be an \`*int64\`.

7. **Generate Mappers**:
   - Running \`dev/tasks/generate-types-and-mappers\` will generate the mapper code once the \`apis/${resource_group,,}/v1beta1/${resource_name,,}_types.go\` file is generating an equivalent CRD.
   - Run \`make all-binary\` to ensure the generated mapper code compiles. Please fix any issue discovered by this compilation.

This issue is part of Epic #5954."
else
	echo "As part of moving resources from ${CONTROLLER_TYPE} controllers to direct controllers (Epic #5954), we need to create the Go types for \`${resource_group}${resource_name}\`.

Currently, \`${resource_group}${resource_name}\` is managed by the ${CONTROLLER_TYPE} controller (marked with \`${LABEL_NAME}=true\`). The goal is to create the Go types in \`apis/${resource_group,,}/v1beta1/\` so that we can eventually migrate the controller implementation to the \"direct\" approach.

### Instructions

1.  **Create a generate.sh**:
    Create \`apis/${resource_group,,}/v1beta1/generate.sh\` which includes \`${resource_group}${resource_name}\`.
    It likely maps to something like \`${PACKAGE}\`.
    Example:
    \`\`\`bash
    go run . generate-types \\
      --service ${PACKAGE} \\
      --api-version ${resource_group,,}.cnrm.cloud.google.com/v1beta1 \\
      --resource ${resource_group}${resource_name}:${resource_name} \\
      --include-skipped-output

    go run . generate-mapper \\
      --service ${PACKAGE} \\
      --api-version ${resource_group,,}.cnrm.cloud.google.com/v1beta1 \\
      --include-skipped-output
    \`\`\`

2.  Set the write permission on the new \`apis/${resource_group,,}/v1beta1/generate.sh\` file. You should do this by running both \`chmod +x apis/${resource_group,,}/v1beta1/generate.sh\` and \`git add --chmod=+x apis/${resource_group,,}/v1beta1/generate.sh\`.

3.  **Generate Scaffolding**:
    Run \`apis/${resource_group,,}/v1beta1/generate.sh\`. This should create \`apis/${resource_group,,}/v1beta1/${resource_name,,}_types.go\`.

4.  **Iterate on Types**:
    Compare the generated CRD with the existing one using \`dev/tasks/diff-crds\`.
    Modify \`apis/${resource_group,,}/v1beta1/${resource_name,,}_types.go\` until the CRD matches the existing one at \`config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_${resource_group,,}${resource_name,,}s.${resource_group,,}.cnrm.cloud.google.com.yaml\`.

    **Acceptance Criteria:**
    - Running \`dev/tasks/diff-crds\` should not show differences (or minimal acceptable ones like descriptions).
    - Ensure that running the check_crd_equivalence MCP on the CRD should return EQUIVALENT.
    - Changes to the schema (fields added/removed) are NOT acceptable.

5.  **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    \`\`\`go
    // Copyright 2026 Google LLC
    \`\`\`
    Please do not change the copyright on existing files.

6.  **Labels**:
    Ensure the controller-runtime annotations match the existing CRD labels, including:
    \`\`\`go
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/managed-by-kcc=true\"
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/system=true\"
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/stability-level=stable\"
    // +kubebuilder:metadata:labels=\"cnrm.cloud.google.com/${LABEL_NAME}=true\"
    \`\`\`
    The goal is to maintain these annotations, not add an annotation if it is missing.

7.  **Status**:
    \`status.observedGeneration\` should be an \`*int64\`.

8. **Generate Mappers**:
   - Running \`dev/tasks/generate-types-and-mappers\` will generate the mapper code once the \`apis/${resource_group,,}/v1beta1/${resource_name,,}_types.go\` file is generating an equivalent CRD.
   - Run \`make all-binary\` to ensure the generated mapper code compiles. Please fix any issue discovered by this compilation.

This issue is part of Epic #5954."
fi
