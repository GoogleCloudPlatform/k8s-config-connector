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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SQLInstanceRef struct {
	/* The SQLInstance selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `SQLInstance` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `SQLInstance` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type SQLInstance struct {
	ProjectID       string
	SQLInstanceName string
}

func (s *SQLInstance) String() string {
	return "projects/" + s.ProjectID + "/instances/" + s.SQLInstanceName
}

func (r *SQLInstanceRef) ResolveRef(ctx context.Context, reader client.Reader, src client.Object) (*SQLInstance, error) {
	if r == nil {
		return nil, nil
	}

	if r.External != "" && r.Name != "" {
		return nil, fmt.Errorf("cannot specify both name and external on project reference")
	}

	tokens := strings.Split(r.External, "/")
	if len(tokens) != 4 && tokens[0] != "projects" && tokens[2] != "instances" {
		return nil, fmt.Errorf("format of sqlinstance external=%q was not known (use projects/<projectId>/instances/<instanceName>)", r.External)
	}

	key := types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	sqlinstance := &unstructured.Unstructured{}
	sqlinstance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sqlinstance.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SQLInstance",
	})
	if err := reader.Get(ctx, key, sqlinstance); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced SQLInstance %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced SQLInstance %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(sqlinstance.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from SQLInstance %s/%s: %w", sqlinstance.GetNamespace(), sqlinstance.GetName(), err)
	}
	if resourceID == "" {
		resourceID = sqlinstance.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, sqlinstance)
	if err != nil {
		return nil, err
	}
	return &SQLInstance{
		ProjectID:       projectID,
		SQLInstanceName: resourceID,
	}, nil
}

// Example
/*

func CloudSqlPropertiesSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlPropertiesSpec) *pb.CloudSqlProperties {
    if in == nil {
        return nil
    }
    out := &pb.CloudSqlProperties{}
    // out.InstanceId = direct.ValueOf(in.InstanceID)
    out.Database = direct.ValueOf(in.Database)
    out.Type = direct.Enum_ToProto[pb.CloudSqlProperties_DatabaseType](mapCtx, in.Type)
    out.Credential = CloudSqlCredential_ToProto(mapCtx, in.Credential)
    if in.InstanceRef != nil {
		sqlinstance, err := in.InstanceRef.ResolveRef(ctx, reader, obj)
        out.InstanceId = sqlinstance.String()
    }
*/
