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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KMSAutokeyConfigIdentity struct {
	parent *KMSAutokeyConfigParent
}

func (i *KMSAutokeyConfigIdentity) String() string {
	return i.parent.String() + "/autokeyConfig"
}

func (r *KMSAutokeyConfigIdentity) Parent() *KMSAutokeyConfigParent {
	return r.parent
}

type KMSAutokeyConfigParent struct {
	FolderID string
}

func (p *KMSAutokeyConfigParent) String() string {
	return "folders/" + p.FolderID
}

func NewAutokeyConfigIdentity(ctx context.Context, reader client.Reader, obj *KMSAutokeyConfig) (*KMSAutokeyConfigIdentity, error) {
	// Get Parent
	folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)

	if err != nil {
		return nil, err
	}
	folderID := folderRef.FolderID
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity, err := ParseKMSAutokeyConfigExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.parent.FolderID != folderID {
			return nil, fmt.Errorf("spec.folderRef changed, expect %s, got %s", actualIdentity.parent.FolderID, folderID)
		}
	}

	return &KMSAutokeyConfigIdentity{
		parent: &KMSAutokeyConfigParent{FolderID: folderID},
	}, nil
}

func ParseKMSAutokeyConfigExternal(external string) (parent *KMSAutokeyConfigIdentity, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 3 || tokens[0] != "folders" || tokens[2] != "autokeyConfig" {
		return nil, fmt.Errorf("format of KMSAutokeyConfig external=%q was not known (use folders/<folderID>/autokeyConfig)", external)
	}
	return &KMSAutokeyConfigIdentity{parent: &KMSAutokeyConfigParent{
		FolderID: tokens[1],
	}}, nil
}
