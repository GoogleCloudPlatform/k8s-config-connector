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

package preflight

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"

	"github.com/blang/semver"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var (
	ulog       = ctrl.Log.WithName("UpgradeChecker")
	devVersion = semver.MustParse("0.0.0-dev")
)

// NewUpgradeChecker provides an implementation of declarative.Preflight that
// does version comparison between the version of the existing KCC and the version to deploy.
// If it's a major version change, it returns error and surface the error status on the DeclarativeObject.
func NewUpgradeChecker(client client.Client, repo manifest.Repository) *UpgradeChecker {
	return &UpgradeChecker{client: client, repo: repo}
}

type UpgradeChecker struct {
	client client.Client
	repo   manifest.Repository
}

func (u *UpgradeChecker) Preflight(ctx context.Context, o declarative.DeclarativeObject) error {
	ulog.Info("preflight check before reconciling the object", "kind", o.GetObjectKind().GroupVersionKind().Kind, "name", o.GetName(), "namespace", o.GetNamespace())
	if !o.GetDeletionTimestamp().IsZero() {
		return nil
	}
	ns := &corev1.Namespace{}
	if err := u.client.Get(ctx, types.NamespacedName{Name: k8s.CNRMSystemNamespace}, ns); err != nil {
		if apierrors.IsNotFound(err) {
			ulog.Info(fmt.Sprintf("%v namespace is not found. Continue the reconciliation.", k8s.CNRMSystemNamespace))
			return nil
		}
		return err
	}

	currentVersionRaw := ns.GetAnnotations()[k8s.VersionAnnotation]
	if currentVersionRaw == "" {
		ulog.Info(fmt.Sprintf("WARNING: No ConfigConnector version is annotated with '%v' namespace. Attempt to deploy ConfigConnector bundle with best effort", k8s.CNRMSystemNamespace))
		return nil
	}

	channel, err := u.repo.LoadChannel(ctx, manifest.StableChannel)
	if err != nil {
		return fmt.Errorf("preflight check failed loading the channel %v: %w", manifest.StableChannel, err)
	}
	version, err := channel.Latest(ctx, manifest.ConfigConnectorComponentName)
	if err != nil {
		return fmt.Errorf("preflight check failed resolving the version to deploy: %w", err)
	}
	if version == nil {
		return fmt.Errorf("could not find the latest version in channel %v", manifest.StableChannel)
	}
	versionToDeployRaw := version.Version
	currentVersion, err := semver.ParseTolerant(currentVersionRaw)
	if err != nil {
		return fmt.Errorf("current version %v is not a valid semantic version: %w", currentVersionRaw, err)
	}
	ulog.Info("Checking version", "current version", currentVersion)
	versionToDeploy, err := semver.ParseTolerant(versionToDeployRaw)
	if err != nil {
		return fmt.Errorf("the version to deploy %v is not a valid semantic version: %w", versionToDeployRaw, err)
	}
	ulog.Info("Checking version", "version to deploy", versionToDeploy)
	if compareMajorOnly(currentVersion, versionToDeploy) != 0 {
		return fmt.Errorf("incompatible version: stop reconciling the existing ConfigConnector of version %v to version %v since it's a major version change. Please kubectl delete the existing ConfigConnector object and recreate it", currentVersion, versionToDeploy)
	}
	return nil
}

func compareMajorOnly(v, w semver.Version) int {
	if v.Equals(devVersion) {
		// If we are using a dev controller ignore semver drift.
		return 0
	}
	if v.Major != w.Major {
		if v.Major > w.Major {
			return 1
		}
		return -1
	}
	return 0
}
