package status

import (
	"context"
	"fmt"
	"reflect"

	semver "github.com/blang/semver/v4"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/utils"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// NewVersionCheck provides an implementation of declarative.Reconciled that
// checks the version of the operator if it is up to the version required by the manifest
func NewVersionCheck(client client.Client, operatorVersionString string) (*versionCheck, error) {
	operatorVersion, err := semver.Parse(operatorVersionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse operator version %q: %v", operatorVersionString, err)
	}
	return &versionCheck{client: client, operatorVersion: operatorVersion}, nil
}

type versionCheck struct {
	client          client.Client
	operatorVersion semver.Version
}

func (p *versionCheck) VersionCheck(
	ctx context.Context,
	src declarative.DeclarativeObject,
	objs *manifest.Objects,
) (bool, error) {
	log := log.FromContext(ctx)
	var minOperatorVersion semver.Version

	// Look for annotation from any resource with the max version
	for _, obj := range objs.Items {
		annotations := obj.UnstructuredObject().GetAnnotations()
		if versionNeededStr, ok := annotations["addons.k8s.io/min-operator-version"]; ok {
			log.WithValues("min-operator-version", versionNeededStr).Info("Got version requirement addons.k8s.io/operator-version")

			versionNeeded, err := semver.Parse(versionNeededStr)
			if err != nil {
				log.WithValues("version", versionNeededStr).Error(err, "Unable to parse version restriction")
				return false, err
			}

			if versionNeeded.GT(minOperatorVersion) {
				minOperatorVersion = versionNeeded
			}
		}
	}

	if p.operatorVersion.GTE(minOperatorVersion) {
		return true, nil
	}

	errors := []string{
		fmt.Sprintf("manifest needs operator version >= %v, this operator is version %v", minOperatorVersion.String(),
			p.operatorVersion.String()),
	}

	currentStatus, err := utils.GetCommonStatus(src)
	if err != nil {
		log.Error(err, "getting status")
		return false, err
	}

	status := currentStatus
	status.Healthy = false
	status.Errors = errors

	if !reflect.DeepEqual(status, currentStatus) {
		err := utils.SetCommonStatus(src, status)
		if err != nil {
			log.Error(err, "unable to update status")
		}
	}

	return false, fmt.Errorf("operator not qualified, manifest needs operator version >= %v", minOperatorVersion.String())
}
