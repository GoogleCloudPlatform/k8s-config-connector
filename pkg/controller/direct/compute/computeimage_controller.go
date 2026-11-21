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

// +tool:controller
// proto.service: google.cloud.compute.v1.Images
// proto.message: google.cloud.compute.v1.Image
// crd.type: ComputeImage
// crd.version: v1beta1

package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeImageGVK, NewComputeImageModel)
}

func NewComputeImageModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeImageModel{config: config}, nil
}

var _ directbase.Model = &computeImageModel{}

type computeImageModel struct {
	config *config.ControllerConfig
}

func (m *computeImageModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeImage{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	imagesClient, err := gcpClient.newImagesClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	desired, err := desiredComputeImage(obj)
	if err != nil {
		return nil, err
	}

	return &ComputeImageAdapter{
		gcpClient: imagesClient,
		id:        id.(*krm.ComputeImageIdentity),
		desired:   desired,
		reader:    reader,
	}, nil
}

// desiredComputeImage builds the image that obj asks for. The references in obj must already be
// resolved.
func desiredComputeImage(obj *krm.ComputeImage) (*computepb.Image, error) {
	mapCtx := &direct.MapContext{}
	desired := ComputeImageSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// GCPLabels adds managed-by-cnrm and drops the Kubernetes labels whose keys contain a "/",
	// which GCE does not allow.
	desired.Labels = label.GCPLabels(obj)
	return desired, nil
}

func (m *computeImageModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeImageIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	imagesClient, err := gcpClient.newImagesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeImageAdapter{
		gcpClient: imagesClient,
		id:        id,
	}, nil
}

type ComputeImageAdapter struct {
	gcpClient *compute.ImagesClient
	id        *krm.ComputeImageIdentity
	desired   *computepb.Image
	actual    *computepb.Image
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeImageAdapter{}

func (a *ComputeImageAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeImage", "name", a.id)

	req := &computepb.GetImageRequest{
		Project: a.id.Project,
		Image:   a.id.Image,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeImage %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeImageAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeImage", "name", a.id)
	mapCtx := &direct.MapContext{}

	a.desired.Name = direct.LazyPtr(a.id.Image)

	req := &computepb.InsertImageRequest{
		Project:       a.id.Project,
		ImageResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeImage %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeImage %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeImage in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeImage %s: %w", a.id, err)
	}

	status := ComputeImageStatus_v1beta1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeImageAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeImage", "name", a.id)
	mapCtx := &direct.MapContext{}

	plan, err := a.planUpdate(ctx)
	if err != nil {
		return err
	}

	updated := a.actual
	if plan.setLabels == nil {
		log.V(2).Info("no field needs update", "name", a.id.String())
	} else {
		plan.diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, plan.diffs)

		log.Info("updating labels", "name", a.id)
		req := &computepb.SetLabelsImageRequest{
			Project:                        a.id.Project,
			Resource:                       a.id.Image,
			GlobalSetLabelsRequestResource: plan.setLabels,
		}
		op, err := a.gcpClient.SetLabels(ctx, req)
		if err != nil {
			return fmt.Errorf("setting labels: %w", err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for labels update: %w", err)
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeImage %s: %w", a.id, err)
		}
	}

	status := ComputeImageStatus_v1beta1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// computeImageUpdate is what Update has to send to GCE to bring a.actual in line with a.desired.
type computeImageUpdate struct {
	// diffs lists the changed fields, for structured reporting.
	diffs *structuredreporting.Diff

	// setLabels is the images.setLabels body, or nil when the labels are already right.
	setLabels *computepb.GlobalSetLabelsRequest
}

// planUpdate works out how to bring a.actual in line with a.desired.
// Every ComputeImageSpec field is immutable: the CRD marks them all Immutable, and the Terraform
// controller recreates the image when one changes. So a spec difference returns an
// ImmutableFieldsMutationError instead of an API call. Only the labels can change, and only with
// images.setLabels: images.patch rejects any body that contains labels, even unchanged ones
// ("Updating 'labels' is not supported").
func (a *ComputeImageAdapter) planUpdate(ctx context.Context) (*computeImageUpdate, error) {
	desired, actual, err := a.normalizedSpecs()
	if err != nil {
		return nil, err
	}

	diffs, _, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), actual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	if diffs.HasDiff() {
		return nil, immutableFieldsError(diffs, desired, actual)
	}

	plan := &computeImageUpdate{diffs: diffs}
	// a.desired.Labels comes from label.GCPLabels (see desiredComputeImage).
	if !maps.Equal(a.desired.GetLabels(), a.actual.GetLabels()) {
		diffs.AddField("labels", a.actual.GetLabels(), a.desired.GetLabels())
		plan.setLabels = &computepb.GlobalSetLabelsRequest{
			LabelFingerprint: a.actual.LabelFingerprint,
			Labels:           a.desired.GetLabels(),
		}
	}
	return plan, nil
}

// normalizedSpecs returns the spec fields of a.desired and a.actual, normalized so that they can be
// compared field by field (see normalizeComputeImageForCompare).
//
// It does not use common.CompareBrownfieldSpec: that helper normalizes each side on its own, but
// three ComputeImage rules need both sides at once. Inherited guestOsFeatures and licenses are a
// superset of what was asked for; a sourceImage family is resolved to an image; and rawDisk can
// only be skipped when both sides have one.
func (a *ComputeImageAdapter) normalizedSpecs() (*computepb.Image, *computepb.Image, error) {
	maskedActual, err := mappers.OnlySpecFields(a.actual, ComputeImageSpec_v1beta1_FromProto, ComputeImageSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}

	desired, actual := normalizeComputeImageForCompare(a.desired, maskedActual)
	return desired, actual, nil
}

// normalizeComputeImageForCompare returns copies of desired and actual that can be compared field
// by field. GCE returns several fields differently from how they were sent:
//   - sourceDisk, sourceImage, sourceSnapshot and licenses come back as full URLs.
//   - A sourceImage family comes back as the image the family pointed to when the image was created.
//   - rawDisk.source is input-only and comes back empty.
//   - imageEncryptionKey.kmsKeyName comes back with a /cryptoKeyVersions/N suffix.
//   - guestOsFeatures and licenses also contain the values inherited from the source.
//   - storageLocations come back in lower case.
//
// Fields that desired leaves unset are not compared, because GCE fills them in (for example
// diskSizeGb and storageLocations). A change to the contents of rawDisk, or to the family that an
// image was built from, cannot be detected here; the immutable-fields webhook blocks those edits.
func normalizeComputeImageForCompare(desired, actual *computepb.Image) (*computepb.Image, *computepb.Image) {
	d := proto.CloneOf(desired)
	a := proto.CloneOf(actual)

	// Labels are reconciled with images.setLabels, and the name is the identity.
	d.Labels, a.Labels = nil, nil
	d.Name, a.Name = nil, nil

	d.SourceDisk, a.SourceDisk = trimComputeURL(d.SourceDisk), trimComputeURL(a.SourceDisk)
	d.SourceSnapshot, a.SourceSnapshot = trimComputeURL(d.SourceSnapshot), trimComputeURL(a.SourceSnapshot)
	if isComputeImageFamily(d.GetSourceImage()) && a.SourceImage != nil {
		d.SourceImage, a.SourceImage = nil, nil
	} else {
		d.SourceImage, a.SourceImage = trimComputeURL(d.SourceImage), trimComputeURL(a.SourceImage)
	}

	if d.RawDisk != nil && a.RawDisk != nil {
		d.RawDisk, a.RawDisk = nil, nil
	}

	if d.ImageEncryptionKey != nil && a.ImageEncryptionKey != nil {
		d.ImageEncryptionKey.KmsKeyName = trimKMSKeyVersion(d.ImageEncryptionKey.KmsKeyName)
		a.ImageEncryptionKey.KmsKeyName = trimKMSKeyVersion(a.ImageEncryptionKey.KmsKeyName)
		if d.ImageEncryptionKey.KmsKeyServiceAccount == nil {
			a.ImageEncryptionKey.KmsKeyServiceAccount = nil
		}
	}

	d.StorageLocations, a.StorageLocations = lowerAndSort(d.StorageLocations), lowerAndSort(a.StorageLocations)

	d.Licenses, a.Licenses = trimComputeURLs(d.Licenses), trimComputeURLs(a.Licenses)
	if containsAll(a.Licenses, d.Licenses) {
		a.Licenses = d.Licenses
	}
	if containsAllGuestOSFeatures(a.GuestOsFeatures, d.GuestOsFeatures) {
		a.GuestOsFeatures = d.GuestOsFeatures
	}

	desiredMsg, actualMsg := d.ProtoReflect(), a.ProtoReflect()
	fields := actualMsg.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		if field := fields.Get(i); !desiredMsg.Has(field) {
			actualMsg.Clear(field)
		}
	}

	return d, a
}

// trimComputeURL turns a full Compute Engine URL into a projects/... path.
func trimComputeURL(ref *string) *string {
	if ref == nil {
		return nil
	}
	trimmed := apirefs.TrimComputeURIPrefix(*ref)
	return &trimmed
}

// trimComputeURLs turns full Compute Engine URLs into projects/... paths.
func trimComputeURLs(refs []string) []string {
	trimmed := make([]string, 0, len(refs))
	for _, ref := range refs {
		trimmed = append(trimmed, apirefs.TrimComputeURIPrefix(ref))
	}
	return trimmed
}

// isComputeImageFamily reports whether ref points to an image family rather than to an image.
func isComputeImageFamily(ref string) bool {
	id := &krm.ComputeImageIdentity{}
	return id.FromExternal(ref) == nil && id.IsFamily()
}

// trimKMSKeyVersion drops the /cryptoKeyVersions/N suffix from a KMS key name.
func trimKMSKeyVersion(name *string) *string {
	if name == nil {
		return nil
	}
	key, _, _ := strings.Cut(*name, "/cryptoKeyVersions/")
	return &key
}

// lowerAndSort returns values in lower case and sorted.
func lowerAndSort(values []string) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, strings.ToLower(value))
	}
	slices.Sort(out)
	return out
}

