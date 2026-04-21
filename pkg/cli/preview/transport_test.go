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

package preview

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockTransport struct {
	called bool
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.called = true
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func TestReadOnlyTransport(t *testing.T) {
	mock := &mockTransport{}
	transport := &readOnlyTransport{delegate: mock}

	tests := []struct {
		method         string
		expectedStatus int
		shouldCall     bool
	}{
		{"GET", http.StatusOK, true},
		{"HEAD", http.StatusOK, true},
		{"OPTIONS", http.StatusOK, true},
		{"POST", http.StatusForbidden, false},
		{"PUT", http.StatusForbidden, false},
		{"DELETE", http.StatusForbidden, false},
		{"PATCH", http.StatusForbidden, false},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			mock.called = false
			req := httptest.NewRequest(tt.method, "http://example.com", nil)
			resp, err := transport.RoundTrip(req)
			if err != nil {
				t.Fatalf("RoundTrip failed: %v", err)
			}
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			if mock.called != tt.shouldCall {
				t.Errorf("expected delegate called=%v, got %v", tt.shouldCall, mock.called)
			}
		})
	}
}
