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
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type errorClient struct {
}

// Some packages, like 'gcpclient' make use of krmtotf which is tightly coupled with the controller-runtime client.
// However, krmtotf does not actually need the client if all the resources passed in have all references resolved, etc.
// To enable usage of the library but also to avoid panics, this erroring client can be passed to krmtotf methods.
func NewErroringClient() client.Client {
	return &errorClient{}
}

func (e *errorClient) Get(_ context.Context, key client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
	return fmt.Errorf("unexpected call to client.Get(...) for %v", key)
}

func (e *errorClient) List(_ context.Context, _ client.ObjectList, _ ...client.ListOption) error {
	return fmt.Errorf("unexpected call to client.List(...)")
}

func (e *errorClient) Create(_ context.Context, obj client.Object, _ ...client.CreateOption) error {
	return fmt.Errorf("unexpected call to client.Create(...) for object with kind %v", obj.GetObjectKind())
}

func (e *errorClient) Delete(_ context.Context, obj client.Object, _ ...client.DeleteOption) error {
	return fmt.Errorf("unexpected call to client.Delete(...) for object with kind %v", obj.GetObjectKind())
}

func (e *errorClient) Update(_ context.Context, obj client.Object, _ ...client.UpdateOption) error {
	return fmt.Errorf("unexpected call to client.Update(...) for object with kind %v", obj.GetObjectKind())
}

func (e *errorClient) Patch(_ context.Context, obj client.Object, _ client.Patch, _ ...client.PatchOption) error {
	return fmt.Errorf("unexpected call to client.Patch(...) for object with kind %v", obj.GetObjectKind())
}

func (e *errorClient) DeleteAllOf(_ context.Context, obj client.Object, _ ...client.DeleteAllOfOption) error {
	return fmt.Errorf("unexpected call to client.DeleteAllOf(...) for object with kind %v", obj.GetObjectKind())
}

func (e *errorClient) Scheme() *runtime.Scheme {
	panic("unexpected call to client.Scheme(...)")
}

func (e *errorClient) RESTMapper() meta.RESTMapper {
	panic("unexpected call to client.RESTMapper(...)")
}

func (e *errorClient) GroupVersionKindFor(_ runtime.Object) (schema.GroupVersionKind, error) {
	panic("unexpected call to client.GroupVersionKindFor(...)")
}

func (e *errorClient) IsObjectNamespaced(_ runtime.Object) (bool, error) {
	panic("unexpected call to client.IsObjectNamespaced(...)")
}

func (e *errorClient) SubResource(_ string) client.SubResourceClient {
	panic("unexpected call to client.SubResource(...)")
}

func (e *errorClient) SubResourceWriter(_ string) client.SubResourceClient {
	panic("unexpected call to client.SubResource(...)")
}

func (e *errorClient) Status() client.SubResourceWriter {
	panic("unexpected call to client.Status(...)")
}
