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

package documentai

import (
	"context"
	"fmt"
	"regexp"

	gcp "cloud.google.com/go/documentai/apiv1"
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krmv1alpha1.DocumentAIProcessorDefaultVersionGVK, NewProcessorDefaultVersionModel, registry.CannotBeDeleted())
}

func NewProcessorDefaultVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelProcessorDefaultVersion{config: *config}, nil
}

var _ directbase.Model = &modelProcessorDefaultVersion{}

type modelProcessorDefaultVersion struct {
	config config.ControllerConfig
}

func (m *modelProcessorDefaultVersion) client(ctx context.Context, location string) (*gcp.DocumentProcessorClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-documentai.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewDocumentProcessorRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Processor client: %w", err)
	}
	return gcpClient, err
}

func (m *modelProcessorDefaultVersion) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krmv1alpha1.DocumentAIProcessorDefaultVersion{}
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
	processorID, ok := id.(*krmv1alpha1.DocumentAIProcessorDefaultVersionIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", id)
	}

	gcpClient, err := m.client(ctx, processorID.Location)
	if err != nil {
		return nil, err
	}

	return &ProcessorDefaultVersionAdapter{
		id:        processorID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelProcessorDefaultVersion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ProcessorDefaultVersionAdapter struct {
	id        *krmv1alpha1.DocumentAIProcessorDefaultVersionIdentity
	gcpClient *gcp.DocumentProcessorClient
	desired   *krmv1alpha1.DocumentAIProcessorDefaultVersion
	actual    *pb.Processor
}

var _ directbase.Adapter = &ProcessorDefaultVersionAdapter{}

func (a *ProcessorDefaultVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Processor for DocumentAIProcessorDefaultVersion", "name", a.id.String())

	req := &pb.GetProcessorRequest{Name: a.id.String()}
	processor, err := a.gcpClient.GetProcessor(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Processor %q: %w", a.id.String(), err)
	}

	a.actual = processor
	return true, nil
}

func (a *ProcessorDefaultVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DocumentAIProcessorDefaultVersion", "name", a.id.String())

	req := &pb.SetDefaultProcessorVersionRequest{
		Processor:               a.id.String(),
		DefaultProcessorVersion: a.desired.Spec.Version,
	}
	op, err := a.gcpClient.SetDefaultProcessorVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("setting default processor version for %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for setting default processor version for %s: %w", a.id.String(), err)
	}

	updated, err := a.gcpClient.GetProcessor(ctx, &pb.GetProcessorRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting Processor after setting default version %s: %w", a.id.String(), err)
	}
	a.actual = updated

	return a.updateStatus(ctx, createOp)
}

func (a *ProcessorDefaultVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DocumentAIProcessorDefaultVersion", "name", a.id.String())

	diffs, err := a.diff(ctx)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.SetDefaultProcessorVersionRequest{
		Processor:               a.id.String(),
		DefaultProcessorVersion: a.desired.Spec.Version,
	}
	op, err := a.gcpClient.SetDefaultProcessorVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("setting default processor version for %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for setting default processor version for %s: %w", a.id.String(), err)
	}

	updated, err := a.gcpClient.GetProcessor(ctx, &pb.GetProcessorRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting Processor after setting default version %s: %w", a.id.String(), err)
	}
	a.actual = updated

	return a.updateStatus(ctx, updateOp)
}

func (a *ProcessorDefaultVersionAdapter) diff(ctx context.Context) (*structuredreporting.Diff, error) {
	desiredVersion := a.desired.Spec.Version
	actualVersion := a.actual.GetDefaultProcessorVersion()

	report := &structuredreporting.Diff{}
	if projectNumberDiffSuppress(desiredVersion, actualVersion) {
		return report, nil
	}

	// Check if desired matches any alias that resolves to actualVersion
	for _, alias := range a.actual.GetProcessorVersionAliases() {
		aliasName := alias.GetAlias()
		targetVersion := alias.GetProcessorVersion()
		if (projectNumberDiffSuppress(desiredVersion, aliasName) || desiredVersion == alias.GetAlias()) &&
			projectNumberDiffSuppress(targetVersion, actualVersion) {
			return report, nil
		}
	}

	report.AddField("spec.version", actualVersion, desiredVersion)
	return report, nil
}

func projectNumberDiffSuppress(old, new string) bool {
	if old == new {
		return true
	}
	re := regexp.MustCompile(`projects/[^/]+`)
	replacement := "projects/equal"
	a2 := re.ReplaceAllString(old, replacement)
	b2 := re.ReplaceAllString(new, replacement)
	return a2 == b2
}

func (a *ProcessorDefaultVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *ProcessorDefaultVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DocumentAIProcessorDefaultVersion is a no-op because processor default versions cannot be deleted in GCP", "name", a.id.String())
	return true, nil
}

func (a *ProcessorDefaultVersionAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	status := &krmv1alpha1.DocumentAIProcessorDefaultVersionStatus{}
	return op.UpdateStatus(ctx, status, nil)
}
