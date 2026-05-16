package apigee

import (
        api "google.golang.org/api/apigee/v1"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
        fuzztesting.RegisterKRMFuzzer_NoProto(instanceFuzzer())
}

func instanceFuzzer() fuzztesting.KRMFuzzer_NoProto {
        f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1Instance{},
                ApigeeInstanceSpec_FromAPI, ApigeeInstanceSpec_ToAPI,
                ApigeeInstanceObservedState_FromAPI, ApigeeInstanceObservedState_ToAPI,
        )

        f.SpecField(".AccessLoggingConfig")
        f.SpecField(".ConsumerAcceptList")
        f.SpecField(".Description")
        f.SpecField(".DiskEncryptionKeyName")
        f.SpecField(".DisplayName")
        f.SpecField(".IpRange")
        f.SpecField(".Location")
        f.SpecField(".PeeringCidrRange")

        f.StatusField(".CreatedAt")
        f.StatusField(".Host")
        f.StatusField(".LastModifiedAt")
        f.StatusField(".Port")
        f.StatusField(".RuntimeVersion")
        f.StatusField(".ServiceAttachment")
        f.StatusField(".State")

        f.Unimplemented_NotYetTriaged(".MaintenanceUpdatePolicy")
        f.Unimplemented_NotYetTriaged(".ScheduledMaintenance")
        f.Unimplemented_NotYetTriaged(".IsVersionLocked")
        f.Unimplemented_NotYetTriaged(".Name")
        f.Unimplemented_NotYetTriaged(".ForceSendFields")
        f.Unimplemented_NotYetTriaged(".NullFields")
        f.Unimplemented_NotYetTriaged(".ServerResponse")
        f.Unimplemented_NotYetTriaged(".AccessLoggingConfig.ForceSendFields")
        f.Unimplemented_NotYetTriaged(".AccessLoggingConfig.NullFields")

        return f
}
