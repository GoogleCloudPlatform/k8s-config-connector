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

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ListCRDs returns the list of KCC CRDs on the API Server. The function returns
// the CRDs in paginated fashion, i.e. one chunk of at most 100 CRDs at a time,
// and the nextPageToken to be used by the caller to fetch the next chunk of CRDs.
// When there are no more CRDs to list, the function returns an empty string nextPageToken.
// Callers are expected to pass an empty string for pageToken when initiating a
// new list operation.
func ListCRDs(ctx context.Context, kubeClient client.Client, pageToken string) (crds []v1.CustomResourceDefinition, nextPageToken string, err error) {
	list := v1.CustomResourceDefinitionList{}
	labelSelector, err := labels.Parse(KCCSystemLabelSelectorRaw)
	if err != nil {
		return nil, "", fmt.Errorf("error parsing '%v' as a label selector: %w", KCCSystemLabelSelectorRaw, err)
	}
	opts := &client.ListOptions{
		Limit:         100,
		Raw:           &metav1.ListOptions{},
		LabelSelector: labelSelector,
		Continue:      pageToken,
	}
	if err := kubeClient.List(ctx, &list, opts); err != nil {
		return nil, "", fmt.Errorf("error listing CRDs: %w", err)
	}

	for _, crd := range list.Items {
		if _, ok := IgnoredCRDList[crd.Name]; ok {
			continue
		}
		crds = append(crds, crd)
	}
	return crds, list.GetContinue(), nil
}
