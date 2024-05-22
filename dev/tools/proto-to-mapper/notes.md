REPO_ROOT="$(git rev-parse --show-toplevel)"

protoc --include_imports --include_source_info \
  -I ${REPO_ROOT}/mockgcp/third_party/googleapis/ \
  ${REPO_ROOT}/mockgcp/third_party/googleapis/mockgcp/monitoring/dashboard/v1/*.proto \
  -o monitoring.pb

protoc --include_imports --include_source_info \
  -I ${REPO_ROOT}/mockgcp/third_party/googleapis/ \
  ${REPO_ROOT}/mockgcp/third_party/googleapis/mockgcp/cloud/redis/cluster/v1/*.proto \
  -o redis.pb

protoc --include_imports --include_source_info \
  -I ${REPO_ROOT}/mockgcp/third_party/googleapis/ \
  ${REPO_ROOT}/mockgcp/third_party/googleapis/mockgcp/cloud/networkconnectivity/v1/*.proto \
  -o networkconnectivity.pb
