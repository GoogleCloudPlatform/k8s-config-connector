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
	// ConfigureVisitor sets up simple replacements
	ConfigureVisitor(url string, replacements NormalizingVisitor)

	// Previsit visits each object.  This is used to extract values to replace with placeholders
	Previsit(event Event, visitor NormalizingVisitor)
}

type NormalizingVisitor interface {
	// ReplacePath replaces values at the given path with newValue
	ReplacePath(path string, newValue any)

	// RemovePath removes values at the given path
	RemovePath(path string)

	// ReplaceStringValue replaces the given string value with the provided string value
	ReplaceStringValue(oldValue string, newValue string)

	// SortSlice will sort the slice at the given path
	SortSlice(path string)
}

type Normalizer interface {
	ConfigureVisitor(url string, replacements NormalizingVisitor)

	// Previsit visits each request, and is used to find placeholder values that may span events
	Previsit(entry Event, replacements NormalizingVisitor)
}

type Event interface {
	// URL returns the URL of the request
	URL() string

	// VisitRequestStringValues calls the callback for each string-typed value found in the request object (if any)
	VisitRequestStringValues(callback func(path string, value string))

	// VisitResponseStringValues calls the callback for each string-typed value found in the response object (if any)
	VisitResponseStringValues(callback func(path string, value string))
}
