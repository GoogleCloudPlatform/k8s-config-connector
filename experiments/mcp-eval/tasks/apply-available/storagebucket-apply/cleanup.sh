# # #!/usr/bin/env bash
# REPO_ROOT="$(git rev-parse --show-toplevel)"
kubectl delete storagebuckets/storagebucket-apply -n storagebucket-apply --ignore-not-found=true

gcloud storage buckets delete gs://storagebucket-apply
