// Copyright 2025 Google LLC
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

package controllers

import (
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	corev1v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
)

func BuildScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()

	schemeFuncs := []func(*runtime.Scheme) error{
		clientgoscheme.AddToScheme,
		v1alpha1.AddToScheme,
		corev1v1beta1.AddToScheme,
	}
	for _, schemeFunc := range schemeFuncs {
		if err := schemeFunc(scheme); err != nil {
			klog.Fatalf("failed to add to runtime scheme: %v", err)
		}
	}
	return scheme
}
