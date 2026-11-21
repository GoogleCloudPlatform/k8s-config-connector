// Copyright 2026 Google LLC
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

package sql

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"google.golang.org/api/option"
	api "google.golang.org/api/sqladmin/v1beta4"
)

type reconcileKey struct{}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func adapterWithTransport(t *testing.T, transport http.RoundTripper) *sqlInstanceAdapter {
	t.Helper()
	service, err := api.NewService(context.Background(),
		option.WithHTTPClient(&http.Client{Transport: transport}),
		option.WithEndpoint("https://sqladmin.test/"))
	if err != nil {
		t.Fatalf("NewService: %v", err)
	}
	return &sqlInstanceAdapter{
		projectID:           "project-a",
		resourceID:          "instance-a",
		sqlOperationsClient: api.NewOperationsService(service),
	}
}

func TestPollForLROCompletionUsesTheContext(t *testing.T) {
	adapter := adapterWithTransport(t, roundTripFunc(func(r *http.Request) (*http.Response, error) {
		// Credentials can travel on the context, so the poll must use it.
		if r.Context().Value(reconcileKey{}) == nil {
			return nil, errors.New("operation polled without the reconcile context")
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"name": "op-1", "status": "DONE"}`)),
		}, nil
	}))
	ctx := context.WithValue(context.Background(), reconcileKey{}, true)

	if err := adapter.pollForLROCompletion(ctx, &api.Operation{Name: "op-1", Status: "RUNNING"}, "update"); err != nil {
		t.Fatalf("pollForLROCompletion: %v", err)
	}
}

func TestPollForLROCompletionReportsPollErrors(t *testing.T) {
	adapter := adapterWithTransport(t, roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("no credentials")
	}))

	err := adapter.pollForLROCompletion(context.Background(), &api.Operation{Name: "op-1", Status: "RUNNING"}, "update")
	if err == nil || !strings.Contains(err.Error(), "op-1") || !strings.Contains(err.Error(), "no credentials") {
		t.Errorf("got error %v, want one naming the operation and the cause", err)
	}
}
