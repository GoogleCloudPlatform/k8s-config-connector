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

package logparser

import (
	"reflect"
	"testing"
)

func TestEntry_Predicates(t *testing.T) {
	tests := []struct {
		name              string
		entry             *Entry
		wantIsMutation    bool
		wantIsReadOnly    bool
		wantIsDelete      bool
		wantIsOpsPoll     bool
		wantRespIsSuccess bool
		wantStatusCode    int
	}{
		{
			name: "GET 200",
			entry: &Entry{
				Request:  HTTPRequest{Method: "GET", URL: "https://example.googleapis.com/v1/projects/p/locations/l/widgets/w1"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    false,
			wantIsReadOnly:    true,
			wantIsDelete:      false,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "POST create",
			entry: &Entry{
				Request:  HTTPRequest{Method: "POST", URL: "https://example.googleapis.com/v1/projects/p/locations/l/widgets?widgetId=w1"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    true,
			wantIsReadOnly:    false,
			wantIsDelete:      false,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "PATCH update",
			entry: &Entry{
				Request:  HTTPRequest{Method: "PATCH", URL: "https://example.googleapis.com/v1/projects/p/locations/l/widgets/w1?updateMask=description"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    true,
			wantIsReadOnly:    false,
			wantIsDelete:      false,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "DELETE",
			entry: &Entry{
				Request:  HTTPRequest{Method: "DELETE", URL: "https://example.googleapis.com/v1/projects/p/locations/l/widgets/w1"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    true,
			wantIsReadOnly:    false,
			wantIsDelete:      true,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "POST :destroy custom delete",
			entry: &Entry{
				Request:  HTTPRequest{Method: "POST", URL: "https://example.googleapis.com/v1/projects/p/locations/l/keys/k1:destroy"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    true,
			wantIsReadOnly:    false,
			wantIsDelete:      true,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "POST :getIamPolicy read-only",
			entry: &Entry{
				Request:  HTTPRequest{Method: "POST", URL: "https://example.googleapis.com/v1/projects/p/locations/l/widgets/w1:getIamPolicy"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    false,
			wantIsReadOnly:    true,
			wantIsDelete:      false,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "POST :setIamPolicy mutation",
			entry: &Entry{
				Request:  HTTPRequest{Method: "POST", URL: "https://example.googleapis.com/v1/projects/p/locations/l/widgets/w1:setIamPolicy"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200},
			},
			wantIsMutation:    true,
			wantIsReadOnly:    false,
			wantIsDelete:      false,
			wantIsOpsPoll:     false,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "GET operations/op-1 poll",
			entry: &Entry{
				Request:  HTTPRequest{Method: "GET", URL: "https://example.googleapis.com/v1/projects/p/locations/l/operations/op-1"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200, Body: `{"done": true}`},
			},
			wantIsMutation:    false,
			wantIsReadOnly:    true,
			wantIsDelete:      false,
			wantIsOpsPoll:     true,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
		{
			name: "POST operations/op-1:wait poll",
			entry: &Entry{
				Request:  HTTPRequest{Method: "POST", URL: "https://example.googleapis.com/v1/projects/p/locations/l/operations/op-1:wait"},
				Response: HTTPResponse{Status: "200 OK", StatusCode: 200, Body: `{"done": true}`},
			},
			wantIsMutation:    false,
			wantIsReadOnly:    false,
			wantIsDelete:      false,
			wantIsOpsPoll:     true,
			wantRespIsSuccess: true,
			wantStatusCode:    200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act & Assert
			if got := tc.entry.IsMutation(); got != tc.wantIsMutation {
				t.Errorf("IsMutation() = %v, want %v", got, tc.wantIsMutation)
			}
			if got := tc.entry.IsReadOnly(); got != tc.wantIsReadOnly {
				t.Errorf("IsReadOnly() = %v, want %v", got, tc.wantIsReadOnly)
			}
			if got := tc.entry.IsDelete(); got != tc.wantIsDelete {
				t.Errorf("IsDelete() = %v, want %v", got, tc.wantIsDelete)
			}
			if got := tc.entry.IsOperationsPoll(); got != tc.wantIsOpsPoll {
				t.Errorf("IsOperationsPoll() = %v, want %v", got, tc.wantIsOpsPoll)
			}
			if got := tc.entry.Response.IsSuccess(); got != tc.wantRespIsSuccess {
				t.Errorf("Response.IsSuccess() = %v, want %v", got, tc.wantRespIsSuccess)
			}
			if got := tc.entry.Response.ExtractStatusNumber(); got != tc.wantStatusCode {
				t.Errorf("Response.ExtractStatusNumber() = %d, want %d", got, tc.wantStatusCode)
			}
		})
	}
}

func TestEntry_URLAndMaskExtraction(t *testing.T) {
	tests := []struct {
		name              string
		entry             *Entry
		wantShortEndpoint string
		wantResIDSegment  string
		wantUpdateMask    []string
	}{
		{
			name: "query param updateMask and location short endpoint",
			entry: &Entry{
				Request: HTTPRequest{
					Method: "PATCH",
					URL:    "https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets/w1?%24alt=json&updateMask=description,displayName&widgetId=w1",
				},
			},
			wantShortEndpoint: ".../locations/us-central1/widgets/w1?widgetId=w1",
			wantResIDSegment:  "w1",
			wantUpdateMask:    []string{"description", "displayName"},
		},
		{
			name: "body string updateMask and custom action suffix stripped from ID segment",
			entry: &Entry{
				Request: HTTPRequest{
					Method: "POST",
					URL:    "https://example.googleapis.com/v1/projects/${projectId}/widgets/w2:customAction",
					Body:   `{"updateMask": "labels, priority"}`,
				},
			},
			wantShortEndpoint: ".../widgets/w2:customAction",
			wantResIDSegment:  "w2",
			wantUpdateMask:    []string{"labels", "priority"},
		},
		{
			name: "body paths object updateMask",
			entry: &Entry{
				Request: HTTPRequest{
					Method: "PATCH",
					URL:    "https://example.googleapis.com/v1/projects/${projectNumber}/global/networks/net-1",
					Body:   `{"updateMask": {"paths": ["autoCreateSubnetworks", "routingConfig.routingMode"]}}`,
				},
			},
			wantShortEndpoint: ".../global/networks/net-1",
			wantResIDSegment:  "net-1",
			wantUpdateMask:    []string{"autoCreateSubnetworks", "routingConfig.routingMode"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act & Assert
			if got := tc.entry.ShortEndpoint(); got != tc.wantShortEndpoint {
				t.Errorf("ShortEndpoint() = %q, want %q", got, tc.wantShortEndpoint)
			}
			if got := tc.entry.ExtractResourceIDSegment(); got != tc.wantResIDSegment {
				t.Errorf("ExtractResourceIDSegment() = %q, want %q", got, tc.wantResIDSegment)
			}
			if got := tc.entry.ExtractUpdateMask(); !reflect.DeepEqual(got, tc.wantUpdateMask) {
				t.Errorf("ExtractUpdateMask() = %v, want %v", got, tc.wantUpdateMask)
			}
		})
	}
}
