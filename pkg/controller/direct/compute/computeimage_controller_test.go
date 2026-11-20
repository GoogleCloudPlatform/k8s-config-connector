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

package compute

import (
	"context"
	"errors"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestComputeImagePlanUpdate(t *testing.T) {
	tests := []struct {
		name      string
		spec      *krm.ComputeImageSpec
		k8sLabels map[string]string
		// actual is what images.get returns. Most cases start from imageFromDisk.
		actual *computepb.Image
		// wantSetLabels is the images.setLabels body, or nil when images.setLabels must not be called.
		wantSetLabels *computepb.GlobalSetLabelsRequest
		// wantDiffFields are the fields reported to structured reporting.
		wantDiffFields []string
	}{
		{
			name:           "unchanged image from a disk needs no calls",
			spec:           specFromDisk(),
			k8sLabels:      map[string]string{"cnrm-test": "true"},
			actual:         imageFromDisk(),
			wantDiffFields: []string{},
		},
		{
			name: "changed labels use setLabels only, keep managed-by-cnrm and drop keys with a slash",
			spec: specFromDisk(),
			k8sLabels: map[string]string{
				"cnrm-test":                    "true",
				"label-one":                    "value-one",
				"app.kubernetes.io/managed-by": "configmanagement.gke.io",
			},
			actual: imageFromDisk(),
			wantSetLabels: &computepb.GlobalSetLabelsRequest{
				LabelFingerprint: proto.String("abcdef0123A="),
				Labels: map[string]string{
					"cnrm-test":       "true",
					"label-one":       "value-one",
					"managed-by-cnrm": "true",
				},
			},
			wantDiffFields: []string{"labels"},
		},
		{
			name:      "a missing managed-by-cnrm label is restored",
			spec:      specFromDisk(),
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.Labels = map[string]string{"cnrm-test": "true"}
			}),
			wantSetLabels: &computepb.GlobalSetLabelsRequest{
				LabelFingerprint: proto.String("abcdef0123A="),
				Labels: map[string]string{
					"cnrm-test":       "true",
					"managed-by-cnrm": "true",
				},
			},
			wantDiffFields: []string{"labels"},
		},
		{
			name: "an image family resolved to an image is not a change",
			spec: &krm.ComputeImageSpec{
				SourceImageRef: &krm.ComputeImageRef{External: "projects/debian-cloud/global/images/family/debian-12"},
			},
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.SourceDisk = nil
				img.SourceImage = proto.String("https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-12-bookworm-v20261023")
			}),
			wantDiffFields: []string{},
		},
		{
			name: "a snapshot returned as a full URL is not a change",
			spec: &krm.ComputeImageSpec{
				SourceSnapshotRef: &krm.ComputeSnapshotRef{External: "projects/p/global/snapshots/s"},
			},
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.SourceDisk = nil
				img.SourceSnapshot = proto.String("https://www.googleapis.com/compute/v1/projects/p/global/snapshots/s")
			}),
			wantDiffFields: []string{},
		},
		{
			name: "the input-only rawDisk.source is not a change",
			spec: &krm.ComputeImageSpec{
				RawDisk: &krm.ImageRawDisk{Source: "https://storage.googleapis.com/bucket/disk.tar.gz"},
			},
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.SourceDisk = nil
				img.RawDisk = &computepb.RawDisk{ContainerType: proto.String("TAR"), Source: proto.String("")}
			}),
			wantDiffFields: []string{},
		},
		{
			name: "the KMS key version suffix is not a change",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.ImageEncryptionKey = &krm.ImageImageEncryptionKey{
					KmsKeySelfLinkRef: &kmsv1beta1.KMSCryptoKeyRef{External: "projects/p/locations/us/keyRings/r/cryptoKeys/k"},
				}
			}),
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.ImageEncryptionKey = &computepb.CustomerEncryptionKey{
					KmsKeyName:           proto.String("projects/p/locations/us/keyRings/r/cryptoKeys/k/cryptoKeyVersions/3"),
					KmsKeyServiceAccount: proto.String("sa@p.iam.gserviceaccount.com"),
				}
			}),
			wantDiffFields: []string{},
		},
		{
			name: "a KMS service account that matches is not a change",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.ImageEncryptionKey = &krm.ImageImageEncryptionKey{
					KmsKeySelfLinkRef:       &kmsv1beta1.KMSCryptoKeyRef{External: "projects/p/locations/us/keyRings/r/cryptoKeys/k"},
					KmsKeyServiceAccountRef: &refsv1beta1.IAMServiceAccountRef{External: "sa@p.iam.gserviceaccount.com"},
				}
			}),
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.ImageEncryptionKey = &computepb.CustomerEncryptionKey{
					KmsKeyName:           proto.String("projects/p/locations/us/keyRings/r/cryptoKeys/k/cryptoKeyVersions/3"),
					KmsKeyServiceAccount: proto.String("sa@p.iam.gserviceaccount.com"),
				}
			}),
			wantDiffFields: []string{},
		},
		{
			name: "inherited guestOsFeatures and licenses, full license URLs and upper-case storageLocations are not changes",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.GuestOsFeatures = []krm.ImageGuestOsFeatures{{Type: "UEFI_COMPATIBLE"}}
				spec.Licenses = []string{"projects/debian-cloud/global/licenses/debian-12-bookworm"}
				spec.StorageLocations = []string{"US"}
			}),
			k8sLabels: map[string]string{"cnrm-test": "true"},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.GuestOsFeatures = []*computepb.GuestOsFeature{
					{Type: proto.String("VIRTIO_SCSI_MULTIQUEUE")},
					{Type: proto.String("UEFI_COMPATIBLE")},
					{Type: proto.String("GVNIC")},
				}
				img.Licenses = []string{
					"https://www.googleapis.com/compute/v1/projects/p/global/licenses/inherited",
					"https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-12-bookworm",
				}
			}),
			wantDiffFields: []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := newTestComputeImageAdapter(t, tc.spec, tc.k8sLabels, tc.actual)
			desiredBefore := proto.CloneOf(a.desired)
			actualBefore := proto.CloneOf(a.actual)

			plan, err := a.planUpdate(context.Background())

			if err != nil {
				t.Fatalf("planUpdate() failed: %v", err)
			}
			if !proto.Equal(a.desired, desiredBefore) || !proto.Equal(a.actual, actualBefore) {
				t.Errorf("planUpdate() modified a.desired or a.actual")
			}
			if diff := cmp.Diff(tc.wantSetLabels, plan.setLabels, protocmp.Transform()); diff != "" {
				t.Errorf("planUpdate() setLabels mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tc.wantDiffFields, plan.diffs.FieldIDs()); diff != "" {
				t.Errorf("planUpdate() diff fields mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestComputeImagePlanUpdateImmutableFields(t *testing.T) {
	tests := []struct {
		name   string
		spec   *krm.ComputeImageSpec
		actual *computepb.Image
		// wantFields are the entries of the ImmutableFieldsMutationError.
		wantFields []string
	}{
		{
			name: "a different description",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.Description = direct.LazyPtr("new description")
			}),
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.Description = proto.String("old description")
			}),
			wantFields: []string{`Field Name: spec.description, Live: "old description", Desired: "new description"`},
		},
		{
			name: "a different family",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.Family = direct.LazyPtr("new-family")
			}),
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.Family = proto.String("old-family")
			}),
			wantFields: []string{`Field Name: spec.family, Live: "old-family", Desired: "new-family"`},
		},
		{
			name: "a different diskSizeGb",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.DiskSizeGb = direct.LazyPtr(int64(20))
			}),
			actual:     imageFromDisk(),
			wantFields: []string{`Field Name: spec.diskSizeGb, Live: "10", Desired: "20"`},
		},
		{
			name: "a different source disk",
			spec: &krm.ComputeImageSpec{
				DiskRef: &krm.ComputeDiskRef{External: "projects/p/zones/us-central1-a/disks/other"},
			},
			actual:     imageFromDisk(),
			wantFields: []string{`Field Name: spec.diskRef, Live: "projects/p/zones/us-central1-a/disks/d", Desired: "projects/p/zones/us-central1-a/disks/other"`},
		},
		{
			name: "a different source image",
			spec: &krm.ComputeImageSpec{
				SourceImageRef: &krm.ComputeImageRef{External: "projects/p/global/images/other"},
			},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.SourceDisk = nil
				img.SourceImage = proto.String("https://www.googleapis.com/compute/v1/projects/p/global/images/img")
			}),
			wantFields: []string{`Field Name: spec.sourceImageRef, Live: "projects/p/global/images/img", Desired: "projects/p/global/images/other"`},
		},
		{
			name: "a different source snapshot",
			spec: &krm.ComputeImageSpec{
				SourceSnapshotRef: &krm.ComputeSnapshotRef{External: "projects/p/global/snapshots/other"},
			},
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.SourceDisk = nil
				img.SourceSnapshot = proto.String("https://www.googleapis.com/compute/v1/projects/p/global/snapshots/s")
			}),
			wantFields: []string{`Field Name: spec.sourceSnapshotRef, Live: "projects/p/global/snapshots/s", Desired: "projects/p/global/snapshots/other"`},
		},
		{
			name: "switching the source from a disk to rawDisk",
			spec: &krm.ComputeImageSpec{
				RawDisk: &krm.ImageRawDisk{Source: "https://storage.googleapis.com/bucket/disk.tar.gz"},
			},
			actual:     imageFromDisk(),
			wantFields: []string{`Field Name: spec.rawDisk, Live: <unset>, Desired: {"source":"https://storage.googleapis.com/bucket/disk.tar.gz"}`},
		},
		{
			name: "switching the source from a disk to an image family",
			spec: &krm.ComputeImageSpec{
				SourceImageRef: &krm.ComputeImageRef{External: "projects/debian-cloud/global/images/family/debian-12"},
			},
			actual:     imageFromDisk(),
			wantFields: []string{`Field Name: spec.sourceImageRef, Live: <unset>, Desired: "projects/debian-cloud/global/images/family/debian-12"`},
		},
		{
			name: "a different KMS key",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.ImageEncryptionKey = &krm.ImageImageEncryptionKey{
					KmsKeySelfLinkRef: &kmsv1beta1.KMSCryptoKeyRef{External: "projects/p/locations/us/keyRings/r/cryptoKeys/other"},
				}
			}),
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.ImageEncryptionKey = &computepb.CustomerEncryptionKey{
					KmsKeyName: proto.String("projects/p/locations/us/keyRings/r/cryptoKeys/k/cryptoKeyVersions/3"),
				}
			}),
			wantFields: []string{`Field Name: spec.imageEncryptionKey, Live: {"kmsKeyName":"projects/p/locations/us/keyRings/r/cryptoKeys/k"}, Desired: {"kmsKeyName":"projects/p/locations/us/keyRings/r/cryptoKeys/other"}`},
		},
		{
			name: "a different KMS service account",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.ImageEncryptionKey = &krm.ImageImageEncryptionKey{
					KmsKeySelfLinkRef:       &kmsv1beta1.KMSCryptoKeyRef{External: "projects/p/locations/us/keyRings/r/cryptoKeys/k"},
					KmsKeyServiceAccountRef: &refsv1beta1.IAMServiceAccountRef{External: "other@p.iam.gserviceaccount.com"},
				}
			}),
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.ImageEncryptionKey = &computepb.CustomerEncryptionKey{
					KmsKeyName:           proto.String("projects/p/locations/us/keyRings/r/cryptoKeys/k/cryptoKeyVersions/3"),
					KmsKeyServiceAccount: proto.String("sa@p.iam.gserviceaccount.com"),
				}
			}),
			wantFields: []string{`Field Name: spec.imageEncryptionKey, Live: {"kmsKeyName":"projects/p/locations/us/keyRings/r/cryptoKeys/k","kmsKeyServiceAccount":"sa@p.iam.gserviceaccount.com"}, Desired: {"kmsKeyName":"projects/p/locations/us/keyRings/r/cryptoKeys/k","kmsKeyServiceAccount":"other@p.iam.gserviceaccount.com"}`},
		},
		{
			name: "a license that the image lacks",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.Licenses = []string{"projects/p/global/licenses/missing"}
			}),
			actual:     imageFromDisk(),
			wantFields: []string{`Field Name: spec.licenses, Live: ["projects/debian-cloud/global/licenses/debian-12-bookworm"], Desired: ["projects/p/global/licenses/missing"]`},
		},
		{
			name: "a requested guestOsFeature that the image lacks",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.GuestOsFeatures = []krm.ImageGuestOsFeatures{{Type: "SEV_CAPABLE"}}
			}),
			actual:     imageFromDisk(),
			wantFields: []string{`Field Name: spec.guestOsFeatures, Live: [{"type":"VIRTIO_SCSI_MULTIQUEUE"},{"type":"GVNIC"}], Desired: [{"type":"SEV_CAPABLE"}]`},
		},
		{
			name: "several changed fields are all reported",
			spec: withComputeImageSpec(specFromDisk(), func(spec *krm.ComputeImageSpec) {
				spec.Description = direct.LazyPtr("new description")
				spec.DiskSizeGb = direct.LazyPtr(int64(20))
			}),
			actual: withComputeImage(imageFromDisk(), func(img *computepb.Image) {
				img.Description = proto.String("old description")
			}),
			wantFields: []string{
				`Field Name: spec.description, Live: "old description", Desired: "new description"`,
				`Field Name: spec.diskSizeGb, Live: "10", Desired: "20"`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := newTestComputeImageAdapter(t, tc.spec, map[string]string{"cnrm-test": "true"}, tc.actual)
			want := k8s.NewImmutableFieldsMutationError(tc.wantFields)

			_, err := a.planUpdate(context.Background())

			var got *k8s.ImmutableFieldsMutationError
			if !errors.As(err, &got) {
				t.Fatalf("planUpdate() error = %v, want an ImmutableFieldsMutationError", err)
			}
			if diff := cmp.Diff(want.Error(), got.Error()); diff != "" {
				t.Errorf("planUpdate() error mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// newTestComputeImageAdapter builds an adapter the way AdapterForObject and Find do, from a spec
// whose references are already resolved, the object's Kubernetes labels, and the image GCE returns.
func newTestComputeImageAdapter(t *testing.T, spec *krm.ComputeImageSpec, k8sLabels map[string]string, actual *computepb.Image) *ComputeImageAdapter {
	t.Helper()
	obj := &krm.ComputeImage{
		ObjectMeta: metav1.ObjectMeta{Labels: k8sLabels},
		Spec:       *spec,
	}
	desired, err := desiredComputeImage(obj)
	if err != nil {
		t.Fatalf("desiredComputeImage() failed: %v", err)
	}
	return &ComputeImageAdapter{desired: desired, actual: actual}
}

// specFromDisk is the spec of an image created from a disk, after its diskRef was resolved.
func specFromDisk() *krm.ComputeImageSpec {
	return &krm.ComputeImageSpec{
		DiskRef: &krm.ComputeDiskRef{External: "projects/p/zones/us-central1-a/disks/d"},
	}
}

// imageFromDisk is what images.get returns for the image created from specFromDisk, including the
// fields that GCE fills in.
func imageFromDisk() *computepb.Image {
	return &computepb.Image{
		Name:             proto.String("img"),
		SourceDisk:       proto.String("https://www.googleapis.com/compute/v1/projects/p/zones/us-central1-a/disks/d"),
		SourceDiskId:     proto.String("1234567890"),
		SourceType:       proto.String("RAW"),
		Status:           proto.String("READY"),
		DiskSizeGb:       proto.Int64(10),
		ArchiveSizeBytes: proto.Int64(1234567),
		StorageLocations: []string{"us"},
		GuestOsFeatures: []*computepb.GuestOsFeature{
			{Type: proto.String("VIRTIO_SCSI_MULTIQUEUE")},
			{Type: proto.String("GVNIC")},
		},
		Licenses:         []string{"https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-12-bookworm"},
		Labels:           map[string]string{"cnrm-test": "true", "managed-by-cnrm": "true"},
		LabelFingerprint: proto.String("abcdef0123A="),
		SelfLink:         proto.String("https://www.googleapis.com/compute/v1/projects/p/global/images/img"),
	}
}

func withComputeImageSpec(spec *krm.ComputeImageSpec, edit func(*krm.ComputeImageSpec)) *krm.ComputeImageSpec {
	edit(spec)
	return spec
}

func withComputeImage(img *computepb.Image, edit func(*computepb.Image)) *computepb.Image {
	edit(img)
	return img
}
