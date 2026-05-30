// Copyright 2026 Google LLC
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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &KMSAutokeyConfigIdentity{}
	_ identity.Resource   = &KMSAutokeyConfig{}
)

var KMSAutokeyConfigIdentityFormat = gcpurls.Template[KMSAutokeyConfigIdentity]("cloudkms.googleapis.com", "folders/{folder}/autokeyConfig")

// +k8s:deepcopy-gen=false
type KMSAutokeyConfigIdentity struct {
	Folder string
}

func (i *KMSAutokeyConfigIdentity) String() string {
	return KMSAutokeyConfigIdentityFormat.ToString(*i)
}

func (i *KMSAutokeyConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := KMSAutokeyConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of KMSAutokeyConfig external=%q was not known (use %s): %w", ref, KMSAutokeyConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of KMSAutokeyConfig external=%q was not known (use %s)", ref, KMSAutokeyConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *KMSAutokeyConfigIdentity) Host() string {
	return KMSAutokeyConfigIdentityFormat.Host()
}

func getIdentityFromKMSAutokeyConfigSpec(ctx context.Context, reader client.Reader, obj client.Object) (*KMSAutokeyConfigIdentity, error) {
	var unstructuredObj *unstructured.Unstructured
	if u, ok := obj.(*unstructured.Unstructured); ok {
		unstructuredObj = u
	} else {
		uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, err
		}
		unstructuredObj = &unstructured.Unstructured{Object: uObj}
	}

	folderID, err := refs.ResolveFolderID(ctx, reader, unstructuredObj)
	if err != nil {
		return nil, err
	}

	return &KMSAutokeyConfigIdentity{
		Folder: folderID,
	}, nil
}

func (obj *KMSAutokeyConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromKMSAutokeyConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &KMSAutokeyConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change KMSAutokeyConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
