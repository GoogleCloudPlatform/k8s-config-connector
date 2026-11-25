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

package webhook

import (
	"context"
	"encoding/json"
	"testing"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	admissionv1 "k8s.io/api/admission/v1"
	authenticationv1 "k8s.io/api/authentication/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func newTestImmutableFieldsHandler(t *testing.T) admission.Handler {
	t.Helper()
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatal(err)
	}
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatal(err)
	}
	return NewImmutableFieldsValidatorHandler(smLoader, dclSchemaLoader, dclmetadata.New())(nil)
}

func marshalAdmissionObject(t *testing.T, gvk k8sschema.GroupVersionKind, annotations map[string]string, spec map[string]interface{}) runtime.RawExtension {
	t.Helper()
	metadata := map[string]interface{}{
		"name":      "an-object",
		"namespace": "a-namespace",
	}
	if len(annotations) > 0 {
		ann := make(map[string]interface{}, len(annotations))
		for k, v := range annotations {
			ann[k] = v
		}
		metadata["annotations"] = ann
	}
	u := &unstructured.Unstructured{Object: map[string]interface{}{
		"metadata": metadata,
		"spec":     runtime.DeepCopyJSON(spec),
	}}
	u.SetGroupVersionKind(gvk)
	raw, err := json.Marshal(u.Object)
	if err != nil {
		t.Fatal(err)
	}
	return runtime.RawExtension{Raw: raw}
}

