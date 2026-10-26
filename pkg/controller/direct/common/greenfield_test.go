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

package common

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	videostitcherpb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/structpb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCompareGreenfieldSpec_PreservesServerDefaultedSiblingSubfields(t *testing.T) {
	desiredKRM := &sampleKRMSpec{
		StrField: stringPtr("my-resource"),
		SubStruct: &sampleKRMSub{
			StrVal: stringPtr("updated-service-name"),
			// BoolVal left nil by user; defaulted to true by GCP on actualProto
		},
	}
	actualProto := &descriptorpb.FileDescriptorProto{
		Name: stringPtr("my-resource"),
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name: stringPtr("old-service-name"),
				Options: &descriptorpb.ServiceOptions{
					Deprecated: boolPtr(true),
				},
			},
		},
	}

	mergedDesired, diff, fieldMask, err := CompareGreenfieldSpec(
		t.Context(),
		desiredKRM,
		actualProto,
		greenfieldSampleSpecFromProto,
		greenfieldSampleSpecToProto,
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff == nil || !diff.HasDiff() {
		t.Fatalf("expected diff on service, got none")
	}
	if got := sets.New(fieldMask.GetPaths()...); !got.Equal(sets.New("service")) {
		t.Errorf("fieldMask paths = %v, want [service]", got)
	}
	if len(mergedDesired.GetService()) != 1 {
		t.Fatalf("expected 1 service in mergedDesired, got %d", len(mergedDesired.GetService()))
	}
	if got := mergedDesired.GetService()[0].GetName(); got != "updated-service-name" {
		t.Errorf("mergedDesired.Service[0].Name = %q, want %q", got, "updated-service-name")
	}
	if !mergedDesired.GetService()[0].GetOptions().GetDeprecated() {
		t.Errorf("mergedDesired.Service[0].Options.Deprecated = false, want true (server-defaulted sibling field was lost)")
	}
}

func TestCompareGreenfieldSpec_DetectsRepeatedFieldReordering(t *testing.T) {
	desiredKRM := &sampleKRMSpec{
		StrField:    stringPtr("my-resource"),
		StringSlice: []string{"Zebra", "Alpha"},
	}
	actualProto := &descriptorpb.FileDescriptorProto{
		Name: stringPtr("my-resource"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: stringPtr("Alpha")},
			{Name: stringPtr("Zebra")},
		},
	}

	_, diff, fieldMask, err := CompareGreenfieldSpec(
		t.Context(),
		desiredKRM,
		actualProto,
		greenfieldSampleSpecFromProto,
		greenfieldSampleSpecToProto,
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff == nil || !diff.HasDiff() {
		t.Fatalf("expected diff when repeated field order changes, got no diff")
	}
	if got := sets.New(fieldMask.GetPaths()...); !got.Equal(sets.New("message_type")) {
		t.Errorf("fieldMask paths = %v, want [message_type]", got)
	}
}

func TestCompareGreenfieldSpec_FieldBehaviors_OutputOnlyAndImmutable(t *testing.T) {
	t.Run("ignores OUTPUT_ONLY diffs and adopts unset IMMUTABLE fields", func(t *testing.T) {
		_, diff, fieldMask, err := CompareGreenfieldSpec(
			t.Context(),
			&filestoreTestSpec{Description: stringPtr("same")},
			&pb.Instance{
				Description: "same",
				Protocol:    pb.Instance_NFS_V3,
				State:       pb.Instance_READY,
			},
			filestoreTestSpecFromProto,
			filestoreTestSpecToProto,
			nil,
		)
		if err != nil {
			t.Fatalf("unexpected error when IMMUTABLE field is unset in desired: %v", err)
		}
		if diff.HasDiff() || len(fieldMask.GetPaths()) != 0 {
			t.Errorf("expected no diff when only OUTPUT_ONLY field differs, got paths %v", fieldMask.GetPaths())
		}
	})

	t.Run("returns error when user mutates an IMMUTABLE field", func(t *testing.T) {
		_, _, _, err := CompareGreenfieldSpec(
			t.Context(),
			&filestoreTestSpec{
				Description: stringPtr("same"),
				Protocol:    stringPtr(pb.Instance_NFS_V4_1.String()),
			},
			&pb.Instance{
				Description: "same",
				Protocol:    pb.Instance_NFS_V3,
			},
			filestoreTestSpecFromProto,
			filestoreTestSpecToProto,
			nil,
		)
		if err == nil || !strings.Contains(err.Error(), "change to immutable field protocol") {
			t.Fatalf("expected immutable field error for protocol, got: %v", err)
		}
	})
}

