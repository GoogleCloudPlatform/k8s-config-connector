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

package preview

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

// BlockedGCPError is an error that occurs when a GCP API call is blocked.
type BlockedGCPError struct {
	Method string
	URL    string
	Body   string
}

var _ error = &BlockedGCPError{}

// Error implements the error interface.
func (e BlockedGCPError) Error() string {
	// encode in json so that we can unwrap even if this is (incorrectly) wrapped as a string (terraform)
	j, _ := json.Marshal(e)
	return fmt.Sprintf("call to GCP blocked (method=%v, url=%v) [jsonstart:BlockedGCPError:%v:jsonend]", e.Method, e.URL, string(j))
}

// ExtractBlockedGCPError will unwrap a BlockedGCPError.
// To tolerate terraform using string-wrapping of error messages, we also parse a json-encoded form.
func ExtractBlockedGCPError(err error) (*BlockedGCPError, bool) {
	var e *BlockedGCPError
	if errors.As(err, &e) {
		return e, true
	}

	// Look for a string-wrapped message
	s := err.Error()
	for {
		_, tail, ok := strings.Cut(s, " [jsonstart:BlockedGCPError:")
		if !ok {
			break
		}
		body, tail, ok := strings.Cut(tail, ":jsonend]")
		if !ok {
			break
		}

		s = tail
		var e BlockedGCPError
		if err := json.Unmarshal([]byte(body), &e); err != nil {
			klog.Warningf("error parsing json-encoded error %q: %v", body, err)
			continue
		} else {
			return &e, true
		}
	}
	return nil, false
}

// interceptingGCPClient is a GCP client that intercepts GCP API calls.
// It forwards read-only operations "upstream" to real GCP.
// It returns a BlockedGCPError on any write operations.
type interceptingGCPClient struct {
	upstreamGCPClient *http.Client
	authorization     oauth2.TokenSource
}

// newInterceptingGCPClient creates a new interceptingGCPClient.
func newInterceptingGCPClient(upstreamGCPClient *http.Client, authorization oauth2.TokenSource) *interceptingGCPClient {
	return &interceptingGCPClient{
		upstreamGCPClient: upstreamGCPClient,
		authorization:     authorization,
	}
}

// HTTPClient is the HTTP client that should be used for GCP API calls.
func (c *interceptingGCPClient) HTTPClient() *http.Client {
	return &http.Client{
		Transport: c.HTTPRoundTripper(),
	}
}

// HTTPRoundTripper is the round tripper that should be used for GCP API calls.
func (c *interceptingGCPClient) HTTPRoundTripper() http.RoundTripper {
	return c
}

// blockedHTTPMethod is called when a write operation is attempted.
// It returns a BlockedGCPError.
func (c *interceptingGCPClient) blockedHTTPMethod(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	log := klog.FromContext(ctx)
	var body []byte
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading body: %w", err)
		} else {
			body = b
		}
	}
	log.Info("blockedHTTPMethod", "req.method", req.Method, "req.url", req.URL.String())
	return nil, BlockedGCPError{
		Method: req.Method,
		URL:    req.URL.String(),
		Body:   string(body),
	}
}

// RoundTrip is the round tripper that implements the http.RoundTripper interface.
func (c *interceptingGCPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	log := klog.FromContext(ctx)

	requestIsAllowed := false
	if req.Method == "GET" {
		requestIsAllowed = true
	}

	if requestIsAllowed {
		if c.authorization != nil {
			token, err := c.authorization.Token()
			if err != nil {
				return nil, err
			}
			req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		}
		response, err := c.upstreamGCPClient.Do(req)
		if response != nil {
			log.Info("forwarded request", "req.method", req.Method, "req.url", req.URL, "response.status", response.Status)
		} else if err != nil {
			log.Error(err, "error forwarding request", "req.method", req.Method, "req.url", req.URL)
		}

		return response, err
	}

	return c.blockedHTTPMethod(req)
}

// GRPCUnaryClientInterceptor intercepts GCP GRPC API calls.
func (c *interceptingGCPClient) GRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// TODO: Add this back in to support GRPC resources (e.g. bigtable)
		// entry := &test.LogEntry{}

		// entry.Request.URL = method
		// entry.Request.Method = "GRPC"

		// if req != nil {
		// 	requestBytes, _ := protojson.Marshal(req.(proto.Message))
		// 	entry.Request.Body = string(requestBytes)
		// }

		// if mockCloudGRPCClientConnection != nil {
		// 	cc = mockCloudGRPCClientConnection
		// }
		// err := invoker(ctx, method, req, reply, cc, opts...)

		// if reply != nil {
		// 	replyBytes, _ := protojson.Marshal(reply.(proto.Message))
		// 	entry.Response.Body = string(replyBytes)
		// }

		// if err != nil {
		// 	entry.Response.Status = fmt.Sprintf("error: %v", err)
		// } else {
		// 	entry.Response.Status = "OK"
		// }

		// for _, eventSink := range eventSinks {
		// 	eventSink.AddHTTPEvent(ctx, entry)
		// }
		// return err
		return fmt.Errorf("GRPC method blocked by InterceptingGCPClient")
	}
}
