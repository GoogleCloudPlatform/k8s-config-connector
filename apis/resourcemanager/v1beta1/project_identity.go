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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ProjectIdentity{}
	_ identity.Resource   = &Project{}
)

var ProjectIdentityFormat = gcpurls.Template[ProjectIdentity]("cloudresourcemanager.googleapis.com", "projects/{project}")

// ProjectIdentity is the identity of a Google Cloud Project resource.
// +k8s:deepcopy-gen=false
type ProjectIdentity struct {
	Project string
}

func (i *ProjectIdentity) String() string {
	return ProjectIdentityFormat.ToString(*i)
}

func (i *ProjectIdentity) FromExternal(ref string) error {
	if !strings.Contains(ref, "/") {
		ref = "projects/" + ref
	}

	parsed, match, err := ProjectIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Project external=%q was not known (use %s): %w", ref, ProjectIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Project external=%q was not known (use %s)", ref, ProjectIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ProjectIdentity) Host() string {
	return ProjectIdentityFormat.Host()
}

func getIdentityFromProjectSpec(ctx context.Context, reader client.Reader, obj *Project) (*ProjectIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &ProjectIdentity{
		Project: resourceID,
	}
	return identity, nil
}

func (obj *Project) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromProjectSpec(ctx, reader, obj)
}