func TestCompareGreenfieldSpec_SkipsNestedInputOnlyFields(t *testing.T) {
	// GCP Get response omits private_key because it is marked [(google.api.field_behavior) = INPUT_ONLY]
	mergedDesired, diff, fieldMask, err := CompareGreenfieldSpec(
		t.Context(),
		&cdnKeyTestSpec{
			Hostname: stringPtr("cdn.example.com"),
			GoogleCdnKey: &googleCdnKeyTestSpec{
				KeyName:    stringPtr("my-key"),
				PrivateKey: stringPtr("top-secret-key"),
			},
		},
		&videostitcherpb.CdnKey{
			Hostname: "cdn.example.com",
			CdnKeyConfig: &videostitcherpb.CdnKey_GoogleCdnKey{
				GoogleCdnKey: &videostitcherpb.GoogleCdnKey{
					KeyName: "my-key",
					// PrivateKey is empty in GCP Get response
				},
			},
		},
		cdnKeyTestSpecFromProto,
		cdnKeyTestSpecToProto,
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff.HasDiff() || len(fieldMask.GetPaths()) != 0 {
		t.Errorf("expected no diff when only nested INPUT_ONLY field private_key differs, got paths %v", fieldMask.GetPaths())
	}
	if got := string(mergedDesired.GetGoogleCdnKey().GetPrivateKey()); got != "top-secret-key" {
		t.Errorf("mergedDesired.GoogleCdnKey.PrivateKey = %q, want %q", got, "top-secret-key")
	}
}

func TestCompareGreenfieldSpec_MapOfStructsAndPtrToStructs(t *testing.T) {
	mergedDesired, diff, fieldMask, err := CompareGreenfieldSpec(
		t.Context(),
		&mapOfStructsTestSpec{
			Description: stringPtr("desc"),
			Versions: map[string]versionDetailsTestSpec{
				"v1": {
					ImageURI: stringPtr("gcr.io/my-image:v1"),
					// CreateTime left nil by user; populated by server in actualProto
				},
			},
			PtrVersions: map[string]*versionDetailsTestSpec{
				"nil-entry": nil,
				"v2": {
					ImageURI: stringPtr("gcr.io/my-image:v2"),
				},
			},
		},
		&pb.Instance{
			Description: "desc",
			Labels: map[string]string{
				"image":      "gcr.io/my-image:v1",
				"createTime": "2026-01-01T00:00:00Z",
				"ptrImage":   "gcr.io/my-image:v2",
				"ptrCreate":  "2026-01-02T00:00:00Z",
				"nilEntry":   "adopted-from-actual",
			},
		},
		mapOfStructsTestSpecFromProto,
		mapOfStructsTestSpecToProto,
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff.HasDiff() || len(fieldMask.GetPaths()) != 0 {
		t.Errorf("expected no diff when map of structs and map of *structs have server defaults, got paths %v", fieldMask.GetPaths())
	}
	if got := mergedDesired.GetLabels()["nilEntry"]; got != "adopted-from-actual" {
		t.Errorf("mergedDesired nilEntry = %q, want %q", got, "adopted-from-actual")
	}
}

func TestMergeUnsetFieldsGreenfield_UnexportedAndNonSettableDoNotPanic(t *testing.T) {
	type specWithUnexported struct {
		Exported   *string
		unexported *string
	}
	desired := &specWithUnexported{}
	actual := &specWithUnexported{
		Exported:   stringPtr("adopted"),
		unexported: stringPtr("secret-internal"),
	}
	mergeUnsetFieldsGreenfield(reflect.ValueOf(desired), reflect.ValueOf(actual))
	if direct.ValueOf(desired.Exported) != "adopted" {
		t.Errorf("Exported = %q, want %q", direct.ValueOf(desired.Exported), "adopted")
	}
	if desired.unexported != nil {
		t.Errorf("expected unexported field to remain nil, got %v", desired.unexported)
	}
	// Also verify non-settable top-level pointer does not panic
	var nilSpec *specWithUnexported
	mergeUnsetFieldsGreenfield(reflect.ValueOf(nilSpec), reflect.ValueOf(actual))
}

func TestDiffForTopLevelFieldsGreenfield_DirectCallAndNilSide(t *testing.T) {
	ctx := t.Context()

	t.Run("direct call skips nested INPUT_ONLY without mutating caller message", func(t *testing.T) {
		desired := &videostitcherpb.CdnKey{
			Hostname: "cdn.example.com",
			CdnKeyConfig: &videostitcherpb.CdnKey_GoogleCdnKey{
				GoogleCdnKey: &videostitcherpb.GoogleCdnKey{
					KeyName:    "my-key",
					PrivateKey: []byte("top-secret-key"),
				},
			},
		}
		actual := &videostitcherpb.CdnKey{
			Hostname: "cdn.example.com",
			CdnKeyConfig: &videostitcherpb.CdnKey_GoogleCdnKey{
				GoogleCdnKey: &videostitcherpb.GoogleCdnKey{
					KeyName: "my-key",
				},
			},
		}
		diff, mask, err := DiffForTopLevelFieldsGreenfield(ctx, desired.ProtoReflect(), actual.ProtoReflect())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff.HasDiff() || len(mask.GetPaths()) != 0 {
			t.Errorf("expected no diff on nested INPUT_ONLY field, got %v", mask.GetPaths())
		}
		if got := string(desired.GetGoogleCdnKey().GetPrivateKey()); got != "top-secret-key" {
			t.Errorf("caller desired message was mutated: got PrivateKey=%q", got)
		}
	})

	t.Run("handles nil actual with non-nil desired without panic and diffs only populated fields", func(t *testing.T) {
		desired := &pb.Instance{
			Description: "only-description-set",
		}
		diff, mask, err := DiffForTopLevelFieldsGreenfield(ctx, desired.ProtoReflect(), nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		wantPaths := []string{"description"}
		if diffPaths := cmp.Diff(wantPaths, mask.GetPaths()); diffPaths != "" {
			t.Errorf("mask.Paths mismatch (-want +got):\n%s", diffPaths)
		}
		if !diff.HasDiff() {
			t.Errorf("expected diff on description")
		}
	})
}

func TestCompareGreenfieldSpec_NilInputs(t *testing.T) {
	mergedDesired, diff, fieldMask, err := CompareGreenfieldSpec(
		t.Context(),
		nil,
		(*descriptorpb.FileDescriptorProto)(nil),
		greenfieldSampleSpecFromProto,
		greenfieldSampleSpecToProto,
		func(ctx context.Context, pb *descriptorpb.FileDescriptorProto) error {
			return nil
		},
	)
	if err != nil {
		t.Fatalf("unexpected error on nil inputs: %v", err)
	}
	if mergedDesired != nil {
		t.Errorf("expected nil mergedDesired, got %v", mergedDesired)
	}
	if diff.HasDiff() || len(fieldMask.GetPaths()) != 0 {
		t.Errorf("expected empty diff and fieldMask on nil inputs, got %v", fieldMask.GetPaths())
	}
}

func TestCompareGreenfieldSpecAndLabels(t *testing.T) {
	type greenfieldLabelTestSpec struct {
		Description string
		Tier        string
	}
	specFromProto := func(mapCtx *direct.MapContext, in *pb.Instance) *greenfieldLabelTestSpec {
		if in == nil {
			return nil
		}
		return &greenfieldLabelTestSpec{Description: in.Description, Tier: in.Tier.String()}
	}

	specToProto := func(mapCtx *direct.MapContext, in *greenfieldLabelTestSpec) *pb.Instance {
		if in == nil {
			return nil
		}
		out := &pb.Instance{Description: in.Description}
		if in.Tier != "" {
			out.Tier = pb.Instance_Tier(pb.Instance_Tier_value[in.Tier])
		}
		return out
	}

	u := &unstructured.Unstructured{}
	u.SetLabels(map[string]string{"env": "staging"})

	mergedDesired, diff, fieldMask, err := CompareGreenfieldSpecAndLabels(
		t.Context(),
		u,
		&greenfieldLabelTestSpec{Description: "new description", Tier: pb.Instance_ENTERPRISE.String()},
		&pb.Instance{
			Description: "old description",
			Tier:        pb.Instance_BASIC_HDD,
			Labels: map[string]string{
				"env":             "prod",
				"managed-by-cnrm": "true",
			},
		},
		"labels",
		specFromProto,
		specToProto,
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !diff.HasDiff() {
		t.Fatalf("expected diff, got none")
	}
	wantPaths := []string{"description", "labels", "tier"}
	if diffPaths := cmp.Diff(wantPaths, fieldMask.GetPaths()); diffPaths != "" {
		t.Errorf("fieldMask.Paths not sorted alphabetically (-want +got):\n%s", diffPaths)
	}
	wantLabels := map[string]string{
		"env":             "staging",
		"managed-by-cnrm": "true",
	}
	if diff := cmp.Diff(wantLabels, mergedDesired.GetLabels()); diff != "" {
		t.Errorf("mergedDesired.Labels mismatch (-want +got):\n%s", diff)
	}
}

func TestCompareGreenfieldSpec_OmittedVsExplicitlyEmptyCollections(t *testing.T) {
	ctx := t.Context()
	actual := &pb.Instance{
		Description: "my-instance",
		FileShares:  []*pb.FileShareConfig{{Name: "default-share"}},
		Labels:      map[string]string{"default-key": "default-val"},
	}

	t.Run("omitted nil slice and map adopt actual and produce zero diff", func(t *testing.T) {
		desired := &collectionsTestSpec{
			Description: direct.LazyPtr("my-instance"),
			ShareNames:  nil,
			Tags:        nil,
		}
		merged, diff, mask, err := CompareGreenfieldSpec(
			ctx, desired, actual, collectionsTestSpecFromProto, collectionsTestSpecToProto, nil,
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff.HasDiff() {
			t.Errorf("expected zero diff for omitted collections, got: %v", diff.Fields)
		}
		if len(mask.GetPaths()) != 0 {
			t.Errorf("expected empty updateMask, got: %v", mask.GetPaths())
		}
		if len(merged.GetFileShares()) != 1 || merged.GetFileShares()[0].GetName() != "default-share" {
			t.Errorf("expected merged FileShares to preserve actual default-share, got: %v", merged.GetFileShares())
		}
	})

	t.Run("explicitly empty slice and map (including nested inside map entry) clear actual and produce diff", func(t *testing.T) {
		desired := &collectionsTestSpec{
			Description: direct.LazyPtr("my-instance"),
			ShareNames:  []string{},
			Tags:        map[string]string{},
			NestedGroups: map[string]nestedCollectionGroup{
				"g1": {Items: []string{}},
			},
		}
		cloned, err := cloneKRMSpec(desired)
		if err != nil {
			t.Fatalf("cloneKRMSpec error: %v", err)
		}
		if cloned.NestedGroups["g1"].Items == nil {
			t.Errorf("expected preserveNonNilEmptyCollections to preserve non-nil empty slice inside map struct entry")
		}
		merged, diff, mask, err := CompareGreenfieldSpec(
			ctx, desired, actual, collectionsTestSpecFromProto, collectionsTestSpecToProto, nil,
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !diff.HasDiff() {
			t.Fatalf("expected diff for explicitly cleared slice and map, got none")
		}
		wantPaths := []string{"file_shares", "labels"}
		if diffPaths := cmp.Diff(wantPaths, mask.GetPaths()); diffPaths != "" {
			t.Errorf("mask.Paths mismatch (-want +got):\n%s", diffPaths)
		}
		if len(merged.GetFileShares()) != 0 {
			t.Errorf("expected merged FileShares to be cleared, got: %v", merged.GetFileShares())
		}
		if len(merged.GetLabels()) != 0 {
			t.Errorf("expected merged Labels to be cleared, got: %v", merged.GetLabels())
		}
	})
}

func TestCompareGreenfieldSpec_PreservesExplicitOneofSelection(t *testing.T) {
	ctx := t.Context()

	t.Run("switching from later-mapped oneof field to earlier-mapped oneof field is not clobbered", func(t *testing.T) {
		actual := &structpb.Value{
			Kind: &structpb.Value_BoolValue{BoolValue: true},
		}
		desired := &oneofValueTestSpec{
			StringValue: direct.LazyPtr("switched-to-string"),
		}
		merged, diff, mask, err := CompareGreenfieldSpec(
			ctx, desired, actual, oneofValueTestSpecFromProto, oneofValueTestSpecToProto, nil,
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !diff.HasDiff() {
			t.Fatalf("expected diff when switching oneof from bool_value to string_value, got none")
		}
		wantPaths := []string{"bool_value", "string_value"}
		if diffPaths := cmp.Diff(wantPaths, mask.GetPaths()); diffPaths != "" {
			t.Errorf("mask.Paths mismatch (-want +got):\n%s", diffPaths)
		}
		if merged.GetStringValue() != "switched-to-string" {
			t.Errorf("expected merged StringValue=%q, got Kind=%T (%v)", "switched-to-string", merged.GetKind(), merged)
		}
	})
}

// --- Test Fixtures & Helper Mappers (Footnotes) ---

func greenfieldSampleSpecFromProto(mapCtx *direct.MapContext, in *descriptorpb.FileDescriptorProto) *sampleKRMSpec {
	if in == nil {
		return nil
	}
	res := &sampleKRMSpec{
		StrField:  in.Name,
		NonPtrStr: in.GetPackage(),
	}
	if len(in.Service) > 0 {
		res.SubStruct = &sampleKRMSub{
			StrVal: in.Service[0].Name,
		}
		if in.Service[0].Options != nil && in.Service[0].Options.Deprecated != nil {
			val := *in.Service[0].Options.Deprecated
			res.SubStruct.BoolVal = &val
		}
	}
	for _, m := range in.MessageType {
		res.StringSlice = append(res.StringSlice, m.GetName())
	}
	return res
}

func greenfieldSampleSpecToProto(mapCtx *direct.MapContext, in *sampleKRMSpec) *descriptorpb.FileDescriptorProto {
	if in == nil {
		return nil
	}
	res := &descriptorpb.FileDescriptorProto{
		Name:    in.StrField,
		Package: stringPtr(in.NonPtrStr),
	}
	if in.SubStruct != nil && (in.SubStruct.StrVal != nil || in.SubStruct.BoolVal != nil) {
		svc := &descriptorpb.ServiceDescriptorProto{
			Name: in.SubStruct.StrVal,
		}
		if in.SubStruct.BoolVal != nil {
			svc.Options = &descriptorpb.ServiceOptions{
				Deprecated: in.SubStruct.BoolVal,
			}
		}
		res.Service = []*descriptorpb.ServiceDescriptorProto{svc}
	}
	for _, s := range in.StringSlice {
		res.MessageType = append(res.MessageType, &descriptorpb.DescriptorProto{
			Name: stringPtr(s),
		})
	}
	return res
}

type filestoreTestSpec struct {
	Description *string `json:"description,omitempty"`
	Protocol    *string `json:"protocol,omitempty"`
}

func filestoreTestSpecFromProto(mapCtx *direct.MapContext, in *pb.Instance) *filestoreTestSpec {
	if in == nil {
		return nil
	}
	protocol := in.GetProtocol().String()
	desc := in.GetDescription()
	return &filestoreTestSpec{
		Description: &desc,
		Protocol:    &protocol,
	}
}

func filestoreTestSpecToProto(mapCtx *direct.MapContext, in *filestoreTestSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{
		Description: direct.ValueOf(in.Description),
		// Simulate an OUTPUT_ONLY field (State) differing between desired and actual
		State: pb.Instance_CREATING,
	}
	if in.Protocol != nil {
		out.Protocol = pb.Instance_FileProtocol(pb.Instance_FileProtocol_value[*in.Protocol])
	}
	return out
}

type googleCdnKeyTestSpec struct {
	KeyName    *string `json:"keyName,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
}

type cdnKeyTestSpec struct {
	Hostname     *string               `json:"hostname,omitempty"`
	GoogleCdnKey *googleCdnKeyTestSpec `json:"googleCdnKey,omitempty"`
}

func cdnKeyTestSpecFromProto(mapCtx *direct.MapContext, in *videostitcherpb.CdnKey) *cdnKeyTestSpec {
	if in == nil {
		return nil
	}
	out := &cdnKeyTestSpec{
		Hostname: direct.LazyPtr(in.GetHostname()),
	}
	if g := in.GetGoogleCdnKey(); g != nil {
		out.GoogleCdnKey = &googleCdnKeyTestSpec{
			KeyName:    direct.LazyPtr(g.GetKeyName()),
			PrivateKey: direct.LazyPtr(string(g.GetPrivateKey())),
		}
	}
	return out
}

func cdnKeyTestSpecToProto(mapCtx *direct.MapContext, in *cdnKeyTestSpec) *videostitcherpb.CdnKey {
	if in == nil {
		return nil
	}
	out := &videostitcherpb.CdnKey{
		Hostname: direct.ValueOf(in.Hostname),
	}
	if in.GoogleCdnKey != nil {
		g := &videostitcherpb.GoogleCdnKey{
			KeyName: direct.ValueOf(in.GoogleCdnKey.KeyName),
		}
		if in.GoogleCdnKey.PrivateKey != nil {
			g.PrivateKey = []byte(*in.GoogleCdnKey.PrivateKey)
		}
		out.CdnKeyConfig = &videostitcherpb.CdnKey_GoogleCdnKey{GoogleCdnKey: g}
	}
	return out
}

type versionDetailsTestSpec struct {
	ImageURI   *string `json:"imageURI,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
}

type mapOfStructsTestSpec struct {
	Description *string                            `json:"description,omitempty"`
	Versions    map[string]versionDetailsTestSpec  `json:"versions,omitempty"`
	PtrVersions map[string]*versionDetailsTestSpec `json:"ptrVersions,omitempty"`
}

func mapOfStructsTestSpecFromProto(mapCtx *direct.MapContext, in *pb.Instance) *mapOfStructsTestSpec {
	if in == nil {
		return nil
	}
	out := &mapOfStructsTestSpec{
		Description: direct.LazyPtr(in.GetDescription()),
	}
	if len(in.GetLabels()) > 0 {
		out.Versions = map[string]versionDetailsTestSpec{
			"v1": {
				ImageURI:   direct.LazyPtr(in.GetLabels()["image"]),
				CreateTime: direct.LazyPtr(in.GetLabels()["createTime"]),
			},
		}
		out.PtrVersions = map[string]*versionDetailsTestSpec{
			"nil-entry": {
				ImageURI: direct.LazyPtr(in.GetLabels()["nilEntry"]),
			},
			"v2": {
				ImageURI:   direct.LazyPtr(in.GetLabels()["ptrImage"]),
				CreateTime: direct.LazyPtr(in.GetLabels()["ptrCreate"]),
			},
		}
	}
	return out
}

func mapOfStructsTestSpecToProto(mapCtx *direct.MapContext, in *mapOfStructsTestSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{
		Description: direct.ValueOf(in.Description),
		Labels:      map[string]string{},
	}
	if v, ok := in.Versions["v1"]; ok {
		out.Labels["image"] = direct.ValueOf(v.ImageURI)
		out.Labels["createTime"] = direct.ValueOf(v.CreateTime)
	}
	if v := in.PtrVersions["v2"]; v != nil {
		out.Labels["ptrImage"] = direct.ValueOf(v.ImageURI)
		out.Labels["ptrCreate"] = direct.ValueOf(v.CreateTime)
	}
	if v := in.PtrVersions["nil-entry"]; v != nil {
		out.Labels["nilEntry"] = direct.ValueOf(v.ImageURI)
	}
	if len(out.Labels) == 0 {
		out.Labels = nil
	}
	return out
}

type nestedCollectionGroup struct {
	Items []string `json:"items,omitempty"`
}

type collectionsTestSpec struct {
	Description  *string                          `json:"description,omitempty"`
	ShareNames   []string                         `json:"shareNames,omitempty"`
	Tags         map[string]string                `json:"tags,omitempty"`
	NestedGroups map[string]nestedCollectionGroup `json:"nestedGroups,omitempty"`
}

func collectionsTestSpecFromProto(mapCtx *direct.MapContext, in *pb.Instance) *collectionsTestSpec {
	if in == nil {
		return nil
	}
	out := &collectionsTestSpec{
		Description: direct.LazyPtr(in.GetDescription()),
	}
	for _, fs := range in.GetFileShares() {
		out.ShareNames = append(out.ShareNames, fs.GetName())
	}
	if len(in.GetLabels()) > 0 {
		out.Tags = make(map[string]string, len(in.GetLabels()))
		for k, v := range in.GetLabels() {
			out.Tags[k] = v
		}
	}
	return out
}

func collectionsTestSpecToProto(mapCtx *direct.MapContext, in *collectionsTestSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{
		Description: direct.ValueOf(in.Description),
	}
	for _, name := range in.ShareNames {
		out.FileShares = append(out.FileShares, &pb.FileShareConfig{Name: name})
	}
	if len(in.Tags) > 0 {
		out.Labels = make(map[string]string, len(in.Tags))
		for k, v := range in.Tags {
			out.Labels[k] = v
		}
	}
	return out
}

type oneofValueTestSpec struct {
	StringValue *string `json:"stringValue,omitempty"`
	BoolValue   *bool   `json:"boolValue,omitempty"`
}

func oneofValueTestSpecFromProto(mapCtx *direct.MapContext, in *structpb.Value) *oneofValueTestSpec {
	if in == nil {
		return nil
	}
	out := &oneofValueTestSpec{}
	switch k := in.GetKind().(type) {
	case *structpb.Value_StringValue:
		out.StringValue = &k.StringValue
	case *structpb.Value_BoolValue:
		out.BoolValue = &k.BoolValue
	}
	return out
}

func oneofValueTestSpecToProto(mapCtx *direct.MapContext, in *oneofValueTestSpec) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{StringValue: *in.StringValue}
	}
	if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{BoolValue: *in.BoolValue}
	}
	return out
}

type mockLegacyExternalNormalizer struct {
	Name      string
	Namespace string
	External  string
	calls     int
}

var _ refs.ExternalNormalizer = &mockLegacyExternalNormalizer{}

func (r *mockLegacyExternalNormalizer) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	r.calls++
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external")
	}
	if r.External == "" && r.Name != "" {
		r.External = "projects/test-proj/locations/us-central1/items/" + r.Name
	}
	return r.External, nil
}

type testLegacyNormObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              struct {
		ParentRef *mockLegacyExternalNormalizer
		EmptyRef  *mockLegacyExternalNormalizer
		SliceRefs []*mockLegacyExternalNormalizer
	}
}

func (t *testLegacyNormObj) DeepCopyObject() runtime.Object {
	return t
}

func TestNormalizeGreenfieldReferences_ExternalNormalizerIdempotent(t *testing.T) {
	ctx := context.Background()
	kube := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build()

	obj := &testLegacyNormObj{}
	// Simulate New*Identity having already normalized ParentRef (setting External while leaving Name non-empty)
	obj.Spec.ParentRef = &mockLegacyExternalNormalizer{
		Name:     "my-parent",
		External: "projects/test-proj/locations/us-central1/items/my-parent",
	}
	obj.Spec.EmptyRef = &mockLegacyExternalNormalizer{}
	obj.Spec.SliceRefs = []*mockLegacyExternalNormalizer{
		{Name: "item-1", Namespace: "custom-ns"},
	}

	if err := NormalizeGreenfieldReferences(ctx, kube, obj, nil); err != nil {
		t.Fatalf("NormalizeGreenfieldReferences failed: %v", err)
	}
	if obj.Spec.ParentRef.calls != 0 {
		t.Errorf("expected already-normalized ParentRef not to be called again, got %d calls", obj.Spec.ParentRef.calls)
	}
	if obj.Spec.EmptyRef.calls != 0 {
		t.Errorf("expected zero-value EmptyRef not to be called, got %d calls", obj.Spec.EmptyRef.calls)
	}
	if obj.Spec.SliceRefs[0].calls != 1 || obj.Spec.SliceRefs[0].External == "" {
		t.Errorf("expected SliceRefs[0] to be normalized once, got calls=%d external=%q", obj.Spec.SliceRefs[0].calls, obj.Spec.SliceRefs[0].External)
	}
}

type testBackupRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	External  string `json:"external,omitempty"`
}