// Who may set, change, or clear spec.resourceID.
//
// A user may not. The controller manager may: it writes the generated ID there. The KRM syncer
// of multi-cluster leader election (service account kcc-syncer in cnrm-system, installed by
// `config-connector ha setup`) may too: on a standby cluster it copies the leader's
// spec.resourceID, without which a standby that becomes the leader creates the resource again.
func TestWhoMaySetServiceGeneratedResourceID(t *testing.T) {
	handler := newTestImmutableFieldsHandler(t)

	kinds := []struct {
		name string
		gvk  k8sschema.GroupVersionKind
		spec map[string]interface{}
	}{
		{
			name: "DCL-based kind",
			gvk:  k8sschema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringGroup"},
			spec: map[string]interface{}{
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/a-project"},
			},
		},
		{
			name: "Terraform-based kind",
			gvk:  k8sschema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringNotificationChannel"},
			spec: map[string]interface{}{
				"type":   "email",
				"labels": map[string]interface{}{"email_address": "nobody@example.com"},
			},
		},
		{
			name: "LoggingLogMetric",
			gvk:  k8sschema.GroupVersionKind{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogMetric"},
			spec: map[string]interface{}{
				"description": "a metric",
				"projectRef":  map[string]interface{}{"external": "projects/a-project"},
			},
		},
	}
	users := []struct {
		name     string
		username string
		allowed  bool
	}{
		{"a user", "jane@example.com", false},
		{"the controller manager", "system:serviceaccount:cnrm-system:cnrm-controller-manager", true},
		{"the controller manager of a namespace", "system:serviceaccount:cnrm-system:cnrm-controller-manager-team-a", true},
		{"the KRM syncer", "system:serviceaccount:cnrm-system:kcc-syncer", true},
		// Only the syncer that `config-connector ha setup` installs, in cnrm-system.
		{"a kcc-syncer service account in another namespace", "system:serviceaccount:team-a:kcc-syncer", false},
		{"the MultiClusterLease controller", "system:serviceaccount:cnrm-system:kcc-mcl", false},
	}
	changes := []struct {
		name  string
		oldID interface{}
		newID interface{}
	}{
		{"set", nil, "1234567890"},
		{"change", "1234567890", "9876543210"},
		{"clear", "1234567890", nil},
	}

	for _, kind := range kinds {
		for _, user := range users {
			for _, change := range changes {
				t.Run(kind.name+"/"+user.name+"/"+change.name, func(t *testing.T) {
					specWithID := func(id interface{}) map[string]interface{} {
						s := runtime.DeepCopyJSON(kind.spec)
						if id != nil {
							s[k8s.ResourceIDFieldName] = id
						}
						return s
					}
					req := admission.Request{AdmissionRequest: admissionv1.AdmissionRequest{
						UserInfo:  authenticationv1.UserInfo{Username: user.username},
						OldObject: marshalAdmissionObject(t, kind.gvk, nil, specWithID(change.oldID)),
						Object:    marshalAdmissionObject(t, kind.gvk, nil, specWithID(change.newID)),
					}}
					resp := handler.Handle(context.Background(), req)
					if resp.Allowed != user.allowed {
						message := ""
						if resp.Result != nil {
							message = resp.Result.Message
						}
						t.Errorf("allowed: got %t, want %t (%s)", resp.Allowed, user.allowed, message)
					}
				})
			}
		}
	}
}

// The KRM syncer's exemption is narrow: it may change spec.resourceID, and nothing else that is
// immutable. Other immutable fields, immutable annotations, and IAM specs are still rejected when
// the syncer tries to change them (alone or together with spec.resourceID), while the controller
// manager remains exempt.
func TestSyncerCannotChangeOtherImmutableFields(t *testing.T) {
	handler := newTestImmutableFieldsHandler(t)

	monitoringGroupGVK := k8sschema.GroupVersionKind{
		Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringGroup",
	}
	notificationChannelGVK := k8sschema.GroupVersionKind{
		Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringNotificationChannel",
	}
	storageNotificationGVK := k8sschema.GroupVersionKind{
		Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageNotification",
	}
	computeAddressGVK := k8sschema.GroupVersionKind{
		Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeAddress",
	}
	loggingLogMetricGVK := k8sschema.GroupVersionKind{
		Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogMetric",
	}
	iamPolicyMemberGVK := k8sschema.GroupVersionKind{
		Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicyMember",
	}
	gkehubFeatureMembershipGVK := k8sschema.GroupVersionKind{
		Group: "gkehub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "GKEHubFeatureMembership",
	}

	cases := []struct {
		name           string
		gvk            k8sschema.GroupVersionKind
		oldAnnotations map[string]string
		newAnnotations map[string]string
		oldSpec        map[string]interface{}
		newSpec        map[string]interface{}
	}{
		{
			name: "DCL immutable field (spec.projectRef)",
			gvk:  monitoringGroupGVK,
			oldSpec: map[string]interface{}{
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/project-a"},
			},
			newSpec: map[string]interface{}{
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/project-b"},
			},
		},
		{
			name: "DCL spec.resourceID together with immutable field (spec.projectRef)",
			gvk:  monitoringGroupGVK,
			oldSpec: map[string]interface{}{
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/project-a"},
			},
			newSpec: map[string]interface{}{
				"resourceID":  "1234567890",
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/project-b"},
			},
		},
		{
			name: "Terraform immutable field (spec.payloadFormat)",
			gvk:  storageNotificationGVK,
			oldSpec: map[string]interface{}{
				"bucketRef":     map[string]interface{}{"external": "bucket-a"},
				"topicRef":      map[string]interface{}{"external": "projects/project-a/topics/topic-a"},
				"payloadFormat": "JSON_API_V1",
			},
			newSpec: map[string]interface{}{
				"bucketRef":     map[string]interface{}{"external": "bucket-a"},
				"topicRef":      map[string]interface{}{"external": "projects/project-a/topics/topic-a"},
				"payloadFormat": "NONE",
			},
		},
		{
			name: "Terraform spec.resourceID together with immutable field (spec.payloadFormat)",
			gvk:  storageNotificationGVK,
			oldSpec: map[string]interface{}{
				"bucketRef":     map[string]interface{}{"external": "bucket-a"},
				"topicRef":      map[string]interface{}{"external": "projects/project-a/topics/topic-a"},
				"payloadFormat": "JSON_API_V1",
			},
			newSpec: map[string]interface{}{
				"resourceID":    "1234567890",
				"bucketRef":     map[string]interface{}{"external": "bucket-a"},
				"topicRef":      map[string]interface{}{"external": "projects/project-a/topics/topic-a"},
				"payloadFormat": "NONE",
			},
		},
		{
			name: "Terraform spec.resourceID together with immutable location",
			gvk:  computeAddressGVK,
			oldSpec: map[string]interface{}{
				"location": "us-central1",
			},
			newSpec: map[string]interface{}{
				"resourceID": "other-name",
				"location":   "us-east1",
			},
		},
		{
			name: "LoggingLogMetric spec.resourceID together with immutable field (spec.projectRef)",
			gvk:  loggingLogMetricGVK,
			oldSpec: map[string]interface{}{
				"description": "a metric",
				"projectRef":  map[string]interface{}{"external": "projects/project-a"},
			},
			newSpec: map[string]interface{}{
				"resourceID":  "1234567890",
				"description": "a metric",
				"projectRef":  map[string]interface{}{"external": "projects/project-b"},
			},
		},
		{
			name:           "state-into-spec annotation together with spec.resourceID",
			gvk:            monitoringGroupGVK,
			oldAnnotations: map[string]string{k8s.StateIntoSpecAnnotation: k8s.StateAbsentInSpec},
			newAnnotations: map[string]string{k8s.StateIntoSpecAnnotation: k8s.StateMergeIntoSpec},
			oldSpec: map[string]interface{}{
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/project-a"},
			},
			newSpec: map[string]interface{}{
				"resourceID":  "1234567890",
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"x\")",
				"projectRef":  map[string]interface{}{"external": "projects/project-a"},
			},
		},
		{
			name:           "container annotation together with spec.resourceID",
			gvk:            notificationChannelGVK,
			oldAnnotations: map[string]string{k8s.ProjectIDAnnotation: "project-a"},
			newAnnotations: map[string]string{k8s.ProjectIDAnnotation: "project-b"},
			oldSpec: map[string]interface{}{
				"type":   "email",
				"labels": map[string]interface{}{"email_address": "nobody@example.com"},
			},
			newSpec: map[string]interface{}{
				"resourceID": "1234567890",
				"type":       "email",
				"labels":     map[string]interface{}{"email_address": "nobody@example.com"},
			},
		},
		{
			name: "IAMPolicyMember spec",
			gvk:  iamPolicyMemberGVK,
			oldSpec: map[string]interface{}{
				"member":      "serviceAccount:a@project-a.iam.gserviceaccount.com",
				"role":        "roles/viewer",
				"resourceRef": map[string]interface{}{"kind": "Project", "external": "projects/project-a"},
			},
			newSpec: map[string]interface{}{
				"member":      "serviceAccount:a@project-a.iam.gserviceaccount.com",
				"role":        "roles/editor",
				"resourceRef": map[string]interface{}{"kind": "Project", "external": "projects/project-a"},
			},
		},
		{
			name: "GKEHubFeatureMembership immutable field (spec.location)",
			gvk:  gkehubFeatureMembershipGVK,
			oldSpec: map[string]interface{}{
				"location": "global",
			},
			newSpec: map[string]interface{}{
				"location": "us-central1",
			},
		},
	}

	users := []struct {
		name     string
		username string
		allowed  bool
	}{
		{"the KRM syncer", "system:serviceaccount:cnrm-system:kcc-syncer", false},
		{"the controller manager", "system:serviceaccount:cnrm-system:cnrm-controller-manager", true},
	}

	for _, tc := range cases {
		for _, user := range users {
			t.Run(tc.name+"/"+user.name, func(t *testing.T) {
				req := admission.Request{AdmissionRequest: admissionv1.AdmissionRequest{
					UserInfo:  authenticationv1.UserInfo{Username: user.username},
					OldObject: marshalAdmissionObject(t, tc.gvk, tc.oldAnnotations, tc.oldSpec),
					Object:    marshalAdmissionObject(t, tc.gvk, tc.newAnnotations, tc.newSpec),
				}}
				resp := handler.Handle(context.Background(), req)
				if resp.Allowed != user.allowed {
					message := ""
					if resp.Result != nil {
						message = resp.Result.Message
					}
					t.Errorf("allowed: got %t, want %t (%s)", resp.Allowed, user.allowed, message)
				}
			})
		}
	}
}
