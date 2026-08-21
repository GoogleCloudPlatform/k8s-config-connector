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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &FirestoreBackupScheduleIdentity{}
	_ identity.Resource   = &FirestoreBackupSchedule{}
)

var FirestoreBackupScheduleIdentityFormat = gcpurls.Template[FirestoreBackupScheduleIdentity]("firestore.googleapis.com", "projects/{project}/databases/{database}/backupSchedules/{backupschedule}")

// +k8s:deepcopy-gen=false
type FirestoreBackupScheduleIdentity struct {
	Project        string
	Database       string
	BackupSchedule string
}

func (i *FirestoreBackupScheduleIdentity) String() string {
	return FirestoreBackupScheduleIdentityFormat.ToString(*i)
}

func (i *FirestoreBackupScheduleIdentity) FromExternal(ref string) error {
	parsed, match, err := FirestoreBackupScheduleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirestoreBackupSchedule external=%q was not known (use %s): %w", ref, FirestoreBackupScheduleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirestoreBackupSchedule external=%q was not known (use %s)", ref, FirestoreBackupScheduleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirestoreBackupScheduleIdentity) Host() string {
	return FirestoreBackupScheduleIdentityFormat.Host()
}

func (i *FirestoreBackupScheduleIdentity) Parent() *v1beta1.FirestoreDatabaseIdentity {
	return &v1beta1.FirestoreDatabaseIdentity{
		Parent: &parent.ProjectParent{
			ProjectID: i.Project,
		},
		Database: i.Database,
	}
}

func getIdentityFromFirestoreBackupScheduleSpec(ctx context.Context, reader client.Reader, obj client.Object) (*FirestoreBackupScheduleIdentity, error) {
	var databaseRef v1beta1.FirestoreDatabaseRef

	switch t := obj.(type) {
	case *FirestoreBackupSchedule:
		databaseRef = t.Spec.DatabaseRef
	case *unstructured.Unstructured:
		var spec struct {
			DatabaseRef v1beta1.FirestoreDatabaseRef `json:"databaseRef"`
		}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(t.Object["spec"].(map[string]interface{}), &spec); err != nil {
			return nil, fmt.Errorf("error parsing spec from unstructured: %w", err)
		}
		databaseRef = spec.DatabaseRef
	default:
		return nil, fmt.Errorf("unrecognized object type %T", obj)
	}

	if err := databaseRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("normalizing databaseRef: %w", err)
	}

	dbIdentity := &v1beta1.FirestoreDatabaseIdentity{}
	if err := dbIdentity.FromExternal(databaseRef.External); err != nil {
		return nil, fmt.Errorf("parsing databaseRef.external=%q: %w", databaseRef.External, err)
	}

	identity := &FirestoreBackupScheduleIdentity{
		Project:        dbIdentity.Parent.ProjectID,
		Database:       dbIdentity.Database,
		BackupSchedule: "", // Will be populated from status or creation
	}

	return identity, nil
}

func (obj *FirestoreBackupSchedule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromFirestoreBackupScheduleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" && obj.Status.ObservedState != nil {
		externalRef = common.ValueOf(obj.Status.ObservedState.Name)
	}

	if externalRef != "" {
		statusIdentity := &FirestoreBackupScheduleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Project != specIdentity.Project || statusIdentity.Database != specIdentity.Database {
			return nil, fmt.Errorf("cannot change FirestoreBackupSchedule identity (old=%q, new parent=%q)", externalRef, specIdentity.Parent().String())
		}

		specIdentity.BackupSchedule = statusIdentity.BackupSchedule
	}

	return specIdentity, nil
}
