// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	conreg "github.com/GoogleCloudPlatform/declarative-resource-client-library/connector/server_registration"
	connectorpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/connector_go_proto"
)

const (
	// CredentialsMetadataKey contains the user-provided GCP credentials.
	CredentialsMetadataKey = "X-Call-Credentials"

	// UserAgentMetadataKey is an optional value used to override the default when making GCP API calls.
	UserAgentMetadataKey = "User-Agent"
)

// Set by InitializeServer().
var grpcServer *grpc.Server

// BodyWriter implements io.Writer and io.Flusher as required by http.ResponseWriter
type BodyWriter struct {
	header http.Header
	b      bytes.Buffer
	code   int
}

// Header returns the current http headers for writing.
func (b *BodyWriter) Header() http.Header {
	return b.header
}

// Write appends data to the response body.
func (b *BodyWriter) Write(data []byte) (int, error) {
	return b.b.Write(data)
}

// WriteHeader writes the http status code.
func (b *BodyWriter) WriteHeader(statusCode int) {
	b.code = statusCode
}

// Flush is required by the io.Flusher interface but is unused here.
func (b *BodyWriter) Flush() {}

// InitializeServer prepares the server for future RPC requests. It must be called before
// attempting to response to any requests.
func InitializeServer() *connectorpb.InitializeResponse {
	grpcServer = grpc.NewServer()
	return conreg.InitializeServer(grpcServer)
}

// UnaryCall directs the gRPC call from the client library to its handler. Requires a prior call to InitializeServer().
func UnaryCall(request *connectorpb.UnaryCallRequest) *connectorpb.UnaryCallResponse {
	if grpcServer == nil {
		return makeErrorResponse(codes.Unavailable, "call the initialize function first")
	}

	r, err := makeHTTPRequest(request)
	if err != nil {
		return makeErrorResponse(codes.Internal, "could not prepare an HTTP request")
	}

	w := &BodyWriter{
		header: make(http.Header),
	}

	grpcServer.ServeHTTP(w, r)

	// While the server sets common HTTP header (e.g., Content-Type) we only care
	// about Grpc-Status and Grpc-Message.
	grpcStatus := codes.Internal
	grpcMessage := "grpc-message field not found"
	for h, v := range w.header {
		if len(v) == 1 {
			if strings.ToLower(h) == "grpc-status" {
				i, err := strconv.Atoi(v[0])
				if err != nil {
					return makeErrorResponse(codes.Internal, "could not parse Grpc-Status header")
				}
				grpcStatus = codes.Code((i))
			}
			if strings.ToLower(h) == "grpc-message" {
				grpcMessage = v[0]
			}
		}
	}

	if grpcStatus != codes.OK {
		return makeErrorResponse(grpcStatus, grpcMessage)
	}

	// Parse the length-prefixed message response body:
	// <1 byte compression flag><4 byte big-endian length><serialized response>
	lpm := w.b.Bytes()
	if len(lpm) < 5 {
		return makeErrorResponse(codes.Internal, "missing response body")
	}

	length := (int(lpm[1]) << 24) | (int(lpm[2]) << 16) | (int(lpm[3]) << 8) | int(lpm[4])
	if len(lpm) < 5+length {
		return makeErrorResponse(codes.Internal, "truncated response body")
	}

	return connectorpb.UnaryCallResponse_builder{
		Response: lpm[5:],
		Status: &statuspb.Status{
			Code:    int32(codes.OK),
			Message: grpcMessage,
		},
	}.Build()
}

func makeErrorResponse(code codes.Code, message string) *connectorpb.UnaryCallResponse {
	return connectorpb.UnaryCallResponse_builder{
		Status: &statuspb.Status{
			Code:    int32(code),
			Message: message,
		},
	}.Build()
}

func makeHTTPRequest(request *connectorpb.UnaryCallRequest) (*http.Request, error) {
	// Convert the serialized request proto to a length-prefixed message
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
	lpm := []byte{
		0, // Not compressed
		byte(len(request.GetRequest()) >> 24),
		byte(len(request.GetRequest()) >> 16),
		byte(len(request.GetRequest()) >> 8),
		byte(len(request.GetRequest())),
	}
	lpm = append(lpm, request.GetRequest()...)

	// Wrap the request body in a suitable HTTP request
	u, err := url.Parse("https://localhost" + request.GetMethod())
	if err != nil {
		return nil, err
	}

	return &http.Request{
		Method:     "POST",
		RequestURI: request.GetMethod(),
		URL:        u,
		Proto:      "HTTP/2",
		ProtoMajor: 2,
		ProtoMinor: 0,
		Header: http.Header{
			"Content-Type":         {"application/grpc+proto"},
			CredentialsMetadataKey: {request.GetCredentials()},
			UserAgentMetadataKey:   {request.GetUserAgent()},
		},
		Trailer:       make(http.Header),
		ContentLength: -1,
		Body:          io.NopCloser(bytes.NewReader(lpm)),
	}, nil
}
