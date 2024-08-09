# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.21.3

RUN apt-get update && apt-get install zip wget

WORKDIR /tmp/protoc
RUN wget -O protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-linux-x86_64.zip
RUN unzip protoc.zip && cp bin/protoc /usr/bin/protoc
RUN rm -rf /tmp/protoc

RUN GOBIN=/usr/bin/ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.3 \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.3
RUN GOBIN=/usr/bin/ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN GOBIN=/usr/bin/ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN mkdir /src
WORKDIR /src

ENTRYPOINT ["/bin/protoc"]