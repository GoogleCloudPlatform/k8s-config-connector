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

package bigquery

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryDatasetAccessGVK, NewDatasetAccessModel)
}

func NewDatasetAccessModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &datasetAccessModel{config: config}, nil
}

type datasetAccessModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &datasetAccessModel{}

func (m *datasetAccessModel) datasetsService(ctx context.Context) (*api.DatasetsService, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpService, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigQuery client: %w", err)
	}
	return gcpService.Datasets, nil
}

func (m *datasetAccessModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryDatasetAccess{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	accessID, ok := id.(*krm.BigQueryDatasetAccessIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", id)
	}

	datasetsService, err := m.datasetsService(ctx)
	if err != nil {
		return nil, err
	}

	return &datasetAccessAdapter{
		datasetsService: datasetsService,
		id:              accessID,
		desired:         obj,
	}, nil
}

func (m *datasetAccessModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type datasetAccessAdapter struct {
	datasetsService  *api.DatasetsService
	id               *krm.BigQueryDatasetAccessIdentity
	desired          *krm.BigQueryDatasetAccess
	actual           *api.DatasetAccess
	apiUpdatedMember bool
}

var _ directbase.Adapter = &datasetAccessAdapter{}

func (a *datasetAccessAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigQueryDatasetAccess", "id", a.id.String())

	dataset, err := a.datasetsService.Get(a.id.Project, a.id.Dataset).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataset %s/%s: %w", a.id.Project, a.id.Dataset, err)
	}

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	for _, entry := range dataset.Access {
		if matched, apiUpdated := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
			a.actual = entry
			a.apiUpdatedMember = apiUpdated
			return true, nil
		}
	}

	return false, nil
}

func matchesAccessEntry(entry *api.DatasetAccess, desired *api.DatasetAccess, desiredIamMember string) (bool, bool) {
	if entry == nil || desired == nil {
		return false, false
	}

	// 1. View
	if desired.View != nil {
		if entry.View == nil {
			return false, false
		}
		if entry.View.ProjectId == desired.View.ProjectId &&
			entry.View.DatasetId == desired.View.DatasetId &&
			entry.View.TableId == desired.View.TableId {
			return true, false
		}
		return false, false
	}

	// 2. Routine
	if desired.Routine != nil {
		if entry.Routine == nil {
			return false, false
		}
		if entry.Routine.ProjectId == desired.Routine.ProjectId &&
			entry.Routine.DatasetId == desired.Routine.DatasetId &&
			entry.Routine.RoutineId == desired.Routine.RoutineId {
			return true, false
		}
		return false, false
	}

	// 3. Dataset
	if desired.Dataset != nil {
		if entry.Dataset == nil || entry.Dataset.Dataset == nil || desired.Dataset.Dataset == nil {
			return false, false
		}
		if entry.Dataset.Dataset.ProjectId == desired.Dataset.Dataset.ProjectId &&
			entry.Dataset.Dataset.DatasetId == desired.Dataset.Dataset.DatasetId &&
			reflect.DeepEqual(entry.Dataset.TargetTypes, desired.Dataset.TargetTypes) {
			return true, false
		}
		return false, false
	}

	// 4. For all other access entries, the role must match.
	if canonicalRole(entry.Role) != canonicalRole(desired.Role) {
		return false, false
	}

	if desired.Domain != "" {
		if strings.EqualFold(entry.Domain, desired.Domain) {
			return true, false
		}
	}
	if desired.GroupByEmail != "" {
		if strings.EqualFold(entry.GroupByEmail, desired.GroupByEmail) {
			return true, false
		}
	}
	if desired.UserByEmail != "" {
		if strings.EqualFold(entry.UserByEmail, desired.UserByEmail) {
			return true, false
		}
	}
	if desired.SpecialGroup != "" {
		if entry.SpecialGroup == desired.SpecialGroup {
			return true, false
		}
	}

	if desiredIamMember != "" {
		if entry.IamMember != "" {
			if strings.EqualFold(entry.IamMember, desiredIamMember) ||
				strings.EqualFold(stripTombstone(entry.IamMember), stripTombstone(desiredIamMember)) {
				return true, false
			}
		}
		memberValue := extractMemberValue(desiredIamMember)
		if memberValue != "" {
			if entry.UserByEmail != "" && strings.EqualFold(entry.UserByEmail, memberValue) {
				return true, true
			}
			if entry.GroupByEmail != "" && strings.EqualFold(entry.GroupByEmail, memberValue) {
				return true, true
			}
			if entry.Domain != "" && strings.EqualFold(entry.Domain, memberValue) {
				return true, true
			}
			if entry.SpecialGroup != "" && strings.EqualFold(entry.SpecialGroup, memberValue) {
				return true, true
			}
		}
	}

	return false, false
}

