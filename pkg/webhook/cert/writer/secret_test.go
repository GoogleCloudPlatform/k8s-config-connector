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

package writer_test

import (
	"context"
	"testing"

	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/writer"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var mgr manager.Manager

func TestSecretCertWriter(t *testing.T) {
	ctx := context.TODO()

	namespaceName := "my-namespace"
	testcontroller.EnsureNamespaceExistsT(t, mgr.GetClient(), namespaceName)
	opts := writer.SecretCertWriterOptions{
		Client: mgr.GetClient(),
		Secret: &types.NamespacedName{
			Name:      "my-name",
			Namespace: namespaceName,
		},
	}
	secretWriter, err := writer.NewSecretCertWriter(opts)
	if err != nil {
		t.Fatalf("error creating secret writer: %v", err)
	}
	artifacts1, changed, err := secretWriter.EnsureCert(ctx, "localhost")
	if err != nil {
		t.Fatalf("error ensuring certificate: %v", err)
	}
	if !changed {
		t.Fatalf("unexpected value for changed: got '%v', want '%v'", changed, true)
	}
	if artifacts1 == nil {
		t.Fatalf("unexpected nil value for artifacts")
	}
	artifacts2, changed, err := secretWriter.EnsureCert(ctx, "localhost")
	if changed {
		t.Fatalf("unexpected value for changed: got '%v', want '%v'", changed, false)
	}
	if !cmp.Equal(artifacts1, artifacts2) {
		t.Fatalf("expected certs to be equal")
	}
}

func TestMain(m *testing.M) {
	testmain.ForUnitTests(m, &mgr)
}
