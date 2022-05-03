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

package controller

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func EnsureNamespaceExists(c client.Client, name string) {
	ns := &corev1.Namespace{}
	ns.SetName(name)
	if err := c.Create(context.Background(), ns); err != nil {
		if !errors.IsAlreadyExists(err) {
			log.Fatalf("error creating namespace %v: %v", name, err)
		}
	}
}

func HasOperatorFinalizer(o metav1.Object) bool {
	for _, f := range o.GetFinalizers() {
		if f == k8s.OperatorFinalizer {
			return true
		}
	}
	return false
}
