// Copyright 2025 Google LLC
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

package mockgcpregistry

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type FactoryFunc func(env *common.MockEnvironment, storage storage.Storage) MockService

func Register(factory FactoryFunc) {
	factories = append(factories, factory)
}

var factories []FactoryFunc

func BuildAllServices(env *common.MockEnvironment, storage storage.Storage) (*Services, error) {
	ret := &Services{}
	for _, factory := range factories {
		service := factory(env, storage)
		ret.Services = append(ret.Services, service)
	}
	return ret, nil
}

type Services struct {
	Services []MockService
}

func (s *Services) ConfigureVisitor(url string, visitor NormalizingVisitor) {
	for _, service := range s.Services {
		supportsNormalization, ok := service.(SupportsNormalization)
		if !ok {
			continue
		}
		supportsNormalization.ConfigureVisitor(url, visitor)
	}
}

func (s *Services) Previsit(event Event, visitor NormalizingVisitor) {
	for _, service := range s.Services {
		supportsNormalization, ok := service.(SupportsNormalization)
		if !ok {
			continue
		}
		supportsNormalization.Previsit(event, visitor)
	}
}

func (s *Services) ConfigureKRMObjectVisitor(u *unstructured.Unstructured, visitor NormalizingVisitor) {
	for _, service := range s.Services {
		supportsKRMObjectNormalization, ok := service.(SupportsKRMObjectNormalization)
		if !ok {
			continue
		}
		supportsKRMObjectNormalization.ConfigureKRMObjectVisitor(u, visitor)
	}
}
