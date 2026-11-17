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

package mockfirebase

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".apiKeyId", "6d1ee53e-9817-4331-ac63-7fdfba6d5e07")
	replacements.ReplacePath(".webId", "MGEzYWJkY2EtZDQ3OS00Y2VjLTgxNjktNzA1YTdhOTZjY2E4")
	replacements.ReplacePath(".etag", "abcdef0123A=")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "firebase.googleapis.com") {
		return
	}

	previsitValue := func(val string) {
		if strings.Contains(val, "/operations/workflows/") {
			tokens := strings.Split(val, "/operations/workflows/")
			if len(tokens) == 2 {
				id := tokens[1]
				if idx := strings.Index(id, "?"); idx != -1 {
					id = id[:idx]
				}
				replacements.ReplaceStringValue(id, "ZTE1YTc1ZDgtMTM3Yy00YTMzLTgzNzUtY2E0NGFjN2EzYTkw")
			}
		}
		if strings.Contains(val, "/androidApps/") {
			tokens := strings.Split(val, "/androidApps/")
			if len(tokens) == 2 {
				id := tokens[1]
				if idx := strings.Index(id, "?"); idx != -1 {
					id = id[:idx]
				}
				if idx := strings.Index(id, ":remove"); idx != -1 {
					id = id[:idx]
				}
				replacements.ReplaceStringValue(id, "1:${projectNumber}:android:e0ec16c940d75f1d8b7d90")
			}
		}
		if strings.Contains(val, "/webApps/") {
			tokens := strings.Split(val, "/webApps/")
			if len(tokens) == 2 {
				id := tokens[1]
				if idx := strings.Index(id, "?"); idx != -1 {
					id = id[:idx]
				}
				if idx := strings.Index(id, ":remove"); idx != -1 {
					id = id[:idx]
				}
				replacements.ReplaceStringValue(id, "1:${projectNumber}:web:d845db9c6118f7948b7d90")
			}
		}
	}

	previsitValue(event.URL())

	event.VisitRequestStringValues(func(path string, value string) {
		previsitValue(value)
	})

	event.VisitResponseStringValues(func(path string, value string) {
		previsitValue(value)
	})
}
