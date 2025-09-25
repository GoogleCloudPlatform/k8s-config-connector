// Copyright 2024 Google LLC
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

package alloydb

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datepb "google.golang.org/genproto/googleapis/type/date"
)

func Date_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *krm.Date {
	mapCtx.NotImplemented()
	return nil
}

func Date_ToProto(mapCtx *direct.MapContext, in *krm.Date) *datepb.Date {
	mapCtx.NotImplemented()
	return nil
}
