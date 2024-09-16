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

package testiam

import (
	"testing"
)

type IAMResourceContext struct {
	Kind              string
	Name              string
	CreateBindingRole string
	UpdateBindingRole string
}

var (
	// Resource kinds with customized binding roles for IAM integration test.
	ResourceContexts = []IAMResourceContext{
		{
			Kind:              "CloudFunctionsFunction",
			Name:              "httpsfunction",
			CreateBindingRole: "roles/cloudfunctions.invoker",
			UpdateBindingRole: "roles/cloudfunctions.invoker",
		},
		{
			Kind:              "BigtableInstance",
			CreateBindingRole: "roles/bigtable.viewer",
			UpdateBindingRole: "roles/bigtable.user",
		},
		{
			Kind:              "KMSKeyRing",
			CreateBindingRole: "roles/cloudkms.publicKeyViewer",
			UpdateBindingRole: "roles/cloudkms.admin",
		},
		{
			Kind:              "Project",
			Name:              "projectinfolder",
			CreateBindingRole: "roles/storage.objectAdmin",
			UpdateBindingRole: "roles/storage.admin",
		},
		{
			Kind:              "Folder",
			Name:              "folderinfolder",
			CreateBindingRole: "roles/storage.objectAdmin",
			UpdateBindingRole: "roles/storage.admin",
		},
		{
			Kind:              "PubSubTopic",
			CreateBindingRole: "roles/pubsub.viewer",
			UpdateBindingRole: "roles/pubsub.editor",
		},
		{
			Kind:              "PubSubSubscription",
			CreateBindingRole: "roles/pubsub.subscriber",
			UpdateBindingRole: "roles/pubsub.viewer",
		},
		{
			Kind:              "SpannerInstance",
			CreateBindingRole: "roles/spanner.databaseReader",
			UpdateBindingRole: "roles/spanner.databaseUser",
		},
		{
			Kind:              "StorageBucket",
			CreateBindingRole: "roles/storage.admin",
			UpdateBindingRole: "roles/storage.objectAdmin",
		},
		{
			Kind:              "IAMServiceAccount",
			CreateBindingRole: "roles/iam.serviceAccountUser",
			UpdateBindingRole: "roles/iam.serviceAccountAdmin",
		},
		{
			Kind:              "DataprocCluster",
			CreateBindingRole: "roles/dataproc.editor",
			UpdateBindingRole: "roles/dataproc.admin",
		},
	}
)

func GetResourceContext(t *testing.T, kind string) *IAMResourceContext {
	t.Helper()
	for _, rc := range ResourceContexts {
		if rc.Kind == kind {
			return &rc
		}
	}
	return &IAMResourceContext{
		Kind:              kind,
		CreateBindingRole: "roles/viewer",
		UpdateBindingRole: "roles/editor",
	}
}
