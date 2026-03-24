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

package sql

import (
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func SQLUserKRMToGCP(in *krm.SQLUser) (*api.User, error) {
	if in == nil {
		return nil, fmt.Errorf("cannot convert nil KRM SQLUser to GCP User")
	}

	out := &api.User{
		Name:     direct.ValueOf(in.Spec.ResourceID),
		Host:     direct.ValueOf(in.Spec.Host),
		Kind:     "sql#user",
		Type:     direct.ValueOf(in.Spec.Type),
		Password: SQLUserPasswordKRMToGCP(in.Spec.Password),
	}

	if in.Spec.PasswordPolicy != nil {
		out.PasswordPolicy = SQLUserPasswordPolicyKRMToGCP(in.Spec.PasswordPolicy)
	}

	/*NOTYET
	out.DatabaseRoles = in.Spec.DatabaseRoles
	*/

	return out, nil
}

func SQLUserPasswordKRMToGCP(in *krm.SQLUserPassword) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func SQLUserPasswordPolicyKRMToGCP(in *krm.SQLUserPasswordPolicy) *api.UserPasswordValidationPolicy {
	if in == nil {
		return nil
	}
	out := &api.UserPasswordValidationPolicy{
		AllowedFailedAttempts:      direct.ValueOf(in.AllowedFailedAttempts),
		EnableFailedAttemptsCheck:  direct.ValueOf(in.EnableFailedAttemptsCheck),
		EnablePasswordVerification: direct.ValueOf(in.EnablePasswordVerification),
		PasswordExpirationDuration: direct.ValueOf(in.PasswordExpirationDuration),
	}
	return out
}

func SQLUserGCPToKRM(in *api.User) *krm.SQLUserSpec {
	if in == nil {
		return nil
	}

	out := &krm.SQLUserSpec{
		Host:       direct.PtrTo(in.Host),
		ResourceID: direct.PtrTo(in.Name),
		Type:       direct.PtrTo(in.Type),
	}

	if in.PasswordPolicy != nil {
		out.PasswordPolicy = SQLUserPasswordPolicyGCPToKRM(in.PasswordPolicy)
	}

	/*NOTYET
	if len(in.DatabaseRoles) > 0 {
		out.DatabaseRoles = in.DatabaseRoles
	}
	*/

	return out
}

func SQLUserPasswordPolicyGCPToKRM(in *api.UserPasswordValidationPolicy) *krm.SQLUserPasswordPolicy {
	if in == nil {
		return nil
	}

	out := &krm.SQLUserPasswordPolicy{}

	if in.AllowedFailedAttempts != 0 {
		out.AllowedFailedAttempts = direct.PtrTo(in.AllowedFailedAttempts)
	}
	if in.EnableFailedAttemptsCheck {
		out.EnableFailedAttemptsCheck = direct.PtrTo(in.EnableFailedAttemptsCheck)
	}
	if in.EnablePasswordVerification {
		out.EnablePasswordVerification = direct.PtrTo(in.EnablePasswordVerification)
	}
	if in.PasswordExpirationDuration != "" {
		out.PasswordExpirationDuration = direct.PtrTo(in.PasswordExpirationDuration)
	}
	if in.Status != nil {
		out.Status = []krm.SQLUserPasswordStatus{
			{
				Locked:                 direct.PtrTo(in.Status.Locked),
				PasswordExpirationTime: direct.PtrTo(in.Status.PasswordExpirationTime),
			},
		}
	}

	return out
}

func SQLUserStatusGCPToKRM(in *api.User) *krm.SQLUserSqlServerUserDetailsStatus {
	if in == nil || in.SqlserverUserDetails == nil {
		return nil
	}
	return &krm.SQLUserSqlServerUserDetailsStatus{
		Disabled:    direct.PtrTo(in.SqlserverUserDetails.Disabled),
		ServerRoles: in.SqlserverUserDetails.ServerRoles,
	}
}
