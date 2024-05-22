.PHONY: generate-pb
generate-pb:
	mkdir -p third_party
	mkdir -p build
	git clone https://github.com/googleapis/googleapis.git third_party/googleapis || (cd third_party/googleapis && git reset --hard origin/master && git pull)
	protoc --include_imports --include_source_info \
		-I ./third_party/googleapis/ \
		./third_party/googleapis/google/api/*.proto \
		./third_party/googleapis/google/api/*/*/*.proto \
		./third_party/googleapis/google/cloud/*/*/*.proto \
		./third_party/googleapis/google/iam/v1/*.proto \
		./third_party/googleapis/google/logging/v2/*.proto \
		./third_party/googleapis/google/monitoring/dashboard/v1/*.proto \
		-o build/googleapis.pb
	go run . -apis ../../../apis/ --api-packages github.com/GoogleCloudPlatform/apis
