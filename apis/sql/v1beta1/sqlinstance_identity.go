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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &SQLInstanceIdentity{}
	_ identity.Resource   = &SQLInstance{}
)

var SQLInstanceIdentityFormat = gcpurls.Template[SQLInstanceIdentity]("cloudsql.googleapis.com", "projects/{project}/instances/{instance}")

// +k8s:deepcopy-gen=false
type SQLInstanceIdentity struct {
	Project  string
	Instance string
	// Location is not part of the canonical identity URI, but is needed for connectionName
	Location string
}

func (i *SQLInstanceIdentity) String() string {
	return SQLInstanceIdentityFormat.ToString(*i)
}

func (i *SQLInstanceIdentity) ConnectionName() string {
	return fmt.Sprintf("%s:%s:%s", i.Project, i.Location, i.Instance)
}

func (i *SQLInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := SQLInstanceIdentityFormat.Parse(ref)
	if err == nil && match {
		*i = *parsed
		return nil
	}

	// Try legacy format with 'locations': projects/{project}/locations/{location}/instances/{instance}
	tokens := strings.Split(ref, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		i.Project = tokens[1]
		i.Location = tokens[3]
		i.Instance = tokens[5]
		return nil
	}

	if err != nil {
		return fmt.Errorf("format of SQLInstance external=%q was not known (use %s): %w", ref, SQLInstanceIdentityFormat.CanonicalForm(), err)
	}
	return fmt.Errorf("format of SQLInstance external=%q was not known (use %s)", ref, SQLInstanceIdentityFormat.CanonicalForm())
}

func (i *SQLInstanceIdentity) Host() string {
	return SQLInstanceIdentityFormat.Host()
}

func GetIdentityFromSQLInstanceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*SQLInstanceIdentity, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, fmt.Errorf("cannot convert to unstructured: %w", err)
		}
		u = &unstructured.Unstructured{Object: m}
	}

	resourceID, err := refs.GetResourceID(u)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location, _, err := unstructured.NestedString(u.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}

	identity := &SQLInstanceIdentity{
		Project:  projectID,
		Instance: resourceID,
		Location: location,
	}
	return identity, nil
}

func (obj *SQLInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := GetIdentityFromSQLInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
