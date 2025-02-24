// Copyright 2022 Google LLC. All Rights Reserved.
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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/alpha/osconfig_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
)

// PatchDeploymentServer implements the gRPC interface for PatchDeployment.
type PatchDeploymentServer struct{}

// ProtoToPatchDeploymentPatchConfigRebootConfigEnum converts a PatchDeploymentPatchConfigRebootConfigEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum) *alpha.PatchDeploymentPatchConfigRebootConfigEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigRebootConfigEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigAptTypeEnum converts a PatchDeploymentPatchConfigAptTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum) *alpha.PatchDeploymentPatchConfigAptTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigAptTypeEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum converts a PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum) *alpha.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum) *alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum) *alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum) *alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(e alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum) *alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(n[len("OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleFrequencyEnum converts a PatchDeploymentRecurringScheduleFrequencyEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum(e alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum) *alpha.PatchDeploymentRecurringScheduleFrequencyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentRecurringScheduleFrequencyEnum(n[len("OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum converts a PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(e alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum) *alpha.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(n[len("OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(e alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum) *alpha.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(n[len("OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRolloutModeEnum converts a PatchDeploymentRolloutModeEnum enum from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRolloutModeEnum(e alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum) *alpha.PatchDeploymentRolloutModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum_name[int32(e)]; ok {
		e := alpha.PatchDeploymentRolloutModeEnum(n[len("OsconfigAlphaPatchDeploymentRolloutModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentInstanceFilter converts a PatchDeploymentInstanceFilter object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentInstanceFilter(p *alphapb.OsconfigAlphaPatchDeploymentInstanceFilter) *alpha.PatchDeploymentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentInstanceFilter{
		All: dcl.Bool(p.GetAll()),
	}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigAlphaPatchDeploymentInstanceFilterGroupLabels(r))
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, r)
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, r)
	}
	for _, r := range p.GetInstanceNamePrefixes() {
		obj.InstanceNamePrefixes = append(obj.InstanceNamePrefixes, r)
	}
	return obj
}

// ProtoToPatchDeploymentInstanceFilterGroupLabels converts a PatchDeploymentInstanceFilterGroupLabels object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentInstanceFilterGroupLabels(p *alphapb.OsconfigAlphaPatchDeploymentInstanceFilterGroupLabels) *alpha.PatchDeploymentInstanceFilterGroupLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentInstanceFilterGroupLabels{}
	return obj
}

