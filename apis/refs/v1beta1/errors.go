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

package v1beta1

import (
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type ReferenceNotReadyError struct {
	RefResourceGVK schema.GroupVersionKind
	RefResource    types.NamespacedName
}

func (e *ReferenceNotReadyError) Error() string {
	return fmt.Sprintf("reference %v %v is not ready", e.RefResourceGVK.Kind, e.RefResource)
}

func NewReferenceNotReadyError(refResourceGVK schema.GroupVersionKind, refResource types.NamespacedName) *ReferenceNotReadyError {
	return &ReferenceNotReadyError{
		RefResourceGVK: refResourceGVK,
		RefResource:    refResource,
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
	return &ReferenceNotFoundError{
		RefResourceGVK: refResourceGVK,
		RefResource:    refResource,
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

type SecretNotFoundError struct {
	Secret types.NamespacedName
}

func (e *SecretNotFoundError) Error() string {
	return fmt.Sprintf("Secret %v was not found", e.Secret)
}

func NewSecretNotFoundError(secret types.NamespacedName) *SecretNotFoundError {
	return &SecretNotFoundError{Secret: secret}
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
	return &KeyInSecretNotFoundError{key: key, secret: secret}
}

func AsKeyInSecretNotFoundError(err error) (unwrappedErr *KeyInSecretNotFoundError, ok bool) {
	ok = errors.As(err, &unwrappedErr)
	return unwrappedErr, ok
}

func IsKeyInSecretNotFoundError(err error) bool {
	_, ok := AsKeyInSecretNotFoundError(err)
	return ok
}
