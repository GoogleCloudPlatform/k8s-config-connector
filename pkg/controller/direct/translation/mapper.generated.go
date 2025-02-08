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

package translation

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Example_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.Example {
	if in == nil {
		return nil
	}
	out := &krm.Example{}
	// MISSING: Name
	out.SourceText = direct.LazyPtr(in.GetSourceText())
	out.TargetText = direct.LazyPtr(in.GetTargetText())
	// MISSING: Usage
	return out
}
func Example_ToProto(mapCtx *direct.MapContext, in *krm.Example) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	out.SourceText = direct.ValueOf(in.SourceText)
	out.TargetText = direct.ValueOf(in.TargetText)
	// MISSING: Usage
	return out
}
func ExampleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.ExampleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExampleObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SourceText
	// MISSING: TargetText
	out.Usage = direct.LazyPtr(in.GetUsage())
	return out
}
func ExampleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExampleObservedState) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SourceText
	// MISSING: TargetText
	out.Usage = direct.ValueOf(in.Usage)
	return out
}
func TranslationExampleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.TranslationExampleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationExampleObservedState{}
	// MISSING: Name
	// MISSING: SourceText
	// MISSING: TargetText
	// MISSING: Usage
	return out
}
func TranslationExampleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationExampleObservedState) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	// MISSING: SourceText
	// MISSING: TargetText
	// MISSING: Usage
	return out
}
func TranslationExampleSpec_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.TranslationExampleSpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationExampleSpec{}
	// MISSING: Name
	// MISSING: SourceText
	// MISSING: TargetText
	// MISSING: Usage
	return out
}
func TranslationExampleSpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationExampleSpec) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	// MISSING: SourceText
	// MISSING: TargetText
	// MISSING: Usage
	return out
}
