// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dlp

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type DeidentifyTemplate struct{}

func DeidentifyTemplateToUnstructured(r *dclService.DeidentifyTemplate) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dlp",
			Version: "beta",
			Type:    "DeidentifyTemplate",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeidentifyConfig != nil && r.DeidentifyConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfig {
		rDeidentifyConfig := make(map[string]interface{})
		if r.DeidentifyConfig.InfoTypeTransformations != nil && r.DeidentifyConfig.InfoTypeTransformations != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformations {
			rDeidentifyConfigInfoTypeTransformations := make(map[string]interface{})
			var rDeidentifyConfigInfoTypeTransformationsTransformations []interface{}
			for _, rDeidentifyConfigInfoTypeTransformationsTransformationsVal := range r.DeidentifyConfig.InfoTypeTransformations.Transformations {
				rDeidentifyConfigInfoTypeTransformationsTransformationsObject := make(map[string]interface{})
				var rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypes []interface{}
				for _, rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypesVal := range rDeidentifyConfigInfoTypeTransformationsTransformationsVal.InfoTypes {
					rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypesObject := make(map[string]interface{})
					if rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypesVal.Name != nil {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypesObject["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypesVal.Name
					}
					rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypes = append(rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypes, rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypesObject)
				}
				rDeidentifyConfigInfoTypeTransformationsTransformationsObject["infoTypes"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValInfoTypes
				if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation {
					rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation := make(map[string]interface{})
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.BucketingConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.BucketingConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfig := make(map[string]interface{})
						var rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets []interface{}
						for _, rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal := range rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.BucketingConfig.Buckets {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.BooleanValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["booleanValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.BooleanValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue := make(map[string]interface{})
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Day != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["day"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Day
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Month != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["month"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Month
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Year != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["year"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Year
									}
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["dateValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DayOfWeekValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["dayOfWeekValue"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DayOfWeekValue)
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.FloatValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["floatValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.FloatValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.IntegerValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["integerValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.IntegerValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.StringValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["stringValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.StringValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue := make(map[string]interface{})
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Hours != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["hours"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Hours
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Minutes != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["minutes"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Minutes
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Nanos != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["nanos"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Nanos
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Seconds != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["seconds"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Seconds
									}
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["timeValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimestampValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["timestampValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimestampValue
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["max"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.BooleanValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["booleanValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.BooleanValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue := make(map[string]interface{})
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Day != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["day"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Day
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Month != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["month"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Month
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Year != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["year"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Year
									}
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["dateValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DayOfWeekValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["dayOfWeekValue"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DayOfWeekValue)
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.FloatValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["floatValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.FloatValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.IntegerValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["integerValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.IntegerValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.StringValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["stringValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.StringValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue := make(map[string]interface{})
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Hours != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["hours"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Hours
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Minutes != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["minutes"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Minutes
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Nanos != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["nanos"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Nanos
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Seconds != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["seconds"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Seconds
									}
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["timeValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimestampValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["timestampValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimestampValue
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["min"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.BooleanValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["booleanValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.BooleanValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue := make(map[string]interface{})
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Day != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["day"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Day
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Month != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["month"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Month
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Year != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["year"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Year
									}
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["dateValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DayOfWeekValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["dayOfWeekValue"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DayOfWeekValue)
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.FloatValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["floatValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.FloatValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.IntegerValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["integerValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.IntegerValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.StringValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["stringValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.StringValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue := make(map[string]interface{})
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Hours != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["hours"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Hours
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Minutes != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["minutes"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Minutes
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Nanos != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["nanos"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Nanos
									}
									if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Seconds != nil {
										rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["seconds"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Seconds
									}
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["timeValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimestampValue != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["timestampValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimestampValue
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["replacementValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets = append(rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets, rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject)
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfig["buckets"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["bucketingConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig := make(map[string]interface{})
						var rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore []interface{}
						for _, rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal := range rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CharactersToSkip != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject["charactersToSkip"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CharactersToSkip
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CommonCharactersToIgnore != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject["commonCharactersToIgnore"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CommonCharactersToIgnore)
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore = append(rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore, rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject)
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["maskingCharacter"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.NumberToMask != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["numberToMask"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.NumberToMask
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["reverseOrder"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["characterMaskConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig["context"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["cryptoDeterministicConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["transient"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfig["cryptoKey"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["cryptoHashConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet)
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["cryptoReplaceFfxFpeConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigContext := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context.Name != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigContext["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context.Name
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["context"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigContext
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["transient"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["cryptoKey"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.LowerBoundDays != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["lowerBoundDays"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.LowerBoundDays
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.UpperBoundDays != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["upperBoundDays"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.UpperBoundDays
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["dateShiftConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue)
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound
						}
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue)
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["upperBound"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["fixedSizeBucketingConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.RedactConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.RedactConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationRedactConfig := make(map[string]interface{})
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["redactConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationRedactConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue := make(map[string]interface{})
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["booleanValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["day"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["month"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["year"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["dateValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue)
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["floatValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["integerValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["stringValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue := make(map[string]interface{})
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos
								}
								if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds != nil {
									rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds
								}
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["timeValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue
							}
							if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue != nil {
								rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["timestampValue"] = *rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue
							}
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfig["newValue"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["replaceConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceWithInfoTypeConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceWithInfoTypeConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceWithInfoTypeConfig := make(map[string]interface{})
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["replaceWithInfoTypeConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceWithInfoTypeConfig
					}
					if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig != nil && rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig {
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationTimePartConfig := make(map[string]interface{})
						if rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig.PartToExtract != nil {
							rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationTimePartConfig["partToExtract"] = string(*rDeidentifyConfigInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig.PartToExtract)
						}
						rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation["timePartConfig"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformationTimePartConfig
					}
					rDeidentifyConfigInfoTypeTransformationsTransformationsObject["primitiveTransformation"] = rDeidentifyConfigInfoTypeTransformationsTransformationsValPrimitiveTransformation
				}
				rDeidentifyConfigInfoTypeTransformationsTransformations = append(rDeidentifyConfigInfoTypeTransformationsTransformations, rDeidentifyConfigInfoTypeTransformationsTransformationsObject)
			}
			rDeidentifyConfigInfoTypeTransformations["transformations"] = rDeidentifyConfigInfoTypeTransformationsTransformations
			rDeidentifyConfig["infoTypeTransformations"] = rDeidentifyConfigInfoTypeTransformations
		}
		if r.DeidentifyConfig.RecordTransformations != nil && r.DeidentifyConfig.RecordTransformations != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformations {
			rDeidentifyConfigRecordTransformations := make(map[string]interface{})
			var rDeidentifyConfigRecordTransformationsFieldTransformations []interface{}
			for _, rDeidentifyConfigRecordTransformationsFieldTransformationsVal := range r.DeidentifyConfig.RecordTransformations.FieldTransformations {
				rDeidentifyConfigRecordTransformationsFieldTransformationsObject := make(map[string]interface{})
				if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition {
					rDeidentifyConfigRecordTransformationsFieldTransformationsValCondition := make(map[string]interface{})
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressions := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions.Conditions != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions.Conditions != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditions := make(map[string]interface{})
							var rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditions []interface{}
							for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions.Conditions.Conditions {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsObject := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Field != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Field != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValField := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Field.Name != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValField["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Field.Name
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsObject["field"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValField
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Operator != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsObject["operator"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Operator)
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.BooleanValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.BooleanValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueDateValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Day != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Day
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Month != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Month
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Year != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Year
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueDateValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DayOfWeekValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.DayOfWeekValue)
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.FloatValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.FloatValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.IntegerValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.IntegerValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.StringValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.StringValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueTimeValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Hours != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Hours
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Minutes != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Minutes
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Nanos != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Nanos
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Seconds != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Seconds
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValueTimeValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimestampValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsVal.Value.TimestampValue
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsObject["value"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsValValue
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditions = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditions, rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditionsObject)
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditions["conditions"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditionsConditions
							rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressions["conditions"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressionsConditions
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions.LogicalOperator != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressions["logicalOperator"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Condition.Expressions.LogicalOperator)
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValCondition["expressions"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValConditionExpressions
					}
					rDeidentifyConfigRecordTransformationsFieldTransformationsObject["condition"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValCondition
				}
				var rDeidentifyConfigRecordTransformationsFieldTransformationsValFields []interface{}
				for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValFieldsVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsVal.Fields {
					rDeidentifyConfigRecordTransformationsFieldTransformationsValFieldsObject := make(map[string]interface{})
					if rDeidentifyConfigRecordTransformationsFieldTransformationsValFieldsVal.Name != nil {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValFieldsObject["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValFieldsVal.Name
					}
					rDeidentifyConfigRecordTransformationsFieldTransformationsValFields = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValFields, rDeidentifyConfigRecordTransformationsFieldTransformationsValFieldsObject)
				}
				rDeidentifyConfigRecordTransformationsFieldTransformationsObject["fields"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValFields
				if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.InfoTypeTransformations != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.InfoTypeTransformations != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations {
					rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformations := make(map[string]interface{})
					var rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformations []interface{}
					for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsVal.InfoTypeTransformations.Transformations {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsObject := make(map[string]interface{})
						var rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypes []interface{}
						for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypesVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.InfoTypes {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypesObject := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypesVal.Name != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypesObject["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypesVal.Name
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypes = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypes, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypesObject)
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsObject["infoTypes"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValInfoTypes
						if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.BucketingConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.BucketingConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfig := make(map[string]interface{})
								var rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets []interface{}
								for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.BucketingConfig.Buckets {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.BooleanValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.BooleanValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue := make(map[string]interface{})
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Day != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Day
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Month != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Month
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Year != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Year
											}
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DayOfWeekValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DayOfWeekValue)
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.FloatValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.FloatValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.IntegerValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.IntegerValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.StringValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.StringValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue := make(map[string]interface{})
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Hours != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Hours
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Minutes != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Minutes
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Nanos != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Nanos
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Seconds != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Seconds
											}
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimestampValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimestampValue
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["max"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.BooleanValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.BooleanValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue := make(map[string]interface{})
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Day != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Day
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Month != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Month
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Year != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Year
											}
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DayOfWeekValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DayOfWeekValue)
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.FloatValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.FloatValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.IntegerValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.IntegerValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.StringValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.StringValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue := make(map[string]interface{})
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Hours != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Hours
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Minutes != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Minutes
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Nanos != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Nanos
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Seconds != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Seconds
											}
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimestampValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimestampValue
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["min"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.BooleanValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.BooleanValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue := make(map[string]interface{})
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Day != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Day
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Month != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Month
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Year != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Year
											}
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DayOfWeekValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DayOfWeekValue)
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.FloatValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.FloatValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.IntegerValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.IntegerValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.StringValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.StringValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue := make(map[string]interface{})
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Hours != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Hours
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Minutes != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Minutes
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Nanos != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Nanos
											}
											if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Seconds != nil {
												rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Seconds
											}
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimestampValue != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimestampValue
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["replacementValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBucketsObject)
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfig["buckets"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfigBuckets
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["bucketingConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationBucketingConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig := make(map[string]interface{})
								var rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore []interface{}
								for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CharactersToSkip != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject["charactersToSkip"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CharactersToSkip
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CommonCharactersToIgnore != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject["commonCharactersToIgnore"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CommonCharactersToIgnore)
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject)
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["maskingCharacter"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.NumberToMask != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["numberToMask"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.NumberToMask
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig["reverseOrder"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["characterMaskConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCharacterMaskConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig["context"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["cryptoDeterministicConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoDeterministicConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["cryptoHashConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoHashConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet)
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["cryptoReplaceFfxFpeConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigContext := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context.Name != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigContext["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context.Name
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["context"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigContext
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.LowerBoundDays != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["lowerBoundDays"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.LowerBoundDays
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.UpperBoundDays != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig["upperBoundDays"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.DateShiftConfig.UpperBoundDays
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["dateShiftConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationDateShiftConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue)
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue)
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["upperBound"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["fixedSizeBucketingConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationFixedSizeBucketingConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.RedactConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.RedactConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationRedactConfig := make(map[string]interface{})
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["redactConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationRedactConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue)
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos
										}
										if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds != nil {
											rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds
										}
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfig["newValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfigNewValue
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["replaceConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceWithInfoTypeConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.ReplaceWithInfoTypeConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceWithInfoTypeConfig := make(map[string]interface{})
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["replaceWithInfoTypeConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationReplaceWithInfoTypeConfig
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationTimePartConfig := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig.PartToExtract != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationTimePartConfig["partToExtract"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsVal.PrimitiveTransformation.TimePartConfig.PartToExtract)
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation["timePartConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformationTimePartConfig
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsObject["primitiveTransformation"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsValPrimitiveTransformation
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformations = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformations, rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformationsObject)
					}
					rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformations["transformations"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformationsTransformations
					rDeidentifyConfigRecordTransformationsFieldTransformationsObject["infoTypeTransformations"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValInfoTypeTransformations
				}
				if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation {
					rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation := make(map[string]interface{})
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.BucketingConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.BucketingConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfig := make(map[string]interface{})
						var rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBuckets []interface{}
						for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.BucketingConfig.Buckets {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsObject := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.BooleanValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.BooleanValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Day != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Day
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Month != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Month
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Year != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DateValue.Year
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxDateValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DayOfWeekValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.DayOfWeekValue)
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.FloatValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.FloatValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.IntegerValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.IntegerValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.StringValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.StringValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Hours != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Hours
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Minutes != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Minutes
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Nanos != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Nanos
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Seconds != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimeValue.Seconds
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMaxTimeValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimestampValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Max.TimestampValue
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["max"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMax
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.BooleanValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.BooleanValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Day != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Day
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Month != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Month
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Year != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DateValue.Year
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinDateValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DayOfWeekValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.DayOfWeekValue)
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.FloatValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.FloatValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.IntegerValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.IntegerValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.StringValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.StringValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Hours != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Hours
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Minutes != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Minutes
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Nanos != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Nanos
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Seconds != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimeValue.Seconds
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMinTimeValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimestampValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.Min.TimestampValue
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["min"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValMin
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.BooleanValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.BooleanValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Day != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Day
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Month != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Month
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Year != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DateValue.Year
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueDateValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DayOfWeekValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.DayOfWeekValue)
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.FloatValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.FloatValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.IntegerValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.IntegerValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.StringValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.StringValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Hours != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Hours
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Minutes != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Minutes
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Nanos != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Nanos
									}
									if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Seconds != nil {
										rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimeValue.Seconds
									}
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValueTimeValue
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimestampValue != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsVal.ReplacementValue.TimestampValue
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsObject["replacementValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsValReplacementValue
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBuckets = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBuckets, rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBucketsObject)
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfig["buckets"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfigBuckets
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["bucketingConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationBucketingConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfig := make(map[string]interface{})
						var rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore []interface{}
						for _, rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal := range rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CharactersToSkip != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject["charactersToSkip"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CharactersToSkip
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CommonCharactersToIgnore != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject["commonCharactersToIgnore"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreVal.CommonCharactersToIgnore)
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore = append(rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore, rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreObject)
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfigCharactersToIgnore
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfig["maskingCharacter"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.NumberToMask != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfig["numberToMask"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.NumberToMask
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfig["reverseOrder"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["characterMaskConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCharacterMaskConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfig["context"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigContext
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigCryptoKey
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["cryptoDeterministicConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoDeterministicConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyTransient
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfigCryptoKey
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["cryptoHashConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoHashConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet)
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigContext
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["cryptoReplaceFfxFpeConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationCryptoReplaceFfxFpeConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigContext := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context.Name != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigContext["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.Context.Name
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfig["context"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigContext
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["transient"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyTransient
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfig["cryptoKey"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfigCryptoKey
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.LowerBoundDays != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfig["lowerBoundDays"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.LowerBoundDays
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.UpperBoundDays != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfig["upperBoundDays"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.DateShiftConfig.UpperBoundDays
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["dateShiftConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationDateShiftConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue)
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigLowerBound
						}
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue)
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfig["upperBound"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfigUpperBound
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["fixedSizeBucketingConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationFixedSizeBucketingConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.RedactConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.RedactConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationRedactConfig := make(map[string]interface{})
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["redactConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationRedactConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue := make(map[string]interface{})
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["booleanValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["day"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["month"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue["year"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["dateValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueDateValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue)
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["floatValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["integerValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["stringValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos
								}
								if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds != nil {
									rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds
								}
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["timeValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValueTimeValue
							}
							if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue != nil {
								rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue["timestampValue"] = *rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue
							}
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfig["newValue"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfigNewValue
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["replaceConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceWithInfoTypeConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.ReplaceWithInfoTypeConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceWithInfoTypeConfig := make(map[string]interface{})
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["replaceWithInfoTypeConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationReplaceWithInfoTypeConfig
					}
					if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.TimePartConfig != nil && rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.TimePartConfig != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig {
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationTimePartConfig := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.TimePartConfig.PartToExtract != nil {
							rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationTimePartConfig["partToExtract"] = string(*rDeidentifyConfigRecordTransformationsFieldTransformationsVal.PrimitiveTransformation.TimePartConfig.PartToExtract)
						}
						rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation["timePartConfig"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformationTimePartConfig
					}
					rDeidentifyConfigRecordTransformationsFieldTransformationsObject["primitiveTransformation"] = rDeidentifyConfigRecordTransformationsFieldTransformationsValPrimitiveTransformation
				}
				rDeidentifyConfigRecordTransformationsFieldTransformations = append(rDeidentifyConfigRecordTransformationsFieldTransformations, rDeidentifyConfigRecordTransformationsFieldTransformationsObject)
			}
			rDeidentifyConfigRecordTransformations["fieldTransformations"] = rDeidentifyConfigRecordTransformationsFieldTransformations
			var rDeidentifyConfigRecordTransformationsRecordSuppressions []interface{}
			for _, rDeidentifyConfigRecordTransformationsRecordSuppressionsVal := range r.DeidentifyConfig.RecordTransformations.RecordSuppressions {
				rDeidentifyConfigRecordTransformationsRecordSuppressionsObject := make(map[string]interface{})
				if rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition {
					rDeidentifyConfigRecordTransformationsRecordSuppressionsValCondition := make(map[string]interface{})
					if rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions {
						rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressions := make(map[string]interface{})
						if rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions.Conditions != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions.Conditions != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions {
							rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditions := make(map[string]interface{})
							var rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditions []interface{}
							for _, rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal := range rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions.Conditions.Conditions {
								rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsObject := make(map[string]interface{})
								if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Field != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Field != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField {
									rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValField := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Field.Name != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValField["name"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Field.Name
									}
									rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsObject["field"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValField
								}
								if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Operator != nil {
									rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsObject["operator"] = string(*rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Operator)
								}
								if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue {
									rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue := make(map[string]interface{})
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.BooleanValue != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["booleanValue"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.BooleanValue
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueDateValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Day != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueDateValue["day"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Day
										}
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Month != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueDateValue["month"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Month
										}
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Year != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueDateValue["year"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DateValue.Year
										}
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["dateValue"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueDateValue
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DayOfWeekValue != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["dayOfWeekValue"] = string(*rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.DayOfWeekValue)
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.FloatValue != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["floatValue"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.FloatValue
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.IntegerValue != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["integerValue"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.IntegerValue
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.StringValue != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["stringValue"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.StringValue
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue != nil && rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue != dclService.EmptyDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueTimeValue := make(map[string]interface{})
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Hours != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueTimeValue["hours"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Hours
										}
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Minutes != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueTimeValue["minutes"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Minutes
										}
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Nanos != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueTimeValue["nanos"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Nanos
										}
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Seconds != nil {
											rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueTimeValue["seconds"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimeValue.Seconds
										}
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["timeValue"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValueTimeValue
									}
									if rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimestampValue != nil {
										rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue["timestampValue"] = *rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsVal.Value.TimestampValue
									}
									rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsObject["value"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsValValue
								}
								rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditions = append(rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditions, rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditionsObject)
							}
							rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditions["conditions"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditionsConditions
							rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressions["conditions"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressionsConditions
						}
						if rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions.LogicalOperator != nil {
							rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressions["logicalOperator"] = string(*rDeidentifyConfigRecordTransformationsRecordSuppressionsVal.Condition.Expressions.LogicalOperator)
						}
						rDeidentifyConfigRecordTransformationsRecordSuppressionsValCondition["expressions"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValConditionExpressions
					}
					rDeidentifyConfigRecordTransformationsRecordSuppressionsObject["condition"] = rDeidentifyConfigRecordTransformationsRecordSuppressionsValCondition
				}
				rDeidentifyConfigRecordTransformationsRecordSuppressions = append(rDeidentifyConfigRecordTransformationsRecordSuppressions, rDeidentifyConfigRecordTransformationsRecordSuppressionsObject)
			}
			rDeidentifyConfigRecordTransformations["recordSuppressions"] = rDeidentifyConfigRecordTransformationsRecordSuppressions
			rDeidentifyConfig["recordTransformations"] = rDeidentifyConfigRecordTransformations
		}
		if r.DeidentifyConfig.TransformationErrorHandling != nil && r.DeidentifyConfig.TransformationErrorHandling != dclService.EmptyDeidentifyTemplateDeidentifyConfigTransformationErrorHandling {
			rDeidentifyConfigTransformationErrorHandling := make(map[string]interface{})
			if r.DeidentifyConfig.TransformationErrorHandling.LeaveUntransformed != nil && r.DeidentifyConfig.TransformationErrorHandling.LeaveUntransformed != dclService.EmptyDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed {
				rDeidentifyConfigTransformationErrorHandlingLeaveUntransformed := make(map[string]interface{})
				rDeidentifyConfigTransformationErrorHandling["leaveUntransformed"] = rDeidentifyConfigTransformationErrorHandlingLeaveUntransformed
			}
			if r.DeidentifyConfig.TransformationErrorHandling.ThrowError != nil && r.DeidentifyConfig.TransformationErrorHandling.ThrowError != dclService.EmptyDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError {
				rDeidentifyConfigTransformationErrorHandlingThrowError := make(map[string]interface{})
				rDeidentifyConfigTransformationErrorHandling["throwError"] = rDeidentifyConfigTransformationErrorHandlingThrowError
			}
			rDeidentifyConfig["transformationErrorHandling"] = rDeidentifyConfigTransformationErrorHandling
		}
		u.Object["deidentifyConfig"] = rDeidentifyConfig
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.LocationId != nil {
		u.Object["locationId"] = *r.LocationId
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToDeidentifyTemplate(u *unstructured.Resource) (*dclService.DeidentifyTemplate, error) {
	r := &dclService.DeidentifyTemplate{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deidentifyConfig"]; ok {
		if rDeidentifyConfig, ok := u.Object["deidentifyConfig"].(map[string]interface{}); ok {
			r.DeidentifyConfig = &dclService.DeidentifyTemplateDeidentifyConfig{}
			if _, ok := rDeidentifyConfig["infoTypeTransformations"]; ok {
				if rDeidentifyConfigInfoTypeTransformations, ok := rDeidentifyConfig["infoTypeTransformations"].(map[string]interface{}); ok {
					r.DeidentifyConfig.InfoTypeTransformations = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformations{}
					if _, ok := rDeidentifyConfigInfoTypeTransformations["transformations"]; ok {
						if s, ok := rDeidentifyConfigInfoTypeTransformations["transformations"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rDeidentifyConfigInfoTypeTransformationsTransformations dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations
									if _, ok := objval["infoTypes"]; ok {
										if s, ok := objval["infoTypes"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes.Name: expected string")
														}
													}
													rDeidentifyConfigInfoTypeTransformationsTransformations.InfoTypes = append(rDeidentifyConfigInfoTypeTransformationsTransformations.InfoTypes, rDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes)
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.InfoTypes: expected []interface{}")
										}
									}
									if _, ok := objval["primitiveTransformation"]; ok {
										if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation, ok := objval["primitiveTransformation"].(map[string]interface{}); ok {
											rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation{}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["bucketingConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["bucketingConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig["buckets"]; ok {
														if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig["buckets"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets
																	if _, ok := objval["max"]; ok {
																		if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax, ok := objval["max"].(map[string]interface{}); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax{}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["booleanValue"]; ok {
																				if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["booleanValue"].(bool); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.BooleanValue = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.BooleanValue: expected bool")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dateValue"]; ok {
																				if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dateValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue{}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["day"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["day"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Day = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Day: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["month"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["month"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Month = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Month: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["year"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["year"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Year = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Year: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dayOfWeekValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dayOfWeekValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DayOfWeekValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["floatValue"]; ok {
																				if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["floatValue"].(float64); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.FloatValue = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.FloatValue: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["integerValue"]; ok {
																				if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["integerValue"].(int64); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.IntegerValue = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.IntegerValue: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["stringValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["stringValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.StringValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.StringValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timeValue"]; ok {
																				if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timeValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue{}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["hours"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["hours"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Hours = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Hours: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["minutes"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["minutes"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Minutes = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Minutes: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["nanos"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["nanos"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Nanos = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Nanos: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["seconds"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["seconds"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Seconds = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Seconds: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timestampValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timestampValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimestampValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimestampValue: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max: expected map[string]interface{}")
																		}
																	}
																	if _, ok := objval["min"]; ok {
																		if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin, ok := objval["min"].(map[string]interface{}); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin{}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["booleanValue"]; ok {
																				if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["booleanValue"].(bool); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.BooleanValue = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.BooleanValue: expected bool")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dateValue"]; ok {
																				if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dateValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue{}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["day"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["day"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Day = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Day: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["month"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["month"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Month = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Month: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["year"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["year"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Year = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Year: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dayOfWeekValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dayOfWeekValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DayOfWeekValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["floatValue"]; ok {
																				if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["floatValue"].(float64); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.FloatValue = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.FloatValue: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["integerValue"]; ok {
																				if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["integerValue"].(int64); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.IntegerValue = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.IntegerValue: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["stringValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["stringValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.StringValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.StringValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timeValue"]; ok {
																				if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timeValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue{}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["hours"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["hours"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Hours = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Hours: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["minutes"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["minutes"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Minutes = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Minutes: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["nanos"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["nanos"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Nanos = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Nanos: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["seconds"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["seconds"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Seconds = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Seconds: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timestampValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timestampValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimestampValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimestampValue: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min: expected map[string]interface{}")
																		}
																	}
																	if _, ok := objval["replacementValue"]; ok {
																		if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue, ok := objval["replacementValue"].(map[string]interface{}); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue{}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["booleanValue"]; ok {
																				if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["booleanValue"].(bool); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.BooleanValue = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.BooleanValue: expected bool")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dateValue"]; ok {
																				if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dateValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue{}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["day"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["day"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Day = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Day: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["month"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["month"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Month = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Month: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["year"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["year"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Year = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Year: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dayOfWeekValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dayOfWeekValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DayOfWeekValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["floatValue"]; ok {
																				if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["floatValue"].(float64); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.FloatValue = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.FloatValue: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["integerValue"]; ok {
																				if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["integerValue"].(int64); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.IntegerValue = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.IntegerValue: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["stringValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["stringValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.StringValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.StringValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timeValue"]; ok {
																				if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timeValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue{}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["hours"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["hours"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Hours = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Hours: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["minutes"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["minutes"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Minutes = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Minutes: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["nanos"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["nanos"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Nanos = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Nanos: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["seconds"]; ok {
																						if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["seconds"].(int64); ok {
																							rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Seconds = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Seconds: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timestampValue"]; ok {
																				if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timestampValue"].(string); ok {
																					rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimestampValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimestampValue: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue: expected map[string]interface{}")
																		}
																	}
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig.Buckets = append(rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig.Buckets, rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets)
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig.Buckets: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["characterMaskConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["characterMaskConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"]; ok {
														if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore
																	if _, ok := objval["charactersToSkip"]; ok {
																		if s, ok := objval["charactersToSkip"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CharactersToSkip = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CharactersToSkip: expected string")
																		}
																	}
																	if _, ok := objval["commonCharactersToIgnore"]; ok {
																		if s, ok := objval["commonCharactersToIgnore"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CommonCharactersToIgnore = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CommonCharactersToIgnore: expected string")
																		}
																	}
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore = append(rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore, rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore)
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore: expected []interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["maskingCharacter"]; ok {
														if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["maskingCharacter"].(string); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter: expected string")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["numberToMask"]; ok {
														if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["numberToMask"].(int64); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.NumberToMask = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.NumberToMask: expected int64")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["reverseOrder"]; ok {
														if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["reverseOrder"].(bool); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder: expected bool")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoDeterministicConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoDeterministicConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["context"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["context"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext["name"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext["name"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoHashConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoHashConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig["cryptoKey"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoReplaceFfxFpeConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoReplaceFfxFpeConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"]; ok {
														if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"].(string); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnumRef(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet: expected string")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"]; ok {
														if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"].(string); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet: expected string")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"]; ok {
														if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"].(int64); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix: expected int64")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["dateShiftConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["dateShiftConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["context"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["context"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext["name"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext["name"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["cryptoKey"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["lowerBoundDays"]; ok {
														if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["lowerBoundDays"].(int64); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.LowerBoundDays = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.LowerBoundDays: expected int64")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["upperBoundDays"]; ok {
														if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["upperBoundDays"].(int64); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.UpperBoundDays = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.UpperBoundDays: expected int64")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["fixedSizeBucketingConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["fixedSizeBucketingConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"]; ok {
														if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"].(float64); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize: expected float64")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"]; ok {
																if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"].(bool); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue: expected bool")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"]; ok {
																if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"].(float64); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue: expected float64")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"]; ok {
																if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"].(int64); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue: expected int64")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["upperBound"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["upperBound"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"]; ok {
																if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"].(bool); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue: expected bool")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"]; ok {
																if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"].(float64); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue: expected float64")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"]; ok {
																if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"].(int64); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue: expected int64")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["redactConfig"]; ok {
												if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["redactConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.RedactConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig{}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.RedactConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig["newValue"]; ok {
														if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig["newValue"].(map[string]interface{}); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue{}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["booleanValue"]; ok {
																if b, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["booleanValue"].(bool); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue: expected bool")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dateValue"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dateValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["day"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["day"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["month"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["month"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["year"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["year"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["floatValue"]; ok {
																if f, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["floatValue"].(float64); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue: expected float64")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["integerValue"]; ok {
																if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["integerValue"].(int64); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue: expected int64")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["stringValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["stringValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timeValue"]; ok {
																if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timeValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue{}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"]; ok {
																		if i, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"].(int64); ok {
																			rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timestampValue"]; ok {
																if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timestampValue"].(string); ok {
																	rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceWithInfoTypeConfig"]; ok {
												if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceWithInfoTypeConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceWithInfoTypeConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig{}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceWithInfoTypeConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["timePartConfig"]; ok {
												if rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation["timePartConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig = &dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig{}
													if _, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig["partToExtract"]; ok {
														if s, ok := rDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig["partToExtract"].(string); ok {
															rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig.PartToExtract = dclService.DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnumRef(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig.PartToExtract: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigInfoTypeTransformationsTransformations.PrimitiveTransformation: expected map[string]interface{}")
										}
									}
									r.DeidentifyConfig.InfoTypeTransformations.Transformations = append(r.DeidentifyConfig.InfoTypeTransformations.Transformations, rDeidentifyConfigInfoTypeTransformationsTransformations)
								}
							}
						} else {
							return nil, fmt.Errorf("r.DeidentifyConfig.InfoTypeTransformations.Transformations: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DeidentifyConfig.InfoTypeTransformations: expected map[string]interface{}")
				}
			}
			if _, ok := rDeidentifyConfig["recordTransformations"]; ok {
				if rDeidentifyConfigRecordTransformations, ok := rDeidentifyConfig["recordTransformations"].(map[string]interface{}); ok {
					r.DeidentifyConfig.RecordTransformations = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformations{}
					if _, ok := rDeidentifyConfigRecordTransformations["fieldTransformations"]; ok {
						if s, ok := rDeidentifyConfigRecordTransformations["fieldTransformations"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rDeidentifyConfigRecordTransformationsFieldTransformations dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations
									if _, ok := objval["condition"]; ok {
										if rDeidentifyConfigRecordTransformationsFieldTransformationsCondition, ok := objval["condition"].(map[string]interface{}); ok {
											rDeidentifyConfigRecordTransformationsFieldTransformations.Condition = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition{}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsCondition["expressions"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsCondition["expressions"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions["conditions"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions["conditions"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.Conditions = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions["conditions"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions["conditions"].([]interface{}); ok {
																	for _, o := range s {
																		if objval, ok := o.(map[string]interface{}); ok {
																			var rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions
																			if _, ok := objval["field"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField, ok := objval["field"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Field = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Field.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Field.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Field: expected map[string]interface{}")
																				}
																			}
																			if _, ok := objval["operator"]; ok {
																				if s, ok := objval["operator"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Operator = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Operator: expected string")
																				}
																			}
																			if _, ok := objval["value"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue, ok := objval["value"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["booleanValue"]; ok {
																						if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["booleanValue"].(bool); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.BooleanValue = dcl.Bool(b)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.BooleanValue: expected bool")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["dateValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["dateValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue["day"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue["day"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue.Day = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue.Day: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue["month"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue["month"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue.Month = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue.Month: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue["year"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue["year"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue.Year = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue.Year: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DateValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["dayOfWeekValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["dayOfWeekValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.DayOfWeekValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["floatValue"]; ok {
																						if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["floatValue"].(float64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.FloatValue = dcl.Float64(f)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.FloatValue: expected float64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["integerValue"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["integerValue"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.IntegerValue = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.IntegerValue: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["stringValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["stringValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.StringValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.StringValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["timeValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["timeValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["hours"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["hours"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Hours = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Hours: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["minutes"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["minutes"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Minutes = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Minutes: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["nanos"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["nanos"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Nanos = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Nanos: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["seconds"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue["seconds"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Seconds = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue.Seconds: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimeValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["timestampValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue["timestampValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimestampValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value.TimestampValue: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.Value: expected map[string]interface{}")
																				}
																			}
																			rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.Conditions.Conditions = append(rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.Conditions.Conditions, rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.Conditions.Conditions: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.Conditions: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions["logicalOperator"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions["logicalOperator"].(string); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.LogicalOperator = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnumRef(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions.LogicalOperator: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.Condition.Expressions: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.Condition: expected map[string]interface{}")
										}
									}
									if _, ok := objval["fields"]; ok {
										if s, ok := objval["fields"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rDeidentifyConfigRecordTransformationsFieldTransformationsFields dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFields
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformationsFields.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsFields.Name: expected string")
														}
													}
													rDeidentifyConfigRecordTransformationsFieldTransformations.Fields = append(rDeidentifyConfigRecordTransformationsFieldTransformations.Fields, rDeidentifyConfigRecordTransformationsFieldTransformationsFields)
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.Fields: expected []interface{}")
										}
									}
									if _, ok := objval["infoTypeTransformations"]; ok {
										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations, ok := objval["infoTypeTransformations"].(map[string]interface{}); ok {
											rDeidentifyConfigRecordTransformationsFieldTransformations.InfoTypeTransformations = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations{}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations["transformations"]; ok {
												if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations["transformations"].([]interface{}); ok {
													for _, o := range s {
														if objval, ok := o.(map[string]interface{}); ok {
															var rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations
															if _, ok := objval["infoTypes"]; ok {
																if s, ok := objval["infoTypes"].([]interface{}); ok {
																	for _, o := range s {
																		if objval, ok := o.(map[string]interface{}); ok {
																			var rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes
																			if _, ok := objval["name"]; ok {
																				if s, ok := objval["name"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes.Name = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes.Name: expected string")
																				}
																			}
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.InfoTypes = append(rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.InfoTypes, rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.InfoTypes: expected []interface{}")
																}
															}
															if _, ok := objval["primitiveTransformation"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation, ok := objval["primitiveTransformation"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["bucketingConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["bucketingConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig["buckets"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig["buckets"].([]interface{}); ok {
																					for _, o := range s {
																						if objval, ok := o.(map[string]interface{}); ok {
																							var rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets
																							if _, ok := objval["max"]; ok {
																								if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax, ok := objval["max"].(map[string]interface{}); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax{}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["booleanValue"]; ok {
																										if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["booleanValue"].(bool); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.BooleanValue = dcl.Bool(b)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.BooleanValue: expected bool")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dateValue"]; ok {
																										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dateValue"].(map[string]interface{}); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue{}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["day"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["day"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Day = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Day: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["month"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["month"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Month = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Month: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["year"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["year"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Year = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Year: expected int64")
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue: expected map[string]interface{}")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dayOfWeekValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dayOfWeekValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DayOfWeekValue: expected string")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["floatValue"]; ok {
																										if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["floatValue"].(float64); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.FloatValue = dcl.Float64(f)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.FloatValue: expected float64")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["integerValue"]; ok {
																										if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["integerValue"].(int64); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.IntegerValue = dcl.Int64(i)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.IntegerValue: expected int64")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["stringValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["stringValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.StringValue = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.StringValue: expected string")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timeValue"]; ok {
																										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timeValue"].(map[string]interface{}); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue{}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["hours"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["hours"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Hours = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Hours: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["minutes"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["minutes"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Minutes = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Minutes: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["nanos"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["nanos"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Nanos = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Nanos: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["seconds"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["seconds"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Seconds = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Seconds: expected int64")
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue: expected map[string]interface{}")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timestampValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timestampValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimestampValue = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimestampValue: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Max: expected map[string]interface{}")
																								}
																							}
																							if _, ok := objval["min"]; ok {
																								if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin, ok := objval["min"].(map[string]interface{}); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin{}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["booleanValue"]; ok {
																										if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["booleanValue"].(bool); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.BooleanValue = dcl.Bool(b)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.BooleanValue: expected bool")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dateValue"]; ok {
																										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dateValue"].(map[string]interface{}); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue{}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["day"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["day"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Day = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Day: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["month"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["month"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Month = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Month: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["year"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["year"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Year = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Year: expected int64")
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue: expected map[string]interface{}")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dayOfWeekValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dayOfWeekValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DayOfWeekValue: expected string")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["floatValue"]; ok {
																										if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["floatValue"].(float64); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.FloatValue = dcl.Float64(f)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.FloatValue: expected float64")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["integerValue"]; ok {
																										if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["integerValue"].(int64); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.IntegerValue = dcl.Int64(i)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.IntegerValue: expected int64")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["stringValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["stringValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.StringValue = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.StringValue: expected string")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timeValue"]; ok {
																										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timeValue"].(map[string]interface{}); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue{}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["hours"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["hours"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Hours = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Hours: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["minutes"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["minutes"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Minutes = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Minutes: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["nanos"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["nanos"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Nanos = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Nanos: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["seconds"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["seconds"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Seconds = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Seconds: expected int64")
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue: expected map[string]interface{}")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timestampValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timestampValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimestampValue = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimestampValue: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.Min: expected map[string]interface{}")
																								}
																							}
																							if _, ok := objval["replacementValue"]; ok {
																								if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue, ok := objval["replacementValue"].(map[string]interface{}); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue{}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["booleanValue"]; ok {
																										if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["booleanValue"].(bool); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.BooleanValue = dcl.Bool(b)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.BooleanValue: expected bool")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dateValue"]; ok {
																										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dateValue"].(map[string]interface{}); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue{}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["day"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["day"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Day = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Day: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["month"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["month"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Month = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Month: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["year"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["year"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Year = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Year: expected int64")
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue: expected map[string]interface{}")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dayOfWeekValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dayOfWeekValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DayOfWeekValue: expected string")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["floatValue"]; ok {
																										if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["floatValue"].(float64); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.FloatValue = dcl.Float64(f)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.FloatValue: expected float64")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["integerValue"]; ok {
																										if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["integerValue"].(int64); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.IntegerValue = dcl.Int64(i)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.IntegerValue: expected int64")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["stringValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["stringValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.StringValue = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.StringValue: expected string")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timeValue"]; ok {
																										if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timeValue"].(map[string]interface{}); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue{}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["hours"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["hours"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Hours = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Hours: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["minutes"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["minutes"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Minutes = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Minutes: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["nanos"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["nanos"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Nanos = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Nanos: expected int64")
																												}
																											}
																											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["seconds"]; ok {
																												if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["seconds"].(int64); ok {
																													rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Seconds = dcl.Int64(i)
																												} else {
																													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Seconds: expected int64")
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue: expected map[string]interface{}")
																										}
																									}
																									if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timestampValue"]; ok {
																										if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timestampValue"].(string); ok {
																											rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimestampValue = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimestampValue: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue: expected map[string]interface{}")
																								}
																							}
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig.Buckets = append(rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig.Buckets, rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig.Buckets: expected []interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.BucketingConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["characterMaskConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["characterMaskConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"].([]interface{}); ok {
																					for _, o := range s {
																						if objval, ok := o.(map[string]interface{}); ok {
																							var rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore
																							if _, ok := objval["charactersToSkip"]; ok {
																								if s, ok := objval["charactersToSkip"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CharactersToSkip = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CharactersToSkip: expected string")
																								}
																							}
																							if _, ok := objval["commonCharactersToIgnore"]; ok {
																								if s, ok := objval["commonCharactersToIgnore"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CommonCharactersToIgnore = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CommonCharactersToIgnore: expected string")
																								}
																							}
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore = append(rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore, rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore: expected []interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["maskingCharacter"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["maskingCharacter"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["numberToMask"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["numberToMask"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.NumberToMask = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.NumberToMask: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["reverseOrder"]; ok {
																				if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig["reverseOrder"].(bool); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder: expected bool")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CharacterMaskConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoDeterministicConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoDeterministicConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["context"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["context"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoDeterministicConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoHashConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoHashConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig["cryptoKey"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig["cryptoKey"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["transient"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["transient"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoHashConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoReplaceFfxFpeConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["cryptoReplaceFfxFpeConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["dateShiftConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["dateShiftConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["context"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["context"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.Context: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["cryptoKey"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["cryptoKey"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["transient"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["transient"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"]; ok {
																								if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"].(string); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["lowerBoundDays"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["lowerBoundDays"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.LowerBoundDays = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.LowerBoundDays: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["upperBoundDays"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig["upperBoundDays"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.UpperBoundDays = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig.UpperBoundDays: expected int64")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.DateShiftConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["fixedSizeBucketingConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["fixedSizeBucketingConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"]; ok {
																				if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"].(float64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"]; ok {
																						if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"].(bool); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue = dcl.Bool(b)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue: expected bool")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"]; ok {
																						if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"].(float64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue = dcl.Float64(f)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue: expected float64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["upperBound"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig["upperBound"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"]; ok {
																						if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"].(bool); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue = dcl.Bool(b)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue: expected bool")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"]; ok {
																						if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"].(float64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue = dcl.Float64(f)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue: expected float64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.FixedSizeBucketingConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["redactConfig"]; ok {
																		if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["redactConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.RedactConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig{}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.RedactConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig["newValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig["newValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["booleanValue"]; ok {
																						if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["booleanValue"].(bool); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue = dcl.Bool(b)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue: expected bool")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dateValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dateValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["day"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["day"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["month"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["month"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["year"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["year"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["floatValue"]; ok {
																						if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["floatValue"].(float64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue = dcl.Float64(f)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue: expected float64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["integerValue"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["integerValue"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["stringValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["stringValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timeValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timeValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timestampValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue["timestampValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig.NewValue: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceWithInfoTypeConfig"]; ok {
																		if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["replaceWithInfoTypeConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceWithInfoTypeConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig{}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.ReplaceWithInfoTypeConfig: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["timePartConfig"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation["timePartConfig"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig["partToExtract"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig["partToExtract"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig.PartToExtract = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig.PartToExtract: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation.TimePartConfig: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.PrimitiveTransformation: expected map[string]interface{}")
																}
															}
															rDeidentifyConfigRecordTransformationsFieldTransformations.InfoTypeTransformations.Transformations = append(rDeidentifyConfigRecordTransformationsFieldTransformations.InfoTypeTransformations.Transformations, rDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations)
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.InfoTypeTransformations.Transformations: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.InfoTypeTransformations: expected map[string]interface{}")
										}
									}
									if _, ok := objval["primitiveTransformation"]; ok {
										if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation, ok := objval["primitiveTransformation"].(map[string]interface{}); ok {
											rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation{}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["bucketingConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["bucketingConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.BucketingConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig["buckets"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig["buckets"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets
																	if _, ok := objval["max"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax, ok := objval["max"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["booleanValue"]; ok {
																				if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["booleanValue"].(bool); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.BooleanValue = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.BooleanValue: expected bool")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dateValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dateValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["day"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["day"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Day = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Day: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["month"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["month"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Month = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Month: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["year"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue["year"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Year = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue.Year: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DateValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dayOfWeekValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["dayOfWeekValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.DayOfWeekValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["floatValue"]; ok {
																				if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["floatValue"].(float64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.FloatValue = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.FloatValue: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["integerValue"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["integerValue"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.IntegerValue = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.IntegerValue: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["stringValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["stringValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.StringValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.StringValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timeValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timeValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["hours"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["hours"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Hours = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Hours: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["minutes"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["minutes"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Minutes = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Minutes: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["nanos"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["nanos"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Nanos = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Nanos: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["seconds"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue["seconds"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Seconds = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue.Seconds: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimeValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timestampValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax["timestampValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimestampValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max.TimestampValue: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Max: expected map[string]interface{}")
																		}
																	}
																	if _, ok := objval["min"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin, ok := objval["min"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["booleanValue"]; ok {
																				if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["booleanValue"].(bool); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.BooleanValue = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.BooleanValue: expected bool")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dateValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dateValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["day"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["day"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Day = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Day: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["month"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["month"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Month = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Month: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["year"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue["year"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Year = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue.Year: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DateValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dayOfWeekValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["dayOfWeekValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.DayOfWeekValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["floatValue"]; ok {
																				if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["floatValue"].(float64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.FloatValue = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.FloatValue: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["integerValue"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["integerValue"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.IntegerValue = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.IntegerValue: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["stringValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["stringValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.StringValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.StringValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timeValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timeValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["hours"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["hours"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Hours = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Hours: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["minutes"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["minutes"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Minutes = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Minutes: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["nanos"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["nanos"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Nanos = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Nanos: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["seconds"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue["seconds"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Seconds = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue.Seconds: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimeValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timestampValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin["timestampValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimestampValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min.TimestampValue: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.Min: expected map[string]interface{}")
																		}
																	}
																	if _, ok := objval["replacementValue"]; ok {
																		if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue, ok := objval["replacementValue"].(map[string]interface{}); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue{}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["booleanValue"]; ok {
																				if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["booleanValue"].(bool); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.BooleanValue = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.BooleanValue: expected bool")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dateValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dateValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["day"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["day"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Day = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Day: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["month"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["month"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Month = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Month: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["year"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue["year"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Year = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue.Year: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DateValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dayOfWeekValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["dayOfWeekValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.DayOfWeekValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["floatValue"]; ok {
																				if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["floatValue"].(float64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.FloatValue = dcl.Float64(f)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.FloatValue: expected float64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["integerValue"]; ok {
																				if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["integerValue"].(int64); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.IntegerValue = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.IntegerValue: expected int64")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["stringValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["stringValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.StringValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.StringValue: expected string")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timeValue"]; ok {
																				if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timeValue"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["hours"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["hours"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Hours = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Hours: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["minutes"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["minutes"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Minutes = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Minutes: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["nanos"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["nanos"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Nanos = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Nanos: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["seconds"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue["seconds"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Seconds = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue.Seconds: expected int64")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimeValue: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timestampValue"]; ok {
																				if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue["timestampValue"].(string); ok {
																					rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimestampValue = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue.TimestampValue: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.ReplacementValue: expected map[string]interface{}")
																		}
																	}
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.BucketingConfig.Buckets = append(rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.BucketingConfig.Buckets, rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets)
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.BucketingConfig.Buckets: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.BucketingConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["characterMaskConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["characterMaskConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["charactersToIgnore"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore
																	if _, ok := objval["charactersToSkip"]; ok {
																		if s, ok := objval["charactersToSkip"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CharactersToSkip = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CharactersToSkip: expected string")
																		}
																	}
																	if _, ok := objval["commonCharactersToIgnore"]; ok {
																		if s, ok := objval["commonCharactersToIgnore"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CommonCharactersToIgnore = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.CommonCharactersToIgnore: expected string")
																		}
																	}
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore = append(rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore, rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore)
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.CharactersToIgnore: expected []interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["maskingCharacter"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["maskingCharacter"].(string); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.MaskingCharacter: expected string")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["numberToMask"]; ok {
														if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["numberToMask"].(int64); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.NumberToMask = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.NumberToMask: expected int64")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["reverseOrder"]; ok {
														if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig["reverseOrder"].(bool); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig.ReverseOrder: expected bool")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CharacterMaskConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["cryptoDeterministicConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["cryptoDeterministicConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig["context"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig["context"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext["name"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext["name"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.Context: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.CryptoKey: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig["surrogateInfoType"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType["name"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig.SurrogateInfoType: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoDeterministicConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["cryptoHashConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["cryptoHashConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig["cryptoKey"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig.CryptoKey: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoHashConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["cryptoReplaceFfxFpeConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["cryptoReplaceFfxFpeConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["commonAlphabet"].(string); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnumRef(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CommonAlphabet: expected string")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["context"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext["name"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Context: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CryptoKey: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["customAlphabet"].(string); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.CustomAlphabet: expected string")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"]; ok {
														if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["radix"].(int64); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.Radix: expected int64")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig["surrogateInfoType"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType["name"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig.SurrogateInfoType: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.CryptoReplaceFfxFpeConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["dateShiftConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["dateShiftConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["context"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["context"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.Context = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext["name"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext["name"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.Context.Name = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.Context.Name: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.Context: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["cryptoKey"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["cryptoKey"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["kmsWrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["cryptoKeyName"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.CryptoKeyName: expected string")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped["wrappedKey"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped.WrappedKey: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.KmsWrapped: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["transient"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["transient"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient["name"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Transient: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey["unwrapped"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"]; ok {
																		if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped["key"].(string); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped.Key: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey.Unwrapped: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.CryptoKey: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["lowerBoundDays"]; ok {
														if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["lowerBoundDays"].(int64); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.LowerBoundDays = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.LowerBoundDays: expected int64")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["upperBoundDays"]; ok {
														if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig["upperBoundDays"].(int64); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.UpperBoundDays = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig.UpperBoundDays: expected int64")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.DateShiftConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["fixedSizeBucketingConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["fixedSizeBucketingConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"]; ok {
														if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig["bucketSize"].(float64); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.BucketSize: expected float64")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig["lowerBound"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"]; ok {
																if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["booleanValue"].(bool); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.BooleanValue: expected bool")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dateValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["day"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Day: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["month"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Month: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue["year"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue.Year: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DateValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["dayOfWeekValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.DayOfWeekValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"]; ok {
																if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["floatValue"].(float64); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.FloatValue: expected float64")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"]; ok {
																if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["integerValue"].(int64); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.IntegerValue: expected int64")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["stringValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.StringValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timeValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["hours"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Hours: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["minutes"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Minutes: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["nanos"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Nanos: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue["seconds"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue.Seconds: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimeValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound["timestampValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound.TimestampValue: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.LowerBound: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig["upperBound"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig["upperBound"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"]; ok {
																if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["booleanValue"].(bool); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.BooleanValue: expected bool")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dateValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["day"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Day: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["month"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Month: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue["year"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue.Year: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DateValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["dayOfWeekValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.DayOfWeekValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"]; ok {
																if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["floatValue"].(float64); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.FloatValue: expected float64")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"]; ok {
																if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["integerValue"].(int64); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.IntegerValue: expected int64")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["stringValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.StringValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timeValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["hours"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Hours: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["minutes"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Minutes: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["nanos"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Nanos: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue["seconds"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue.Seconds: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimeValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound["timestampValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound.TimestampValue: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig.UpperBound: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.FixedSizeBucketingConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["redactConfig"]; ok {
												if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["redactConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.RedactConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig{}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.RedactConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["replaceConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["replaceConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig["newValue"]; ok {
														if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig["newValue"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue{}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["booleanValue"]; ok {
																if b, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["booleanValue"].(bool); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.BooleanValue: expected bool")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["dateValue"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["dateValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["day"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["day"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Day: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["month"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["month"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Month: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["year"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue["year"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue.Year: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DateValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["dayOfWeekValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.DayOfWeekValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["floatValue"]; ok {
																if f, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["floatValue"].(float64); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.FloatValue: expected float64")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["integerValue"]; ok {
																if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["integerValue"].(int64); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.IntegerValue: expected int64")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["stringValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["stringValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.StringValue: expected string")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["timeValue"]; ok {
																if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["timeValue"].(map[string]interface{}); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue{}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["hours"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Hours: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["minutes"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Minutes: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["nanos"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Nanos: expected int64")
																		}
																	}
																	if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"]; ok {
																		if i, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue["seconds"].(int64); ok {
																			rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue.Seconds: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimeValue: expected map[string]interface{}")
																}
															}
															if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["timestampValue"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue["timestampValue"].(string); ok {
																	rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue.TimestampValue: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig.NewValue: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["replaceWithInfoTypeConfig"]; ok {
												if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["replaceWithInfoTypeConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceWithInfoTypeConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig{}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.ReplaceWithInfoTypeConfig: expected map[string]interface{}")
												}
											}
											if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["timePartConfig"]; ok {
												if rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation["timePartConfig"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.TimePartConfig = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig{}
													if _, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig["partToExtract"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig["partToExtract"].(string); ok {
															rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.TimePartConfig.PartToExtract = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnumRef(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.TimePartConfig.PartToExtract: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation.TimePartConfig: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsFieldTransformations.PrimitiveTransformation: expected map[string]interface{}")
										}
									}
									r.DeidentifyConfig.RecordTransformations.FieldTransformations = append(r.DeidentifyConfig.RecordTransformations.FieldTransformations, rDeidentifyConfigRecordTransformationsFieldTransformations)
								}
							}
						} else {
							return nil, fmt.Errorf("r.DeidentifyConfig.RecordTransformations.FieldTransformations: expected []interface{}")
						}
					}
					if _, ok := rDeidentifyConfigRecordTransformations["recordSuppressions"]; ok {
						if s, ok := rDeidentifyConfigRecordTransformations["recordSuppressions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rDeidentifyConfigRecordTransformationsRecordSuppressions dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressions
									if _, ok := objval["condition"]; ok {
										if rDeidentifyConfigRecordTransformationsRecordSuppressionsCondition, ok := objval["condition"].(map[string]interface{}); ok {
											rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition{}
											if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsCondition["expressions"]; ok {
												if rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsCondition["expressions"].(map[string]interface{}); ok {
													rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions{}
													if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions["conditions"]; ok {
														if rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions["conditions"].(map[string]interface{}); ok {
															rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.Conditions = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions{}
															if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions["conditions"]; ok {
																if s, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions["conditions"].([]interface{}); ok {
																	for _, o := range s {
																		if objval, ok := o.(map[string]interface{}); ok {
																			var rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions
																			if _, ok := objval["field"]; ok {
																				if rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField, ok := objval["field"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Field = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField{}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField["name"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField["name"].(string); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Field.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Field.Name: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Field: expected map[string]interface{}")
																				}
																			}
																			if _, ok := objval["operator"]; ok {
																				if s, ok := objval["operator"].(string); ok {
																					rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Operator = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Operator: expected string")
																				}
																			}
																			if _, ok := objval["value"]; ok {
																				if rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue, ok := objval["value"].(map[string]interface{}); ok {
																					rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue{}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["booleanValue"]; ok {
																						if b, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["booleanValue"].(bool); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.BooleanValue = dcl.Bool(b)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.BooleanValue: expected bool")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["dateValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["dateValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue["day"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue["day"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue.Day = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue.Day: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue["month"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue["month"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue.Month = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue.Month: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue["year"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue["year"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue.Year = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue.Year: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DateValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["dayOfWeekValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["dayOfWeekValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DayOfWeekValue = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.DayOfWeekValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["floatValue"]; ok {
																						if f, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["floatValue"].(float64); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.FloatValue = dcl.Float64(f)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.FloatValue: expected float64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["integerValue"]; ok {
																						if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["integerValue"].(int64); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.IntegerValue = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.IntegerValue: expected int64")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["stringValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["stringValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.StringValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.StringValue: expected string")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["timeValue"]; ok {
																						if rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["timeValue"].(map[string]interface{}); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue = &dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue{}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["hours"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["hours"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Hours = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Hours: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["minutes"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["minutes"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Minutes = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Minutes: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["nanos"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["nanos"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Nanos = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Nanos: expected int64")
																								}
																							}
																							if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["seconds"]; ok {
																								if i, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue["seconds"].(int64); ok {
																									rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Seconds = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue.Seconds: expected int64")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimeValue: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["timestampValue"]; ok {
																						if s, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue["timestampValue"].(string); ok {
																							rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimestampValue = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value.TimestampValue: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.Value: expected map[string]interface{}")
																				}
																			}
																			rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.Conditions.Conditions = append(rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.Conditions.Conditions, rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.Conditions.Conditions: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.Conditions: expected map[string]interface{}")
														}
													}
													if _, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions["logicalOperator"]; ok {
														if s, ok := rDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions["logicalOperator"].(string); ok {
															rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.LogicalOperator = dclService.DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnumRef(s)
														} else {
															return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions.LogicalOperator: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition.Expressions: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rDeidentifyConfigRecordTransformationsRecordSuppressions.Condition: expected map[string]interface{}")
										}
									}
									r.DeidentifyConfig.RecordTransformations.RecordSuppressions = append(r.DeidentifyConfig.RecordTransformations.RecordSuppressions, rDeidentifyConfigRecordTransformationsRecordSuppressions)
								}
							}
						} else {
							return nil, fmt.Errorf("r.DeidentifyConfig.RecordTransformations.RecordSuppressions: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DeidentifyConfig.RecordTransformations: expected map[string]interface{}")
				}
			}
			if _, ok := rDeidentifyConfig["transformationErrorHandling"]; ok {
				if rDeidentifyConfigTransformationErrorHandling, ok := rDeidentifyConfig["transformationErrorHandling"].(map[string]interface{}); ok {
					r.DeidentifyConfig.TransformationErrorHandling = &dclService.DeidentifyTemplateDeidentifyConfigTransformationErrorHandling{}
					if _, ok := rDeidentifyConfigTransformationErrorHandling["leaveUntransformed"]; ok {
						if _, ok := rDeidentifyConfigTransformationErrorHandling["leaveUntransformed"].(map[string]interface{}); ok {
							r.DeidentifyConfig.TransformationErrorHandling.LeaveUntransformed = &dclService.DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed{}
						} else {
							return nil, fmt.Errorf("r.DeidentifyConfig.TransformationErrorHandling.LeaveUntransformed: expected map[string]interface{}")
						}
					}
					if _, ok := rDeidentifyConfigTransformationErrorHandling["throwError"]; ok {
						if _, ok := rDeidentifyConfigTransformationErrorHandling["throwError"].(map[string]interface{}); ok {
							r.DeidentifyConfig.TransformationErrorHandling.ThrowError = &dclService.DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError{}
						} else {
							return nil, fmt.Errorf("r.DeidentifyConfig.TransformationErrorHandling.ThrowError: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DeidentifyConfig.TransformationErrorHandling: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DeidentifyConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["locationId"]; ok {
		if s, ok := u.Object["locationId"].(string); ok {
			r.LocationId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LocationId: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetDeidentifyTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeidentifyTemplate(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDeidentifyTemplate(ctx, r)
	if err != nil {
		return nil, err
	}
	return DeidentifyTemplateToUnstructured(r), nil
}

func ListDeidentifyTemplate(ctx context.Context, config *dcl.Config, location string, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDeidentifyTemplate(ctx, location, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DeidentifyTemplateToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyDeidentifyTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeidentifyTemplate(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDeidentifyTemplate(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDeidentifyTemplate(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DeidentifyTemplateToUnstructured(r), nil
}

func DeidentifyTemplateHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeidentifyTemplate(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDeidentifyTemplate(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDeidentifyTemplate(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDeidentifyTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeidentifyTemplate(u)
	if err != nil {
		return err
	}
	return c.DeleteDeidentifyTemplate(ctx, r)
}

func DeidentifyTemplateID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDeidentifyTemplate(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *DeidentifyTemplate) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dlp",
		"DeidentifyTemplate",
		"beta",
	}
}

func (r *DeidentifyTemplate) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeidentifyTemplate) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeidentifyTemplate) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *DeidentifyTemplate) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeidentifyTemplate) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeidentifyTemplate) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeidentifyTemplate) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDeidentifyTemplate(ctx, config, resource)
}

func (r *DeidentifyTemplate) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDeidentifyTemplate(ctx, config, resource, opts...)
}

func (r *DeidentifyTemplate) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DeidentifyTemplateHasDiff(ctx, config, resource, opts...)
}

func (r *DeidentifyTemplate) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDeidentifyTemplate(ctx, config, resource)
}

func (r *DeidentifyTemplate) ID(resource *unstructured.Resource) (string, error) {
	return DeidentifyTemplateID(resource)
}

func init() {
	unstructured.Register(&DeidentifyTemplate{})
}
