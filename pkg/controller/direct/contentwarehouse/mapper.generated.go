// Copyright 2025 Google LLC
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

package contentwarehouse

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contentwarehouse/apiv1/contentwarehousepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContentwarehouseSynonymSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SynonymSet) *krm.ContentwarehouseSynonymSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseSynonymSetObservedState{}
	// MISSING: Name
	// MISSING: Context
	// MISSING: Synonyms
	return out
}
func ContentwarehouseSynonymSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseSynonymSetObservedState) *pb.SynonymSet {
	if in == nil {
		return nil
	}
	out := &pb.SynonymSet{}
	// MISSING: Name
	// MISSING: Context
	// MISSING: Synonyms
	return out
}
func ContentwarehouseSynonymSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.SynonymSet) *krm.ContentwarehouseSynonymSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseSynonymSetSpec{}
	// MISSING: Name
	// MISSING: Context
	// MISSING: Synonyms
	return out
}
func ContentwarehouseSynonymSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseSynonymSetSpec) *pb.SynonymSet {
	if in == nil {
		return nil
	}
	out := &pb.SynonymSet{}
	// MISSING: Name
	// MISSING: Context
	// MISSING: Synonyms
	return out
}
func SynonymSet_FromProto(mapCtx *direct.MapContext, in *pb.SynonymSet) *krm.SynonymSet {
	if in == nil {
		return nil
	}
	out := &krm.SynonymSet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Context = direct.LazyPtr(in.GetContext())
	out.Synonyms = direct.Slice_FromProto(mapCtx, in.Synonyms, SynonymSet_Synonym_FromProto)
	return out
}
func SynonymSet_ToProto(mapCtx *direct.MapContext, in *krm.SynonymSet) *pb.SynonymSet {
	if in == nil {
		return nil
	}
	out := &pb.SynonymSet{}
	out.Name = direct.ValueOf(in.Name)
	out.Context = direct.ValueOf(in.Context)
	out.Synonyms = direct.Slice_ToProto(mapCtx, in.Synonyms, SynonymSet_Synonym_ToProto)
	return out
}
func SynonymSet_Synonym_FromProto(mapCtx *direct.MapContext, in *pb.SynonymSet_Synonym) *krm.SynonymSet_Synonym {
	if in == nil {
		return nil
	}
	out := &krm.SynonymSet_Synonym{}
	out.Words = in.Words
	return out
}
func SynonymSet_Synonym_ToProto(mapCtx *direct.MapContext, in *krm.SynonymSet_Synonym) *pb.SynonymSet_Synonym {
	if in == nil {
		return nil
	}
	out := &pb.SynonymSet_Synonym{}
	out.Words = in.Words
	return out
}
