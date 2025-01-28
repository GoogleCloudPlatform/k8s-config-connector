// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockgcpregistry

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

// MockService is the interface implemented by all services
type MockService interface {
	// Register initializes the service, normally registering the GRPC service.
	Register(grpcServer *grpc.Server)

	// NewHTTPMux creates an HTTP mux for serving http traffic
	NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error)

	// ExpectedHosts is the hostname(s) we serve on e.g. foo.googleapis.com
	// We also support patterns like `{region}-foo.googleapis.com`
	ExpectedHosts() []string
}

// SupportsNormalization can be implemented to support normalization
type SupportsNormalization interface {
	ConfigureVisitor(url string, replacements NormalizingVisitor)
}

type NormalizingVisitor interface {
	ReplacePath(path string, newValue any)
}

type Normalizer interface {
	ConfigureVisitor(url string, replacements NormalizingVisitor)
}
