package apigee

import (
        api "google.golang.org/api/apigee/v1"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
        krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
)

func init() {
        fuzztesting.RegisterKRMFuzzer_NoProto(envgroupFuzzer())
}

func envgroupFuzzer() fuzztesting.KRMFuzzer_NoProto {
        f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1EnvironmentGroup{},
                func(ctx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroup) *krm.ApigeeEnvgroupSpec {
                        return ApigeeEnvgroupSpec_FromApi(ctx, in)
                },
                func(ctx *direct.MapContext, in *krm.ApigeeEnvgroupSpec) *api.GoogleCloudApigeeV1EnvironmentGroup {
                        resourceID := ""
                        if in.ResourceID != nil {
                                resourceID = *in.ResourceID
                        }
                        return ApigeeEnvgroupSpec_ToApi(ctx, in, resourceID)
                },
                ApigeeEnvgroupObservedState_FromApi, ApigeeEnvgroupObservedState_ToApi,
        )

        f.SpecField(".Hostnames")
        
        f.StatusField(".CreatedAt")
        f.StatusField(".LastModifiedAt")
        f.StatusField(".State")
        f.StatusField(".Name")

        f.Unimplemented_NotYetTriaged(".ForceSendFields")
        f.Unimplemented_NotYetTriaged(".NullFields")
        f.Unimplemented_NotYetTriaged(".ServerResponse")

        return f
}
