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

package inputstream_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/inputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
)

const (
	targetProjectID      = "target-project"
	targetProjectNumber  = "111111111111"
	defaultProjectID     = "default-project-from-adc"
	defaultProjectNumber = "999999999999"
)

type recordedRequest struct {
	Method      string
	Host        string
	Path        string
	Query       url.Values
	UserProject string
}

// fakeGCP stands in for the Cloud Storage, Resource Manager and Cloud Asset
// Inventory APIs used by an asset inventory export, and records every request.
type fakeGCP struct {
	mu       sync.Mutex
	requests []recordedRequest
}

func (f *fakeGCP) RoundTrip(req *http.Request) (*http.Response, error) {
	f.mu.Lock()
	f.requests = append(f.requests, recordedRequest{
		Method:      req.Method,
		Host:        req.URL.Host,
		Path:        req.URL.Path,
		Query:       req.URL.Query(),
		UserProject: req.Header.Get("X-Goog-User-Project"),
	})
	f.mu.Unlock()

	status, body := http.StatusOK, "{}"
	switch {
	case req.Method == http.MethodPost && req.URL.Path == "/storage/v1/b": // buckets.insert
		body = `{"name":"export-bucket"}`
	case req.URL.Host == "cloudresourcemanager.googleapis.com": // projects.get
		id := strings.TrimPrefix(req.URL.Path, "/v1/projects/")
		number := map[string]string{targetProjectID: targetProjectNumber, defaultProjectID: defaultProjectNumber}[id]
		body = fmt.Sprintf(`{"projectId":%q,"projectNumber":%q}`, id, number)
	case strings.HasSuffix(req.URL.Path, ":exportAssets"), strings.Contains(req.URL.Path, "/operations/"):
		body = `{"name":"projects/1/operations/ExportAssets/RESOURCE/1","done":true}`
	case req.Method == http.MethodGet && req.URL.Query().Get("alt") == "media": // objects.get (download)
		body = `{"name":"//storage.googleapis.com/projects/_/buckets/b1","asset_type":"storage.googleapis.com/Bucket"}` + "\n"
	case req.Method == http.MethodDelete:
		status, body = http.StatusNoContent, ""
	}
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    req,
	}, nil
}

func (f *fakeGCP) recorded() []recordedRequest {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]recordedRequest(nil), f.requests...)
}

// setUp routes all HTTP traffic to a fake and makes the local "default
// project" differ from the project being exported. The default project comes
// from a service account application default credentials file, which
// gcp.GetDefaultProjectID() reads before it falls back to gcloud. No token is
// ever minted from that file.
func setUp(t *testing.T) *fakeGCP {
	t.Helper()
	adc := filepath.Join(t.TempDir(), "adc.json")
	content := fmt.Sprintf(`{"type":"service_account","project_id":%q,"client_email":"fake@%s.iam.gserviceaccount.com","private_key":"unused","token_uri":"https://oauth2.googleapis.com/token"}`,
		defaultProjectID, defaultProjectID)
	if err := os.WriteFile(adc, []byte(content), 0o600); err != nil {
		t.Fatalf("error writing %v: %v", adc, err)
	}
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", adc)

	fake := &fakeGCP{}
	original := http.DefaultTransport
	http.DefaultTransport = fake
	t.Cleanup(func() { http.DefaultTransport = original })
	return fake
}

// newDevNull returns a non-piped stdin so that NewAssetStream performs an export.
func newDevNull(t *testing.T) *os.File {
	t.Helper()
	f, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("error opening %v: %v", os.DevNull, err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}

func findRequest(t *testing.T, requests []recordedRequest, desc string, match func(recordedRequest) bool) recordedRequest {
	t.Helper()
	for _, r := range requests {
		if match(r) {
			return r
		}
	}
	t.Fatalf("no %v request was sent; requests: %+v", desc, requests)
	return recordedRequest{}
}

func isBucketInsert(r recordedRequest) bool {
	return r.Method == http.MethodPost && r.Path == "/storage/v1/b"
}

func isExportAssets(r recordedRequest) bool {
	return r.Method == http.MethodPost && strings.HasSuffix(r.Path, ":exportAssets")
}

func TestNewAssetStreamExportUsesProjectParameter(t *testing.T) {
	fake := setUp(t)
	params := &parameters.Parameters{ProjectID: targetProjectID, OAuth2Token: "fake-token"}

	stream, err := inputstream.NewAssetStream(params, newDevNull(t))
	if err != nil {
		t.Fatalf("NewAssetStream() returned error: %v", err)
	}
	defer stream.Close()

	requests := fake.recorded()
	bucketInsert := findRequest(t, requests, "buckets.insert", isBucketInsert)
	if got := bucketInsert.Query.Get("project"); got != targetProjectID {
		t.Errorf("temporary bucket was created in project %q, want %q (the --project value)", got, targetProjectID)
	}
	exportAssets := findRequest(t, requests, "exportAssets", isExportAssets)
	if want := "/v1/projects/" + targetProjectID + ":exportAssets"; exportAssets.Path != want {
		t.Errorf("exportAssets path = %q, want %q", exportAssets.Path, want)
	}
	if exportAssets.UserProject != targetProjectNumber {
		t.Errorf("exportAssets X-Goog-User-Project = %q, want %q (the --project number)", exportAssets.UserProject, targetProjectNumber)
	}
	for _, r := range requests {
		if strings.Contains(r.Path+"?"+r.Query.Encode(), defaultProjectID) || r.UserProject == defaultProjectNumber {
			t.Errorf("request used the default project instead of --project=%v: %+v", targetProjectID, r)
		}
	}
}

func TestNewAssetStreamFolderExportUsesDefaultProject(t *testing.T) {
	fake := setUp(t)
	params := &parameters.Parameters{FolderID: 123, OAuth2Token: "fake-token"}

	stream, err := inputstream.NewAssetStream(params, newDevNull(t))
	if err != nil {
		t.Fatalf("NewAssetStream() returned error: %v", err)
	}
	defer stream.Close()

	// There is no --project for folder and organization exports, so the
	// temporary bucket and the quota project are still the default project.
	requests := fake.recorded()
	bucketInsert := findRequest(t, requests, "buckets.insert", isBucketInsert)
	if got := bucketInsert.Query.Get("project"); got != defaultProjectID {
		t.Errorf("temporary bucket was created in project %q, want %q (the default project)", got, defaultProjectID)
	}
	exportAssets := findRequest(t, requests, "exportAssets", isExportAssets)
	if want := "/v1/folders/123:exportAssets"; exportAssets.Path != want {
		t.Errorf("exportAssets path = %q, want %q", exportAssets.Path, want)
	}
	if exportAssets.UserProject != defaultProjectNumber {
		t.Errorf("exportAssets X-Goog-User-Project = %q, want %q (the default project number)", exportAssets.UserProject, defaultProjectNumber)
	}
}
