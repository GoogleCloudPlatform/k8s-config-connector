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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/beta/osconfig_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
)

// PatchDeploymentServer implements the gRPC interface for PatchDeployment.
type PatchDeploymentServer struct{}

// ProtoToPatchDeploymentPatchConfigRebootConfigEnum converts a PatchDeploymentPatchConfigRebootConfigEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum) *beta.PatchDeploymentPatchConfigRebootConfigEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigRebootConfigEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigAptTypeEnum converts a PatchDeploymentPatchConfigAptTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigAptTypeEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum) *beta.PatchDeploymentPatchConfigAptTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigAptTypeEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum converts a PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum) *beta.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum) *beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum) *beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum) *beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(e betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum) *beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(n[len("OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleFrequencyEnum converts a PatchDeploymentRecurringScheduleFrequencyEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum(e betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum) *beta.PatchDeploymentRecurringScheduleFrequencyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentRecurringScheduleFrequencyEnum(n[len("OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum converts a PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(e betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum) *beta.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(n[len("OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(e betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum) *beta.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(n[len("OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentRolloutModeEnum converts a PatchDeploymentRolloutModeEnum enum from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRolloutModeEnum(e betapb.OsconfigBetaPatchDeploymentRolloutModeEnum) *beta.PatchDeploymentRolloutModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaPatchDeploymentRolloutModeEnum_name[int32(e)]; ok {
		e := beta.PatchDeploymentRolloutModeEnum(n[len("OsconfigBetaPatchDeploymentRolloutModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPatchDeploymentInstanceFilter converts a PatchDeploymentInstanceFilter object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentInstanceFilter(p *betapb.OsconfigBetaPatchDeploymentInstanceFilter) *beta.PatchDeploymentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentInstanceFilter{
		All: dcl.Bool(p.GetAll()),
	}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigBetaPatchDeploymentInstanceFilterGroupLabels(r))
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
func ProtoToOsconfigBetaPatchDeploymentInstanceFilterGroupLabels(p *betapb.OsconfigBetaPatchDeploymentInstanceFilterGroupLabels) *beta.PatchDeploymentInstanceFilterGroupLabels {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentInstanceFilterGroupLabels{}
	return obj
}

// ProtoToPatchDeploymentPatchConfig converts a PatchDeploymentPatchConfig object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfig(p *betapb.OsconfigBetaPatchDeploymentPatchConfig) *beta.PatchDeploymentPatchConfig {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfig{
		RebootConfig:  ProtoToOsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum(p.GetRebootConfig()),
		Apt:           ProtoToOsconfigBetaPatchDeploymentPatchConfigApt(p.GetApt()),
		Yum:           ProtoToOsconfigBetaPatchDeploymentPatchConfigYum(p.GetYum()),
		Goo:           ProtoToOsconfigBetaPatchDeploymentPatchConfigGoo(p.GetGoo()),
		Zypper:        ProtoToOsconfigBetaPatchDeploymentPatchConfigZypper(p.GetZypper()),
		WindowsUpdate: ProtoToOsconfigBetaPatchDeploymentPatchConfigWindowsUpdate(p.GetWindowsUpdate()),
		PreStep:       ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStep(p.GetPreStep()),
		PostStep:      ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStep(p.GetPostStep()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigApt converts a PatchDeploymentPatchConfigApt object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigApt(p *betapb.OsconfigBetaPatchDeploymentPatchConfigApt) *beta.PatchDeploymentPatchConfigApt {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigApt{
		Type: ProtoToOsconfigBetaPatchDeploymentPatchConfigAptTypeEnum(p.GetType()),
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
func ProtoToOsconfigBetaPatchDeploymentPatchConfigYum(p *betapb.OsconfigBetaPatchDeploymentPatchConfigYum) *beta.PatchDeploymentPatchConfigYum {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigYum{
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
func ProtoToOsconfigBetaPatchDeploymentPatchConfigGoo(p *betapb.OsconfigBetaPatchDeploymentPatchConfigGoo) *beta.PatchDeploymentPatchConfigGoo {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigGoo{}
	return obj
}

// ProtoToPatchDeploymentPatchConfigZypper converts a PatchDeploymentPatchConfigZypper object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigZypper(p *betapb.OsconfigBetaPatchDeploymentPatchConfigZypper) *beta.PatchDeploymentPatchConfigZypper {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigZypper{
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
func ProtoToOsconfigBetaPatchDeploymentPatchConfigWindowsUpdate(p *betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdate) *beta.PatchDeploymentPatchConfigWindowsUpdate {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigWindowsUpdate{}
	for _, r := range p.GetClassifications() {
		obj.Classifications = append(obj.Classifications, *ProtoToOsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(r))
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
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStep(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStep) *beta.PatchDeploymentPatchConfigPreStep {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPreStep{
		LinuxExecStepConfig:   ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig(p.GetLinuxExecStepConfig()),
		WindowsExecStepConfig: ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig(p.GetWindowsExecStepConfig()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfig converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfig object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig) *beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) *beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfig converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfig object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig) *beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) *beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStep converts a PatchDeploymentPatchConfigPostStep object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStep(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStep) *beta.PatchDeploymentPatchConfigPostStep {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPostStep{
		LinuxExecStepConfig:   ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig(p.GetLinuxExecStepConfig()),
		WindowsExecStepConfig: ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig(p.GetWindowsExecStepConfig()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfig converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfig object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig) *beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) *beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfig converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfig object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig) *beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig{
		LocalPath:   dcl.StringOrNil(p.GetLocalPath()),
		GcsObject:   ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(p.GetGcsObject()),
		Interpreter: ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedSuccessCodes() {
		obj.AllowedSuccessCodes = append(obj.AllowedSuccessCodes, r)
	}
	return obj
}

// ProtoToPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(p *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject) *beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{
		Bucket:           dcl.StringOrNil(p.GetBucket()),
		Object:           dcl.StringOrNil(p.GetObject()),
		GenerationNumber: dcl.Int64OrNil(p.GetGenerationNumber()),
	}
	return obj
}

// ProtoToPatchDeploymentOneTimeSchedule converts a PatchDeploymentOneTimeSchedule object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentOneTimeSchedule(p *betapb.OsconfigBetaPatchDeploymentOneTimeSchedule) *beta.PatchDeploymentOneTimeSchedule {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentOneTimeSchedule{
		ExecuteTime: dcl.StringOrNil(p.GetExecuteTime()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringSchedule converts a PatchDeploymentRecurringSchedule object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringSchedule(p *betapb.OsconfigBetaPatchDeploymentRecurringSchedule) *beta.PatchDeploymentRecurringSchedule {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRecurringSchedule{
		TimeZone:        ProtoToOsconfigBetaPatchDeploymentRecurringScheduleTimeZone(p.GetTimeZone()),
		StartTime:       dcl.StringOrNil(p.GetStartTime()),
		EndTime:         dcl.StringOrNil(p.GetEndTime()),
		TimeOfDay:       ProtoToOsconfigBetaPatchDeploymentRecurringScheduleTimeOfDay(p.GetTimeOfDay()),
		Frequency:       ProtoToOsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum(p.GetFrequency()),
		Weekly:          ProtoToOsconfigBetaPatchDeploymentRecurringScheduleWeekly(p.GetWeekly()),
		Monthly:         ProtoToOsconfigBetaPatchDeploymentRecurringScheduleMonthly(p.GetMonthly()),
		LastExecuteTime: dcl.StringOrNil(p.GetLastExecuteTime()),
		NextExecuteTime: dcl.StringOrNil(p.GetNextExecuteTime()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleTimeZone converts a PatchDeploymentRecurringScheduleTimeZone object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleTimeZone(p *betapb.OsconfigBetaPatchDeploymentRecurringScheduleTimeZone) *beta.PatchDeploymentRecurringScheduleTimeZone {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRecurringScheduleTimeZone{
		Id:      dcl.StringOrNil(p.GetId()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleTimeOfDay converts a PatchDeploymentRecurringScheduleTimeOfDay object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleTimeOfDay(p *betapb.OsconfigBetaPatchDeploymentRecurringScheduleTimeOfDay) *beta.PatchDeploymentRecurringScheduleTimeOfDay {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRecurringScheduleTimeOfDay{
		Hours:   dcl.Int64OrNil(p.GetHours()),
		Minutes: dcl.Int64OrNil(p.GetMinutes()),
		Seconds: dcl.Int64OrNil(p.GetSeconds()),
		Nanos:   dcl.Int64OrNil(p.GetNanos()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleWeekly converts a PatchDeploymentRecurringScheduleWeekly object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleWeekly(p *betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeekly) *beta.PatchDeploymentRecurringScheduleWeekly {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRecurringScheduleWeekly{
		DayOfWeek: ProtoToOsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(p.GetDayOfWeek()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleMonthly converts a PatchDeploymentRecurringScheduleMonthly object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleMonthly(p *betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthly) *beta.PatchDeploymentRecurringScheduleMonthly {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRecurringScheduleMonthly{
		WeekDayOfMonth: ProtoToOsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(p.GetWeekDayOfMonth()),
		MonthDay:       dcl.Int64OrNil(p.GetMonthDay()),
	}
	return obj
}

// ProtoToPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(p *betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) *beta.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{
		WeekOrdinal: dcl.Int64OrNil(p.GetWeekOrdinal()),
		DayOfWeek:   ProtoToOsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(p.GetDayOfWeek()),
	}
	return obj
}

// ProtoToPatchDeploymentRollout converts a PatchDeploymentRollout object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRollout(p *betapb.OsconfigBetaPatchDeploymentRollout) *beta.PatchDeploymentRollout {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRollout{
		Mode:             ProtoToOsconfigBetaPatchDeploymentRolloutModeEnum(p.GetMode()),
		DisruptionBudget: ProtoToOsconfigBetaPatchDeploymentRolloutDisruptionBudget(p.GetDisruptionBudget()),
	}
	return obj
}

// ProtoToPatchDeploymentRolloutDisruptionBudget converts a PatchDeploymentRolloutDisruptionBudget object from its proto representation.
func ProtoToOsconfigBetaPatchDeploymentRolloutDisruptionBudget(p *betapb.OsconfigBetaPatchDeploymentRolloutDisruptionBudget) *beta.PatchDeploymentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &beta.PatchDeploymentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.GetFixed()),
		Percent: dcl.Int64OrNil(p.GetPercent()),
	}
	return obj
}

// ProtoToPatchDeployment converts a PatchDeployment resource from its proto representation.
func ProtoToPatchDeployment(p *betapb.OsconfigBetaPatchDeployment) *beta.PatchDeployment {
	obj := &beta.PatchDeployment{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		InstanceFilter:    ProtoToOsconfigBetaPatchDeploymentInstanceFilter(p.GetInstanceFilter()),
		PatchConfig:       ProtoToOsconfigBetaPatchDeploymentPatchConfig(p.GetPatchConfig()),
		Duration:          dcl.StringOrNil(p.GetDuration()),
		OneTimeSchedule:   ProtoToOsconfigBetaPatchDeploymentOneTimeSchedule(p.GetOneTimeSchedule()),
		RecurringSchedule: ProtoToOsconfigBetaPatchDeploymentRecurringSchedule(p.GetRecurringSchedule()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		LastExecuteTime:   dcl.StringOrNil(p.GetLastExecuteTime()),
		Rollout:           ProtoToOsconfigBetaPatchDeploymentRollout(p.GetRollout()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// PatchDeploymentPatchConfigRebootConfigEnumToProto converts a PatchDeploymentPatchConfigRebootConfigEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnumToProto(e *beta.PatchDeploymentPatchConfigRebootConfigEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum_value["PatchDeploymentPatchConfigRebootConfigEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum(0)
}

// PatchDeploymentPatchConfigAptTypeEnumToProto converts a PatchDeploymentPatchConfigAptTypeEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigAptTypeEnumToProto(e *beta.PatchDeploymentPatchConfigAptTypeEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum_value["PatchDeploymentPatchConfigAptTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum(0)
}

// PatchDeploymentPatchConfigWindowsUpdateClassificationsEnumToProto converts a PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnumToProto(e *beta.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_value["PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(0)
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto(e *beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto(e *beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto(e *beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto(e *beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum) betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum_value["PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(0)
}

// PatchDeploymentRecurringScheduleFrequencyEnumToProto converts a PatchDeploymentRecurringScheduleFrequencyEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnumToProto(e *beta.PatchDeploymentRecurringScheduleFrequencyEnum) betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum_value["PatchDeploymentRecurringScheduleFrequencyEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum(0)
}

// PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto converts a PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto(e *beta.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum) betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum_value["PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(0)
}

// PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto(e *beta.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum) betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum_value["PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(0)
}

// PatchDeploymentRolloutModeEnumToProto converts a PatchDeploymentRolloutModeEnum enum to its proto representation.
func OsconfigBetaPatchDeploymentRolloutModeEnumToProto(e *beta.PatchDeploymentRolloutModeEnum) betapb.OsconfigBetaPatchDeploymentRolloutModeEnum {
	if e == nil {
		return betapb.OsconfigBetaPatchDeploymentRolloutModeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaPatchDeploymentRolloutModeEnum_value["PatchDeploymentRolloutModeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaPatchDeploymentRolloutModeEnum(v)
	}
	return betapb.OsconfigBetaPatchDeploymentRolloutModeEnum(0)
}

// PatchDeploymentInstanceFilterToProto converts a PatchDeploymentInstanceFilter object to its proto representation.
func OsconfigBetaPatchDeploymentInstanceFilterToProto(o *beta.PatchDeploymentInstanceFilter) *betapb.OsconfigBetaPatchDeploymentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentInstanceFilter{}
	p.SetAll(dcl.ValueOrEmptyBool(o.All))
	sGroupLabels := make([]*betapb.OsconfigBetaPatchDeploymentInstanceFilterGroupLabels, len(o.GroupLabels))
	for i, r := range o.GroupLabels {
		sGroupLabels[i] = OsconfigBetaPatchDeploymentInstanceFilterGroupLabelsToProto(&r)
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
func OsconfigBetaPatchDeploymentInstanceFilterGroupLabelsToProto(o *beta.PatchDeploymentInstanceFilterGroupLabels) *betapb.OsconfigBetaPatchDeploymentInstanceFilterGroupLabels {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentInstanceFilterGroupLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// PatchDeploymentPatchConfigToProto converts a PatchDeploymentPatchConfig object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigToProto(o *beta.PatchDeploymentPatchConfig) *betapb.OsconfigBetaPatchDeploymentPatchConfig {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfig{}
	p.SetRebootConfig(OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnumToProto(o.RebootConfig))
	p.SetApt(OsconfigBetaPatchDeploymentPatchConfigAptToProto(o.Apt))
	p.SetYum(OsconfigBetaPatchDeploymentPatchConfigYumToProto(o.Yum))
	p.SetGoo(OsconfigBetaPatchDeploymentPatchConfigGooToProto(o.Goo))
	p.SetZypper(OsconfigBetaPatchDeploymentPatchConfigZypperToProto(o.Zypper))
	p.SetWindowsUpdate(OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateToProto(o.WindowsUpdate))
	p.SetPreStep(OsconfigBetaPatchDeploymentPatchConfigPreStepToProto(o.PreStep))
	p.SetPostStep(OsconfigBetaPatchDeploymentPatchConfigPostStepToProto(o.PostStep))
	return p
}

// PatchDeploymentPatchConfigAptToProto converts a PatchDeploymentPatchConfigApt object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigAptToProto(o *beta.PatchDeploymentPatchConfigApt) *betapb.OsconfigBetaPatchDeploymentPatchConfigApt {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigApt{}
	p.SetType(OsconfigBetaPatchDeploymentPatchConfigAptTypeEnumToProto(o.Type))
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
func OsconfigBetaPatchDeploymentPatchConfigYumToProto(o *beta.PatchDeploymentPatchConfigYum) *betapb.OsconfigBetaPatchDeploymentPatchConfigYum {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigYum{}
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
func OsconfigBetaPatchDeploymentPatchConfigGooToProto(o *beta.PatchDeploymentPatchConfigGoo) *betapb.OsconfigBetaPatchDeploymentPatchConfigGoo {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigGoo{}
	return p
}

// PatchDeploymentPatchConfigZypperToProto converts a PatchDeploymentPatchConfigZypper object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigZypperToProto(o *beta.PatchDeploymentPatchConfigZypper) *betapb.OsconfigBetaPatchDeploymentPatchConfigZypper {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigZypper{}
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
func OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateToProto(o *beta.PatchDeploymentPatchConfigWindowsUpdate) *betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdate {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdate{}
	sClassifications := make([]betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum, len(o.Classifications))
	for i, r := range o.Classifications {
		sClassifications[i] = betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(betapb.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum_value[string(r)])
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
func OsconfigBetaPatchDeploymentPatchConfigPreStepToProto(o *beta.PatchDeploymentPatchConfigPreStep) *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStep {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPreStep{}
	p.SetLinuxExecStepConfig(OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto(o.LinuxExecStepConfig))
	p.SetWindowsExecStepConfig(OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto(o.WindowsExecStepConfig))
	return p
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfig object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigToProto(o *beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig) *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectToProto(o *beta.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfig object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigToProto(o *beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig) *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectToProto(o *beta.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) *betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPostStepToProto converts a PatchDeploymentPatchConfigPostStep object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepToProto(o *beta.PatchDeploymentPatchConfigPostStep) *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStep {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPostStep{}
	p.SetLinuxExecStepConfig(OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto(o.LinuxExecStepConfig))
	p.SetWindowsExecStepConfig(OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto(o.WindowsExecStepConfig))
	return p
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfig object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigToProto(o *beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig) *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectToProto(o *beta.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfig object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigToProto(o *beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig) *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig{}
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetGcsObject(OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto(o.GcsObject))
	p.SetInterpreter(OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumToProto(o.Interpreter))
	sAllowedSuccessCodes := make([]int64, len(o.AllowedSuccessCodes))
	for i, r := range o.AllowedSuccessCodes {
		sAllowedSuccessCodes[i] = r
	}
	p.SetAllowedSuccessCodes(sAllowedSuccessCodes)
	return p
}

// PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto converts a PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject object to its proto representation.
func OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectToProto(o *beta.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject) *betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGenerationNumber(dcl.ValueOrEmptyInt64(o.GenerationNumber))
	return p
}

// PatchDeploymentOneTimeScheduleToProto converts a PatchDeploymentOneTimeSchedule object to its proto representation.
func OsconfigBetaPatchDeploymentOneTimeScheduleToProto(o *beta.PatchDeploymentOneTimeSchedule) *betapb.OsconfigBetaPatchDeploymentOneTimeSchedule {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentOneTimeSchedule{}
	p.SetExecuteTime(dcl.ValueOrEmptyString(o.ExecuteTime))
	return p
}

// PatchDeploymentRecurringScheduleToProto converts a PatchDeploymentRecurringSchedule object to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleToProto(o *beta.PatchDeploymentRecurringSchedule) *betapb.OsconfigBetaPatchDeploymentRecurringSchedule {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRecurringSchedule{}
	p.SetTimeZone(OsconfigBetaPatchDeploymentRecurringScheduleTimeZoneToProto(o.TimeZone))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyString(o.EndTime))
	p.SetTimeOfDay(OsconfigBetaPatchDeploymentRecurringScheduleTimeOfDayToProto(o.TimeOfDay))
	p.SetFrequency(OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnumToProto(o.Frequency))
	p.SetWeekly(OsconfigBetaPatchDeploymentRecurringScheduleWeeklyToProto(o.Weekly))
	p.SetMonthly(OsconfigBetaPatchDeploymentRecurringScheduleMonthlyToProto(o.Monthly))
	p.SetLastExecuteTime(dcl.ValueOrEmptyString(o.LastExecuteTime))
	p.SetNextExecuteTime(dcl.ValueOrEmptyString(o.NextExecuteTime))
	return p
}

// PatchDeploymentRecurringScheduleTimeZoneToProto converts a PatchDeploymentRecurringScheduleTimeZone object to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleTimeZoneToProto(o *beta.PatchDeploymentRecurringScheduleTimeZone) *betapb.OsconfigBetaPatchDeploymentRecurringScheduleTimeZone {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRecurringScheduleTimeZone{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// PatchDeploymentRecurringScheduleTimeOfDayToProto converts a PatchDeploymentRecurringScheduleTimeOfDay object to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleTimeOfDayToProto(o *beta.PatchDeploymentRecurringScheduleTimeOfDay) *betapb.OsconfigBetaPatchDeploymentRecurringScheduleTimeOfDay {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRecurringScheduleTimeOfDay{}
	p.SetHours(dcl.ValueOrEmptyInt64(o.Hours))
	p.SetMinutes(dcl.ValueOrEmptyInt64(o.Minutes))
	p.SetSeconds(dcl.ValueOrEmptyInt64(o.Seconds))
	p.SetNanos(dcl.ValueOrEmptyInt64(o.Nanos))
	return p
}

// PatchDeploymentRecurringScheduleWeeklyToProto converts a PatchDeploymentRecurringScheduleWeekly object to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleWeeklyToProto(o *beta.PatchDeploymentRecurringScheduleWeekly) *betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeekly {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRecurringScheduleWeekly{}
	p.SetDayOfWeek(OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumToProto(o.DayOfWeek))
	return p
}

// PatchDeploymentRecurringScheduleMonthlyToProto converts a PatchDeploymentRecurringScheduleMonthly object to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleMonthlyToProto(o *beta.PatchDeploymentRecurringScheduleMonthly) *betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthly {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthly{}
	p.SetWeekDayOfMonth(OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto(o.WeekDayOfMonth))
	p.SetMonthDay(dcl.ValueOrEmptyInt64(o.MonthDay))
	return p
}

// PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto converts a PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth object to its proto representation.
func OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthToProto(o *beta.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) *betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{}
	p.SetWeekOrdinal(dcl.ValueOrEmptyInt64(o.WeekOrdinal))
	p.SetDayOfWeek(OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumToProto(o.DayOfWeek))
	return p
}

// PatchDeploymentRolloutToProto converts a PatchDeploymentRollout object to its proto representation.
func OsconfigBetaPatchDeploymentRolloutToProto(o *beta.PatchDeploymentRollout) *betapb.OsconfigBetaPatchDeploymentRollout {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRollout{}
	p.SetMode(OsconfigBetaPatchDeploymentRolloutModeEnumToProto(o.Mode))
	p.SetDisruptionBudget(OsconfigBetaPatchDeploymentRolloutDisruptionBudgetToProto(o.DisruptionBudget))
	return p
}

// PatchDeploymentRolloutDisruptionBudgetToProto converts a PatchDeploymentRolloutDisruptionBudget object to its proto representation.
func OsconfigBetaPatchDeploymentRolloutDisruptionBudgetToProto(o *beta.PatchDeploymentRolloutDisruptionBudget) *betapb.OsconfigBetaPatchDeploymentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaPatchDeploymentRolloutDisruptionBudget{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	return p
}

// PatchDeploymentToProto converts a PatchDeployment resource to its proto representation.
func PatchDeploymentToProto(resource *beta.PatchDeployment) *betapb.OsconfigBetaPatchDeployment {
	p := &betapb.OsconfigBetaPatchDeployment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInstanceFilter(OsconfigBetaPatchDeploymentInstanceFilterToProto(resource.InstanceFilter))
	p.SetPatchConfig(OsconfigBetaPatchDeploymentPatchConfigToProto(resource.PatchConfig))
	p.SetDuration(dcl.ValueOrEmptyString(resource.Duration))
	p.SetOneTimeSchedule(OsconfigBetaPatchDeploymentOneTimeScheduleToProto(resource.OneTimeSchedule))
	p.SetRecurringSchedule(OsconfigBetaPatchDeploymentRecurringScheduleToProto(resource.RecurringSchedule))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetLastExecuteTime(dcl.ValueOrEmptyString(resource.LastExecuteTime))
	p.SetRollout(OsconfigBetaPatchDeploymentRolloutToProto(resource.Rollout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyPatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Apply() method.
func (s *PatchDeploymentServer) applyPatchDeployment(ctx context.Context, c *beta.Client, request *betapb.ApplyOsconfigBetaPatchDeploymentRequest) (*betapb.OsconfigBetaPatchDeployment, error) {
	p := ProtoToPatchDeployment(request.GetResource())
	res, err := c.ApplyPatchDeployment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PatchDeploymentToProto(res)
	return r, nil
}

// applyOsconfigBetaPatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Apply() method.
func (s *PatchDeploymentServer) ApplyOsconfigBetaPatchDeployment(ctx context.Context, request *betapb.ApplyOsconfigBetaPatchDeploymentRequest) (*betapb.OsconfigBetaPatchDeployment, error) {
	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPatchDeployment(ctx, cl, request)
}

// DeletePatchDeployment handles the gRPC request by passing it to the underlying PatchDeployment Delete() method.
func (s *PatchDeploymentServer) DeleteOsconfigBetaPatchDeployment(ctx context.Context, request *betapb.DeleteOsconfigBetaPatchDeploymentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePatchDeployment(ctx, ProtoToPatchDeployment(request.GetResource()))

}

// ListOsconfigBetaPatchDeployment handles the gRPC request by passing it to the underlying PatchDeploymentList() method.
func (s *PatchDeploymentServer) ListOsconfigBetaPatchDeployment(ctx context.Context, request *betapb.ListOsconfigBetaPatchDeploymentRequest) (*betapb.ListOsconfigBetaPatchDeploymentResponse, error) {
	cl, err := createConfigPatchDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPatchDeployment(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.OsconfigBetaPatchDeployment
	for _, r := range resources.Items {
		rp := PatchDeploymentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListOsconfigBetaPatchDeploymentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPatchDeployment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