func canonicalRole(r string) string {
	switch r {
	case "roles/bigquery.dataOwner":
		return "OWNER"
	case "roles/bigquery.dataEditor":
		return "WRITER"
	case "roles/bigquery.dataViewer":
		return "READER"
	default:
		return r
	}
}

func stripTombstone(member string) string {
	if strings.HasPrefix(strings.ToLower(member), "deleted:") {
		member = member[len("deleted:"):]
	}
	if i := strings.Index(member, "?"); i != -1 {
		member = member[:i]
	}
	return member
}

func extractMemberValue(iamMember string) string {
	s := stripTombstone(iamMember)
	if idx := strings.Index(s, ":"); idx != -1 {
		s = s[idx+1:]
	}
	return s
}

func (a *datasetAccessAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigQueryDatasetAccess", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	backoff := wait.Backoff{
		Steps:    5,
		Duration: 100 * time.Millisecond,
		Factor:   2.0,
		Jitter:   0.1,
	}

	err := retry.OnError(backoff, isConflictOrPreconditionFailed, func() error {
		dataset, err := a.datasetsService.Get(a.id.Project, a.id.Dataset).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting dataset %s/%s for access creation: %w", a.id.Project, a.id.Dataset, err)
		}

		for _, entry := range dataset.Access {
			if matched, apiUpdated := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
				a.actual = entry
				a.apiUpdatedMember = apiUpdated
				return nil
			}
		}

		newAccess := append(dataset.Access, desiredAccess)
		patchReq := &api.Dataset{
			Access: newAccess,
			Etag:   dataset.Etag,
		}

		patchedDataset, err := a.datasetsService.Patch(a.id.Project, a.id.Dataset, patchReq).Context(ctx).Do()
		if err != nil {
			return err
		}

		for _, entry := range patchedDataset.Access {
			if matched, apiUpdated := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
				a.actual = entry
				a.apiUpdatedMember = apiUpdated
				break
			}
		}
		if a.actual == nil {
			a.actual = desiredAccess
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("creating BigQueryDatasetAccess: %w", err)
	}

	return a.updateStatus(ctx, op)
}

func isConflictOrPreconditionFailed(err error) bool {
	if err == nil {
		return false
	}
	if direct.HasHTTPCode(err, http.StatusPreconditionFailed) || direct.HasHTTPCode(err, http.StatusConflict) {
		return true
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "precondition") || strings.Contains(msg, "conflict")
}

func (a *datasetAccessAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigQueryDatasetAccess", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	matched, apiUpdated := matchesAccessEntry(a.actual, desiredAccess, desiredIamMember)
	if !matched {
		diff := describeAccessDiff(a.actual, desiredAccess, desiredIamMember)
		return fmt.Errorf("BigQueryDatasetAccess is immutable and cannot be updated: %s", diff)
	}
	a.apiUpdatedMember = apiUpdated

	return a.updateStatus(ctx, op)
}

