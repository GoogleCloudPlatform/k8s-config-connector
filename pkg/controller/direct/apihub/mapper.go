package apihub

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Attribute_AllowedValue_FromProto(mapCtx *direct.MapContext, in *apihubpb.Attribute_AllowedValue) *krm.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &krm.Attribute_AllowedValue{}
	out.ID = &in.Id
        out.DisplayName = &in.DisplayName
        out.Description = &in.Description
        out.Immutable = &in.Immutable
	return out
}

func Attribute_AllowedValue_ToProto(mapCtx *direct.MapContext, in *krm.Attribute_AllowedValue) *apihubpb.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	id := ""
        if in.ID != nil {
            id = *in.ID
        }
        displayName := ""
        if in.DisplayName != nil {
            displayName = *in.DisplayName
        }
        description := ""
        if in.Description != nil {
            description = *in.Description
        }
        immutable := false
        if in.Immutable != nil {
            immutable = *in.Immutable
        }

	out := &apihubpb.Attribute_AllowedValue{
		Id: id,
                DisplayName: displayName,
                Description: description,
                Immutable: immutable,
	}
	return out
}
