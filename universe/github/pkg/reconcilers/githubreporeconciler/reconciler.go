// Copyright 2023 Google LLC
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

package githubreporeconciler

import (
	"context"
	"fmt"
	"reflect"
	"sort"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	api "github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/reconcilers"
	github "github.com/google/go-github/v53/github"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

type GithubRepoReconciler struct {
}

func (r *GithubRepoReconciler) GetWatchInfo() reconcilers.WatchInfo {
	return reconcilers.WatchInfo{
		Primary: &api.GithubRepo{},
	}
}

func (r *GithubRepoReconciler) NewOp(req *reconcilers.ReconcileRequest) reconcilers.ReconciliationOp {
	return &reconcileOp{
		ReconcileRequest: req,
	}
}

type reconcileOp struct {
	*reconcilers.ReconcileRequest

	tagProtections []*github.TagProtection
}

func (op *reconcileOp) Reconcile(ctx context.Context) (*reconcilers.ReconcileResult, error) {
	log := klog.FromContext(ctx)

	id := op.ID.NamespacedName

	log.Info("reconcile request for repo", "id", id)

	want := &api.GithubRepo{}
	if err := op.Client.Get(ctx, id, want); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	orgName := want.Spec.Org
	repoName := want.Spec.Repo

	actual, err := op.find(ctx, orgName, repoName)
	if err != nil {
		return nil, err
	}

	want = normalize(want.DeepCopy())
	actual = normalize(actual.DeepCopy())

	if reflect.DeepEqual(actual.Spec, want.Spec) {
		return nil, nil
	}

	if !reflect.DeepEqual(actual.Spec.ProtectedTags, want.Spec.ProtectedTags) {
		actualTags := collectSet(actual.Spec.ProtectedTags, func(t api.ProtectedTag) string { return t.Pattern })
		wantTags := collectSet(want.Spec.ProtectedTags, func(t api.ProtectedTag) string { return t.Pattern })

		tagsToRemove := actualTags.Difference(wantTags)
		tagsToAdd := wantTags.Difference(actualTags)

		for _, pattern := range tagsToAdd.UnsortedList() {
			if _, err := op.ExternalClient.CreateRepoTagProtection(ctx, orgName, repoName, pattern); err != nil {
				return nil, fmt.Errorf("creating repo tag protection: %w", err)
			}
		}
		for _, pattern := range tagsToRemove.UnsortedList() {
			var found *github.TagProtection
			for _, tagProtection := range op.tagProtections {
				if tagProtection.GetPattern() == pattern {
					found = tagProtection
				}
			}
			if found == nil {
				// shouldn't be possible, it was there before
				return nil, fmt.Errorf("tagPermission not found for %q", pattern)
			}
			if err := op.ExternalClient.DeleteRepoTagProtection(ctx, orgName, repoName, found.GetID()); err != nil {
				return nil, fmt.Errorf("deleting repo tag protection: %w", err)
			}
		}
	}

	return nil, nil
}

func (op *reconcileOp) find(ctx context.Context, orgName string, repoName string) (*api.GithubRepo, error) {
	_, err := op.ExternalClient.GetRepo(ctx, orgName, repoName)
	if err != nil {
		return nil, fmt.Errorf("error fetching github repo %s/%s: %v", orgName, repoName, err)
	}

	found := &api.GithubRepo{}
	found.Spec.Org = orgName
	found.Spec.Repo = repoName

	tagProtections, err := op.ExternalClient.ListRepoTagProtection(ctx, orgName, repoName)
	if err != nil {
		return nil, fmt.Errorf("error listing tag protection for github repo %s/%s: %v", orgName, repoName, err)
	}

	// Save so we have the string -> id map for later (for deletion, in particular)
	op.tagProtections = tagProtections

	for _, tagProtection := range tagProtections {
		found.Spec.ProtectedTags = append(found.Spec.ProtectedTags, api.ProtectedTag{Pattern: tagProtection.GetPattern()})
	}

	return found, nil

}

func normalize(obj *api.GithubRepo) *api.GithubRepo {
	sort.Slice(obj.Spec.ProtectedTags, func(i, j int) bool {
		return obj.Spec.ProtectedTags[i].Pattern < obj.Spec.ProtectedTags[j].Pattern
	})
	return obj
}

func collectSet[T any, V comparable](values []T, fn func(t T) V) sets.Set[V] {
	result := sets.New[V]()
	for _, value := range values {
		result.Insert(fn(value))
	}
	return result
}