func describeAccessDiff(actual, desired *api.DatasetAccess, desiredIamMember string) string {
	if actual == nil {
		return "existing access entry not found on dataset"
	}
	if desired == nil {
		return "desired access entry is nil"
	}

	var diffs []string

	// Check View diff
	if desired.View != nil || actual.View != nil {
		if actual.View == nil {
			diffs = append(diffs, fmt.Sprintf("view added (%s/%s/%s)", desired.View.ProjectId, desired.View.DatasetId, desired.View.TableId))
		} else if desired.View == nil {
			diffs = append(diffs, fmt.Sprintf("view removed (was %s/%s/%s)", actual.View.ProjectId, actual.View.DatasetId, actual.View.TableId))
		} else if actual.View.ProjectId != desired.View.ProjectId || actual.View.DatasetId != desired.View.DatasetId || actual.View.TableId != desired.View.TableId {
			diffs = append(diffs, fmt.Sprintf("view changed from %s/%s/%s to %s/%s/%s",
				actual.View.ProjectId, actual.View.DatasetId, actual.View.TableId,
				desired.View.ProjectId, desired.View.DatasetId, desired.View.TableId))
		}
	}

	// Check Routine diff
	if desired.Routine != nil || actual.Routine != nil {
		if actual.Routine == nil {
			diffs = append(diffs, fmt.Sprintf("routine added (%s/%s/%s)", desired.Routine.ProjectId, desired.Routine.DatasetId, desired.Routine.RoutineId))
		} else if desired.Routine == nil {
			diffs = append(diffs, fmt.Sprintf("routine removed (was %s/%s/%s)", actual.Routine.ProjectId, actual.Routine.DatasetId, actual.Routine.RoutineId))
		} else if actual.Routine.ProjectId != desired.Routine.ProjectId || actual.Routine.DatasetId != desired.Routine.DatasetId || actual.Routine.RoutineId != desired.Routine.RoutineId {
			diffs = append(diffs, fmt.Sprintf("routine changed from %s/%s/%s to %s/%s/%s",
				actual.Routine.ProjectId, actual.Routine.DatasetId, actual.Routine.RoutineId,
				desired.Routine.ProjectId, desired.Routine.DatasetId, desired.Routine.RoutineId))
		}
	}

	// Check Dataset diff
	if desired.Dataset != nil || actual.Dataset != nil {
		if actual.Dataset == nil {
			diffs = append(diffs, "dataset entry added")
		} else if desired.Dataset == nil {
			diffs = append(diffs, "dataset entry removed")
		} else {
			if actual.Dataset.Dataset != nil && desired.Dataset.Dataset != nil {
				if actual.Dataset.Dataset.ProjectId != desired.Dataset.Dataset.ProjectId || actual.Dataset.Dataset.DatasetId != desired.Dataset.Dataset.DatasetId {
					diffs = append(diffs, fmt.Sprintf("dataset reference changed from %s/%s to %s/%s",
						actual.Dataset.Dataset.ProjectId, actual.Dataset.Dataset.DatasetId,
						desired.Dataset.Dataset.ProjectId, desired.Dataset.Dataset.DatasetId))
				}
			}
			if !reflect.DeepEqual(actual.Dataset.TargetTypes, desired.Dataset.TargetTypes) {
				diffs = append(diffs, fmt.Sprintf("dataset targetTypes changed from %v to %v", actual.Dataset.TargetTypes, desired.Dataset.TargetTypes))
			}
		}
	}

	// Check Role diff
	if canonicalRole(actual.Role) != canonicalRole(desired.Role) {
		diffs = append(diffs, fmt.Sprintf("role changed from %q to %q", actual.Role, desired.Role))
	}

	// Check UserByEmail diff
	if desired.UserByEmail != "" && !strings.EqualFold(actual.UserByEmail, desired.UserByEmail) {
		diffs = append(diffs, fmt.Sprintf("userByEmail changed from %q to %q", actual.UserByEmail, desired.UserByEmail))
	}

	// Check GroupByEmail diff
	if desired.GroupByEmail != "" && !strings.EqualFold(actual.GroupByEmail, desired.GroupByEmail) {
		diffs = append(diffs, fmt.Sprintf("groupByEmail changed from %q to %q", actual.GroupByEmail, desired.GroupByEmail))
	}

	// Check Domain diff
	if desired.Domain != "" && !strings.EqualFold(actual.Domain, desired.Domain) {
		diffs = append(diffs, fmt.Sprintf("domain changed from %q to %q", actual.Domain, desired.Domain))
	}

	// Check SpecialGroup diff
	if desired.SpecialGroup != "" && actual.SpecialGroup != desired.SpecialGroup {
		diffs = append(diffs, fmt.Sprintf("specialGroup changed from %q to %q", actual.SpecialGroup, desired.SpecialGroup))
	}

	// Check IamMember diff
	if desiredIamMember != "" {
		memberVal := extractMemberValue(desiredIamMember)
		if actual.IamMember != "" && !strings.EqualFold(stripTombstone(actual.IamMember), stripTombstone(desiredIamMember)) {
			diffs = append(diffs, fmt.Sprintf("iamMember changed from %q to %q", actual.IamMember, desiredIamMember))
		} else if actual.IamMember == "" {
			actualGrantee := actualGranteeString(actual)
			if actualGrantee != "" && !strings.EqualFold(extractMemberValue(actualGrantee), memberVal) {
				diffs = append(diffs, fmt.Sprintf("grantee changed from %s to iamMember %q", actualGrantee, desiredIamMember))
			}
		}
	}

	if len(diffs) == 0 {
		return "grantee or configuration type changed"
	}

	return strings.Join(diffs, "; ")
}

