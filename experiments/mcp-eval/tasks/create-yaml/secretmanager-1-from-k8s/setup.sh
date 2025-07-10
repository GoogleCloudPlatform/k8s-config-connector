#!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
BUILD_DIR="${REPO_ROOT}/.build/tasks/secretmanager-1-from-k8s"
mkdir -p "${BUILD_DIR}"
rm -f "${BUILD_DIR}"/*

# Create the Kubernetes secret
cat <<EOF > "${BUILD_DIR}/k8s-secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: my-k8s-secret
stringData:
  secret-data: "super-secret-value"
EOF

# We would normally apply this with kubectl, but for this test, 
# the file's existence is enough for the prompt.
echo "Created k8s-secret.yaml for reference."

