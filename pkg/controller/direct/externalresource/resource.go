/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package externalresource

import (
	"reflect"
)

type ExternalResourceReference struct {
	hasSelfLink bool
	externalRef string
}

type GCPResource interface {
	GetName() string
}

// New builds the External Reference for CloudBuildWorkerPool when the object already exists on the GCP server.
// Ideally, it expects to use the GCP object's selfLink value as the external reference. However, if
// the resource does not have `selfLink` (like CloudBuildWorkerPool), it builds the URL from the GCP object.
func New(serviceBaseURL string, gcpObj GCPResource) *ExternalResourceReference {
	baseExternalRef := &ExternalResourceReference{
		hasSelfLink: false,
		externalRef: serviceBaseURL + gcpObj.GetName(),
	}
	// Ignore the actual GCP proto, get the `selfLink` in general.
	s := reflect.ValueOf(gcpObj).Elem()
	if s.Kind() != reflect.Struct {
		return baseExternalRef
	}

	selfLink := s.FieldByName("SelfLink")

	if !selfLink.IsValid() || selfLink.IsZero() {
		return baseExternalRef
	}
	switch selfLink.Kind() {
	case reflect.String:
		baseExternalRef.hasSelfLink = true
		baseExternalRef.externalRef = selfLink.String()
	case reflect.Pointer:
		baseExternalRef.hasSelfLink = true
		baseExternalRef.externalRef = selfLink.Elem().String()
	}
	return baseExternalRef
}

func (e *ExternalResourceReference) Get() *string {
	return &e.externalRef
}