// containsAll reports whether every value in want is also in got.
func containsAll(got, want []string) bool {
	for _, value := range want {
		if !slices.Contains(got, value) {
			return false
		}
	}
	return true
}

// containsAllGuestOSFeatures reports whether every feature type in want is also in got.
func containsAllGuestOSFeatures(got, want []*computepb.GuestOsFeature) bool {
	types := make(map[string]bool, len(got))
	for _, feature := range got {
		types[feature.GetType()] = true
	}
	for _, feature := range want {
		if !types[feature.GetType()] {
			return false
		}
	}
	return true
}

// immutableFieldsError reports each changed spec field by its ComputeImage path, with the live and
// desired values after normalization, so that a user can see what differs.
func immutableFieldsError(diffs *structuredreporting.Diff, desired, actual *computepb.Image) error {
	fields := make([]string, 0, len(diffs.Fields))
	for _, field := range diffs.Fields {
		fd := field.ProtoFieldDescriptor
		if fd == nil {
			fields = append(fields, field.ID)
			continue
		}
		fields = append(fields, fmt.Sprintf("Field Name: %s, Live: %s, Desired: %s",
			computeImageSpecPath(fd), imageFieldValue(actual, fd), imageFieldValue(desired, fd)))
	}
	return k8s.NewImmutableFieldsMutationError(fields)
}

