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

package mockaiplatform

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if strings.Contains(url, "/notebookRuntimeTemplates") {
		replacements.RemovePath(".dataPersistentDiskSpec")
		replacements.RemovePath(".response.dataPersistentDiskSpec")
	}
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "aiplatform.googleapis.com") && !strings.Contains(event.URL(), "google.cloud.aiplatform") {
		return
	}

	if strings.Contains(event.URL(), "PipelineService") || strings.Contains(event.URL(), "pipelineJobs") || strings.Contains(event.URL(), "persistentResources") {
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.DeleteOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.DeleteOperationMetadata")
	}

	if strings.Contains(event.URL(), "PipelineService") || strings.Contains(event.URL(), "pipelineJobs") {
		event.VisitResponseStringValues(func(path string, value string) {
			if strings.HasSuffix(path, `["vertex-ai-pipelines-run-billing-id"]`) || strings.HasSuffix(path, `.vertex-ai-pipelines-run-billing-id`) {
				replacements.ReplaceStringValue(value, "619702208161644544")
			}
		})
	}

	if strings.Contains(event.URL(), "tensorboards") && strings.Contains(event.URL(), "experiments") {
		replacements.ReplaceStringValue("updateMask=description%2CdisplayName%2Clabels%2Csource", "updateMask=description%2CdisplayName%2Clabels")
	}

	if strings.Contains(event.URL(), "ReasoningEngineService") || strings.Contains(event.URL(), "reasoningEngines") {
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.CreateReasoningEngineOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.CreateReasoningEngineOperationMetadata")
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.UpdateReasoningEngineOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.UpdateReasoningEngineOperationMetadata")
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.DeleteOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.DeleteOperationMetadata")
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.ReasoningEngine", "type.googleapis.com/google.cloud.aiplatform.v1.ReasoningEngine")

		previsitReasoningEngine := func(val string) {
			if strings.Contains(val, "/reasoningEngines/") {
				tokens := strings.Split(val, "/")
				for i := 0; i < len(tokens)-1; i++ {
					if tokens[i] == "reasoningEngines" {
						id := tokens[i+1]
						if idx := strings.Index(id, "?"); idx != -1 {
							id = id[:idx]
						}
						if isNumeric(id) {
							replacements.ReplaceStringValue(id, "${reasoningEngineID}")
						}
					}
				}
			}
		}
		previsitReasoningEngine(event.URL())
		event.VisitRequestStringValues(func(path string, value string) {
			previsitReasoningEngine(value)
		})
		event.VisitResponseStringValues(func(path string, value string) {
			previsitReasoningEngine(value)
		})
	}

	if strings.Contains(event.URL(), "SpecialistPoolService") || strings.Contains(event.URL(), "specialistPools") {
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.CreateSpecialistPoolOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.CreateSpecialistPoolOperationMetadata")
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.UpdateSpecialistPoolOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.UpdateSpecialistPoolOperationMetadata")
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.DeleteOperationMetadata", "type.googleapis.com/google.cloud.aiplatform.v1.DeleteOperationMetadata")
		replacements.ReplaceStringValue("type.googleapis.com/google.cloud.aiplatform.v1beta1.SpecialistPool", "type.googleapis.com/google.cloud.aiplatform.v1.SpecialistPool")

		previsitSpecialistPool := func(val string) {
			if strings.Contains(val, "/specialistPools/") {
				tokens := strings.Split(val, "/")
				for i := 0; i < len(tokens)-1; i++ {
					if tokens[i] == "specialistPools" {
						id := tokens[i+1]
						if idx := strings.Index(id, "?"); idx != -1 {
							id = id[:idx]
						}
						if isNumeric(id) {
							replacements.ReplaceStringValue(id, "${specialistPoolID}")
						}
					}
				}
			}
		}
		previsitSpecialistPool(event.URL())
		event.VisitRequestStringValues(func(path string, value string) {
			previsitSpecialistPool(value)
		})
		event.VisitResponseStringValues(func(path string, value string) {
			previsitSpecialistPool(value)
		})
	}
}
