

wget https://raw.githubusercontent.com/googleapis/google-api-go-client/main/servicenetworking/v1/servicenetworking-api.json

mkdir -p ../../apis/mockgcp/cloud/servicenetworking/v1/
go run . servicenetworking-api.json > ../../apis/mockgcp/cloud/servicenetworking/v1/servicenetworking.proto



wget https://raw.githubusercontent.com/googleapis/google-api-go-client/main/bigquery/v2/bigquery-api.json

mkdir -p ../../apis/mockgcp/cloud/bigquery/v2/
go run . bigquery-api.json > ../../apis/mockgcp/cloud/bigquery/v2/api.proto



wget https://raw.githubusercontent.com/googleapis/google-api-go-client/main/storage/v1/storage-api.json
mkdir -p ../../apis/mockgcp/storage/v1/
go run . storage-api.json > ../../apis/mockgcp/storage/v1/generated.proto
