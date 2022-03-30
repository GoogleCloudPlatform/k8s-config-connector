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
	"time"

	"google.golang.org/api/secretmanager/v1"
)

func (s *mockSecretManager) getSecret(request *http.Request, project *projectData, secretName string) (*http.Response, error) {
	project.mutex.Lock()
	defer project.mutex.Unlock()

	secret := project.secrets[secretName]
	if secret == nil {
		return s.ErrorNotFound(request)
	}

	return s.ReplyJSON(secret.data)
}

func (s *mockSecretManager) postSecret(r *http.Request, project *projectData) (*http.Response, error) {
	request := secretmanager.Secret{}
	if response := s.ParseRequest(r, &request); response != nil {
		return response, nil
	}

	secretId := r.Form.Get("secretId")
	if secretId == "" {
		return nil, fmt.Errorf("secretId is required")
	}

	create := &secretData{}
	create.data = request
	create.data.Name = fmt.Sprintf("projects/%d/secrets/%s", project.Number, secretId)
	create.data.CreateTime = time.Now().UTC().Format(time.RFC3339)

	project.mutex.Lock()
	defer project.mutex.Unlock()

	project.secrets[secretId] = create

	return s.ReplyJSON(create.data)
}
