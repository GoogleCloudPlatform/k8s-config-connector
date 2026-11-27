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

package pubsub

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/durationpb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	krmbigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

func TestSubscriptionModel_AdapterForObject_BigQueryTable(t *testing.T) {
	tests := []struct {
		name      string
		tableRef  krmbigqueryv1beta1.BigQueryTableRef
		wantTable string
	}{
		{
			name:      "table referenced by name",
			tableRef:  krmbigqueryv1beta1.BigQueryTableRef{Name: "my-table"},
			wantTable: "my-project.my_dataset.my_table",
		},
		{
			name:      "canonical external table reference",
			tableRef:  krmbigqueryv1beta1.BigQueryTableRef{External: "projects/my-project/datasets/my_dataset/tables/my_table"},
			wantTable: "my-project.my_dataset.my_table",
		},
		{
			name:      "external table reference with the bigquery.googleapis.com host",
			tableRef:  krmbigqueryv1beta1.BigQueryTableRef{External: "//bigquery.googleapis.com/projects/my-project/datasets/my_dataset/tables/my_table"},
			wantTable: "my-project.my_dataset.my_table",
		},
		{
			name:      "legacy external table reference in the Pub/Sub format",
			tableRef:  krmbigqueryv1beta1.BigQueryTableRef{External: "my-project.my_dataset.my_table"},
			wantTable: "my-project.my_dataset.my_table",
		},
		{
			name:      "legacy external table reference with a colon after the project",
			tableRef:  krmbigqueryv1beta1.BigQueryTableRef{External: "my-project:my_dataset.my_table"},
			wantTable: "my-project.my_dataset.my_table",
		},
		{
			name:      "legacy external table reference with a domain-scoped project",
			tableRef:  krmbigqueryv1beta1.BigQueryTableRef{External: "example.com:my-project.my_dataset.my_table"},
			wantTable: "example.com:my-project.my_dataset.my_table",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := &SubscriptionModel{config: config.ControllerConfig{HTTPClient: &http.Client{}}}
			op := &directbase.AdapterForObjectOperation{
				Reader: fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(newBigQueryTable(t)).Build(),
				Object: newSubscriptionWithTableRef(t, tc.tableRef),
			}

			adapter, err := model.AdapterForObject(context.Background(), op)

			if err != nil {
				t.Fatalf("AdapterForObject() error: %v", err)
			}
			if got := adapter.(*pubSubSubscriptionAdapter).desired.GetBigqueryConfig().GetTable(); got != tc.wantTable {
				t.Errorf("desired bigquery_config.table = %q, want %q", got, tc.wantTable)
			}
		})
	}
}

// newBigQueryTable returns a BigQueryTable named "my-table" whose
// status.externalRef is projects/my-project/datasets/my_dataset/tables/my_table.
func newBigQueryTable(t *testing.T) *unstructured.Unstructured {
	t.Helper()
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krmbigqueryv1beta1.BigQueryTableGVK)
	u.SetNamespace("test-ns")
	u.SetName("my-table")
	if err := unstructured.SetNestedField(u.Object, "projects/my-project/datasets/my_dataset/tables/my_table", "status", "externalRef"); err != nil {
		t.Fatalf("setting status.externalRef: %v", err)
	}
	return u
}

// newSubscriptionWithTableRef returns a PubSubSubscription in my-project that
// writes to the BigQuery table referenced by tableRef.
func newSubscriptionWithTableRef(t *testing.T, tableRef krmbigqueryv1beta1.BigQueryTableRef) *unstructured.Unstructured {
	t.Helper()
	obj := &krm.PubSubSubscription{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "my-subscription",
			Namespace:   "test-ns",
			Annotations: map[string]string{"cnrm.cloud.google.com/project-id": "my-project"},
		},
		Spec: krm.PubSubSubscriptionSpec{
			TopicRef:       &krm.PubSubTopicRef{External: "projects/my-project/topics/my-topic"},
			BigqueryConfig: &krm.BigQueryConfig{TableRef: &tableRef},
		},
	}
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		t.Fatalf("converting PubSubSubscription to unstructured: %v", err)
	}
	u := &unstructured.Unstructured{Object: m}
	u.SetGroupVersionKind(krm.PubSubSubscriptionGVK)
	return u
}

func TestSubscriptionModel_AdapterForObject_InvalidBigQueryTable(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  string
	}{
		{
			name:     "table without a project and dataset",
			external: "my_table",
			wantErr:  `format of BigQueryTable external="my_table" was not known`,
		},
		{
			name:     "legacy format with an empty table",
			external: "my-project.my_dataset.",
			wantErr:  `format of BigQueryTable external="my-project.my_dataset." was not known`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := &SubscriptionModel{config: config.ControllerConfig{HTTPClient: &http.Client{}}}
			op := &directbase.AdapterForObjectOperation{
				Reader: fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build(),
				Object: newSubscriptionWithTableRef(t, krmbigqueryv1beta1.BigQueryTableRef{External: tc.external}),
			}

			_, err := model.AdapterForObject(context.Background(), op)

			if err == nil {
				t.Fatalf("AdapterForObject() error = nil, want an error containing %q", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("AdapterForObject() error = %v, want an error containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestCompareSubscription_BigQueryTable(t *testing.T) {
	tests := []struct {
		name         string
		actualTable  string
		desiredTable string
		wantPaths    []string
	}{
		{
			name:         "same table",
			actualTable:  "my-project.my_dataset.my_table",
			desiredTable: "my-project.my_dataset.my_table",
		},
		{
			name:         "same table read back with a colon after the project",
			actualTable:  "my-project:my_dataset.my_table",
			desiredTable: "my-project.my_dataset.my_table",
		},
		{
			name:         "different table",
			actualTable:  "my-project.my_dataset.old_table",
			desiredTable: "my-project.my_dataset.new_table",
			wantPaths:    []string{"bigquery_config"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := newBigQuerySubscriptionFromGCP(tc.actualTable)
			desired := &pb.Subscription{
				Name:           "projects/my-project/subscriptions/my-subscription",
				Topic:          "projects/my-project/topics/my-topic",
				BigqueryConfig: &pb.BigQueryConfig{Table: tc.desiredTable},
			}

			_, mask, err := compareSubscription(context.Background(), actual, desired)

			if err != nil {
				t.Fatalf("compareSubscription() error: %v", err)
			}
			if diff := cmp.Diff(tc.wantPaths, mask.GetPaths(), cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("compareSubscription() update mask mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// newBigQuerySubscriptionFromGCP returns a BigQuery subscription as GetSubscription
// returns it, including the server defaults for unset fields.
func newBigQuerySubscriptionFromGCP(table string) *pb.Subscription {
	return &pb.Subscription{
		Name:                     "projects/my-project/subscriptions/my-subscription",
		Topic:                    "projects/my-project/topics/my-topic",
		PushConfig:               &pb.PushConfig{},
		BigqueryConfig:           &pb.BigQueryConfig{Table: table, State: pb.BigQueryConfig_ACTIVE},
		AckDeadlineSeconds:       10,
		MessageRetentionDuration: durationpb.New(7 * 24 * time.Hour),
		ExpirationPolicy:         &pb.ExpirationPolicy{Ttl: durationpb.New(31 * 24 * time.Hour)},
		State:                    pb.Subscription_ACTIVE,
	}
}
