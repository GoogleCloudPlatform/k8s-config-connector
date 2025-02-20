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

package logging

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/api/distribution"
	metric "google.golang.org/genproto/googleapis/api/metric"
)

func MetricDescriptor_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.MetricDescriptor {
	return nil
}

func Distribution_BucketOptions_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.Distribution_BucketOptions {
	return nil
}

func MetricDescriptor_ToProto(mapCtx *direct.MapContext, in interface{}) *metric.MetricDescriptor {
	return nil
}

func Distribution_BucketOptions_ToProto(mapCtx *direct.MapContext, in interface{}) *pb.Distribution_BucketOptions {
	return nil
}
