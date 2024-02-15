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

package mocksecretmanager

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	cloudresourcemanagerv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	secretmanager "google.golang.org/api/secretmanager/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func TestSecretManagerSecretVersion(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	// v1alpha1.AddToScheme(scheme)
	// v1beta1.AddToScheme(scheme)
	corev1.AddToScheme(scheme)
	k8sClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	mockCloud := mockgcp.NewMockRoundTripper(t, k8sClient, storage.NewInMemoryStorage())

	httpClient := &http.Client{
		Transport: mockCloud,
	}

	t.Logf("creating project")
	crm, err := cloudresourcemanagerv1.NewService(ctx, option.WithHTTPClient(httpClient), option.WithAPIKey("fake"))
	if err != nil {
		t.Fatalf("error building cloudresourcemanagerv1 client: %v", err)
	}
	op, err := crm.Projects.Create(&cloudresourcemanagerv1.Project{ProjectId: "mock-project"}).Context(ctx).Do()
	if err != nil {
		t.Fatalf("error creating project: %v", err)
	}
	if !op.Done {
		t.Fatalf("expected mock create project operation to be done immediately")
	}

	t.Logf("creating secret")
	client, err := secretmanager.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("secretmanager.NewService failed: %v", err)
	}
	parent := "projects/mock-project"
	create := &secretmanager.Secret{}
	secretID := "testsecret"
	created, err := client.Projects.Secrets.Create(parent, create).SecretId(secretID).Context(ctx).Do()
	if err != nil {
		t.Fatalf("secretmanager Create failed: %v", err)
	}

	read, err := client.Projects.Secrets.Get(created.Name).Context(ctx).Do()
	if err != nil {
		t.Fatalf("secretmanager Get(%q) failed: %v", created.Name, err)
	}

	t.Logf("read back created secret: %v", read)

	secretData := base64.StdEncoding.EncodeToString([]byte("thesecret"))

	addedVersion, err := client.Projects.Secrets.AddVersion(created.Name, &secretmanager.AddSecretVersionRequest{
		Payload: &secretmanager.SecretPayload{Data: secretData},
	}).Context(ctx).Do()
	if err != nil {
		t.Fatalf("secretmanager Secrets.AddVersion(%q) failed: %v", created.Name, err)
	}
	t.Logf("created version: %#v", addedVersion)

	readVersion, err := client.Projects.Secrets.Versions.Access(addedVersion.Name).Context(ctx).Do()
	if err != nil {
		t.Fatalf("secretmanager Secrets.Versions.Access(%q) failed: %v", addedVersion.Name, err)
	}
	t.Logf("read version: %v", readVersion)
	if got, want := readVersion.Payload.Data, secretData; got != want {
		t.Errorf("secret version payload data did not match expected; got %q, want %q", got, want)
	}

	enabledVersion, err := client.Projects.Secrets.Versions.Enable(addedVersion.Name, &secretmanager.EnableSecretVersionRequest{}).Context(ctx).Do()
	if err != nil {
		t.Fatalf("secretmanager Secrets.Versions.Enable(%q) failed: %v", addedVersion.Name, err)
	}
	t.Logf("enabled version: %v", enabledVersion)

	var b bytes.Buffer
	req, err := http.NewRequest("POST", "https://secretmanager.googleapis.com/v1/"+addedVersion.Name+":enable?alt=json", &b)
	if err != nil {
		t.Fatalf("error building http request: %v", err)
	}
	response, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("error doing http request: %v", err)
	}
	if response.StatusCode != 200 {
		t.Fatalf("unexpected status from request: %q", response.Status)
	}

}