func actualGranteeString(entry *api.DatasetAccess) string {
	if entry.UserByEmail != "" {
		return "userByEmail: " + entry.UserByEmail
	}
	if entry.GroupByEmail != "" {
		return "groupByEmail: " + entry.GroupByEmail
	}
	if entry.Domain != "" {
		return "domain: " + entry.Domain
	}
	if entry.SpecialGroup != "" {
		return "specialGroup: " + entry.SpecialGroup
	}
	return ""
}

func (a *datasetAccessAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigQueryDatasetAccess", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	backoff := wait.Backoff{
		Steps:    5,
		Duration: 100 * time.Millisecond,
		Factor:   2.0,
		Jitter:   0.1,
	}

	err := retry.OnError(backoff, isConflictOrPreconditionFailed, func() error {
		dataset, err := a.datasetsService.Get(a.id.Project, a.id.Dataset).Context(ctx).Do()
		if err != nil {
			if direct.IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("getting dataset %s/%s for access deletion: %w", a.id.Project, a.id.Dataset, err)
		}

		foundIdx := -1
		for idx, entry := range dataset.Access {
			if matched, _ := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
				foundIdx = idx
				break
			}
		}

		if foundIdx == -1 {
			return nil
		}

		updatedAccess := make([]*api.DatasetAccess, 0, len(dataset.Access)-1)
		for idx, entry := range dataset.Access {
			if idx != foundIdx {
				updatedAccess = append(updatedAccess, entry)
			}
		}

		patchReq := &api.Dataset{
			Access: updatedAccess,
			Etag:   dataset.Etag,
		}

		_, err = a.datasetsService.Patch(a.id.Project, a.id.Dataset, patchReq).Context(ctx).Do()
		if err != nil {
			if direct.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return false, fmt.Errorf("deleting BigQueryDatasetAccess: %w", err)
	}

	return true, nil
}

func (a *datasetAccessAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryDatasetAccessStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.apiUpdatedMember {
		status.ApiUpdatedMember = direct.LazyPtr(true)
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *datasetAccessAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryDatasetAccess{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryDatasetAccessSpec_FromAPI(mapCtx, a.actual))
	obj.Spec.ProjectRef = parent.ProjectRef{External: a.id.Project}
	obj.Spec.DatasetId = a.id.Dataset
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}
