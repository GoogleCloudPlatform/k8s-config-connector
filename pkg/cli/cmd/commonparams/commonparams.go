// Copyright 2022 Google LLC
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

package commonparams

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/valutil"

	"github.com/spf13/cobra"
)

type IAMFormatOption string
type ResourceFormatOption string

const (
	IAMFormatParamName               = "iam-format"
	FilterDeletedIAMMembersParamName = "filter-deleted-iam-members"
	OAuth2TokenParamName             = "oauth2-token"
	OutputParamName                  = "output"
	ResourceFormatParamName          = "resource-format"

	PartialPolicyFormatOption   = "partialpolicy"
	PolicyIAMFormatOption       = "policy"
	PolicyMemberIAMFormatOption = "policymember"
	NoneIAMFormatOption         = "none"

	KRMResourceFormatOption = outputsink.KRMResourceFormat
	HCLResourceFormatOption = outputsink.HCLResourceFormat

	IAMFormatDefault               = PolicyIAMFormatOption
	FilterDeletedIAMMembersDefault = false
	OAuth2TokenDefault             = ""
	OutputDefault                  = ""
	ResourceFormatDefault          = KRMResourceFormatOption

	OAuth2TokenUsage = "an optional OAuth 2.0 access token to be used as the identity for communication with GCP services, can be obtained with 'gcloud auth print-access-token'"
	OutputUsage      = "an optional output file path, disables standard output, when a file the result will contain all of the command output, when a directory, the directory will contain a new file for each resource in the output"
)

var (
	IAMFormatUsage               = fmt.Sprintf("specify the IAM resource format or disable IAM output, options are '%v', '%v', '%v', or '%v'", PartialPolicyFormatOption, PolicyIAMFormatOption, PolicyMemberIAMFormatOption, NoneIAMFormatOption)
	FilterDeletedIAMMembersUsage = fmt.Sprintf("specify whether to filter out deleted IAM members, options are '%v' or '%v', (default: '%v')", true, false, FilterDeletedIAMMembersDefault)
	ResourceFormatUsage          = fmt.Sprintf("specify the format of the outputted resources, options are '%v' or '%v' (default: '%v')", KRMResourceFormatOption, HCLResourceFormatOption, ResourceFormatDefault)
)

func AddOAuth2TokenParam(cmd *cobra.Command, value *string) {
	cmd.Flags().StringVar(value, OAuth2TokenParamName, OAuth2TokenDefault, OAuth2TokenUsage)
}

func AddIAMFormatParam(cmd *cobra.Command, value *string) {
	cmd.Flags().StringVar(value, IAMFormatParamName, IAMFormatDefault, IAMFormatUsage)
}

func AddFilterDeletedIAMMembersParam(cmd *cobra.Command, value *bool) {
	cmd.Flags().BoolVar(value, FilterDeletedIAMMembersParamName, FilterDeletedIAMMembersDefault, FilterDeletedIAMMembersUsage)
}

func ValidateIAMFormat(value string) error {
	iamFormatOptions := []string{PartialPolicyFormatOption, PolicyIAMFormatOption, PolicyMemberIAMFormatOption, NoneIAMFormatOption}
	if valutil.IsDefaultValue(value) {
		return fmt.Errorf("invalid empty value for %v: must be one of {%v}", IAMFormatParamName, strings.Join(iamFormatOptions, ", "))
	}
	for _, o := range iamFormatOptions {
		if value == o {
			return nil
		}
	}
	return fmt.Errorf("invalid %v value of '%v': must be one of {%v}", IAMFormatParamName, value, strings.Join(iamFormatOptions, ", "))
}

// Convert from the IAMPolicyParam supplied on the command line to the stream type
func IAMFormatParamToStreamIAMFormat(iamFormatParam string) (stream.IAMFormat, error) {
	switch iamFormatParam {
	case PartialPolicyFormatOption:
		return stream.IAMFormatPartialPolicy, nil
	case PolicyIAMFormatOption:
		return stream.IAMFormatPolicy, nil
	case PolicyMemberIAMFormatOption:
		return stream.IAMFormatPolicyMember, nil
	}
	return "", fmt.Errorf("invalid policy format option '%v'", iamFormatParam)
}

func AddOutputParam(cmd *cobra.Command, value *string) {
	cmd.Flags().StringVar(value, OutputParamName, OutputDefault, OutputUsage)

}

func AddResourceFormatParam(cmd *cobra.Command, value *string) {
	cmd.Flags().StringVar(value, ResourceFormatParamName, ResourceFormatDefault, ResourceFormatUsage)
	if err := cmd.Flags().MarkHidden(ResourceFormatParamName); err != nil {
		panic(err)
	}
}

func ValidateResourceFormat(resourceFormat, iamFormat string) error {
	if err := validateResourceFormatValue(resourceFormat); err != nil {
		return err
	}

	return validateResourceFormatMutualExclusivity(resourceFormat, iamFormat)
}

func validateResourceFormatValue(value string) error {
	resourceFormatOptions := []string{KRMResourceFormatOption, HCLResourceFormatOption}
	if valutil.IsDefaultValue(value) {
		return fmt.Errorf("invalid empty value for %v: must be one of {%v}", ResourceFormatParamName, strings.Join(resourceFormatOptions, ", "))
	}
	for _, o := range resourceFormatOptions {
		if value == o {
			return nil
		}
	}
	return fmt.Errorf("invalid %v value of '%v': must be one of {%v}", ResourceFormatParamName, value, strings.Join(resourceFormatOptions, ", "))
}

func validateResourceFormatMutualExclusivity(resourceFormat, iamFormat string) error {
	switch resourceFormat {
	case HCLResourceFormatOption:
		return validateHCLResourceFormatMutualExclusivity(iamFormat)
	case KRMResourceFormatOption:
		// all parameters can be used with the KRM format
		return nil
	default:
		return fmt.Errorf("unhandled resource format %v", resourceFormat)
	}
}

func validateHCLResourceFormatMutualExclusivity(iamFormat string) error {
	if iamFormat == NoneIAMFormatOption {
		return nil
	}
	return fmt.Errorf("unsupported value of '%v' for flag '%v': when '%v' is '%v' the '%v' flag must have a value of '%v'",
		iamFormat, IAMFormatParamName, ResourceFormatParamName, HCLResourceFormatOption, IAMFormatParamName, NoneIAMFormatOption)
}
