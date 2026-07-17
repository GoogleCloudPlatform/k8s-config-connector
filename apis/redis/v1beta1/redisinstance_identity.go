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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &RedisInstanceIdentity{}
	_ identity.Resource   = &RedisInstance{}
)

var RedisInstanceIdentityFormat = gcpurls.Template[RedisInstanceIdentity]("redis.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// RedisInstanceIdentity is the identity of a GCP RedisInstance resource.
// +k8s:deepcopy-gen=false
type RedisInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *RedisInstanceIdentity) String() string {
	return RedisInstanceIdentityFormat.ToString(*i)
}

func (i *RedisInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := RedisInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of RedisInstance external=%q was not known (use %s): %w", ref, RedisInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of RedisInstance external=%q was not known (use %s)", ref, RedisInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *RedisInstanceIdentity) Host() string {
	return RedisInstanceIdentityFormat.Host()
}

func (i *RedisInstanceIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromRedisInstanceSpec(ctx context.Context, reader client.Reader, obj *RedisInstance) (*RedisInstanceIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Region
	if location == "" {
		return nil, fmt.Errorf("cannot resolve region")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &RedisInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *RedisInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromRedisInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
