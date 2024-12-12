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

type ReferenceNotReadyError struct {
	RefResourceGVK schema.GroupVersionKind
	RefResource    types.NamespacedName
}

func (e *ReferenceNotReadyError) Error() string {
	return fmt.Sprintf("reference %v %v is not ready", e.RefResourceGVK.Kind, e.RefResource)
}

func NewReferenceNotReadyError(refResourceGVK schema.GroupVersionKind, refResource types.NamespacedName) *ReferenceNotReadyError {
	return &ReferenceNotReadyError{refResourceGVK, refResource}
}

func NewReferenceNotReadyErrorForResource(r *Resource) *ReferenceNotReadyError {
	return &ReferenceNotReadyError{
		r.GroupVersionKind(),
		types.NamespacedName{Namespace: r.GetNamespace(), Name: r.GetName()},
	}
}

func AsReferenceNotReadyError(err error) (unwrappedErr *ReferenceNotReadyError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsReferenceNotReadyError(err error) bool {
	_, ok := AsReferenceNotReadyError(err)
	return ok
}

type ReferenceNotFoundError struct {
	RefResourceGVK schema.GroupVersionKind
	RefResource    types.NamespacedName
}

func (e *ReferenceNotFoundError) Error() string {
	return fmt.Sprintf("reference %v %v is not found", e.RefResourceGVK.Kind, e.RefResource)
}

func NewReferenceNotFoundError(refResourceGVK schema.GroupVersionKind, refResource types.NamespacedName) *ReferenceNotFoundError {
	return &ReferenceNotFoundError{refResourceGVK, refResource}
}

func NewReferenceNotFoundErrorForResource(r *Resource) *ReferenceNotFoundError {
	return &ReferenceNotFoundError{
		r.GroupVersionKind(),
		types.NamespacedName{Namespace: r.GetNamespace(), Name: r.GetName()},
	}
}

func AsReferenceNotFoundError(err error) (unwrappedErr *ReferenceNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsReferenceNotFoundError(err error) bool {
	_, ok := AsReferenceNotFoundError(err)
	return ok
}

var ErrIAMNotFound = fmt.Errorf("IAM resource does not exist")

type SecretNotFoundError struct {
	Secret types.NamespacedName
}

func (e *SecretNotFoundError) Error() string {
	return fmt.Sprintf("Secret %v was not found", e.Secret)
}

func NewSecretNotFoundError(secret types.NamespacedName) *SecretNotFoundError {
	return &SecretNotFoundError{secret}
}

func AsSecretNotFoundError(err error) (unwrappedErr *SecretNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsSecretNotFoundError(err error) bool {
	_, ok := AsSecretNotFoundError(err)
	return ok
}

type KeyInSecretNotFoundError struct {
	key    string
	secret types.NamespacedName
}

func (e *KeyInSecretNotFoundError) Error() string {
	return fmt.Sprintf("key '%v' was not found in Secret %v", e.key, e.secret)
}

func NewKeyInSecretNotFoundError(key string, secret types.NamespacedName) *KeyInSecretNotFoundError {
	return &KeyInSecretNotFoundError{key, secret}
}

func AsKeyInSecretNotFoundError(err error) (unwrappedErr *KeyInSecretNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsKeyInSecretNotFoundError(err error) bool {
	_, ok := AsKeyInSecretNotFoundError(err)
	return ok
}

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
