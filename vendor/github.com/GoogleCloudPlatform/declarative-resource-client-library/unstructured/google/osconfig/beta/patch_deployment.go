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
package osconfig

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type PatchDeployment struct{}

func PatchDeploymentToUnstructured(r *dclService.PatchDeployment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "osconfig",
			Version: "beta",
			Type:    "PatchDeployment",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Duration != nil {
		u.Object["duration"] = *r.Duration
	}
	if r.InstanceFilter != nil && r.InstanceFilter != dclService.EmptyPatchDeploymentInstanceFilter {
		rInstanceFilter := make(map[string]interface{})
		if r.InstanceFilter.All != nil {
			rInstanceFilter["all"] = *r.InstanceFilter.All
		}
		var rInstanceFilterGroupLabels []interface{}
		for _, rInstanceFilterGroupLabelsVal := range r.InstanceFilter.GroupLabels {
			rInstanceFilterGroupLabelsObject := make(map[string]interface{})
			if rInstanceFilterGroupLabelsVal.Labels != nil {
				rInstanceFilterGroupLabelsValLabels := make(map[string]interface{})
				for k, v := range rInstanceFilterGroupLabelsVal.Labels {
					rInstanceFilterGroupLabelsValLabels[k] = v
				}
				rInstanceFilterGroupLabelsObject["labels"] = rInstanceFilterGroupLabelsValLabels
			}
			rInstanceFilterGroupLabels = append(rInstanceFilterGroupLabels, rInstanceFilterGroupLabelsObject)
		}
		rInstanceFilter["groupLabels"] = rInstanceFilterGroupLabels
		var rInstanceFilterInstanceNamePrefixes []interface{}
		for _, rInstanceFilterInstanceNamePrefixesVal := range r.InstanceFilter.InstanceNamePrefixes {
			rInstanceFilterInstanceNamePrefixes = append(rInstanceFilterInstanceNamePrefixes, rInstanceFilterInstanceNamePrefixesVal)
		}
		rInstanceFilter["instanceNamePrefixes"] = rInstanceFilterInstanceNamePrefixes
		var rInstanceFilterInstances []interface{}
		for _, rInstanceFilterInstancesVal := range r.InstanceFilter.Instances {
			rInstanceFilterInstances = append(rInstanceFilterInstances, rInstanceFilterInstancesVal)
		}
		rInstanceFilter["instances"] = rInstanceFilterInstances
		var rInstanceFilterZones []interface{}
		for _, rInstanceFilterZonesVal := range r.InstanceFilter.Zones {
			rInstanceFilterZones = append(rInstanceFilterZones, rInstanceFilterZonesVal)
		}
		rInstanceFilter["zones"] = rInstanceFilterZones
		u.Object["instanceFilter"] = rInstanceFilter
	}
	if r.LastExecuteTime != nil {
		u.Object["lastExecuteTime"] = *r.LastExecuteTime
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.OneTimeSchedule != nil && r.OneTimeSchedule != dclService.EmptyPatchDeploymentOneTimeSchedule {
		rOneTimeSchedule := make(map[string]interface{})
		if r.OneTimeSchedule.ExecuteTime != nil {
			rOneTimeSchedule["executeTime"] = *r.OneTimeSchedule.ExecuteTime
		}
		u.Object["oneTimeSchedule"] = rOneTimeSchedule
	}
	if r.PatchConfig != nil && r.PatchConfig != dclService.EmptyPatchDeploymentPatchConfig {
		rPatchConfig := make(map[string]interface{})
		if r.PatchConfig.Apt != nil && r.PatchConfig.Apt != dclService.EmptyPatchDeploymentPatchConfigApt {
			rPatchConfigApt := make(map[string]interface{})
			var rPatchConfigAptExcludes []interface{}
			for _, rPatchConfigAptExcludesVal := range r.PatchConfig.Apt.Excludes {
				rPatchConfigAptExcludes = append(rPatchConfigAptExcludes, rPatchConfigAptExcludesVal)
			}
			rPatchConfigApt["excludes"] = rPatchConfigAptExcludes
			var rPatchConfigAptExclusivePackages []interface{}
			for _, rPatchConfigAptExclusivePackagesVal := range r.PatchConfig.Apt.ExclusivePackages {
				rPatchConfigAptExclusivePackages = append(rPatchConfigAptExclusivePackages, rPatchConfigAptExclusivePackagesVal)
			}
			rPatchConfigApt["exclusivePackages"] = rPatchConfigAptExclusivePackages
			if r.PatchConfig.Apt.Type != nil {
				rPatchConfigApt["type"] = string(*r.PatchConfig.Apt.Type)
			}
			rPatchConfig["apt"] = rPatchConfigApt
		}
		if r.PatchConfig.Goo != nil && r.PatchConfig.Goo != dclService.EmptyPatchDeploymentPatchConfigGoo {
			rPatchConfigGoo := make(map[string]interface{})
			rPatchConfig["goo"] = rPatchConfigGoo
		}
		if r.PatchConfig.PostStep != nil && r.PatchConfig.PostStep != dclService.EmptyPatchDeploymentPatchConfigPostStep {
			rPatchConfigPostStep := make(map[string]interface{})
			if r.PatchConfig.PostStep.LinuxExecStepConfig != nil && r.PatchConfig.PostStep.LinuxExecStepConfig != dclService.EmptyPatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
				rPatchConfigPostStepLinuxExecStepConfig := make(map[string]interface{})
				var rPatchConfigPostStepLinuxExecStepConfigAllowedSuccessCodes []interface{}
				for _, rPatchConfigPostStepLinuxExecStepConfigAllowedSuccessCodesVal := range r.PatchConfig.PostStep.LinuxExecStepConfig.AllowedSuccessCodes {
					rPatchConfigPostStepLinuxExecStepConfigAllowedSuccessCodes = append(rPatchConfigPostStepLinuxExecStepConfigAllowedSuccessCodes, rPatchConfigPostStepLinuxExecStepConfigAllowedSuccessCodesVal)
				}
				rPatchConfigPostStepLinuxExecStepConfig["allowedSuccessCodes"] = rPatchConfigPostStepLinuxExecStepConfigAllowedSuccessCodes
				if r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject != nil && r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject != dclService.EmptyPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
					rPatchConfigPostStepLinuxExecStepConfigGcsObject := make(map[string]interface{})
					if r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Bucket != nil {
						rPatchConfigPostStepLinuxExecStepConfigGcsObject["bucket"] = *r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Bucket
					}
					if r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.GenerationNumber != nil {
						rPatchConfigPostStepLinuxExecStepConfigGcsObject["generationNumber"] = *r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.GenerationNumber
					}
					if r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Object != nil {
						rPatchConfigPostStepLinuxExecStepConfigGcsObject["object"] = *r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Object
					}
					rPatchConfigPostStepLinuxExecStepConfig["gcsObject"] = rPatchConfigPostStepLinuxExecStepConfigGcsObject
				}
				if r.PatchConfig.PostStep.LinuxExecStepConfig.Interpreter != nil {
					rPatchConfigPostStepLinuxExecStepConfig["interpreter"] = string(*r.PatchConfig.PostStep.LinuxExecStepConfig.Interpreter)
				}
				if r.PatchConfig.PostStep.LinuxExecStepConfig.LocalPath != nil {
					rPatchConfigPostStepLinuxExecStepConfig["localPath"] = *r.PatchConfig.PostStep.LinuxExecStepConfig.LocalPath
				}
				rPatchConfigPostStep["linuxExecStepConfig"] = rPatchConfigPostStepLinuxExecStepConfig
			}
			if r.PatchConfig.PostStep.WindowsExecStepConfig != nil && r.PatchConfig.PostStep.WindowsExecStepConfig != dclService.EmptyPatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
				rPatchConfigPostStepWindowsExecStepConfig := make(map[string]interface{})
				var rPatchConfigPostStepWindowsExecStepConfigAllowedSuccessCodes []interface{}
				for _, rPatchConfigPostStepWindowsExecStepConfigAllowedSuccessCodesVal := range r.PatchConfig.PostStep.WindowsExecStepConfig.AllowedSuccessCodes {
					rPatchConfigPostStepWindowsExecStepConfigAllowedSuccessCodes = append(rPatchConfigPostStepWindowsExecStepConfigAllowedSuccessCodes, rPatchConfigPostStepWindowsExecStepConfigAllowedSuccessCodesVal)
				}
				rPatchConfigPostStepWindowsExecStepConfig["allowedSuccessCodes"] = rPatchConfigPostStepWindowsExecStepConfigAllowedSuccessCodes
				if r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject != nil && r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject != dclService.EmptyPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject {
					rPatchConfigPostStepWindowsExecStepConfigGcsObject := make(map[string]interface{})
					if r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Bucket != nil {
						rPatchConfigPostStepWindowsExecStepConfigGcsObject["bucket"] = *r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Bucket
					}
					if r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.GenerationNumber != nil {
						rPatchConfigPostStepWindowsExecStepConfigGcsObject["generationNumber"] = *r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.GenerationNumber
					}
					if r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Object != nil {
						rPatchConfigPostStepWindowsExecStepConfigGcsObject["object"] = *r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Object
					}
					rPatchConfigPostStepWindowsExecStepConfig["gcsObject"] = rPatchConfigPostStepWindowsExecStepConfigGcsObject
				}
				if r.PatchConfig.PostStep.WindowsExecStepConfig.Interpreter != nil {
					rPatchConfigPostStepWindowsExecStepConfig["interpreter"] = string(*r.PatchConfig.PostStep.WindowsExecStepConfig.Interpreter)
				}
				if r.PatchConfig.PostStep.WindowsExecStepConfig.LocalPath != nil {
					rPatchConfigPostStepWindowsExecStepConfig["localPath"] = *r.PatchConfig.PostStep.WindowsExecStepConfig.LocalPath
				}
				rPatchConfigPostStep["windowsExecStepConfig"] = rPatchConfigPostStepWindowsExecStepConfig
			}
			rPatchConfig["postStep"] = rPatchConfigPostStep
		}
		if r.PatchConfig.PreStep != nil && r.PatchConfig.PreStep != dclService.EmptyPatchDeploymentPatchConfigPreStep {
			rPatchConfigPreStep := make(map[string]interface{})
			if r.PatchConfig.PreStep.LinuxExecStepConfig != nil && r.PatchConfig.PreStep.LinuxExecStepConfig != dclService.EmptyPatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
				rPatchConfigPreStepLinuxExecStepConfig := make(map[string]interface{})
				var rPatchConfigPreStepLinuxExecStepConfigAllowedSuccessCodes []interface{}
				for _, rPatchConfigPreStepLinuxExecStepConfigAllowedSuccessCodesVal := range r.PatchConfig.PreStep.LinuxExecStepConfig.AllowedSuccessCodes {
					rPatchConfigPreStepLinuxExecStepConfigAllowedSuccessCodes = append(rPatchConfigPreStepLinuxExecStepConfigAllowedSuccessCodes, rPatchConfigPreStepLinuxExecStepConfigAllowedSuccessCodesVal)
				}
				rPatchConfigPreStepLinuxExecStepConfig["allowedSuccessCodes"] = rPatchConfigPreStepLinuxExecStepConfigAllowedSuccessCodes
				if r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject != nil && r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject != dclService.EmptyPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
					rPatchConfigPreStepLinuxExecStepConfigGcsObject := make(map[string]interface{})
					if r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Bucket != nil {
						rPatchConfigPreStepLinuxExecStepConfigGcsObject["bucket"] = *r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Bucket
					}
					if r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.GenerationNumber != nil {
						rPatchConfigPreStepLinuxExecStepConfigGcsObject["generationNumber"] = *r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.GenerationNumber
					}
					if r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Object != nil {
						rPatchConfigPreStepLinuxExecStepConfigGcsObject["object"] = *r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Object
					}
					rPatchConfigPreStepLinuxExecStepConfig["gcsObject"] = rPatchConfigPreStepLinuxExecStepConfigGcsObject
				}
				if r.PatchConfig.PreStep.LinuxExecStepConfig.Interpreter != nil {
					rPatchConfigPreStepLinuxExecStepConfig["interpreter"] = string(*r.PatchConfig.PreStep.LinuxExecStepConfig.Interpreter)
				}
				if r.PatchConfig.PreStep.LinuxExecStepConfig.LocalPath != nil {
					rPatchConfigPreStepLinuxExecStepConfig["localPath"] = *r.PatchConfig.PreStep.LinuxExecStepConfig.LocalPath
				}
				rPatchConfigPreStep["linuxExecStepConfig"] = rPatchConfigPreStepLinuxExecStepConfig
			}
			if r.PatchConfig.PreStep.WindowsExecStepConfig != nil && r.PatchConfig.PreStep.WindowsExecStepConfig != dclService.EmptyPatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
				rPatchConfigPreStepWindowsExecStepConfig := make(map[string]interface{})
				var rPatchConfigPreStepWindowsExecStepConfigAllowedSuccessCodes []interface{}
				for _, rPatchConfigPreStepWindowsExecStepConfigAllowedSuccessCodesVal := range r.PatchConfig.PreStep.WindowsExecStepConfig.AllowedSuccessCodes {
					rPatchConfigPreStepWindowsExecStepConfigAllowedSuccessCodes = append(rPatchConfigPreStepWindowsExecStepConfigAllowedSuccessCodes, rPatchConfigPreStepWindowsExecStepConfigAllowedSuccessCodesVal)
				}
				rPatchConfigPreStepWindowsExecStepConfig["allowedSuccessCodes"] = rPatchConfigPreStepWindowsExecStepConfigAllowedSuccessCodes
				if r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject != nil && r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject != dclService.EmptyPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
					rPatchConfigPreStepWindowsExecStepConfigGcsObject := make(map[string]interface{})
					if r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Bucket != nil {
						rPatchConfigPreStepWindowsExecStepConfigGcsObject["bucket"] = *r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Bucket
					}
					if r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.GenerationNumber != nil {
						rPatchConfigPreStepWindowsExecStepConfigGcsObject["generationNumber"] = *r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.GenerationNumber
					}
					if r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Object != nil {
						rPatchConfigPreStepWindowsExecStepConfigGcsObject["object"] = *r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Object
					}
					rPatchConfigPreStepWindowsExecStepConfig["gcsObject"] = rPatchConfigPreStepWindowsExecStepConfigGcsObject
				}
				if r.PatchConfig.PreStep.WindowsExecStepConfig.Interpreter != nil {
					rPatchConfigPreStepWindowsExecStepConfig["interpreter"] = string(*r.PatchConfig.PreStep.WindowsExecStepConfig.Interpreter)
				}
				if r.PatchConfig.PreStep.WindowsExecStepConfig.LocalPath != nil {
					rPatchConfigPreStepWindowsExecStepConfig["localPath"] = *r.PatchConfig.PreStep.WindowsExecStepConfig.LocalPath
				}
				rPatchConfigPreStep["windowsExecStepConfig"] = rPatchConfigPreStepWindowsExecStepConfig
			}
			rPatchConfig["preStep"] = rPatchConfigPreStep
		}
		if r.PatchConfig.RebootConfig != nil {
			rPatchConfig["rebootConfig"] = string(*r.PatchConfig.RebootConfig)
		}
		if r.PatchConfig.WindowsUpdate != nil && r.PatchConfig.WindowsUpdate != dclService.EmptyPatchDeploymentPatchConfigWindowsUpdate {
			rPatchConfigWindowsUpdate := make(map[string]interface{})
			var rPatchConfigWindowsUpdateClassifications []interface{}
			for _, rPatchConfigWindowsUpdateClassificationsVal := range r.PatchConfig.WindowsUpdate.Classifications {
				rPatchConfigWindowsUpdateClassifications = append(rPatchConfigWindowsUpdateClassifications, string(rPatchConfigWindowsUpdateClassificationsVal))
			}
			rPatchConfigWindowsUpdate["classifications"] = rPatchConfigWindowsUpdateClassifications
			var rPatchConfigWindowsUpdateExcludes []interface{}
			for _, rPatchConfigWindowsUpdateExcludesVal := range r.PatchConfig.WindowsUpdate.Excludes {
				rPatchConfigWindowsUpdateExcludes = append(rPatchConfigWindowsUpdateExcludes, rPatchConfigWindowsUpdateExcludesVal)
			}
			rPatchConfigWindowsUpdate["excludes"] = rPatchConfigWindowsUpdateExcludes
			var rPatchConfigWindowsUpdateExclusivePatches []interface{}
			for _, rPatchConfigWindowsUpdateExclusivePatchesVal := range r.PatchConfig.WindowsUpdate.ExclusivePatches {
				rPatchConfigWindowsUpdateExclusivePatches = append(rPatchConfigWindowsUpdateExclusivePatches, rPatchConfigWindowsUpdateExclusivePatchesVal)
			}
			rPatchConfigWindowsUpdate["exclusivePatches"] = rPatchConfigWindowsUpdateExclusivePatches
			rPatchConfig["windowsUpdate"] = rPatchConfigWindowsUpdate
		}
		if r.PatchConfig.Yum != nil && r.PatchConfig.Yum != dclService.EmptyPatchDeploymentPatchConfigYum {
			rPatchConfigYum := make(map[string]interface{})
			var rPatchConfigYumExcludes []interface{}
			for _, rPatchConfigYumExcludesVal := range r.PatchConfig.Yum.Excludes {
				rPatchConfigYumExcludes = append(rPatchConfigYumExcludes, rPatchConfigYumExcludesVal)
			}
			rPatchConfigYum["excludes"] = rPatchConfigYumExcludes
			var rPatchConfigYumExclusivePackages []interface{}
			for _, rPatchConfigYumExclusivePackagesVal := range r.PatchConfig.Yum.ExclusivePackages {
				rPatchConfigYumExclusivePackages = append(rPatchConfigYumExclusivePackages, rPatchConfigYumExclusivePackagesVal)
			}
			rPatchConfigYum["exclusivePackages"] = rPatchConfigYumExclusivePackages
			if r.PatchConfig.Yum.Minimal != nil {
				rPatchConfigYum["minimal"] = *r.PatchConfig.Yum.Minimal
			}
			if r.PatchConfig.Yum.Security != nil {
				rPatchConfigYum["security"] = *r.PatchConfig.Yum.Security
			}
			rPatchConfig["yum"] = rPatchConfigYum
		}
		if r.PatchConfig.Zypper != nil && r.PatchConfig.Zypper != dclService.EmptyPatchDeploymentPatchConfigZypper {
			rPatchConfigZypper := make(map[string]interface{})
			var rPatchConfigZypperCategories []interface{}
			for _, rPatchConfigZypperCategoriesVal := range r.PatchConfig.Zypper.Categories {
				rPatchConfigZypperCategories = append(rPatchConfigZypperCategories, rPatchConfigZypperCategoriesVal)
			}
			rPatchConfigZypper["categories"] = rPatchConfigZypperCategories
			var rPatchConfigZypperExcludes []interface{}
			for _, rPatchConfigZypperExcludesVal := range r.PatchConfig.Zypper.Excludes {
				rPatchConfigZypperExcludes = append(rPatchConfigZypperExcludes, rPatchConfigZypperExcludesVal)
			}
			rPatchConfigZypper["excludes"] = rPatchConfigZypperExcludes
			var rPatchConfigZypperExclusivePatches []interface{}
			for _, rPatchConfigZypperExclusivePatchesVal := range r.PatchConfig.Zypper.ExclusivePatches {
				rPatchConfigZypperExclusivePatches = append(rPatchConfigZypperExclusivePatches, rPatchConfigZypperExclusivePatchesVal)
			}
			rPatchConfigZypper["exclusivePatches"] = rPatchConfigZypperExclusivePatches
			var rPatchConfigZypperSeverities []interface{}
			for _, rPatchConfigZypperSeveritiesVal := range r.PatchConfig.Zypper.Severities {
				rPatchConfigZypperSeverities = append(rPatchConfigZypperSeverities, rPatchConfigZypperSeveritiesVal)
			}
			rPatchConfigZypper["severities"] = rPatchConfigZypperSeverities
			if r.PatchConfig.Zypper.WithOptional != nil {
				rPatchConfigZypper["withOptional"] = *r.PatchConfig.Zypper.WithOptional
			}
			if r.PatchConfig.Zypper.WithUpdate != nil {
				rPatchConfigZypper["withUpdate"] = *r.PatchConfig.Zypper.WithUpdate
			}
			rPatchConfig["zypper"] = rPatchConfigZypper
		}
		u.Object["patchConfig"] = rPatchConfig
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RecurringSchedule != nil && r.RecurringSchedule != dclService.EmptyPatchDeploymentRecurringSchedule {
		rRecurringSchedule := make(map[string]interface{})
		if r.RecurringSchedule.EndTime != nil {
			rRecurringSchedule["endTime"] = *r.RecurringSchedule.EndTime
		}
		if r.RecurringSchedule.Frequency != nil {
			rRecurringSchedule["frequency"] = string(*r.RecurringSchedule.Frequency)
		}
		if r.RecurringSchedule.LastExecuteTime != nil {
			rRecurringSchedule["lastExecuteTime"] = *r.RecurringSchedule.LastExecuteTime
		}
		if r.RecurringSchedule.Monthly != nil && r.RecurringSchedule.Monthly != dclService.EmptyPatchDeploymentRecurringScheduleMonthly {
			rRecurringScheduleMonthly := make(map[string]interface{})
			if r.RecurringSchedule.Monthly.MonthDay != nil {
				rRecurringScheduleMonthly["monthDay"] = *r.RecurringSchedule.Monthly.MonthDay
			}
			if r.RecurringSchedule.Monthly.WeekDayOfMonth != nil && r.RecurringSchedule.Monthly.WeekDayOfMonth != dclService.EmptyPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
				rRecurringScheduleMonthlyWeekDayOfMonth := make(map[string]interface{})
				if r.RecurringSchedule.Monthly.WeekDayOfMonth.DayOfWeek != nil {
					rRecurringScheduleMonthlyWeekDayOfMonth["dayOfWeek"] = string(*r.RecurringSchedule.Monthly.WeekDayOfMonth.DayOfWeek)
				}
				if r.RecurringSchedule.Monthly.WeekDayOfMonth.WeekOrdinal != nil {
					rRecurringScheduleMonthlyWeekDayOfMonth["weekOrdinal"] = *r.RecurringSchedule.Monthly.WeekDayOfMonth.WeekOrdinal
				}
				rRecurringScheduleMonthly["weekDayOfMonth"] = rRecurringScheduleMonthlyWeekDayOfMonth
			}
			rRecurringSchedule["monthly"] = rRecurringScheduleMonthly
		}
		if r.RecurringSchedule.NextExecuteTime != nil {
			rRecurringSchedule["nextExecuteTime"] = *r.RecurringSchedule.NextExecuteTime
		}
		if r.RecurringSchedule.StartTime != nil {
			rRecurringSchedule["startTime"] = *r.RecurringSchedule.StartTime
		}
		if r.RecurringSchedule.TimeOfDay != nil && r.RecurringSchedule.TimeOfDay != dclService.EmptyPatchDeploymentRecurringScheduleTimeOfDay {
			rRecurringScheduleTimeOfDay := make(map[string]interface{})
			if r.RecurringSchedule.TimeOfDay.Hours != nil {
				rRecurringScheduleTimeOfDay["hours"] = *r.RecurringSchedule.TimeOfDay.Hours
			}
			if r.RecurringSchedule.TimeOfDay.Minutes != nil {
				rRecurringScheduleTimeOfDay["minutes"] = *r.RecurringSchedule.TimeOfDay.Minutes
			}
			if r.RecurringSchedule.TimeOfDay.Nanos != nil {
				rRecurringScheduleTimeOfDay["nanos"] = *r.RecurringSchedule.TimeOfDay.Nanos
			}
			if r.RecurringSchedule.TimeOfDay.Seconds != nil {
				rRecurringScheduleTimeOfDay["seconds"] = *r.RecurringSchedule.TimeOfDay.Seconds
			}
			rRecurringSchedule["timeOfDay"] = rRecurringScheduleTimeOfDay
		}
		if r.RecurringSchedule.TimeZone != nil && r.RecurringSchedule.TimeZone != dclService.EmptyPatchDeploymentRecurringScheduleTimeZone {
			rRecurringScheduleTimeZone := make(map[string]interface{})
			if r.RecurringSchedule.TimeZone.Id != nil {
				rRecurringScheduleTimeZone["id"] = *r.RecurringSchedule.TimeZone.Id
			}
			if r.RecurringSchedule.TimeZone.Version != nil {
				rRecurringScheduleTimeZone["version"] = *r.RecurringSchedule.TimeZone.Version
			}
			rRecurringSchedule["timeZone"] = rRecurringScheduleTimeZone
		}
		if r.RecurringSchedule.Weekly != nil && r.RecurringSchedule.Weekly != dclService.EmptyPatchDeploymentRecurringScheduleWeekly {
			rRecurringScheduleWeekly := make(map[string]interface{})
			if r.RecurringSchedule.Weekly.DayOfWeek != nil {
				rRecurringScheduleWeekly["dayOfWeek"] = string(*r.RecurringSchedule.Weekly.DayOfWeek)
			}
			rRecurringSchedule["weekly"] = rRecurringScheduleWeekly
		}
		u.Object["recurringSchedule"] = rRecurringSchedule
	}
	if r.Rollout != nil && r.Rollout != dclService.EmptyPatchDeploymentRollout {
		rRollout := make(map[string]interface{})
		if r.Rollout.DisruptionBudget != nil && r.Rollout.DisruptionBudget != dclService.EmptyPatchDeploymentRolloutDisruptionBudget {
			rRolloutDisruptionBudget := make(map[string]interface{})
			if r.Rollout.DisruptionBudget.Fixed != nil {
				rRolloutDisruptionBudget["fixed"] = *r.Rollout.DisruptionBudget.Fixed
			}
			if r.Rollout.DisruptionBudget.Percent != nil {
				rRolloutDisruptionBudget["percent"] = *r.Rollout.DisruptionBudget.Percent
			}
			rRollout["disruptionBudget"] = rRolloutDisruptionBudget
		}
		if r.Rollout.Mode != nil {
			rRollout["mode"] = string(*r.Rollout.Mode)
		}
		u.Object["rollout"] = rRollout
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToPatchDeployment(u *unstructured.Resource) (*dclService.PatchDeployment, error) {
	r := &dclService.PatchDeployment{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["duration"]; ok {
		if s, ok := u.Object["duration"].(string); ok {
			r.Duration = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Duration: expected string")
		}
	}
	if _, ok := u.Object["instanceFilter"]; ok {
		if rInstanceFilter, ok := u.Object["instanceFilter"].(map[string]interface{}); ok {
			r.InstanceFilter = &dclService.PatchDeploymentInstanceFilter{}
			if _, ok := rInstanceFilter["all"]; ok {
				if b, ok := rInstanceFilter["all"].(bool); ok {
					r.InstanceFilter.All = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.All: expected bool")
				}
			}
			if _, ok := rInstanceFilter["groupLabels"]; ok {
				if s, ok := rInstanceFilter["groupLabels"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInstanceFilterGroupLabels dclService.PatchDeploymentInstanceFilterGroupLabels
							if _, ok := objval["labels"]; ok {
								if rInstanceFilterGroupLabelsLabels, ok := objval["labels"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rInstanceFilterGroupLabelsLabels {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rInstanceFilterGroupLabels.Labels = m
								} else {
									return nil, fmt.Errorf("rInstanceFilterGroupLabels.Labels: expected map[string]interface{}")
								}
							}
							r.InstanceFilter.GroupLabels = append(r.InstanceFilter.GroupLabels, rInstanceFilterGroupLabels)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.GroupLabels: expected []interface{}")
				}
			}
			if _, ok := rInstanceFilter["instanceNamePrefixes"]; ok {
				if s, ok := rInstanceFilter["instanceNamePrefixes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.InstanceFilter.InstanceNamePrefixes = append(r.InstanceFilter.InstanceNamePrefixes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.InstanceNamePrefixes: expected []interface{}")
				}
			}
			if _, ok := rInstanceFilter["instances"]; ok {
				if s, ok := rInstanceFilter["instances"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.InstanceFilter.Instances = append(r.InstanceFilter.Instances, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.Instances: expected []interface{}")
				}
			}
			if _, ok := rInstanceFilter["zones"]; ok {
				if s, ok := rInstanceFilter["zones"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.InstanceFilter.Zones = append(r.InstanceFilter.Zones, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.Zones: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.InstanceFilter: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["lastExecuteTime"]; ok {
		if s, ok := u.Object["lastExecuteTime"].(string); ok {
			r.LastExecuteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LastExecuteTime: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["oneTimeSchedule"]; ok {
		if rOneTimeSchedule, ok := u.Object["oneTimeSchedule"].(map[string]interface{}); ok {
			r.OneTimeSchedule = &dclService.PatchDeploymentOneTimeSchedule{}
			if _, ok := rOneTimeSchedule["executeTime"]; ok {
				if s, ok := rOneTimeSchedule["executeTime"].(string); ok {
					r.OneTimeSchedule.ExecuteTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.OneTimeSchedule.ExecuteTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.OneTimeSchedule: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["patchConfig"]; ok {
		if rPatchConfig, ok := u.Object["patchConfig"].(map[string]interface{}); ok {
			r.PatchConfig = &dclService.PatchDeploymentPatchConfig{}
			if _, ok := rPatchConfig["apt"]; ok {
				if rPatchConfigApt, ok := rPatchConfig["apt"].(map[string]interface{}); ok {
					r.PatchConfig.Apt = &dclService.PatchDeploymentPatchConfigApt{}
					if _, ok := rPatchConfigApt["excludes"]; ok {
						if s, ok := rPatchConfigApt["excludes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Apt.Excludes = append(r.PatchConfig.Apt.Excludes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Apt.Excludes: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigApt["exclusivePackages"]; ok {
						if s, ok := rPatchConfigApt["exclusivePackages"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Apt.ExclusivePackages = append(r.PatchConfig.Apt.ExclusivePackages, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Apt.ExclusivePackages: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigApt["type"]; ok {
						if s, ok := rPatchConfigApt["type"].(string); ok {
							r.PatchConfig.Apt.Type = dclService.PatchDeploymentPatchConfigAptTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Apt.Type: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.Apt: expected map[string]interface{}")
				}
			}
			if _, ok := rPatchConfig["goo"]; ok {
				if _, ok := rPatchConfig["goo"].(map[string]interface{}); ok {
					r.PatchConfig.Goo = &dclService.PatchDeploymentPatchConfigGoo{}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.Goo: expected map[string]interface{}")
				}
			}
			if _, ok := rPatchConfig["postStep"]; ok {
				if rPatchConfigPostStep, ok := rPatchConfig["postStep"].(map[string]interface{}); ok {
					r.PatchConfig.PostStep = &dclService.PatchDeploymentPatchConfigPostStep{}
					if _, ok := rPatchConfigPostStep["linuxExecStepConfig"]; ok {
						if rPatchConfigPostStepLinuxExecStepConfig, ok := rPatchConfigPostStep["linuxExecStepConfig"].(map[string]interface{}); ok {
							r.PatchConfig.PostStep.LinuxExecStepConfig = &dclService.PatchDeploymentPatchConfigPostStepLinuxExecStepConfig{}
							if _, ok := rPatchConfigPostStepLinuxExecStepConfig["allowedSuccessCodes"]; ok {
								if s, ok := rPatchConfigPostStepLinuxExecStepConfig["allowedSuccessCodes"].([]interface{}); ok {
									for _, ss := range s {
										if intval, ok := ss.(int64); ok {
											r.PatchConfig.PostStep.LinuxExecStepConfig.AllowedSuccessCodes = append(r.PatchConfig.PostStep.LinuxExecStepConfig.AllowedSuccessCodes, intval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.AllowedSuccessCodes: expected []interface{}")
								}
							}
							if _, ok := rPatchConfigPostStepLinuxExecStepConfig["gcsObject"]; ok {
								if rPatchConfigPostStepLinuxExecStepConfigGcsObject, ok := rPatchConfigPostStepLinuxExecStepConfig["gcsObject"].(map[string]interface{}); ok {
									r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject = &dclService.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject{}
									if _, ok := rPatchConfigPostStepLinuxExecStepConfigGcsObject["bucket"]; ok {
										if s, ok := rPatchConfigPostStepLinuxExecStepConfigGcsObject["bucket"].(string); ok {
											r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Bucket = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Bucket: expected string")
										}
									}
									if _, ok := rPatchConfigPostStepLinuxExecStepConfigGcsObject["generationNumber"]; ok {
										if i, ok := rPatchConfigPostStepLinuxExecStepConfigGcsObject["generationNumber"].(int64); ok {
											r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.GenerationNumber = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.GenerationNumber: expected int64")
										}
									}
									if _, ok := rPatchConfigPostStepLinuxExecStepConfigGcsObject["object"]; ok {
										if s, ok := rPatchConfigPostStepLinuxExecStepConfigGcsObject["object"].(string); ok {
											r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Object = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject.Object: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.GcsObject: expected map[string]interface{}")
								}
							}
							if _, ok := rPatchConfigPostStepLinuxExecStepConfig["interpreter"]; ok {
								if s, ok := rPatchConfigPostStepLinuxExecStepConfig["interpreter"].(string); ok {
									r.PatchConfig.PostStep.LinuxExecStepConfig.Interpreter = dclService.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.Interpreter: expected string")
								}
							}
							if _, ok := rPatchConfigPostStepLinuxExecStepConfig["localPath"]; ok {
								if s, ok := rPatchConfigPostStepLinuxExecStepConfig["localPath"].(string); ok {
									r.PatchConfig.PostStep.LinuxExecStepConfig.LocalPath = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig.LocalPath: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.PostStep.LinuxExecStepConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rPatchConfigPostStep["windowsExecStepConfig"]; ok {
						if rPatchConfigPostStepWindowsExecStepConfig, ok := rPatchConfigPostStep["windowsExecStepConfig"].(map[string]interface{}); ok {
							r.PatchConfig.PostStep.WindowsExecStepConfig = &dclService.PatchDeploymentPatchConfigPostStepWindowsExecStepConfig{}
							if _, ok := rPatchConfigPostStepWindowsExecStepConfig["allowedSuccessCodes"]; ok {
								if s, ok := rPatchConfigPostStepWindowsExecStepConfig["allowedSuccessCodes"].([]interface{}); ok {
									for _, ss := range s {
										if intval, ok := ss.(int64); ok {
											r.PatchConfig.PostStep.WindowsExecStepConfig.AllowedSuccessCodes = append(r.PatchConfig.PostStep.WindowsExecStepConfig.AllowedSuccessCodes, intval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.AllowedSuccessCodes: expected []interface{}")
								}
							}
							if _, ok := rPatchConfigPostStepWindowsExecStepConfig["gcsObject"]; ok {
								if rPatchConfigPostStepWindowsExecStepConfigGcsObject, ok := rPatchConfigPostStepWindowsExecStepConfig["gcsObject"].(map[string]interface{}); ok {
									r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject = &dclService.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject{}
									if _, ok := rPatchConfigPostStepWindowsExecStepConfigGcsObject["bucket"]; ok {
										if s, ok := rPatchConfigPostStepWindowsExecStepConfigGcsObject["bucket"].(string); ok {
											r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Bucket = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Bucket: expected string")
										}
									}
									if _, ok := rPatchConfigPostStepWindowsExecStepConfigGcsObject["generationNumber"]; ok {
										if i, ok := rPatchConfigPostStepWindowsExecStepConfigGcsObject["generationNumber"].(int64); ok {
											r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.GenerationNumber = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.GenerationNumber: expected int64")
										}
									}
									if _, ok := rPatchConfigPostStepWindowsExecStepConfigGcsObject["object"]; ok {
										if s, ok := rPatchConfigPostStepWindowsExecStepConfigGcsObject["object"].(string); ok {
											r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Object = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject.Object: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.GcsObject: expected map[string]interface{}")
								}
							}
							if _, ok := rPatchConfigPostStepWindowsExecStepConfig["interpreter"]; ok {
								if s, ok := rPatchConfigPostStepWindowsExecStepConfig["interpreter"].(string); ok {
									r.PatchConfig.PostStep.WindowsExecStepConfig.Interpreter = dclService.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.Interpreter: expected string")
								}
							}
							if _, ok := rPatchConfigPostStepWindowsExecStepConfig["localPath"]; ok {
								if s, ok := rPatchConfigPostStepWindowsExecStepConfig["localPath"].(string); ok {
									r.PatchConfig.PostStep.WindowsExecStepConfig.LocalPath = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig.LocalPath: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.PostStep.WindowsExecStepConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.PostStep: expected map[string]interface{}")
				}
			}
			if _, ok := rPatchConfig["preStep"]; ok {
				if rPatchConfigPreStep, ok := rPatchConfig["preStep"].(map[string]interface{}); ok {
					r.PatchConfig.PreStep = &dclService.PatchDeploymentPatchConfigPreStep{}
					if _, ok := rPatchConfigPreStep["linuxExecStepConfig"]; ok {
						if rPatchConfigPreStepLinuxExecStepConfig, ok := rPatchConfigPreStep["linuxExecStepConfig"].(map[string]interface{}); ok {
							r.PatchConfig.PreStep.LinuxExecStepConfig = &dclService.PatchDeploymentPatchConfigPreStepLinuxExecStepConfig{}
							if _, ok := rPatchConfigPreStepLinuxExecStepConfig["allowedSuccessCodes"]; ok {
								if s, ok := rPatchConfigPreStepLinuxExecStepConfig["allowedSuccessCodes"].([]interface{}); ok {
									for _, ss := range s {
										if intval, ok := ss.(int64); ok {
											r.PatchConfig.PreStep.LinuxExecStepConfig.AllowedSuccessCodes = append(r.PatchConfig.PreStep.LinuxExecStepConfig.AllowedSuccessCodes, intval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.AllowedSuccessCodes: expected []interface{}")
								}
							}
							if _, ok := rPatchConfigPreStepLinuxExecStepConfig["gcsObject"]; ok {
								if rPatchConfigPreStepLinuxExecStepConfigGcsObject, ok := rPatchConfigPreStepLinuxExecStepConfig["gcsObject"].(map[string]interface{}); ok {
									r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject = &dclService.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject{}
									if _, ok := rPatchConfigPreStepLinuxExecStepConfigGcsObject["bucket"]; ok {
										if s, ok := rPatchConfigPreStepLinuxExecStepConfigGcsObject["bucket"].(string); ok {
											r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Bucket = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Bucket: expected string")
										}
									}
									if _, ok := rPatchConfigPreStepLinuxExecStepConfigGcsObject["generationNumber"]; ok {
										if i, ok := rPatchConfigPreStepLinuxExecStepConfigGcsObject["generationNumber"].(int64); ok {
											r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.GenerationNumber = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.GenerationNumber: expected int64")
										}
									}
									if _, ok := rPatchConfigPreStepLinuxExecStepConfigGcsObject["object"]; ok {
										if s, ok := rPatchConfigPreStepLinuxExecStepConfigGcsObject["object"].(string); ok {
											r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Object = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject.Object: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.GcsObject: expected map[string]interface{}")
								}
							}
							if _, ok := rPatchConfigPreStepLinuxExecStepConfig["interpreter"]; ok {
								if s, ok := rPatchConfigPreStepLinuxExecStepConfig["interpreter"].(string); ok {
									r.PatchConfig.PreStep.LinuxExecStepConfig.Interpreter = dclService.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.Interpreter: expected string")
								}
							}
							if _, ok := rPatchConfigPreStepLinuxExecStepConfig["localPath"]; ok {
								if s, ok := rPatchConfigPreStepLinuxExecStepConfig["localPath"].(string); ok {
									r.PatchConfig.PreStep.LinuxExecStepConfig.LocalPath = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig.LocalPath: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.PreStep.LinuxExecStepConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rPatchConfigPreStep["windowsExecStepConfig"]; ok {
						if rPatchConfigPreStepWindowsExecStepConfig, ok := rPatchConfigPreStep["windowsExecStepConfig"].(map[string]interface{}); ok {
							r.PatchConfig.PreStep.WindowsExecStepConfig = &dclService.PatchDeploymentPatchConfigPreStepWindowsExecStepConfig{}
							if _, ok := rPatchConfigPreStepWindowsExecStepConfig["allowedSuccessCodes"]; ok {
								if s, ok := rPatchConfigPreStepWindowsExecStepConfig["allowedSuccessCodes"].([]interface{}); ok {
									for _, ss := range s {
										if intval, ok := ss.(int64); ok {
											r.PatchConfig.PreStep.WindowsExecStepConfig.AllowedSuccessCodes = append(r.PatchConfig.PreStep.WindowsExecStepConfig.AllowedSuccessCodes, intval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.AllowedSuccessCodes: expected []interface{}")
								}
							}
							if _, ok := rPatchConfigPreStepWindowsExecStepConfig["gcsObject"]; ok {
								if rPatchConfigPreStepWindowsExecStepConfigGcsObject, ok := rPatchConfigPreStepWindowsExecStepConfig["gcsObject"].(map[string]interface{}); ok {
									r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject = &dclService.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject{}
									if _, ok := rPatchConfigPreStepWindowsExecStepConfigGcsObject["bucket"]; ok {
										if s, ok := rPatchConfigPreStepWindowsExecStepConfigGcsObject["bucket"].(string); ok {
											r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Bucket = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Bucket: expected string")
										}
									}
									if _, ok := rPatchConfigPreStepWindowsExecStepConfigGcsObject["generationNumber"]; ok {
										if i, ok := rPatchConfigPreStepWindowsExecStepConfigGcsObject["generationNumber"].(int64); ok {
											r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.GenerationNumber = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.GenerationNumber: expected int64")
										}
									}
									if _, ok := rPatchConfigPreStepWindowsExecStepConfigGcsObject["object"]; ok {
										if s, ok := rPatchConfigPreStepWindowsExecStepConfigGcsObject["object"].(string); ok {
											r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Object = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject.Object: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.GcsObject: expected map[string]interface{}")
								}
							}
							if _, ok := rPatchConfigPreStepWindowsExecStepConfig["interpreter"]; ok {
								if s, ok := rPatchConfigPreStepWindowsExecStepConfig["interpreter"].(string); ok {
									r.PatchConfig.PreStep.WindowsExecStepConfig.Interpreter = dclService.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.Interpreter: expected string")
								}
							}
							if _, ok := rPatchConfigPreStepWindowsExecStepConfig["localPath"]; ok {
								if s, ok := rPatchConfigPreStepWindowsExecStepConfig["localPath"].(string); ok {
									r.PatchConfig.PreStep.WindowsExecStepConfig.LocalPath = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig.LocalPath: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.PreStep.WindowsExecStepConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.PreStep: expected map[string]interface{}")
				}
			}
			if _, ok := rPatchConfig["rebootConfig"]; ok {
				if s, ok := rPatchConfig["rebootConfig"].(string); ok {
					r.PatchConfig.RebootConfig = dclService.PatchDeploymentPatchConfigRebootConfigEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.PatchConfig.RebootConfig: expected string")
				}
			}
			if _, ok := rPatchConfig["windowsUpdate"]; ok {
				if rPatchConfigWindowsUpdate, ok := rPatchConfig["windowsUpdate"].(map[string]interface{}); ok {
					r.PatchConfig.WindowsUpdate = &dclService.PatchDeploymentPatchConfigWindowsUpdate{}
					if _, ok := rPatchConfigWindowsUpdate["classifications"]; ok {
						if s, ok := rPatchConfigWindowsUpdate["classifications"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.WindowsUpdate.Classifications = append(r.PatchConfig.WindowsUpdate.Classifications, dclService.PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.WindowsUpdate.Classifications: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigWindowsUpdate["excludes"]; ok {
						if s, ok := rPatchConfigWindowsUpdate["excludes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.WindowsUpdate.Excludes = append(r.PatchConfig.WindowsUpdate.Excludes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.WindowsUpdate.Excludes: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigWindowsUpdate["exclusivePatches"]; ok {
						if s, ok := rPatchConfigWindowsUpdate["exclusivePatches"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.WindowsUpdate.ExclusivePatches = append(r.PatchConfig.WindowsUpdate.ExclusivePatches, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.WindowsUpdate.ExclusivePatches: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.WindowsUpdate: expected map[string]interface{}")
				}
			}
			if _, ok := rPatchConfig["yum"]; ok {
				if rPatchConfigYum, ok := rPatchConfig["yum"].(map[string]interface{}); ok {
					r.PatchConfig.Yum = &dclService.PatchDeploymentPatchConfigYum{}
					if _, ok := rPatchConfigYum["excludes"]; ok {
						if s, ok := rPatchConfigYum["excludes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Yum.Excludes = append(r.PatchConfig.Yum.Excludes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Yum.Excludes: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigYum["exclusivePackages"]; ok {
						if s, ok := rPatchConfigYum["exclusivePackages"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Yum.ExclusivePackages = append(r.PatchConfig.Yum.ExclusivePackages, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Yum.ExclusivePackages: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigYum["minimal"]; ok {
						if b, ok := rPatchConfigYum["minimal"].(bool); ok {
							r.PatchConfig.Yum.Minimal = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Yum.Minimal: expected bool")
						}
					}
					if _, ok := rPatchConfigYum["security"]; ok {
						if b, ok := rPatchConfigYum["security"].(bool); ok {
							r.PatchConfig.Yum.Security = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Yum.Security: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.Yum: expected map[string]interface{}")
				}
			}
			if _, ok := rPatchConfig["zypper"]; ok {
				if rPatchConfigZypper, ok := rPatchConfig["zypper"].(map[string]interface{}); ok {
					r.PatchConfig.Zypper = &dclService.PatchDeploymentPatchConfigZypper{}
					if _, ok := rPatchConfigZypper["categories"]; ok {
						if s, ok := rPatchConfigZypper["categories"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Zypper.Categories = append(r.PatchConfig.Zypper.Categories, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Zypper.Categories: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigZypper["excludes"]; ok {
						if s, ok := rPatchConfigZypper["excludes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Zypper.Excludes = append(r.PatchConfig.Zypper.Excludes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Zypper.Excludes: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigZypper["exclusivePatches"]; ok {
						if s, ok := rPatchConfigZypper["exclusivePatches"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Zypper.ExclusivePatches = append(r.PatchConfig.Zypper.ExclusivePatches, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Zypper.ExclusivePatches: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigZypper["severities"]; ok {
						if s, ok := rPatchConfigZypper["severities"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.PatchConfig.Zypper.Severities = append(r.PatchConfig.Zypper.Severities, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Zypper.Severities: expected []interface{}")
						}
					}
					if _, ok := rPatchConfigZypper["withOptional"]; ok {
						if b, ok := rPatchConfigZypper["withOptional"].(bool); ok {
							r.PatchConfig.Zypper.WithOptional = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Zypper.WithOptional: expected bool")
						}
					}
					if _, ok := rPatchConfigZypper["withUpdate"]; ok {
						if b, ok := rPatchConfigZypper["withUpdate"].(bool); ok {
							r.PatchConfig.Zypper.WithUpdate = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.PatchConfig.Zypper.WithUpdate: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PatchConfig.Zypper: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PatchConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["recurringSchedule"]; ok {
		if rRecurringSchedule, ok := u.Object["recurringSchedule"].(map[string]interface{}); ok {
			r.RecurringSchedule = &dclService.PatchDeploymentRecurringSchedule{}
			if _, ok := rRecurringSchedule["endTime"]; ok {
				if s, ok := rRecurringSchedule["endTime"].(string); ok {
					r.RecurringSchedule.EndTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.EndTime: expected string")
				}
			}
			if _, ok := rRecurringSchedule["frequency"]; ok {
				if s, ok := rRecurringSchedule["frequency"].(string); ok {
					r.RecurringSchedule.Frequency = dclService.PatchDeploymentRecurringScheduleFrequencyEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.Frequency: expected string")
				}
			}
			if _, ok := rRecurringSchedule["lastExecuteTime"]; ok {
				if s, ok := rRecurringSchedule["lastExecuteTime"].(string); ok {
					r.RecurringSchedule.LastExecuteTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.LastExecuteTime: expected string")
				}
			}
			if _, ok := rRecurringSchedule["monthly"]; ok {
				if rRecurringScheduleMonthly, ok := rRecurringSchedule["monthly"].(map[string]interface{}); ok {
					r.RecurringSchedule.Monthly = &dclService.PatchDeploymentRecurringScheduleMonthly{}
					if _, ok := rRecurringScheduleMonthly["monthDay"]; ok {
						if i, ok := rRecurringScheduleMonthly["monthDay"].(int64); ok {
							r.RecurringSchedule.Monthly.MonthDay = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.Monthly.MonthDay: expected int64")
						}
					}
					if _, ok := rRecurringScheduleMonthly["weekDayOfMonth"]; ok {
						if rRecurringScheduleMonthlyWeekDayOfMonth, ok := rRecurringScheduleMonthly["weekDayOfMonth"].(map[string]interface{}); ok {
							r.RecurringSchedule.Monthly.WeekDayOfMonth = &dclService.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth{}
							if _, ok := rRecurringScheduleMonthlyWeekDayOfMonth["dayOfWeek"]; ok {
								if s, ok := rRecurringScheduleMonthlyWeekDayOfMonth["dayOfWeek"].(string); ok {
									r.RecurringSchedule.Monthly.WeekDayOfMonth.DayOfWeek = dclService.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.RecurringSchedule.Monthly.WeekDayOfMonth.DayOfWeek: expected string")
								}
							}
							if _, ok := rRecurringScheduleMonthlyWeekDayOfMonth["weekOrdinal"]; ok {
								if i, ok := rRecurringScheduleMonthlyWeekDayOfMonth["weekOrdinal"].(int64); ok {
									r.RecurringSchedule.Monthly.WeekDayOfMonth.WeekOrdinal = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.RecurringSchedule.Monthly.WeekDayOfMonth.WeekOrdinal: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.Monthly.WeekDayOfMonth: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.Monthly: expected map[string]interface{}")
				}
			}
			if _, ok := rRecurringSchedule["nextExecuteTime"]; ok {
				if s, ok := rRecurringSchedule["nextExecuteTime"].(string); ok {
					r.RecurringSchedule.NextExecuteTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.NextExecuteTime: expected string")
				}
			}
			if _, ok := rRecurringSchedule["startTime"]; ok {
				if s, ok := rRecurringSchedule["startTime"].(string); ok {
					r.RecurringSchedule.StartTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.StartTime: expected string")
				}
			}
			if _, ok := rRecurringSchedule["timeOfDay"]; ok {
				if rRecurringScheduleTimeOfDay, ok := rRecurringSchedule["timeOfDay"].(map[string]interface{}); ok {
					r.RecurringSchedule.TimeOfDay = &dclService.PatchDeploymentRecurringScheduleTimeOfDay{}
					if _, ok := rRecurringScheduleTimeOfDay["hours"]; ok {
						if i, ok := rRecurringScheduleTimeOfDay["hours"].(int64); ok {
							r.RecurringSchedule.TimeOfDay.Hours = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.TimeOfDay.Hours: expected int64")
						}
					}
					if _, ok := rRecurringScheduleTimeOfDay["minutes"]; ok {
						if i, ok := rRecurringScheduleTimeOfDay["minutes"].(int64); ok {
							r.RecurringSchedule.TimeOfDay.Minutes = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.TimeOfDay.Minutes: expected int64")
						}
					}
					if _, ok := rRecurringScheduleTimeOfDay["nanos"]; ok {
						if i, ok := rRecurringScheduleTimeOfDay["nanos"].(int64); ok {
							r.RecurringSchedule.TimeOfDay.Nanos = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.TimeOfDay.Nanos: expected int64")
						}
					}
					if _, ok := rRecurringScheduleTimeOfDay["seconds"]; ok {
						if i, ok := rRecurringScheduleTimeOfDay["seconds"].(int64); ok {
							r.RecurringSchedule.TimeOfDay.Seconds = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.TimeOfDay.Seconds: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.TimeOfDay: expected map[string]interface{}")
				}
			}
			if _, ok := rRecurringSchedule["timeZone"]; ok {
				if rRecurringScheduleTimeZone, ok := rRecurringSchedule["timeZone"].(map[string]interface{}); ok {
					r.RecurringSchedule.TimeZone = &dclService.PatchDeploymentRecurringScheduleTimeZone{}
					if _, ok := rRecurringScheduleTimeZone["id"]; ok {
						if s, ok := rRecurringScheduleTimeZone["id"].(string); ok {
							r.RecurringSchedule.TimeZone.Id = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.TimeZone.Id: expected string")
						}
					}
					if _, ok := rRecurringScheduleTimeZone["version"]; ok {
						if s, ok := rRecurringScheduleTimeZone["version"].(string); ok {
							r.RecurringSchedule.TimeZone.Version = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.TimeZone.Version: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.TimeZone: expected map[string]interface{}")
				}
			}
			if _, ok := rRecurringSchedule["weekly"]; ok {
				if rRecurringScheduleWeekly, ok := rRecurringSchedule["weekly"].(map[string]interface{}); ok {
					r.RecurringSchedule.Weekly = &dclService.PatchDeploymentRecurringScheduleWeekly{}
					if _, ok := rRecurringScheduleWeekly["dayOfWeek"]; ok {
						if s, ok := rRecurringScheduleWeekly["dayOfWeek"].(string); ok {
							r.RecurringSchedule.Weekly.DayOfWeek = dclService.PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.RecurringSchedule.Weekly.DayOfWeek: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.RecurringSchedule.Weekly: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RecurringSchedule: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["rollout"]; ok {
		if rRollout, ok := u.Object["rollout"].(map[string]interface{}); ok {
			r.Rollout = &dclService.PatchDeploymentRollout{}
			if _, ok := rRollout["disruptionBudget"]; ok {
				if rRolloutDisruptionBudget, ok := rRollout["disruptionBudget"].(map[string]interface{}); ok {
					r.Rollout.DisruptionBudget = &dclService.PatchDeploymentRolloutDisruptionBudget{}
					if _, ok := rRolloutDisruptionBudget["fixed"]; ok {
						if i, ok := rRolloutDisruptionBudget["fixed"].(int64); ok {
							r.Rollout.DisruptionBudget.Fixed = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Rollout.DisruptionBudget.Fixed: expected int64")
						}
					}
					if _, ok := rRolloutDisruptionBudget["percent"]; ok {
						if i, ok := rRolloutDisruptionBudget["percent"].(int64); ok {
							r.Rollout.DisruptionBudget.Percent = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Rollout.DisruptionBudget.Percent: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Rollout.DisruptionBudget: expected map[string]interface{}")
				}
			}
			if _, ok := rRollout["mode"]; ok {
				if s, ok := rRollout["mode"].(string); ok {
					r.Rollout.Mode = dclService.PatchDeploymentRolloutModeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Rollout.Mode: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Rollout: expected map[string]interface{}")
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

func GetPatchDeployment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPatchDeployment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetPatchDeployment(ctx, r)
	if err != nil {
		return nil, err
	}
	return PatchDeploymentToUnstructured(r), nil
}

func ListPatchDeployment(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListPatchDeployment(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, PatchDeploymentToUnstructured(r))
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

func ApplyPatchDeployment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPatchDeployment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPatchDeployment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyPatchDeployment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return PatchDeploymentToUnstructured(r), nil
}

func PatchDeploymentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPatchDeployment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPatchDeployment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyPatchDeployment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeletePatchDeployment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPatchDeployment(u)
	if err != nil {
		return err
	}
	return c.DeletePatchDeployment(ctx, r)
}

func PatchDeploymentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToPatchDeployment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *PatchDeployment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"osconfig",
		"PatchDeployment",
		"beta",
	}
}

func (r *PatchDeployment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PatchDeployment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PatchDeployment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *PatchDeployment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PatchDeployment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PatchDeployment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PatchDeployment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPatchDeployment(ctx, config, resource)
}

func (r *PatchDeployment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPatchDeployment(ctx, config, resource, opts...)
}

func (r *PatchDeployment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PatchDeploymentHasDiff(ctx, config, resource, opts...)
}

func (r *PatchDeployment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePatchDeployment(ctx, config, resource)
}

func (r *PatchDeployment) ID(resource *unstructured.Resource) (string, error) {
	return PatchDeploymentID(resource)
}

func init() {
	unstructured.Register(&PatchDeployment{})
}
