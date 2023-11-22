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

package crdloader_test

import (
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

const (
	PubSubGroup     = "pubsub.cnrm.cloud.google.com"
	PubSubTopicKind = "PubSubTopic"
	PubSubVersion   = "v1beta1"
)

type CRDTestCase struct {
	Name          string
	ShouldSucceed bool
	Group         string
	Version       string
	Kind          string
}

var crdTestCases = []CRDTestCase{
	{"AllNil", false, "", "", ""},
	{"UnknownKind", false, "", "", "MyNewKind"},
	{"NoKind", false, PubSubGroup, PubSubVersion, ""},
	{"JustKind", true, "", "", PubSubTopicKind},
	{"InvalidVersion", false, "", "v1invalid1", PubSubTopicKind},
	{"KindAndVersion", true, "", PubSubVersion, PubSubTopicKind},
	{"InvalidGroup", false, "invalidgroup.google.com", "", PubSubTopicKind},
	{"GroupAndKind", true, PubSubGroup, "", PubSubTopicKind},
	{"GroupKindAndVersion", true, PubSubGroup, PubSubVersion, PubSubTopicKind},
}

func TestGetCRD(t *testing.T) {
	for _, tc := range crdTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			getCRDAssertResult(t, tc, crdloader.GetCRD)
		})
	}
}

func TestCrdLoader_GetCRD(t *testing.T) {
	env := &envtest.Environment{
		CRDDirectoryPaths:        []string{repo.GetCRDsPath()},
		ControlPlaneStartTimeout: time.Minute,
		ControlPlaneStopTimeout:  time.Minute,
	}
	cfg, err := env.Start()
	if err != nil {
		t.Fatalf("error starting test environment: %v", err)
	}
	if err := apiextensions.SchemeBuilder.AddToScheme(scheme.Scheme); err != nil {
		t.Fatalf("error adding to scheme: %v", err)
	}
	clientOptions := client.Options{
		Scheme: scheme.Scheme,
	}
	kubeClient, err := client.New(cfg, clientOptions)
	if err != nil {
		t.Fatalf("error creating k8s client: %v", err)
	}
	for _, tc := range crdTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			loader := crdloader.New(kubeClient)
			getCRDAssertResult(t, tc, loader.GetCRD)
		})
	}
}

func getCRDAssertResult(t *testing.T, tc CRDTestCase, getCRDFunc func(string, string, string) (*apiextensions.CustomResourceDefinition, error)) {
	crd, err := getCRDFunc(tc.Group, tc.Version, tc.Kind)
	if tc.ShouldSucceed {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	} else {
		if err == nil {
			t.Fatalf("expected error, instead got 'nil'")
		}
		return
	}
	if crd.Spec.Names.Kind != tc.Kind {
		t.Errorf("mismatched value for 'kind': got '%v', want '%v'", crd.Spec.Names.Kind, tc.Kind)
	}
	if tc.Group != "" && crd.Spec.Group != tc.Group {
		t.Errorf("mismatched value for 'group': got '%v', want '%v'", crd.Spec.Group, tc.Group)
	}
	version := k8s.GetVersionFromCRD(crd)
	if tc.Version != "" && version != tc.Version {
		t.Errorf("mismatched value for 'version': got '%v', want '%v'", version, tc.Version)
	}
}
