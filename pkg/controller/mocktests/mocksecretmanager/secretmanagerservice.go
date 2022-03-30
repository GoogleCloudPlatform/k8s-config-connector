// Copyright 2022 Google LLC
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

package mocksecretmanager

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/mocktests/mockbase"
	option "google.golang.org/api/option"
	secretmanager "google.golang.org/api/secretmanager/v1"
	"k8s.io/klog/v2"
)

const ExpectedHost = "secretmanager.googleapis.com"

// mockSecretManager represents a mocked secret manager service.
type mockSecretManager struct {
	svc *secretmanager.Service

	mutex            sync.Mutex
	projectsByID     map[string]*projectData
	projectsByNumber map[int64]*projectData

	mockbase.ServiceBase
}

type projectData struct {
	Number int64
	ID     string

	mutex   sync.Mutex
	secrets map[string]*secretData
}

type secretData struct {
	data secretmanager.Secret

	mutex    sync.Mutex
	versions map[string]*secretVersionData
}

type secretVersionData struct {
	data    secretmanager.SecretVersion
	secrets secretmanager.SecretPayload
}

// NewMock creates a mockSecretManager
func NewMock(expectedHost string) MockService {
	s := &mockSecretManager{
		projectsByID:     make(map[string]*projectData),
		projectsByNumber: make(map[int64]*projectData),
	}
	s.ServiceBase.Init(expectedHost, s)
	return s
}

type MockService interface {
	http.RoundTripper
	NewService(ctx context.Context) (*secretmanager.Service, error)
}

// NewService creates a new mock cloudresourcemanager client.
func NewService(ctx context.Context) (*secretmanager.Service, error) {
	s := NewMock(ExpectedHost)
	return s.NewService(ctx)
}

func (s *mockSecretManager) getProject(projectID string) *projectData {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var project *projectData

	projectNumber, err := strconv.ParseInt(projectID, 10, 64)
	if err == nil {
		project = s.projectsByNumber[projectNumber]
	} else {
		project = s.projectsByID[projectID]
		if project == nil {
			project = &projectData{
				Number:  123, // TODO: Fix projectid
				ID:      projectID,
				secrets: make(map[string]*secretData),
			}
			s.projectsByID[project.ID] = project
			s.projectsByNumber[project.Number] = project
		}
	}
	return project
}

func (s *mockSecretManager) NewService(ctx context.Context) (*secretmanager.Service, error) {
	httpClient := &http.Client{Transport: s}
	svc, err := secretmanager.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("failed to build mock secretmanager service: %w", err)
	}
	s.svc = svc

	return svc, nil
}

func (s *mockSecretManager) ProcessRequest(request *http.Request) (*http.Response, error) {
	url := request.URL

	pathTokens := strings.Split(strings.TrimPrefix(url.Path, "/"), "/")

	if err := request.ParseForm(); err != nil {
		klog.Warningf("error parsing form: %s %s %v", request.Method, request.URL, err)
		httpResponse := &http.Response{
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest,
		}
		return httpResponse, nil
	}
	if len(pathTokens) > 0 && (pathTokens[0] == "v1") {
		// version := pathTokens[0]

		if len(pathTokens) >= 3 && pathTokens[1] == "projects" {
			project := s.getProject(pathTokens[2])
			if project == nil {
				klog.Infof("project %q not found", pathTokens[2])
				httpResponse := &http.Response{
					Status:     http.StatusText(http.StatusNotFound),
					StatusCode: http.StatusNotFound,
				}
				return httpResponse, nil
			}

			// TODO: Lock project here?

			if len(pathTokens) == 5 && pathTokens[3] == "secrets" {
				id := pathTokens[4]
				if strings.HasSuffix(id, ":addVersion") {
					id = strings.TrimSuffix(id, ":addVersion")
					if request.Method == "POST" {
						return s.createSecretVersion(request, project, id)
					}
				} else {
					if request.Method == "GET" {
						return s.getSecret(request, project, id)
					}
				}

				return s.ErrorMethodNotAllowed(request)
			}

			if len(pathTokens) == 4 && pathTokens[3] == "secrets" {

				if request.Method == "POST" {
					return s.postSecret(request, project)
				}
				return s.ErrorMethodNotAllowed(request)
			}

			if len(pathTokens) == 7 && pathTokens[3] == "secrets" && pathTokens[5] == "versions" {
				secretID := pathTokens[4]
				versionID := pathTokens[6]
				if strings.HasSuffix(versionID, ":enable") {
					versionID = strings.TrimSuffix(versionID, ":enable")

					if request.Method == "POST" {
						return s.enableSecretVersion(request, project, secretID, versionID)
					}
				} else if strings.HasSuffix(versionID, ":access") {
					versionID = strings.TrimSuffix(versionID, ":access")

					if request.Method == "GET" {
						return s.accessSecret(request, project, secretID, versionID)
					}
				} else {
					if request.Method == "GET" {
						return s.getSecretVersion(request, project, secretID, versionID)
					}
				}

				return s.ErrorMethodNotAllowed(request)
			}
		}
	}

	klog.Warningf("unhandled request: %s %s %#v", request.Method, request.URL, request)
	httpResponse := &http.Response{
		Status:     http.StatusText(http.StatusNotFound),
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
	}
	return httpResponse, nil
}
