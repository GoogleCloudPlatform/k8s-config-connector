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

package preflight

import (
	"context"
	"fmt"
	"log"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/asserts"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/loaders"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var (
	cfg *rest.Config
)

func init() {
	s := scheme.Scheme
	if err := corev1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering core kcc operator scheme: %v", err)
	}
}

type FakeRepo struct {
	channel *loaders.Channel
}

var _ manifest.Repository = &FakeRepo{}

func (r *FakeRepo) LoadManifest(ctx context.Context, component string, id string, o declarative.DeclarativeObject) (map[string]string, error) {
	panic("implement me")
}

func (r *FakeRepo) LoadChannel(ctx context.Context, name string) (*loaders.Channel, error) {
	return r.channel, nil
}

func (r *FakeRepo) LoadNamespacedComponents(ctx context.Context, componentName string, version string) (map[string]string, error) {
	panic("implement me")
}

func TestUpgradeChecker_Preflight(t *testing.T) {
	t.Parallel()
	curTime := metav1.Now()
	testConfigConnector := &corev1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-configconnector",
		},
		Spec: corev1beta1.ConfigConnectorSpec{
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}
	tests := []struct {
		name    string
		cc      *corev1beta1.ConfigConnector
		ns      *corev1.Namespace
		channel *loaders.Channel
		err     error
	}{
		{
			name:    "no existing instance, can upgrade/deploy the new version",
			cc:      testConfigConnector,
			channel: nil,
			ns:      nil,
			err:     nil,
		},
		{
			name: "new version is compatible, can upgrade it",
			cc:   testConfigConnector,
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.CNRMSystemNamespace,
					Annotations: map[string]string{
						k8s.VersionAnnotation: "1.2.3",
					},
				},
			},
			channel: &loaders.Channel{
				Manifests: []loaders.Version{
					{
						Package: "configconnector",
						Version: "1.4.0",
					},
				},
			},
			err: nil,
		},
		{
			name: "reconcile on the same version",
			cc:   testConfigConnector,
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.CNRMSystemNamespace,
					Annotations: map[string]string{
						k8s.VersionAnnotation: "1.2.3",
					},
				},
			},
			channel: &loaders.Channel{
				Manifests: []loaders.Version{
					{
						Package: "configconnector",
						Version: "1.2.3",
					},
				},
			},
			err: nil,
		},
		{
			name: "major change, no upgrade",
			cc:   testConfigConnector,
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.CNRMSystemNamespace,
					Annotations: map[string]string{
						k8s.VersionAnnotation: "1.2.3",
					},
				},
			},
			channel: &loaders.Channel{
				Manifests: []loaders.Version{
					{
						Package: "configconnector",
						Version: "2.0.0",
					},
				},
			},
			err: fmt.Errorf("incompatible version"),
		},
		{
			name: "major change, no downgrade",
			cc:   testConfigConnector,
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.CNRMSystemNamespace,
					Annotations: map[string]string{
						k8s.VersionAnnotation: "2.0.0",
					},
				},
			},
			channel: &loaders.Channel{
				Manifests: []loaders.Version{
					{
						Package: "configconnector",
						Version: "1.2.0",
					},
				},
			},
			err: fmt.Errorf("incompatible version"),
		},
		{
			name: "delete will always pass preflight check",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-configconnector",
					DeletionTimestamp: &curTime,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.CNRMSystemNamespace,
					Annotations: map[string]string{
						k8s.VersionAnnotation: "1.2.3",
					},
				},
			},
			channel: &loaders.Channel{
				Manifests: []loaders.Version{
					{
						Package: "configconnector",
						Version: "2.0.0",
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			client := mgr.GetClient()
			if tc.ns != nil {
				if err := client.Create(ctx, tc.ns); err != nil {
					t.Fatalf("error creating %v %v: %v", tc.ns.Kind, tc.cc.Name, err)
				}
			}
			repo := FakeRepo{
				channel: tc.channel,
			}
			u := NewUpgradeChecker(client, &repo)
			err := u.Preflight(ctx, tc.cc)
			asserts.AssertErrorIsExpected(t, err, tc.err)
		})
	}
}
