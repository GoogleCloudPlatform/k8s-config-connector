// Copyright 2024 Google LLC
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

package directbase

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type UpdateOperation struct {
	object *unstructured.Unstructured
}

func NewUpdateOperation(object *unstructured.Unstructured) *UpdateOperation {
	return &UpdateOperation{
		object: object,
	}
}

func (o *UpdateOperation) GetUnstructured() *unstructured.Unstructured {
	return o.object
}

type CreateOperation struct {
	object *unstructured.Unstructured
}

func NewCreateOperation(object *unstructured.Unstructured) *CreateOperation {
	return &CreateOperation{
		object: object,
	}
}

func (o *CreateOperation) GetUnstructured() *unstructured.Unstructured {
	return o.object
}

type DeleteOperation struct {
	object *unstructured.Unstructured
}

func NewDeleteOperation(object *unstructured.Unstructured) *DeleteOperation {
	return &DeleteOperation{
		object: object,
	}
}

func (o *DeleteOperation) GetUnstructured() *unstructured.Unstructured {
	return o.object
}
