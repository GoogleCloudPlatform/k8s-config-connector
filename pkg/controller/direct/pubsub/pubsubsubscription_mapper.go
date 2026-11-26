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

// +generated:mapper
// krm.group: pubsub.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.pubsub.v1

package pubsub

import (
	"fmt"

	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	krmbigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExpirationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ExpirationPolicy) *krm.ExpirationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ExpirationPolicy{}
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	return out
}

func ExpirationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ExpirationPolicy) *pb.ExpirationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ExpirationPolicy{}
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	return out
}

func PushConfig_FromProto(mapCtx *direct.MapContext, in *pb.PushConfig) *krm.PushConfig {
	if in == nil {
		return nil
	}
	out := &krm.PushConfig{}
	out.PushEndpoint = direct.LazyPtr(in.GetPushEndpoint())
	out.Attributes = in.Attributes
	out.OidcToken = PushConfigOidcToken_FromProto(mapCtx, in.GetOidcToken())
	out.NoWrapper = PushConfigNoWrapper_FromProto(mapCtx, in.GetNoWrapper())
	return out
}

func PushConfig_ToProto(mapCtx *direct.MapContext, in *krm.PushConfig) *pb.PushConfig {
	if in == nil {
		return nil
	}
	out := &pb.PushConfig{}
	out.PushEndpoint = direct.ValueOf(in.PushEndpoint)
	out.Attributes = in.Attributes
	if in.OidcToken != nil {
		out.AuthenticationMethod = &pb.PushConfig_OidcToken_{OidcToken: PushConfigOidcToken_ToProto(mapCtx, in.OidcToken)}
	}
	if oneof := PushConfigNoWrapper_ToProto(mapCtx, in.NoWrapper); oneof != nil {
		out.Wrapper = &pb.PushConfig_NoWrapper_{NoWrapper: oneof}
	}
	return out
}

// BigQueryConfig_FromProto is hand-written because Pub/Sub returns
// bigquery_config.table in its own "{project}.{dataset}.{table}" (or
// "{project}:{dataset}.{table}") format, while tableRef.external holds the
// canonical BigQueryTable format "projects/{project}/datasets/{dataset}/tables/{table}".
func BigQueryConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryConfig) *krm.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConfig{}
	if in.GetTable() != "" {
		out.TableRef = &krmbigqueryv1beta1.BigQueryTableRef{External: bigQueryTableExternal(in.GetTable())}
	}
	out.UseTopicSchema = direct.LazyPtr(in.GetUseTopicSchema())
	out.WriteMetadata = direct.LazyPtr(in.GetWriteMetadata())
	out.DropUnknownFields = direct.LazyPtr(in.GetDropUnknownFields())
	// MISSING: State
	// MISSING: UseTableSchema
	// MISSING: ServiceAccountEmail
	return out
}

// bigQueryTableExternal converts a Pub/Sub table name to the canonical
// BigQueryTable external format. Pub/Sub validates table names on write, so a
// value that can't be parsed is unexpected; it is kept unchanged so that it shows
// up as a diff (and gets corrected) instead of failing the mapping.
func bigQueryTableExternal(table string) string {
	id, ok := ParsePubSubBigQueryTable(table)
	if !ok {
		return table
	}
	return id.String()
}

// BigQueryConfig_ToProto is hand-written because Pub/Sub only accepts
// bigquery_config.table as "{project}.{dataset}.{table}" or
// "{project}:{dataset}.{table}". It rejects the canonical
// "projects/{project}/datasets/{dataset}/tables/{table}" format that
// tableRef.external holds after reference normalization with
// "Could not parse BigQuery table path".
func BigQueryConfig_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConfig) *pb.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryConfig{}
	if in.TableRef != nil {
		out.Table = pubSubBigQueryTable(in.TableRef.External)
	}
	out.UseTopicSchema = direct.ValueOf(in.UseTopicSchema)
	out.WriteMetadata = direct.ValueOf(in.WriteMetadata)
	out.DropUnknownFields = direct.ValueOf(in.DropUnknownFields)
	// MISSING: State
	// MISSING: UseTableSchema
	// MISSING: ServiceAccountEmail
	return out
}

// pubSubBigQueryTable converts a canonical BigQueryTable external to the
// "{project}.{dataset}.{table}" format used by Pub/Sub. Desired values are
// already canonical after reference normalization; only actual values
// round-tripped by mappers.OnlySpecFields can be non-canonical, and those are
// kept unchanged so they surface as a diff instead of failing the mapping.
func pubSubBigQueryTable(external string) string {
	id := &krmbigqueryv1beta1.BigQueryTableIdentity{}
	if err := id.FromExternal(external); err != nil {
		return external
	}
	return fmt.Sprintf("%s.%s.%s", id.Project, id.Dataset, id.Table)
}