// ProtoToPatchDeploymentPatchConfig converts a PatchDeploymentPatchConfig object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfig(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfig) *alpha.PatchDeploymentPatchConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfig{
		RebootConfig:  ProtoToOsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum(p.GetRebootConfig()),
		Apt:           ProtoToOsconfigAlphaPatchDeploymentPatchConfigApt(p.GetApt()),
		Yum:           ProtoToOsconfigAlphaPatchDeploymentPatchConfigYum(p.GetYum()),
		Goo:           ProtoToOsconfigAlphaPatchDeploymentPatchConfigGoo(p.GetGoo()),
		Zypper:        ProtoToOsconfigAlphaPatchDeploymentPatchConfigZypper(p.GetZypper()),
		WindowsUpdate: ProtoToOsconfigAlphaPatchDeploymentPatchConfigWindowsUpdate(p.GetWindowsUpdate()),
		PreStep:       ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStep(p.GetPreStep()),
		PostStep:      ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStep(p.GetPostStep()),
		RetryStrategy: ProtoToOsconfigAlphaPatchDeploymentPatchConfigRetryStrategy(p.GetRetryStrategy()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigApt converts a PatchDeploymentPatchConfigApt object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigApt(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigApt) *alpha.PatchDeploymentPatchConfigApt {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigApt{
		Type: ProtoToOsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum(p.GetType()),
	}
	for _, r := range p.GetExcludes() {
		obj.Excludes = append(obj.Excludes, r)
	}
	for _, r := range p.GetExclusivePackages() {
		obj.ExclusivePackages = append(obj.ExclusivePackages, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigYum converts a PatchDeploymentPatchConfigYum object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigYum(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigYum) *alpha.PatchDeploymentPatchConfigYum {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigYum{
		Security: dcl.Bool(p.GetSecurity()),
		Minimal:  dcl.Bool(p.GetMinimal()),
	}
	for _, r := range p.GetExcludes() {
		obj.Excludes = append(obj.Excludes, r)
	}
	for _, r := range p.GetExclusivePackages() {
		obj.ExclusivePackages = append(obj.ExclusivePackages, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigGoo converts a PatchDeploymentPatchConfigGoo object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigGoo(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigGoo) *alpha.PatchDeploymentPatchConfigGoo {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigGoo{}
	return obj
}

// ProtoToPatchDeploymentPatchConfigZypper converts a PatchDeploymentPatchConfigZypper object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigZypper(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigZypper) *alpha.PatchDeploymentPatchConfigZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigZypper{
		WithOptional: dcl.Bool(p.GetWithOptional()),
		WithUpdate:   dcl.Bool(p.GetWithUpdate()),
	}
	for _, r := range p.GetCategories() {
		obj.Categories = append(obj.Categories, r)
	}
	for _, r := range p.GetSeverities() {
		obj.Severities = append(obj.Severities, r)
	}
	for _, r := range p.GetExcludes() {
		obj.Excludes = append(obj.Excludes, r)
	}
	for _, r := range p.GetExclusivePatches() {
		obj.ExclusivePatches = append(obj.ExclusivePatches, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigWindowsUpdate converts a PatchDeploymentPatchConfigWindowsUpdate object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigWindowsUpdate(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdate) *alpha.PatchDeploymentPatchConfigWindowsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigWindowsUpdate{}
	for _, r := range p.GetClassifications() {
		obj.Classifications = append(obj.Classifications, *ProtoToOsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(r))
	}
	for _, r := range p.GetExcludes() {
		obj.Excludes = append(obj.Excludes, r)
	}
	for _, r := range p.GetExclusivePatches() {
		obj.ExclusivePatches = append(obj.ExclusivePatches, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStep converts a PatchDeploymentPatchConfigPreStep object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStep(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStep) *alpha.PatchDeploymentPatchConfigPreStep {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPreStep{
		LinuxExecStepConfig:   ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig(p.GetLinuxExecStepConfig()),
		WindowsExecStepConfig: ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig(p.GetWindowsExecStepConfig()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfig converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfig object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig) *alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) *alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfig converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfig object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig) *alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) *alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStep converts a PatchDeploymentPatchConfigPostStep object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStep(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStep) *alpha.PatchDeploymentPatchConfigPostStep {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPostStep{
		LinuxExecStepConfig:   ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig(p.GetLinuxExecStepConfig()),
		WindowsExecStepConfig: ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig(p.GetWindowsExecStepConfig()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfig converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfig object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig) *alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) *alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfig converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfig object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig) *alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject) *alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigRetryStrategy converts a PatchDeploymentPatchConfigRetryStrategy object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentPatchConfigRetryStrategy(p *alphapb.OsconfigAlphaPatchDeploymentPatchConfigRetryStrategy) *alpha.PatchDeploymentPatchConfigRetryStrategy {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentPatchConfigRetryStrategy{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToPatchDeploymentOneTimeSchedule converts a PatchDeploymentOneTimeSchedule object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentOneTimeSchedule(p *alphapb.OsconfigAlphaPatchDeploymentOneTimeSchedule) *alpha.PatchDeploymentOneTimeSchedule {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentOneTimeSchedule{
		ExecuteTime: dcl.StringOrNil(p.GetExecuteTime()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringSchedule converts a PatchDeploymentRecurringSchedule object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringSchedule(p *alphapb.OsconfigAlphaPatchDeploymentRecurringSchedule) *alpha.PatchDeploymentRecurringSchedule {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRecurringSchedule{
		TimeZone:        ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleTimeZone(p.GetTimeZone()),
		StartTime:       dcl.StringOrNil(p.GetStartTime()),
		EndTime:         dcl.StringOrNil(p.GetEndTime()),
		TimeOfDay:       ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDay(p.GetTimeOfDay()),
		Frequency:       ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum(p.GetFrequency()),
		Weekly:          ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleWeekly(p.GetWeekly()),
		Monthly:         ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleMonthly(p.GetMonthly()),
		LastExecuteTime: dcl.StringOrNil(p.GetLastExecuteTime()),
		NextExecuteTime: dcl.StringOrNil(p.GetNextExecuteTime()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleTimeZone converts a PatchDeploymentRecurringScheduleTimeZone object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleTimeZone(p *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleTimeZone) *alpha.PatchDeploymentRecurringScheduleTimeZone {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRecurringScheduleTimeZone{
		Id:      dcl.StringOrNil(p.GetId()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleTimeOfDay converts a PatchDeploymentRecurringScheduleTimeOfDay object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDay(p *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDay) *alpha.PatchDeploymentRecurringScheduleTimeOfDay {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRecurringScheduleTimeOfDay{
		Hours:   dcl.Int64OrNil(p.GetHours()),
		Minutes: dcl.Int64OrNil(p.GetMinutes()),
		Seconds: dcl.Int64OrNil(p.GetSeconds()),
		Nanos:   dcl.Int64OrNil(p.GetNanos()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleWeekly converts a PatchDeploymentRecurringScheduleWeekly object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleWeekly(p *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeekly) *alpha.PatchDeploymentRecurringScheduleWeekly {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRecurringScheduleWeekly{
		DayOfWeek: ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(p.GetDayOfWeek()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleMonthly converts a PatchDeploymentRecurringScheduleMonthly object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleMonthly(p *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthly) *alpha.PatchDeploymentRecurringScheduleMonthly {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRecurringScheduleMonthly{
		WeekDayOfMonth: ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(p.GetWeekDayOfMonth()),
		MonthDay:       dcl.Int64OrNil(p.GetMonthDay()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(p *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) *alpha.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{
		WeekOrdinal: dcl.Int64OrNil(p.GetWeekOrdinal()),
		DayOfWeek:   ProtoToOsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(p.GetDayOfWeek()),
	}
	return obj
}

// ProtoToPatchDeploymentRollout converts a PatchDeploymentRollout object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRollout(p *alphapb.OsconfigAlphaPatchDeploymentRollout) *alpha.PatchDeploymentRollout {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRollout{
		Mode:             ProtoToOsconfigAlphaPatchDeploymentRolloutModeEnum(p.GetMode()),
		DisruptionBudget: ProtoToOsconfigAlphaPatchDeploymentRolloutDisruptionBudget(p.GetDisruptionBudget()),
	}
	return obj
}

// ProtoToPatchDeploymentRolloutDisruptionBudget converts a PatchDeploymentRolloutDisruptionBudget object from its proto representation.
func ProtoToOsconfigAlphaPatchDeploymentRolloutDisruptionBudget(p *alphapb.OsconfigAlphaPatchDeploymentRolloutDisruptionBudget) *alpha.PatchDeploymentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &alpha.PatchDeploymentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.GetFixed()),
		Percent: dcl.Int64OrNil(p.GetPercent()),
	}
	return obj
}

// ProtoToPatchDeployment converts a PatchDeployment resource from its proto representation.
func ProtoToPatchDeployment(p *alphapb.OsconfigAlphaPatchDeployment) *alpha.PatchDeployment {
	obj := &alpha.PatchDeployment{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		InstanceFilter:    ProtoToOsconfigAlphaPatchDeploymentInstanceFilter(p.GetInstanceFilter()),
		PatchConfig:       ProtoToOsconfigAlphaPatchDeploymentPatchConfig(p.GetPatchConfig()),
		Duration:          dcl.StringOrNil(p.GetDuration()),
		OneTimeSchedule:   ProtoToOsconfigAlphaPatchDeploymentOneTimeSchedule(p.GetOneTimeSchedule()),
		RecurringSchedule: ProtoToOsconfigAlphaPatchDeploymentRecurringSchedule(p.GetRecurringSchedule()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		LastExecuteTime:   dcl.StringOrNil(p.GetLastExecuteTime()),
		Rollout:           ProtoToOsconfigAlphaPatchDeploymentRollout(p.GetRollout()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// PatchDeploymentPatchConfigRebootConfigEnumToProto converts a PatchDeploymentPatchConfigRebootConfigEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnumToProto(e *alpha.PatchDeploymentPatchConfigRebootConfigEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum_value["PatchDeploymentPatchConfigRebootConfigEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnum(0)
}

// PatchDeploymentPatchConfigAptTypeEnumToProto converts a PatchDeploymentPatchConfigAptTypeEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnumToProto(e *alpha.PatchDeploymentPatchConfigAptTypeEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum_value["PatchDeploymentPatchConfigAptTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnum(0)
}

// PatchDeploymentPatchConfigWindowsUpdateClassificationsEnumToProto converts a PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnumToProto(e *alpha.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_value["PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(0)
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto(e *alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto(e *alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto(e *alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto(e *alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum) alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentRecurringScheduleFrequencyEnumToProto converts a PatchDeploymentRecurringScheduleFrequencyEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnumToProto(e *alpha.PatchDeploymentRecurringScheduleFrequencyEnum) alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum_value["PatchDeploymentRecurringScheduleFrequencyEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnum(0)
}

// PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto converts a PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto(e *alpha.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum) alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum_value["PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(0)
}

// PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto(e *alpha.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum) alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum_value["PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(0)
}

// PatchDeploymentRolloutModeEnumToProto converts a PatchDeploymentRolloutModeEnum enum to its proto representation.
func OsconfigAlphaPatchDeploymentRolloutModeEnumToProto(e *alpha.PatchDeploymentRolloutModeEnum) alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum_value["PatchDeploymentRolloutModeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum(v)
	}
	return alphapb.OsconfigAlphaPatchDeploymentRolloutModeEnum(0)
}

// PatchDeploymentInstanceFilterToProto converts a PatchDeploymentInstanceFilter object to its proto representation.
func OsconfigAlphaPatchDeploymentInstanceFilterToProto(o *alpha.PatchDeploymentInstanceFilter) *alphapb.OsconfigAlphaPatchDeploymentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentInstanceFilter{}
	p.SetAll(dcl.ValueOrEmptyBool(o.All))
	sGroupLabels := make([]*alphapb.OsconfigAlphaPatchDeploymentInstanceFilterGroupLabels, len(o.GroupLabels))
	for i, r := range o.GroupLabels {
		sGroupLabels[i] = OsconfigAlphaPatchDeploymentInstanceFilterGroupLabelsToProto(&r)
	}
	p.SetGroupLabels(sGroupLabels)
	sZones := make([]string, len(o.Zones))
	for i, r := range o.Zones {
		sZones[i] = r
	}
	p.SetZones(sZones)
	sInstances := make([]string, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = r
	}
	p.SetInstances(sInstances)
	sInstanceNamePrefixes := make([]string, len(o.InstanceNamePrefixes))
	for i, r := range o.InstanceNamePrefixes {
		sInstanceNamePrefixes[i] = r
	}
	p.SetInstanceNamePrefixes(sInstanceNamePrefixes)
	return p
}

// PatchDeploymentInstanceFilterGroupLabelsToProto converts a PatchDeploymentInstanceFilterGroupLabels object to its proto representation.
func OsconfigAlphaPatchDeploymentInstanceFilterGroupLabelsToProto(o *alpha.PatchDeploymentInstanceFilterGroupLabels) *alphapb.OsconfigAlphaPatchDeploymentInstanceFilterGroupLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentInstanceFilterGroupLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// PatchDeploymentPatchConfigToProto converts a PatchDeploymentPatchConfig object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigToProto(o *alpha.PatchDeploymentPatchConfig) *alphapb.OsconfigAlphaPatchDeploymentPatchConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfig{}
	p.SetRebootConfig(OsconfigAlphaPatchDeploymentPatchConfigRebootConfigEnumToProto(o.RebootConfig))
	p.SetApt(OsconfigAlphaPatchDeploymentPatchConfigAptToProto(o.Apt))
	p.SetYum(OsconfigAlphaPatchDeploymentPatchConfigYumToProto(o.Yum))
	p.SetGoo(OsconfigAlphaPatchDeploymentPatchConfigGooToProto(o.Goo))
	p.SetZypper(OsconfigAlphaPatchDeploymentPatchConfigZypperToProto(o.Zypper))
	p.SetWindowsUpdate(OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateToProto(o.WindowsUpdate))
	p.SetPreStep(OsconfigAlphaPatchDeploymentPatchConfigPreStepToProto(o.PreStep))
	p.SetPostStep(OsconfigAlphaPatchDeploymentPatchConfigPostStepToProto(o.PostStep))
	p.SetRetryStrategy(OsconfigAlphaPatchDeploymentPatchConfigRetryStrategyToProto(o.RetryStrategy))
	return p
}

// PatchDeploymentPatchConfigAptToProto converts a PatchDeploymentPatchConfigApt object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigAptToProto(o *alpha.PatchDeploymentPatchConfigApt) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigApt{}
	p.SetType(OsconfigAlphaPatchDeploymentPatchConfigAptTypeEnumToProto(o.Type))
	sExcludes := make([]string, len(o.Excludes))
	for i, r := range o.Excludes {
		sExcludes[i] = r
	}
	p.SetExcludes(sExcludes)
	sExclusivePackages := make([]string, len(o.ExclusivePackages))
	for i, r := range o.ExclusivePackages {
		sExclusivePackages[i] = r
	}
	p.SetExclusivePackages(sExclusivePackages)
	return p
}

// PatchDeploymentPatchConfigYumToProto converts a PatchDeploymentPatchConfigYum object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigYumToProto(o *alpha.PatchDeploymentPatchConfigYum) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigYum{}
	p.SetSecurity(dcl.ValueOrEmptyBool(o.Security))
	p.SetMinimal(dcl.ValueOrEmptyBool(o.Minimal))
	sExcludes := make([]string, len(o.Excludes))
	for i, r := range o.Excludes {
		sExcludes[i] = r
	}
	p.SetExcludes(sExcludes)
	sExclusivePackages := make([]string, len(o.ExclusivePackages))
	for i, r := range o.ExclusivePackages {
		sExclusivePackages[i] = r
	}
	p.SetExclusivePackages(sExclusivePackages)
	return p
}

// PatchDeploymentPatchConfigGooToProto converts a PatchDeploymentPatchConfigGoo object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigGooToProto(o *alpha.PatchDeploymentPatchConfigGoo) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigGoo {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigGoo{}
	return p
}

// PatchDeploymentPatchConfigZypperToProto converts a PatchDeploymentPatchConfigZypper object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigZypperToProto(o *alpha.PatchDeploymentPatchConfigZypper) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigZypper{}
	p.SetWithOptional(dcl.ValueOrEmptyBool(o.WithOptional))
	p.SetWithUpdate(dcl.ValueOrEmptyBool(o.WithUpdate))
	sCategories := make([]string, len(o.Categories))
	for i, r := range o.Categories {
		sCategories[i] = r
	}
	p.SetCategories(sCategories)
	sSeverities := make([]string, len(o.Severities))
	for i, r := range o.Severities {
		sSeverities[i] = r
	}
	p.SetSeverities(sSeverities)
	sExcludes := make([]string, len(o.Excludes))
	for i, r := range o.Excludes {
		sExcludes[i] = r
	}
	p.SetExcludes(sExcludes)
	sExclusivePatches := make([]string, len(o.ExclusivePatches))
	for i, r := range o.ExclusivePatches {
		sExclusivePatches[i] = r
	}
	p.SetExclusivePatches(sExclusivePatches)
	return p
}

// PatchDeploymentPatchConfigWindowsUpdateToProto converts a PatchDeploymentPatchConfigWindowsUpdate object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateToProto(o *alpha.PatchDeploymentPatchConfigWindowsUpdate) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdate{}
	sClassifications := make([]alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum, len(o.Classifications))
	for i, r := range o.Classifications {
		sClassifications[i] = alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(alphapb.OsconfigAlphaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_value[string(r)])
	}
	p.SetClassifications(sClassifications)
	sExcludes := make([]string, len(o.Excludes))
	for i, r := range o.Excludes {
		sExcludes[i] = r
	}
	p.SetExcludes(sExcludes)
	sExclusivePatches := make([]string, len(o.ExclusivePatches))
	for i, r := range o.ExclusivePatches {
		sExclusivePatches[i] = r
	}
	p.SetExclusivePatches(sExclusivePatches)
	return p
}

// PatchDeploymentPatchConfigPreStepToProto converts a PatchDeploymentPatchConfigPreStep object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepToProto(o *alpha.PatchDeploymentPatchConfigPreStep) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStep {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStep{}
	p.SetLinuxExecStepConfig(OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto(o.LinuxExecStepConfig))
	p.SetWindowsExecStepConfig(OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto(o.WindowsExecStepConfig))
	return p
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfig object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto(o *alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto(o *alpha.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfig object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto(o *alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto(o *alpha.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPostStepToProto converts a PatchDeploymentPatchConfigPostStep object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepToProto(o *alpha.PatchDeploymentPatchConfigPostStep) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStep {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStep{}
	p.SetLinuxExecStepConfig(OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto(o.LinuxExecStepConfig))
	p.SetWindowsExecStepConfig(OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto(o.WindowsExecStepConfig))
	return p
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfig object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto(o *alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto(o *alpha.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfig object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto(o *alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto(o *alpha.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigRetryStrategyToProto converts a PatchDeploymentPatchConfigRetryStrategy object to its proto representation.
func OsconfigAlphaPatchDeploymentPatchConfigRetryStrategyToProto(o *alpha.PatchDeploymentPatchConfigRetryStrategy) *alphapb.OsconfigAlphaPatchDeploymentPatchConfigRetryStrategy {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentPatchConfigRetryStrategy{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// PatchDeploymentOneTimeScheduleToProto converts a PatchDeploymentOneTimeSchedule object to its proto representation.
func OsconfigAlphaPatchDeploymentOneTimeScheduleToProto(o *alpha.PatchDeploymentOneTimeSchedule) *alphapb.OsconfigAlphaPatchDeploymentOneTimeSchedule {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentOneTimeSchedule{}
	p.SetExecuteTime(dcl.ValueOrEmptyString(o.ExecuteTime))
	return p
}

// PatchDeploymentRecurringScheduleToProto converts a PatchDeploymentRecurringSchedule object to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleToProto(o *alpha.PatchDeploymentRecurringSchedule) *alphapb.OsconfigAlphaPatchDeploymentRecurringSchedule {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRecurringSchedule{}
	p.SetTimeZone(OsconfigAlphaPatchDeploymentRecurringScheduleTimeZoneToProto(o.TimeZone))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyString(o.EndTime))
	p.SetTimeOfDay(OsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDayToProto(o.TimeOfDay))
	p.SetFrequency(OsconfigAlphaPatchDeploymentRecurringScheduleFrequencyEnumToProto(o.Frequency))
	p.SetWeekly(OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyToProto(o.Weekly))
	p.SetMonthly(OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyToProto(o.Monthly))
	p.SetLastExecuteTime(dcl.ValueOrEmptyString(o.LastExecuteTime))
	p.SetNextExecuteTime(dcl.ValueOrEmptyString(o.NextExecuteTime))
	return p
}

// PatchDeploymentRecurringScheduleTimeZoneToProto converts a PatchDeploymentRecurringScheduleTimeZone object to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleTimeZoneToProto(o *alpha.PatchDeploymentRecurringScheduleTimeZone) *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleTimeZone {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleTimeZone{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// PatchDeploymentRecurringScheduleTimeOfDayToProto converts a PatchDeploymentRecurringScheduleTimeOfDay object to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDayToProto(o *alpha.PatchDeploymentRecurringScheduleTimeOfDay) *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDay {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleTimeOfDay{}
	p.SetHours(dcl.ValueOrEmptyInt64(o.Hours))
	p.SetMinutes(dcl.ValueOrEmptyInt64(o.Minutes))
	p.SetSeconds(dcl.ValueOrEmptyInt64(o.Seconds))
	p.SetNanos(dcl.ValueOrEmptyInt64(o.Nanos))
	return p
}

// PatchDeploymentRecurringScheduleWeeklyToProto converts a PatchDeploymentRecurringScheduleWeekly object to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyToProto(o *alpha.PatchDeploymentRecurringScheduleWeekly) *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeekly {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleWeekly{}
	p.SetDayOfWeek(OsconfigAlphaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto(o.DayOfWeek))
	return p
}

// PatchDeploymentRecurringScheduleMonthlyToProto converts a PatchDeploymentRecurringScheduleMonthly object to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyToProto(o *alpha.PatchDeploymentRecurringScheduleMonthly) *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthly {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthly{}
	p.SetWeekDayOfMonth(OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto(o.WeekDayOfMonth))
	p.SetMonthDay(dcl.ValueOrEmptyInt64(o.MonthDay))
	return p
}

// PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth object to its proto representation.
func OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto(o *alpha.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) *alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{}
	p.SetWeekOrdinal(dcl.ValueOrEmptyInt64(o.WeekOrdinal))
	p.SetDayOfWeek(OsconfigAlphaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto(o.DayOfWeek))
	return p
}

// PatchDeploymentRolloutToProto converts a PatchDeploymentRollout object to its proto representation.
func OsconfigAlphaPatchDeploymentRolloutToProto(o *alpha.PatchDeploymentRollout) *alphapb.OsconfigAlphaPatchDeploymentRollout {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRollout{}
	p.SetMode(OsconfigAlphaPatchDeploymentRolloutModeEnumToProto(o.Mode))
	p.SetDisruptionBudget(OsconfigAlphaPatchDeploymentRolloutDisruptionBudgetToProto(o.DisruptionBudget))
	return p
}

// PatchDeploymentRolloutDisruptionBudgetToProto converts a PatchDeploymentRolloutDisruptionBudget object to its proto representation.
func OsconfigAlphaPatchDeploymentRolloutDisruptionBudgetToProto(o *alpha.PatchDeploymentRolloutDisruptionBudget) *alphapb.OsconfigAlphaPatchDeploymentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaPatchDeploymentRolloutDisruptionBudget{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	return p
}

// PatchDeploymentToProto converts a PatchDeployment resource to its proto representation.
func PatchDeploymentToProto(resource *alpha.PatchDeployment) *alphapb.OsconfigAlphaPatchDeployment {
	p := &alphapb.OsconfigAlphaPatchDeployment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInstanceFilter(OsconfigAlphaPatchDeploymentInstanceFilterToProto(resource.InstanceFilter))
	p.SetPatchConfig(OsconfigAlphaPatchDeploymentPatchConfigToProto(resource.PatchConfig))
	p.SetDuration(dcl.ValueOrEmptyString(resource.Duration))
	p.SetOneTimeSchedule(OsconfigAlphaPatchDeploymentOneTimeScheduleToProto(resource.OneTimeSchedule))
	p.SetRecurringSchedule(OsconfigAlphaPatchDeploymentRecurringScheduleToProto(resource.RecurringSchedule))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetLastExecuteTime(dcl.ValueOrEmptyString(resource.LastExecuteTime))
	p.SetRollout(OsconfigAlphaPatchDeploymentRolloutToProto(resource.Rollout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyPatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Apply() method.
func (s *PatchDeploymentServer) applyPatchDeployment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOsconfigAlphaPatchDeploymentRequest) (*alphapb.OsconfigAlphaPatchDeployment, error) {
	p := ProtoToPatchDeployment(request.GetResource())
	res, err := c.ApplyPatchDeployment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PatchDeploymentToProto(res)
	return r, nil
}

// applyOsconfigAlphaPatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Apply() method.
func (s *PatchDeploymentServer) ApplyOsconfigAlphaPatchDeployment(ctx context.Context, request *alphapb.ApplyOsconfigAlphaPatchDeploymentRequest) (*alphapb.OsconfigAlphaPatchDeployment, error) {
	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPatchDeployment(ctx, cl, request)
}

// DeletePatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Delete() method.
func (s *PatchDeploymentServer) DeleteOsconfigAlphaPatchDeployment(ctx context.Context, request *alphapb.DeleteOsconfigAlphaPatchDeploymentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePatchDeployment(ctx, ProtoToPatchDeployment(request.GetResource()))

}

// ListOsconfigAlphaPatchDeployment handles the gRPC request by passing it to the underlying PatchDeploymentList() method.
func (s *PatchDeploymentServer) ListOsconfigAlphaPatchDeployment(ctx context.Context, request *alphapb.ListOsconfigAlphaPatchDeploymentRequest) (*alphapb.ListOsconfigAlphaPatchDeploymentResponse, error) {
	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPatchDeployment(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OsconfigAlphaPatchDeployment
	for _, r := range resources.Items {
		rp := PatchDeploymentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListOsconfigAlphaPatchDeploymentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPatchDeployment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
