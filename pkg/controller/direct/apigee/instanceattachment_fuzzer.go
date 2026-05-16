package apigee

import (
        api "google.golang.org/api/apigee/v1"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
        fuzztesting.RegisterKRMFuzzer_NoProto(instanceAttachmentFuzzer())
}

func instanceAttachmentFuzzer() fuzztesting.KRMFuzzer_NoProto {
        f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1InstanceAttachment{},
                ApigeeInstanceAttachmentSpec_FromAPI, ApigeeInstanceAttachmentSpec_ToAPI,
                ApigeeInstanceAttachmentObservedState_FromAPI, ApigeeInstanceAttachmentObservedState_ToAPI,
        )

        f.SpecField(".Environment")
        f.StatusField(".CreatedAt")

        f.Unimplemented_NotYetTriaged(".Name")
        f.Unimplemented_NotYetTriaged(".ForceSendFields")
        f.Unimplemented_NotYetTriaged(".NullFields")
        f.Unimplemented_NotYetTriaged(".ServerResponse")

        f.FilterStatus = func(in *api.GoogleCloudApigeeV1InstanceAttachment) {
                // time.RFC3339 format drops the milliseconds, so we zero them to pass roundtrip.
                // also keep it within a reasonable range for time parsing.
                in.CreatedAt = in.CreatedAt % 253402300799000
                if in.CreatedAt < 0 {
                        in.CreatedAt = -in.CreatedAt
                }
                in.CreatedAt = (in.CreatedAt / 1000) * 1000
        }

        return f
}
