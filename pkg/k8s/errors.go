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

package k8s

import (
	"errors"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type ErrorWithReason struct {
	Message string
	Reason  string
}

func (e ErrorWithReason) Error() string {
	return fmt.Sprintf("%v: %v", e.Reason, e.Message)
}

func NewReferenceNotReadyErrorForResource(r *Resource) *refs.ReferenceNotReadyError {
	return refs.NewReferenceNotReadyError(
		r.GroupVersionKind(),
		types.NamespacedName{Namespace: r.GetNamespace(), Name: r.GetName()},
	)
}

func NewReferenceNotFoundErrorForResource(r *Resource) *refs.ReferenceNotFoundError {
	return refs.NewReferenceNotFoundError(
		r.GroupVersionKind(),
		types.NamespacedName{Namespace: r.GetNamespace(), Name: r.GetName()},
	)
}

var ErrIAMNotFound = fmt.Errorf("IAM resource does not exist")

type TransitiveDependencyNotFoundError struct {
	ResourceGVK schema.GroupVersionKind
	Resource    types.NamespacedName
}

func (e *TransitiveDependencyNotFoundError) Error() string {
	return fmt.Sprintf("transitive dependency %v %v is not found", e.ResourceGVK.Kind, e.Resource)
}

func NewTransitiveDependencyNotFoundError(resourceGVK schema.GroupVersionKind, resource types.NamespacedName) *TransitiveDependencyNotFoundError {
	return &TransitiveDependencyNotFoundError{resourceGVK, resource}
}

func AsTransitiveDependencyNotFoundError(err error) (unwrappedErr *TransitiveDependencyNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsTransitiveDependencyNotFoundError(err error) bool {
	_, ok := AsTransitiveDependencyNotFoundError(err)
	return ok
}

type TransitiveDependencyNotReadyError struct {
	ResourceGVK schema.GroupVersionKind
	Resource    types.NamespacedName
}

func (e *TransitiveDependencyNotReadyError) Error() string {
	return fmt.Sprintf("transitive dependency %v %v is not ready", e.ResourceGVK.Kind, e.Resource)
}

func NewTransitiveDependencyNotReadyError(resourceGVK schema.GroupVersionKind, resource types.NamespacedName) *TransitiveDependencyNotReadyError {
	return &TransitiveDependencyNotReadyError{resourceGVK, resource}
}

func AsTransitiveDependencyNotReadyError(err error) (unwrappedErr *TransitiveDependencyNotReadyError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsTransitiveDependencyNotReadyError(err error) bool {
	_, ok := AsTransitiveDependencyNotReadyError(err)
	return ok
}

type ServerGeneratedIDNotFoundError struct {
	resourceGVK schema.GroupVersionKind
	resource    types.NamespacedName
}

func (e *ServerGeneratedIDNotFoundError) Error() string {
	return fmt.Sprintf("the resource %v %v has a server-generated ID that has not yet been set",
		e.resourceGVK.Kind, e.resource)
}

func NewServerGeneratedIDNotFoundError(resourceGVK schema.GroupVersionKind, resource types.NamespacedName) *ServerGeneratedIDNotFoundError {
	return &ServerGeneratedIDNotFoundError{resourceGVK, resource}
}

func AsServerGeneratedIDNotFoundError(err error) (unwrappedErr *ServerGeneratedIDNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

type ResourceIDNotFoundError struct {
	resourceGVK schema.GroupVersionKind
	resource    types.NamespacedName
}

func (e *ResourceIDNotFoundError) Error() string {
	return fmt.Sprintf("'%s' is not set in %v %v", ResourceIDFieldPath,
		e.resourceGVK.Kind, e.resource)
}

func NewResourceIDNotFoundError(resourceGVK schema.GroupVersionKind, resource types.NamespacedName) *ResourceIDNotFoundError {
	return &ResourceIDNotFoundError{resourceGVK, resource}
}

func AsResourceIDNotFoundError(err error) (unwrappedErr *ResourceIDNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

type ImmutableFieldsMutationError struct {
	immutableFields []string
}

func (e *ImmutableFieldsMutationError) Error() string {
	return fmt.Sprintf("cannot make changes to immutable field(s): %v; please refer to our troubleshooting doc: https://cloud.google.com/config-connector/docs/troubleshooting", e.immutableFields)
}

func NewImmutableFieldsMutationError(immutableFields []string) *ImmutableFieldsMutationError {
	return &ImmutableFieldsMutationError{immutableFields}
}