type fileShareTestElem struct {
	Name            *string           `json:"name,omitempty"`
	CapacityGb      *int64            `json:"capacityGb,omitempty"`
	SourceBackup    *string           `json:"sourceBackup,omitempty"`
	SourceBackupRef *testBackupRef    `json:"sourceBackupRef,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
}

type sliceElemTestSpec struct {
	FileShares []fileShareTestElem `json:"fileShares,omitempty"`
}

func sliceElemTestSpecFromProto(mapCtx *direct.MapContext, in *pb.Instance) *sliceElemTestSpec {
	if in == nil {
		return nil
	}
	out := &sliceElemTestSpec{}
	for _, fs := range in.GetFileShares() {
		elem := fileShareTestElem{
			Name:         direct.LazyPtr(fs.GetName()),
			CapacityGb:   direct.LazyPtr(fs.GetCapacityGb()),
			SourceBackup: direct.LazyPtr(fs.GetSourceBackup()),
		}
		if fs.GetSourceBackup() != "" {
			elem.SourceBackupRef = &testBackupRef{External: fs.GetSourceBackup()}
		}
		out.FileShares = append(out.FileShares, elem)
	}
	return out
}

func sliceElemTestSpecToProto(mapCtx *direct.MapContext, in *sliceElemTestSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	for _, fs := range in.FileShares {
		p := &pb.FileShareConfig{
			Name:       direct.ValueOf(fs.Name),
			CapacityGb: direct.ValueOf(fs.CapacityGb),
		}
		if fs.SourceBackupRef != nil && fs.SourceBackupRef.External != "" {
			p.Source = &pb.FileShareConfig_SourceBackup{SourceBackup: fs.SourceBackupRef.External}
		} else if fs.SourceBackup != nil {
			p.Source = &pb.FileShareConfig_SourceBackup{SourceBackup: *fs.SourceBackup}
		}
		out.FileShares = append(out.FileShares, p)
	}
	return out
}

func TestCompareGreenfieldSpec_SliceElementServerDefaults(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		desiredKRM     *sliceElemTestSpec
		actualProto    *pb.Instance
		wantHasDiff    bool
		wantCapacities []int64
	}{
		{
			name: "unmodified slice element adopts server-populated default subfields",
			desiredKRM: &sliceElemTestSpec{
				FileShares: []fileShareTestElem{
					{Name: direct.PtrTo("share-a")},
				},
			},
			actualProto: &pb.Instance{
				FileShares: []*pb.FileShareConfig{
					{Name: "share-a", CapacityGb: 1024},
				},
			},
			wantHasDiff:    false,
			wantCapacities: []int64{1024},
		},
		{
			name: "slice element with KRM Ref set by Name and resolved External adopts server defaults",
			desiredKRM: &sliceElemTestSpec{
				FileShares: []fileShareTestElem{
					{
						Name: direct.PtrTo("share-a"),
						SourceBackupRef: &testBackupRef{
							Name:      "my-backup",
							Namespace: "default",
							External:  "projects/p/locations/l/backups/my-backup",
						},
					},
				},
			},
			actualProto: &pb.Instance{
				FileShares: []*pb.FileShareConfig{
					{
						Name:       "share-a",
						CapacityGb: 1024,
						Source:     &pb.FileShareConfig_SourceBackup{SourceBackup: "projects/p/locations/l/backups/my-backup"},
					},
				},
			},
			wantHasDiff:    false,
			wantCapacities: []int64{1024},
		},
		{
			name: "reordered slice elements do not cross-contaminate defaults between indices",
			desiredKRM: &sliceElemTestSpec{
				FileShares: []fileShareTestElem{
					{Name: direct.PtrTo("share-b")},
					{Name: direct.PtrTo("share-a"), CapacityGb: direct.PtrTo(int64(2048))},
				},
			},
			actualProto: &pb.Instance{
				FileShares: []*pb.FileShareConfig{
					{Name: "share-a", CapacityGb: 2048},
					{Name: "share-b", CapacityGb: 1024},
				},
			},
			wantHasDiff:    true,
			wantCapacities: []int64{0, 2048},
		},
		{
			name: "partial slice modification merges unchanged element while keeping modified element clean",
			desiredKRM: &sliceElemTestSpec{
				FileShares: []fileShareTestElem{
					{Name: direct.PtrTo("share-a")},
					{
						Name:         direct.PtrTo("share-b"),
						SourceBackup: direct.PtrTo("projects/p/locations/l/backups/b2"),
					},
				},
			},
			actualProto: &pb.Instance{
				FileShares: []*pb.FileShareConfig{
					{Name: "share-a", CapacityGb: 1024},
					{
						Name:       "share-b",
						CapacityGb: 2048,
						Source:     &pb.FileShareConfig_SourceBackup{SourceBackup: "projects/p/locations/l/backups/b1"},
					},
				},
			},
			wantHasDiff:    true,
			wantCapacities: []int64{1024, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mergedDesiredProto, diff, fieldMask, err := CompareGreenfieldSpec(
				ctx,
				tc.desiredKRM,
				tc.actualProto,
				sliceElemTestSpecFromProto,
				sliceElemTestSpecToProto,
				nil,
			)
			if err != nil {
				t.Fatalf("CompareGreenfieldSpec failed: %v", err)
			}
			gotHasDiff := diff.HasDiff() || len(fieldMask.GetPaths()) > 0
			if gotHasDiff != tc.wantHasDiff {
				t.Errorf("got hasDiff=%v (paths=%v, diff=%v), want %v", gotHasDiff, fieldMask.GetPaths(), diff, tc.wantHasDiff)
			}
			for i, wantCap := range tc.wantCapacities {
				if gotCap := mergedDesiredProto.GetFileShares()[i].GetCapacityGb(); gotCap != wantCap {
					t.Errorf("mergedDesiredProto.FileShares[%d].CapacityGb = %d, want %d", i, gotCap, wantCap)
				}
			}
		})
	}
}

func TestSliceElemHasNoConflicts_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		d    fileShareTestElem
		a    fileShareTestElem
		want bool
	}{
		{
			name: "map with deleted key reports conflict",
			d:    fileShareTestElem{Labels: map[string]string{"env": "prod"}},
			a:    fileShareTestElem{Labels: map[string]string{"env": "prod", "tier": "backend"}},
			want: false,
		},
		{
			name: "cleared non-nil empty map vs populated map reports conflict",
			d:    fileShareTestElem{Labels: map[string]string{}},
			a:    fileShareTestElem{Labels: map[string]string{"env": "prod"}},
			want: false,
		},
		{
			name: "non-nil empty map vs nil map has no conflict",
			d:    fileShareTestElem{Labels: map[string]string{}},
			a:    fileShareTestElem{Labels: nil},
			want: true,
		},
		{
			name: "KRM Ref pointing to different External reports conflict",
			d: fileShareTestElem{
				SourceBackupRef: &testBackupRef{Name: "b2", External: "projects/p/locations/l/backups/b2"},
			},
			a: fileShareTestElem{
				SourceBackupRef: &testBackupRef{External: "projects/p/locations/l/backups/b1"},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := sliceElemHasNoConflicts(reflect.ValueOf(tc.d), reflect.ValueOf(tc.a)); got != tc.want {
				t.Errorf("sliceElemHasNoConflicts() = %v, want %v", got, tc.want)
			}
		})
	}
}
