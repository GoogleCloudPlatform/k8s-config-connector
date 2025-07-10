# #!/usr/bin/env bash
REPO_ROOT="$(git rev-parse --show-toplevel)"
kubectl delete -f ${REPO_ROOT}/.build/tasks/storagebucket-3-apply/storagebucket-apply.yaml --ignore-not-found=true

gcloud storage buckets delete gs://storagebucket-apply
rm -f ${REPO_ROOT}/.build/tasks/storagebucket-3-apply/*.yaml
