# #!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
# kubectl delete -f ${REPO_ROOT}/.build/tasks/storagebucket-valid/storagebucket-valid.yaml --ignore-not-found=true
rm -f ${REPO_ROOT}/.build/tasks/storagebucket-1-ubla/*.yaml
