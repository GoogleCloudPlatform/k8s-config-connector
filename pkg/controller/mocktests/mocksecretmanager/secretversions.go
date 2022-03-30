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
	"fmt"
	"net/http"
	"strconv"
	"time"

	secretmanager "google.golang.org/api/secretmanager/v1"
)

func (s *mockSecretManager) createSecretVersion(request *http.Request, project *projectData, secretID string) (*http.Response, error) {
	project.mutex.Lock()
	defer project.mutex.Unlock()

	secret := project.secrets[secretID]
	if secret == nil {
		return s.ErrorNotFound(request)
	}

	body := secretmanager.AddSecretVersionRequest{}
	if response := s.ParseRequest(request, &body); response != nil {
		return response, nil
	}

	if body.Payload == nil {
		body.Payload = &secretmanager.SecretPayload{}
	}

	secret.mutex.Lock()
	defer secret.mutex.Unlock()

	versionID := strconv.Itoa(len(secret.versions) + 1)

	secretVersion := &secretVersionData{}
	secretVersion.data.Name = fmt.Sprintf("projects/%d/secrets/%s/versions/%s", project.Number, secretID, versionID)
	secretVersion.data.CreateTime = time.Now().UTC().Format(time.RFC3339)
	secretVersion.data.State = "ENABLED"
	// TODO:
	// "replicationStatus": {
	//   "automatic": {}
	// },
	// "etag": "\"15e33ca226cf60\"",
	// "clientSpecifiedPayloadChecksum": true

	secretVersion.secrets = *body.Payload

	if secret.versions == nil {
		secret.versions = make(map[string]*secretVersionData)
	}
	secret.versions[versionID] = secretVersion

	return s.ReplyJSON(&secretVersion.data)
}

func (s *mockSecretManager) enableSecretVersion(request *http.Request, project *projectData, secretID string, versionID string) (*http.Response, error) {
	project.mutex.Lock()
	defer project.mutex.Unlock()

	secret := project.secrets[secretID]
	if secret == nil {
		return s.ErrorNotFound(request)
	}

	// body is empty
	body := struct{}{}
	if response := s.ParseRequest(request, &body); response != nil {
		return response, nil
	}

	secret.mutex.Lock()
	defer secret.mutex.Unlock()

	secretVersion := secret.versions[versionID]
	if secretVersion == nil {
		return s.ErrorNotFound(request)
	}
	secretVersion.data.State = "ENABLED"

	return s.ReplyJSON(&secretVersion.data)
}

func (s *mockSecretManager) getSecretVersion(request *http.Request, project *projectData, secretID string, versionID string) (*http.Response, error) {
	project.mutex.Lock()
	defer project.mutex.Unlock()

	secret := project.secrets[secretID]
	if secret == nil {
		return s.ErrorNotFound(request)
	}

	secret.mutex.Lock()
	defer secret.mutex.Unlock()

	secretVersion := secret.versions[versionID]
	if secretVersion == nil {
		return s.ErrorNotFound(request)
	}

	return s.ReplyJSON(&secretVersion.data)
}

func (s *mockSecretManager) accessSecret(request *http.Request, project *projectData, secretID string, versionID string) (*http.Response, error) {
	project.mutex.Lock()
	defer project.mutex.Unlock()

	secret := project.secrets[secretID]
	if secret == nil {
		return s.ErrorNotFound(request)
	}

	response := &secretmanager.AccessSecretVersionResponse{}

	secret.mutex.Lock()
	defer secret.mutex.Unlock()

	secretVersion := secret.versions[versionID]
	if secretVersion == nil {
		return s.ErrorNotFound(request)
	}
	response.Name = secretVersion.data.Name
	response.Payload = &secretVersion.secrets

	return s.ReplyJSON(response)
}
