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

package container

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"

	pb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func walkDiff(parent string, m1, m2 protoreflect.Message, diffs map[string]bool) {
	m1.Range(func(fd protoreflect.FieldDescriptor, v1 protoreflect.Value) bool {
		path := parent + "." + string(fd.Name())
		v2 := m2.Get(fd)

		if !m2.Has(fd) {
			diffs[path] = true
			return true
		}

		switch fd.Kind() {
		case protoreflect.MessageKind:
			if fd.IsList() {
				l1 := v1.List()
				l2 := v2.List()
				if l1.Len() != l2.Len() {
					diffs[path] = true
				} else {
					for i := 0; i < l1.Len(); i++ {
						walkDiff(path+"[]", l1.Get(i).Message(), l2.Get(i).Message(), diffs)
					}
				}
			} else if fd.IsMap() {
				diffs[path] = true
			} else {
				walkDiff(path, v1.Message(), v2.Message(), diffs)
			}
		default:
			if v1.Interface() != v2.Interface() {
				diffs[path] = true
			}
		}
		return true
	})
}

func minimizePaths(paths []string) []string {
	var minimized []string
	for _, p := range paths {
		isChild := false
		for _, m := range minimized {
			if strings.HasPrefix(p, m+".") || strings.HasPrefix(p, m+"[") {
				isChild = true
				break
			}
		}
		if !isChild {
			minimized = append(minimized, p)
		}
	}
	return minimized
}

// TestFindAllDifferingFields dynamically walks the protobuf structure and KRM representation,
// generating random data to detect any fields that are not preserved across the FromProto/ToProto translation layer.
// This serves as an automated verification tool to maintain fuzzer coverage by identifying newly added fields,
// fields that aren't mapped, or fields not correctly triaged in UnimplementedFields.
func TestFindAllDifferingFields(t *testing.T) {
	randStream := rand.New(rand.NewSource(42))

	// Collect Spec differences
	specDiffs := make(map[string]bool)
	for i := 0; i < 50; i++ {
		p1 := &pb.Cluster{}
		fuzz.FillWithRandom(t, randStream, p1)

		ctx := &direct.MapContext{}
		krmSpec := ContainerClusterSpec_FromProto(ctx, p1)
		p2 := ContainerClusterSpec_ToProto(ctx, krmSpec)

		walkDiff("", p1.ProtoReflect(), p2.ProtoReflect(), specDiffs)
	}

	var specPaths []string
	for p := range specDiffs {
		specPaths = append(specPaths, p)
	}
	sort.Strings(specPaths)
	specPaths = minimizePaths(specPaths)

	fmt.Println("--- SPEC DIFFERENCES (MINIMIZED) ---")
	for _, p := range specPaths {
		fmt.Printf("f.Unimplemented_NotYetTriaged(%q)\n", p)
	}

	// Collect Status differences
	statusDiffs := make(map[string]bool)
	for i := 0; i < 50; i++ {
		p1 := &pb.Cluster{}
		fuzz.FillWithRandom(t, randStream, p1)

		ctx := &direct.MapContext{}
		krmStatus := ClusterObservedState_FromProto(ctx, p1)
		p2 := ClusterObservedState_ToProto(ctx, krmStatus)

		walkDiff("", p1.ProtoReflect(), p2.ProtoReflect(), statusDiffs)
	}

	var statusPaths []string
	for p := range statusDiffs {
		statusPaths = append(statusPaths, p)
	}
	sort.Strings(statusPaths)
	statusPaths = minimizePaths(statusPaths)

	fmt.Println("--- STATUS DIFFERENCES (MINIMIZED) ---")
	for _, p := range statusPaths {
		fmt.Printf("f.Unimplemented_NotYetTriaged(%q)\n", p)
	}
}

func TestContainerClusterFuzzer(t *testing.T) {
	t.Parallel()
	fuzzer := containerClusterFuzzer()
	for i := int64(0); i < 100; i++ {
		seed := rand.Int63()
		fuzzer.FuzzSpec(t, seed)
		fuzzer.FuzzStatus(t, seed)
	}
}
