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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/api/option"
	api "google.golang.org/api/sqladmin/v1beta4"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
)

func TestSQLOperationError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		op      *api.Operation
		wantErr string // empty means no error expected
	}{
		{
			name: "nil operation",
			op:   nil,
		},
		{
			name: "successful operation",
			op:   &api.Operation{Status: "DONE"},
		},
		{
			name: "stockout",
			op: &api.Operation{
				Status: "DONE",
				Error: &api.OperationErrors{
					Errors: []*api.OperationError{
						{Code: "ERROR_STOCKOUT", Message: "The zone or region does not have sufficient resources to handle the request at the moment."},
					},
				},
			},
			wantErr: "ERROR_STOCKOUT: The zone or region does not have sufficient resources",
		},
		{
			name: "multiple errors",
			op: &api.Operation{
				Status: "DONE",
				Error: &api.OperationErrors{
					Errors: []*api.OperationError{
						{Code: "CODE_A", Message: "first"},
						nil,
						{Code: "CODE_B", Message: "second"},
					},
				},
			},
			wantErr: "CODE_A: first; CODE_B: second",
		},
		{
			name: "error without details",
			op: &api.Operation{
				Status: "DONE",
				Error:  &api.OperationErrors{},
			},
			wantErr: "operation reported an error without details",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := sqlOperationError(tc.op)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("sqlOperationError() = %v, want nil", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("sqlOperationError() = %v, want error containing %q", err, tc.wantErr)
			}
		})
	}
}

// newFakeSQLOperationsService returns an OperationsService backed by an httptest server
// that serves every request with the given status code and JSON body.
func newFakeSQLOperationsService(t *testing.T, statusCode int, body string) *api.OperationsService {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)

	svc, err := api.NewService(context.Background(), option.WithEndpoint(srv.URL+"/"), option.WithHTTPClient(srv.Client()))
	if err != nil {
		t.Fatalf("creating sqladmin service: %v", err)
	}
	return api.NewOperationsService(svc)
}

func TestPollForLROCompletion(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		initialOp  *api.Operation
		statusCode int    // response of the fake operations.get endpoint
		body       string // response of the fake operations.get endpoint
		wantErr    string // empty means no error expected
	}{
		{
			name:      "already done and successful",
			initialOp: &api.Operation{Name: "op-1", Status: "DONE"},
		},
		{
			name: "already done with error",
			initialOp: &api.Operation{Name: "op-1", Status: "DONE", Error: &api.OperationErrors{
				Errors: []*api.OperationError{{Code: "ERROR_STOCKOUT", Message: "no capacity"}},
			}},
			wantErr: "SQLInstance test-instance create operation op-1 failed: ERROR_STOCKOUT: no capacity",
		},
		{
			name:       "polled until done and successful",
			initialOp:  &api.Operation{Name: "op-1", Status: "PENDING"},
			statusCode: http.StatusOK,
			body:       `{"name": "op-1", "status": "DONE"}`,
		},
		{
			name:       "polled until done with error",
			initialOp:  &api.Operation{Name: "op-1", Status: "PENDING"},
			statusCode: http.StatusOK,
			body:       `{"name": "op-1", "status": "DONE", "error": {"kind": "sql#operationErrors", "errors": [{"kind": "sql#operationError", "code": "ERROR_STOCKOUT", "message": "no capacity"}]}}`,
			wantErr:    "SQLInstance test-instance create operation op-1 failed: ERROR_STOCKOUT: no capacity",
		},
		{
			// Previously this dereferenced the nil operation returned by Do() and panicked.
			name:       "getting the operation fails",
			initialOp:  &api.Operation{Name: "op-1", Status: "PENDING"},
			statusCode: http.StatusNotFound,
			body:       `{"error": {"code": 404, "message": "operation not found"}}`,
			wantErr:    "getting SQLInstance test-instance create operation op-1 failed",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := &sqlInstanceAdapter{projectID: "test-project", resourceID: "test-instance"}
			if tc.statusCode != 0 {
				a.sqlOperationsClient = newFakeSQLOperationsService(t, tc.statusCode, tc.body)
			}

			err := a.pollForLROCompletion(context.Background(), tc.initialOp, "create")
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("pollForLROCompletion() = %v, want nil", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("pollForLROCompletion() = %v, want error containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestCheckInstanceCanBeReconciled(t *testing.T) {
	t.Parallel()
	for _, state := range []string{"", "RUNNABLE", "PENDING_CREATE", "MAINTENANCE", "SUSPENDED"} {
		if err := checkInstanceCanBeReconciled("test-project", &api.DatabaseInstance{Name: "test-instance", State: state}); err != nil {
			t.Errorf("checkInstanceCanBeReconciled(state=%q) = %v, want nil", state, err)
		}
	}
	if err := checkInstanceCanBeReconciled("test-project", nil); err != nil {
		t.Errorf("checkInstanceCanBeReconciled(nil) = %v, want nil", err)
	}

	err := checkInstanceCanBeReconciled("test-project", &api.DatabaseInstance{Name: "test-instance", State: "FAILED"})
	if err == nil {
		t.Fatalf("checkInstanceCanBeReconciled(state=FAILED) = nil, want error")
	}
	for _, want := range []string{`instance "test-instance" is in state FAILED`, "--instance=test-instance --project=test-project"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("checkInstanceCanBeReconciled(state=FAILED) = %q, want it to contain %q", err, want)
		}
	}
}

// TestUpdate_FailedInstanceIsNotReported asserts that an instance in FAILED state is never
// reported as Ready/UpToDate, and that no GCP call (e.g. instances.update) is attempted.
// The GCP clients are nil, so any attempted call would panic.
func TestUpdate_FailedInstanceIsNotReported(t *testing.T) {
	t.Parallel()
	a := &sqlInstanceAdapter{
		projectID:  "test-project",
		resourceID: "test-instance",
		desired:    &krm.SQLInstance{},
		actual:     &api.DatabaseInstance{Name: "test-instance", State: "FAILED"},
		fieldMeta:  map[string]*FieldMetadata{},
	}
	u := &unstructured.Unstructured{Object: map[string]any{}}

	err := a.Update(context.Background(), directbase.NewUpdateOperation(lifecyclehandler.LifecycleHandler{}, nil, u))
	if err == nil || !strings.Contains(err.Error(), "is in state FAILED") {
		t.Fatalf("Update() = %v, want FAILED state error", err)
	}
	if _, found := u.Object["status"]; found {
		t.Errorf("Update() set status %v, want status to be left untouched", u.Object["status"])
	}
}
