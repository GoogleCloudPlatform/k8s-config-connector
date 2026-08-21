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
	osconfigpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/osconfig_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig"
)

// PatchDeploymentServer implements the gRPC interface for PatchDeployment.
type PatchDeploymentServer struct{}

// ProtoToPatchDeploymentPatchConfigRebootConfigEnum converts a PatchDeploymentPatchConfigRebootConfigEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigRebootConfigEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum) *osconfig.PatchDeploymentPatchConfigRebootConfigEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigRebootConfigEnum(n[len("OsconfigPatchDeploymentPatchConfigRebootConfigEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigAptTypeEnum converts a PatchDeploymentPatchConfigAptTypeEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigAptTypeEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum) *osconfig.PatchDeploymentPatchConfigAptTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigAptTypeEnum(n[len("OsconfigPatchDeploymentPatchConfigAptTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum converts a PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum) *osconfig.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(n[len("OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum) *osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(n[len("OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum) *osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(n[len("OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum) *osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(n[len("OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(e osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum) *osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(n[len("OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleFrequencyEnum converts a PatchDeploymentRecurringScheduleFrequencyEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleFrequencyEnum(e osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum) *osconfig.PatchDeploymentRecurringScheduleFrequencyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentRecurringScheduleFrequencyEnum(n[len("OsconfigPatchDeploymentRecurringScheduleFrequencyEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum converts a PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(e osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum) *osconfig.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(n[len("OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(e osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum) *osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(n[len("OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRolloutModeEnum converts a PatchDeploymentRolloutModeEnum enum from its proto representation.
func ProtoToOsconfigPatchDeploymentRolloutModeEnum(e osconfigpb.OsconfigPatchDeploymentRolloutModeEnum) *osconfig.PatchDeploymentRolloutModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigPatchDeploymentRolloutModeEnum_name[int32(e)]; ok {
		e := osconfig.PatchDeploymentRolloutModeEnum(n[len("OsconfigPatchDeploymentRolloutModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentInstanceFilter converts a PatchDeploymentInstanceFilter object from its proto representation.
func ProtoToOsconfigPatchDeploymentInstanceFilter(p *osconfigpb.OsconfigPatchDeploymentInstanceFilter) *osconfig.PatchDeploymentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentInstanceFilter{
		All: dcl.Bool(p.GetAll()),
	}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigPatchDeploymentInstanceFilterGroupLabels(r))
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
func ProtoToOsconfigPatchDeploymentInstanceFilterGroupLabels(p *osconfigpb.OsconfigPatchDeploymentInstanceFilterGroupLabels) *osconfig.PatchDeploymentInstanceFilterGroupLabels {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentInstanceFilterGroupLabels{}
	return obj
}

// ProtoToPatchDeploymentPatchConfig converts a PatchDeploymentPatchConfig object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfig(p *osconfigpb.OsconfigPatchDeploymentPatchConfig) *osconfig.PatchDeploymentPatchConfig {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfig{
		RebootConfig:  ProtoToOsconfigPatchDeploymentPatchConfigRebootConfigEnum(p.GetRebootConfig()),
		Apt:           ProtoToOsconfigPatchDeploymentPatchConfigApt(p.GetApt()),
		Yum:           ProtoToOsconfigPatchDeploymentPatchConfigYum(p.GetYum()),
		Goo:           ProtoToOsconfigPatchDeploymentPatchConfigGoo(p.GetGoo()),
		Zypper:        ProtoToOsconfigPatchDeploymentPatchConfigZypper(p.GetZypper()),
		WindowsUpdate: ProtoToOsconfigPatchDeploymentPatchConfigWindowsUpdate(p.GetWindowsUpdate()),
		PreStep:       ProtoToOsconfigPatchDeploymentPatchConfigPreStep(p.GetPreStep()),
		PostStep:      ProtoToOsconfigPatchDeploymentPatchConfigPostStep(p.GetPostStep()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigApt converts a PatchDeploymentPatchConfigApt object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigApt(p *osconfigpb.OsconfigPatchDeploymentPatchConfigApt) *osconfig.PatchDeploymentPatchConfigApt {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigApt{
		Type: ProtoToOsconfigPatchDeploymentPatchConfigAptTypeEnum(p.GetType()),
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
func ProtoToOsconfigPatchDeploymentPatchConfigYum(p *osconfigpb.OsconfigPatchDeploymentPatchConfigYum) *osconfig.PatchDeploymentPatchConfigYum {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigYum{
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
func ProtoToOsconfigPatchDeploymentPatchConfigGoo(p *osconfigpb.OsconfigPatchDeploymentPatchConfigGoo) *osconfig.PatchDeploymentPatchConfigGoo {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigGoo{}
	return obj
}

// ProtoToPatchDeploymentPatchConfigZypper converts a PatchDeploymentPatchConfigZypper object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigZypper(p *osconfigpb.OsconfigPatchDeploymentPatchConfigZypper) *osconfig.PatchDeploymentPatchConfigZypper {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigZypper{
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
func ProtoToOsconfigPatchDeploymentPatchConfigWindowsUpdate(p *osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdate) *osconfig.PatchDeploymentPatchConfigWindowsUpdate {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigWindowsUpdate{}
	for _, r := range p.GetClassifications() {
		obj.Classifications = append(obj.Classifications, *ProtoToOsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(r))
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
func ProtoToOsconfigPatchDeploymentPatchConfigPreStep(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStep) *osconfig.PatchDeploymentPatchConfigPreStep {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPreStep{
		LinuxExecStepConfig:   ProtoToOsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig(p.GetLinuxExecStepConfig()),
		WindowsExecStepConfig: ProtoToOsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig(p.GetWindowsExecStepConfig()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfig converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfig object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig) *osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) *osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfig converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfig object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig) *osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) *osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStep converts a PatchDeploymentPatchConfigPostStep object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStep(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStep) *osconfig.PatchDeploymentPatchConfigPostStep {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPostStep{
		LinuxExecStepConfig:   ProtoToOsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig(p.GetLinuxExecStepConfig()),
		WindowsExecStepConfig: ProtoToOsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig(p.GetWindowsExecStepConfig()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfig converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfig object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig) *osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) *osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfig converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfig object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig) *osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(p *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject) *osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentOneTimeSchedule converts a PatchDeploymentOneTimeSchedule object from its proto representation.
func ProtoToOsconfigPatchDeploymentOneTimeSchedule(p *osconfigpb.OsconfigPatchDeploymentOneTimeSchedule) *osconfig.PatchDeploymentOneTimeSchedule {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentOneTimeSchedule{
		ExecuteTime: dcl.StringOrNil(p.GetExecuteTime()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringSchedule converts a PatchDeploymentRecurringSchedule object from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringSchedule(p *osconfigpb.OsconfigPatchDeploymentRecurringSchedule) *osconfig.PatchDeploymentRecurringSchedule {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRecurringSchedule{
		TimeZone:        ProtoToOsconfigPatchDeploymentRecurringScheduleTimeZone(p.GetTimeZone()),
		StartTime:       dcl.StringOrNil(p.GetStartTime()),
		EndTime:         dcl.StringOrNil(p.GetEndTime()),
		TimeOfDay:       ProtoToOsconfigPatchDeploymentRecurringScheduleTimeOfDay(p.GetTimeOfDay()),
		Frequency:       ProtoToOsconfigPatchDeploymentRecurringScheduleFrequencyEnum(p.GetFrequency()),
		Weekly:          ProtoToOsconfigPatchDeploymentRecurringScheduleWeekly(p.GetWeekly()),
		Monthly:         ProtoToOsconfigPatchDeploymentRecurringScheduleMonthly(p.GetMonthly()),
		LastExecuteTime: dcl.StringOrNil(p.GetLastExecuteTime()),
		NextExecuteTime: dcl.StringOrNil(p.GetNextExecuteTime()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleTimeZone converts a PatchDeploymentRecurringScheduleTimeZone object from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleTimeZone(p *osconfigpb.OsconfigPatchDeploymentRecurringScheduleTimeZone) *osconfig.PatchDeploymentRecurringScheduleTimeZone {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRecurringScheduleTimeZone{
		Id:      dcl.StringOrNil(p.GetId()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleTimeOfDay converts a PatchDeploymentRecurringScheduleTimeOfDay object from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleTimeOfDay(p *osconfigpb.OsconfigPatchDeploymentRecurringScheduleTimeOfDay) *osconfig.PatchDeploymentRecurringScheduleTimeOfDay {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRecurringScheduleTimeOfDay{
		Hours:   dcl.Int64OrNil(p.GetHours()),
		Minutes: dcl.Int64OrNil(p.GetMinutes()),
		Seconds: dcl.Int64OrNil(p.GetSeconds()),
		Nanos:   dcl.Int64OrNil(p.GetNanos()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleWeekly converts a PatchDeploymentRecurringScheduleWeekly object from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleWeekly(p *osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeekly) *osconfig.PatchDeploymentRecurringScheduleWeekly {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRecurringScheduleWeekly{
		DayOfWeek: ProtoToOsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(p.GetDayOfWeek()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleMonthly converts a PatchDeploymentRecurringScheduleMonthly object from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleMonthly(p *osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthly) *osconfig.PatchDeploymentRecurringScheduleMonthly {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRecurringScheduleMonthly{
		WeekDayOfMonth: ProtoToOsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(p.GetWeekDayOfMonth()),
		MonthDay:       dcl.Int64OrNil(p.GetMonthDay()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth object from its proto representation.
func ProtoToOsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(p *osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) *osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{
		WeekOrdinal: dcl.Int64OrNil(p.GetWeekOrdinal()),
		DayOfWeek:   ProtoToOsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(p.GetDayOfWeek()),
	}
	return obj
}

// ProtoToPatchDeploymentRollout converts a PatchDeploymentRollout object from its proto representation.
func ProtoToOsconfigPatchDeploymentRollout(p *osconfigpb.OsconfigPatchDeploymentRollout) *osconfig.PatchDeploymentRollout {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRollout{
		Mode:             ProtoToOsconfigPatchDeploymentRolloutModeEnum(p.GetMode()),
		DisruptionBudget: ProtoToOsconfigPatchDeploymentRolloutDisruptionBudget(p.GetDisruptionBudget()),
	}
	return obj
}

// ProtoToPatchDeploymentRolloutDisruptionBudget converts a PatchDeploymentRolloutDisruptionBudget object from its proto representation.
func ProtoToOsconfigPatchDeploymentRolloutDisruptionBudget(p *osconfigpb.OsconfigPatchDeploymentRolloutDisruptionBudget) *osconfig.PatchDeploymentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &osconfig.PatchDeploymentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.GetFixed()),
		Percent: dcl.Int64OrNil(p.GetPercent()),
	}
	return obj
}

// ProtoToPatchDeployment converts a PatchDeployment resource from its proto representation.
func ProtoToPatchDeployment(p *osconfigpb.OsconfigPatchDeployment) *osconfig.PatchDeployment {
	obj := &osconfig.PatchDeployment{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		InstanceFilter:    ProtoToOsconfigPatchDeploymentInstanceFilter(p.GetInstanceFilter()),
		PatchConfig:       ProtoToOsconfigPatchDeploymentPatchConfig(p.GetPatchConfig()),
		Duration:          dcl.StringOrNil(p.GetDuration()),
		OneTimeSchedule:   ProtoToOsconfigPatchDeploymentOneTimeSchedule(p.GetOneTimeSchedule()),
		RecurringSchedule: ProtoToOsconfigPatchDeploymentRecurringSchedule(p.GetRecurringSchedule()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		LastExecuteTime:   dcl.StringOrNil(p.GetLastExecuteTime()),
		Rollout:           ProtoToOsconfigPatchDeploymentRollout(p.GetRollout()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// PatchDeploymentPatchConfigRebootConfigEnumToProto converts a PatchDeploymentPatchConfigRebootConfigEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigRebootConfigEnumToProto(e *osconfig.PatchDeploymentPatchConfigRebootConfigEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum_value["PatchDeploymentPatchConfigRebootConfigEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigRebootConfigEnum(0)
}

// PatchDeploymentPatchConfigAptTypeEnumToProto converts a PatchDeploymentPatchConfigAptTypeEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigAptTypeEnumToProto(e *osconfig.PatchDeploymentPatchConfigAptTypeEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum_value["PatchDeploymentPatchConfigAptTypeEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigAptTypeEnum(0)
}

// PatchDeploymentPatchConfigWindowsUpdateClassificationsEnumToProto converts a PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnumToProto(e *osconfig.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_value["PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(0)
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto(e *osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto(e *osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto(e *osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto(e *osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum) osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentRecurringScheduleFrequencyEnumToProto converts a PatchDeploymentRecurringScheduleFrequencyEnum enum to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleFrequencyEnumToProto(e *osconfig.PatchDeploymentRecurringScheduleFrequencyEnum) osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum_value["PatchDeploymentRecurringScheduleFrequencyEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentRecurringScheduleFrequencyEnum(0)
}

// PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto converts a PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum enum to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto(e *osconfig.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum) osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum_value["PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(0)
}

// PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum enum to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto(e *osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum) osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum_value["PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(0)
}

// PatchDeploymentRolloutModeEnumToProto converts a PatchDeploymentRolloutModeEnum enum to its proto representation.
func OsconfigPatchDeploymentRolloutModeEnumToProto(e *osconfig.PatchDeploymentRolloutModeEnum) osconfigpb.OsconfigPatchDeploymentRolloutModeEnum {
	if e == nil {
		return osconfigpb.OsconfigPatchDeploymentRolloutModeEnum(0)
	}
	if v, ok := osconfigpb.OsconfigPatchDeploymentRolloutModeEnum_value["PatchDeploymentRolloutModeEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigPatchDeploymentRolloutModeEnum(v)
	}
	return osconfigpb.OsconfigPatchDeploymentRolloutModeEnum(0)
}

// PatchDeploymentInstanceFilterToProto converts a PatchDeploymentInstanceFilter object to its proto representation.
func OsconfigPatchDeploymentInstanceFilterToProto(o *osconfig.PatchDeploymentInstanceFilter) *osconfigpb.OsconfigPatchDeploymentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentInstanceFilter{}
	p.SetAll(dcl.ValueOrEmptyBool(o.All))
	sGroupLabels := make([]*osconfigpb.OsconfigPatchDeploymentInstanceFilterGroupLabels, len(o.GroupLabels))
	for i, r := range o.GroupLabels {
		sGroupLabels[i] = OsconfigPatchDeploymentInstanceFilterGroupLabelsToProto(&r)
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
func OsconfigPatchDeploymentInstanceFilterGroupLabelsToProto(o *osconfig.PatchDeploymentInstanceFilterGroupLabels) *osconfigpb.OsconfigPatchDeploymentInstanceFilterGroupLabels {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentInstanceFilterGroupLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// PatchDeploymentPatchConfigToProto converts a PatchDeploymentPatchConfig object to its proto representation.
func OsconfigPatchDeploymentPatchConfigToProto(o *osconfig.PatchDeploymentPatchConfig) *osconfigpb.OsconfigPatchDeploymentPatchConfig {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfig{}
	p.SetRebootConfig(OsconfigPatchDeploymentPatchConfigRebootConfigEnumToProto(o.RebootConfig))
	p.SetApt(OsconfigPatchDeploymentPatchConfigAptToProto(o.Apt))
	p.SetYum(OsconfigPatchDeploymentPatchConfigYumToProto(o.Yum))
	p.SetGoo(OsconfigPatchDeploymentPatchConfigGooToProto(o.Goo))
	p.SetZypper(OsconfigPatchDeploymentPatchConfigZypperToProto(o.Zypper))
	p.SetWindowsUpdate(OsconfigPatchDeploymentPatchConfigWindowsUpdateToProto(o.WindowsUpdate))
	p.SetPreStep(OsconfigPatchDeploymentPatchConfigPreStepToProto(o.PreStep))
	p.SetPostStep(OsconfigPatchDeploymentPatchConfigPostStepToProto(o.PostStep))
	return p
}

// PatchDeploymentPatchConfigAptToProto converts a PatchDeploymentPatchConfigApt object to its proto representation.
func OsconfigPatchDeploymentPatchConfigAptToProto(o *osconfig.PatchDeploymentPatchConfigApt) *osconfigpb.OsconfigPatchDeploymentPatchConfigApt {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigApt{}
	p.SetType(OsconfigPatchDeploymentPatchConfigAptTypeEnumToProto(o.Type))
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
func OsconfigPatchDeploymentPatchConfigYumToProto(o *osconfig.PatchDeploymentPatchConfigYum) *osconfigpb.OsconfigPatchDeploymentPatchConfigYum {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigYum{}
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
func OsconfigPatchDeploymentPatchConfigGooToProto(o *osconfig.PatchDeploymentPatchConfigGoo) *osconfigpb.OsconfigPatchDeploymentPatchConfigGoo {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigGoo{}
	return p
}

// PatchDeploymentPatchConfigZypperToProto converts a PatchDeploymentPatchConfigZypper object to its proto representation.
func OsconfigPatchDeploymentPatchConfigZypperToProto(o *osconfig.PatchDeploymentPatchConfigZypper) *osconfigpb.OsconfigPatchDeploymentPatchConfigZypper {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigZypper{}
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
func OsconfigPatchDeploymentPatchConfigWindowsUpdateToProto(o *osconfig.PatchDeploymentPatchConfigWindowsUpdate) *osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdate {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdate{}
	sClassifications := make([]osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum, len(o.Classifications))
	for i, r := range o.Classifications {
		sClassifications[i] = osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(osconfigpb.OsconfigPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_value[string(r)])
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
func OsconfigPatchDeploymentPatchConfigPreStepToProto(o *osconfig.PatchDeploymentPatchConfigPreStep) *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStep {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPreStep{}
	p.SetLinuxExecStepConfig(OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto(o.LinuxExecStepConfig))
	p.SetWindowsExecStepConfig(OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto(o.WindowsExecStepConfig))
	return p
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfig object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto(o *osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig) *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto(o *osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfig object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto(o *osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig) *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto(o *osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) *osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPostStepToProto converts a PatchDeploymentPatchConfigPostStep object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepToProto(o *osconfig.PatchDeploymentPatchConfigPostStep) *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStep {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPostStep{}
	p.SetLinuxExecStepConfig(OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto(o.LinuxExecStepConfig))
	p.SetWindowsExecStepConfig(OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto(o.WindowsExecStepConfig))
	return p
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfig object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto(o *osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig) *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto(o *osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfig object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto(o *osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig) *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject object to its proto representation.
func OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto(o *osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject) *osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentOneTimeScheduleToProto converts a PatchDeploymentOneTimeSchedule object to its proto representation.
func OsconfigPatchDeploymentOneTimeScheduleToProto(o *osconfig.PatchDeploymentOneTimeSchedule) *osconfigpb.OsconfigPatchDeploymentOneTimeSchedule {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentOneTimeSchedule{}
	p.SetExecuteTime(dcl.ValueOrEmptyString(o.ExecuteTime))
	return p
}

// PatchDeploymentRecurringScheduleToProto converts a PatchDeploymentRecurringSchedule object to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleToProto(o *osconfig.PatchDeploymentRecurringSchedule) *osconfigpb.OsconfigPatchDeploymentRecurringSchedule {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRecurringSchedule{}
	p.SetTimeZone(OsconfigPatchDeploymentRecurringScheduleTimeZoneToProto(o.TimeZone))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyString(o.EndTime))
	p.SetTimeOfDay(OsconfigPatchDeploymentRecurringScheduleTimeOfDayToProto(o.TimeOfDay))
	p.SetFrequency(OsconfigPatchDeploymentRecurringScheduleFrequencyEnumToProto(o.Frequency))
	p.SetWeekly(OsconfigPatchDeploymentRecurringScheduleWeeklyToProto(o.Weekly))
	p.SetMonthly(OsconfigPatchDeploymentRecurringScheduleMonthlyToProto(o.Monthly))
	p.SetLastExecuteTime(dcl.ValueOrEmptyString(o.LastExecuteTime))
	p.SetNextExecuteTime(dcl.ValueOrEmptyString(o.NextExecuteTime))
	return p
}

// PatchDeploymentRecurringScheduleTimeZoneToProto converts a PatchDeploymentRecurringScheduleTimeZone object to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleTimeZoneToProto(o *osconfig.PatchDeploymentRecurringScheduleTimeZone) *osconfigpb.OsconfigPatchDeploymentRecurringScheduleTimeZone {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRecurringScheduleTimeZone{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// PatchDeploymentRecurringScheduleTimeOfDayToProto converts a PatchDeploymentRecurringScheduleTimeOfDay object to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleTimeOfDayToProto(o *osconfig.PatchDeploymentRecurringScheduleTimeOfDay) *osconfigpb.OsconfigPatchDeploymentRecurringScheduleTimeOfDay {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRecurringScheduleTimeOfDay{}
	p.SetHours(dcl.ValueOrEmptyInt64(o.Hours))
	p.SetMinutes(dcl.ValueOrEmptyInt64(o.Minutes))
	p.SetSeconds(dcl.ValueOrEmptyInt64(o.Seconds))
	p.SetNanos(dcl.ValueOrEmptyInt64(o.Nanos))
	return p
}

// PatchDeploymentRecurringScheduleWeeklyToProto converts a PatchDeploymentRecurringScheduleWeekly object to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleWeeklyToProto(o *osconfig.PatchDeploymentRecurringScheduleWeekly) *osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeekly {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRecurringScheduleWeekly{}
	p.SetDayOfWeek(OsconfigPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto(o.DayOfWeek))
	return p
}

// PatchDeploymentRecurringScheduleMonthlyToProto converts a PatchDeploymentRecurringScheduleMonthly object to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleMonthlyToProto(o *osconfig.PatchDeploymentRecurringScheduleMonthly) *osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthly {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthly{}
	p.SetWeekDayOfMonth(OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto(o.WeekDayOfMonth))
	p.SetMonthDay(dcl.ValueOrEmptyInt64(o.MonthDay))
	return p
}

// PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth object to its proto representation.
func OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto(o *osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) *osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{}
	p.SetWeekOrdinal(dcl.ValueOrEmptyInt64(o.WeekOrdinal))
	p.SetDayOfWeek(OsconfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto(o.DayOfWeek))
	return p
}

// PatchDeploymentRolloutToProto converts a PatchDeploymentRollout object to its proto representation.
func OsconfigPatchDeploymentRolloutToProto(o *osconfig.PatchDeploymentRollout) *osconfigpb.OsconfigPatchDeploymentRollout {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRollout{}
	p.SetMode(OsconfigPatchDeploymentRolloutModeEnumToProto(o.Mode))
	p.SetDisruptionBudget(OsconfigPatchDeploymentRolloutDisruptionBudgetToProto(o.DisruptionBudget))
	return p
}

// PatchDeploymentRolloutDisruptionBudgetToProto converts a PatchDeploymentRolloutDisruptionBudget object to its proto representation.
func OsconfigPatchDeploymentRolloutDisruptionBudgetToProto(o *osconfig.PatchDeploymentRolloutDisruptionBudget) *osconfigpb.OsconfigPatchDeploymentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigPatchDeploymentRolloutDisruptionBudget{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	return p
}

// PatchDeploymentToProto converts a PatchDeployment resource to its proto representation.
func PatchDeploymentToProto(resource *osconfig.PatchDeployment) *osconfigpb.OsconfigPatchDeployment {
	p := &osconfigpb.OsconfigPatchDeployment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInstanceFilter(OsconfigPatchDeploymentInstanceFilterToProto(resource.InstanceFilter))
	p.SetPatchConfig(OsconfigPatchDeploymentPatchConfigToProto(resource.PatchConfig))
	p.SetDuration(dcl.ValueOrEmptyString(resource.Duration))
	p.SetOneTimeSchedule(OsconfigPatchDeploymentOneTimeScheduleToProto(resource.OneTimeSchedule))
	p.SetRecurringSchedule(OsconfigPatchDeploymentRecurringScheduleToProto(resource.RecurringSchedule))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetLastExecuteTime(dcl.ValueOrEmptyString(resource.LastExecuteTime))
	p.SetRollout(OsconfigPatchDeploymentRolloutToProto(resource.Rollout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyPatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Apply() method.
func (s *PatchDeploymentServer) applyPatchDeployment(ctx context.Context, c *osconfig.Client, request *osconfigpb.ApplyOsconfigPatchDeploymentRequest) (*osconfigpb.OsconfigPatchDeployment, error) {
	p := ProtoToPatchDeployment(request.GetResource())
	res, err := c.ApplyPatchDeployment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PatchDeploymentToProto(res)
	return r, nil
}

// applyOsconfigPatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Apply() method.
func (s *PatchDeploymentServer) ApplyOsconfigPatchDeployment(ctx context.Context, request *osconfigpb.ApplyOsconfigPatchDeploymentRequest) (*osconfigpb.OsconfigPatchDeployment, error) {
	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPatchDeployment(ctx, cl, request)
}

// DeletePatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Delete() method.
func (s *PatchDeploymentServer) DeleteOsconfigPatchDeployment(ctx context.Context, request *osconfigpb.DeleteOsconfigPatchDeploymentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePatchDeployment(ctx, ProtoToPatchDeployment(request.GetResource()))

}

// ListOsconfigPatchDeployment handles the gRPC request by passing it to the underlying PatchDeploymentList() method.
func (s *PatchDeploymentServer) ListOsconfigPatchDeployment(ctx context.Context, request *osconfigpb.ListOsconfigPatchDeploymentRequest) (*osconfigpb.ListOsconfigPatchDeploymentResponse, error) {
	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPatchDeployment(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*osconfigpb.OsconfigPatchDeployment
	for _, r := range resources.Items {
		rp := PatchDeploymentToProto(r)
		protos = append(protos, rp)
	}
	p := &osconfigpb.ListOsconfigPatchDeploymentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPatchDeployment(ctx context.Context, service_account_file string) (*osconfig.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return osconfig.NewClient(conf), nil
}