// computeImageSpecPath returns the ComputeImage field that sets the Image field fd. The three
// source fields are references in the KRM spec; all other fields share their JSON name.
func computeImageSpecPath(fd protoreflect.FieldDescriptor) string {
	switch fd.Name() {
	case "source_disk":
		return "spec.diskRef"
	case "source_image":
		return "spec.sourceImageRef"
	case "source_snapshot":
		return "spec.sourceSnapshotRef"
	}
	return "spec." + fd.JSONName()
}

// imageFieldValue returns the value of the field fd of img as compact JSON, or <unset>.
func imageFieldValue(img *computepb.Image, fd protoreflect.FieldDescriptor) string {
	if !img.ProtoReflect().Has(fd) {
		return "<unset>"
	}
	single := &computepb.Image{}
	single.ProtoReflect().Set(fd, img.ProtoReflect().Get(fd))
	// protojson output is deliberately unstable, so pull the value out and compact it.
	b, err := protojson.Marshal(single)
	if err != nil {
		return fmt.Sprintf("%v", img.ProtoReflect().Get(fd).Interface())
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return string(b)
	}
	var out bytes.Buffer
	if err := json.Compact(&out, obj[fd.JSONName()]); err != nil {
		return string(obj[fd.JSONName()])
	}
	return out.String()
}

func (a *ComputeImageAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeImage{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeImageSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Image)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Image)
	u.SetGroupVersionKind(krm.ComputeImageGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)
	return u, nil
}

func (a *ComputeImageAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeImage", "name", a.id)

	req := &computepb.DeleteImageRequest{
		Project: a.id.Project,
		Image:   a.id.Image,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeImage %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeImage", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute ComputeImage %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeImageAdapter) get(ctx context.Context) (*computepb.Image, error) {
	getReq := &computepb.GetImageRequest{
		Project: a.id.Project,
		Image:   a.id.Image,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeImage %s: %w", a.id, err)
	}
	return resource, nil
}
