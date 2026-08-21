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
package dataproc

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type WorkflowTemplate struct{}

func WorkflowTemplateToUnstructured(r *dclService.WorkflowTemplate) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataproc",
			Version: "alpha",
			Type:    "WorkflowTemplate",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DagTimeout != nil {
		u.Object["dagTimeout"] = *r.DagTimeout
	}
	var rJobs []interface{}
	for _, rJobsVal := range r.Jobs {
		rJobsObject := make(map[string]interface{})
		if rJobsVal.HadoopJob != nil && rJobsVal.HadoopJob != dclService.EmptyWorkflowTemplateJobsHadoopJob {
			rJobsValHadoopJob := make(map[string]interface{})
			var rJobsValHadoopJobArchiveUris []interface{}
			for _, rJobsValHadoopJobArchiveUrisVal := range rJobsVal.HadoopJob.ArchiveUris {
				rJobsValHadoopJobArchiveUris = append(rJobsValHadoopJobArchiveUris, rJobsValHadoopJobArchiveUrisVal)
			}
			rJobsValHadoopJob["archiveUris"] = rJobsValHadoopJobArchiveUris
			var rJobsValHadoopJobArgs []interface{}
			for _, rJobsValHadoopJobArgsVal := range rJobsVal.HadoopJob.Args {
				rJobsValHadoopJobArgs = append(rJobsValHadoopJobArgs, rJobsValHadoopJobArgsVal)
			}
			rJobsValHadoopJob["args"] = rJobsValHadoopJobArgs
			var rJobsValHadoopJobFileUris []interface{}
			for _, rJobsValHadoopJobFileUrisVal := range rJobsVal.HadoopJob.FileUris {
				rJobsValHadoopJobFileUris = append(rJobsValHadoopJobFileUris, rJobsValHadoopJobFileUrisVal)
			}
			rJobsValHadoopJob["fileUris"] = rJobsValHadoopJobFileUris
			var rJobsValHadoopJobJarFileUris []interface{}
			for _, rJobsValHadoopJobJarFileUrisVal := range rJobsVal.HadoopJob.JarFileUris {
				rJobsValHadoopJobJarFileUris = append(rJobsValHadoopJobJarFileUris, rJobsValHadoopJobJarFileUrisVal)
			}
			rJobsValHadoopJob["jarFileUris"] = rJobsValHadoopJobJarFileUris
			if rJobsVal.HadoopJob.LoggingConfig != nil && rJobsVal.HadoopJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsHadoopJobLoggingConfig {
				rJobsValHadoopJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.HadoopJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValHadoopJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.HadoopJob.LoggingConfig.DriverLogLevels {
						rJobsValHadoopJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValHadoopJobLoggingConfig["driverLogLevels"] = rJobsValHadoopJobLoggingConfigDriverLogLevels
				}
				rJobsValHadoopJob["loggingConfig"] = rJobsValHadoopJobLoggingConfig
			}
			if rJobsVal.HadoopJob.MainClass != nil {
				rJobsValHadoopJob["mainClass"] = *rJobsVal.HadoopJob.MainClass
			}
			if rJobsVal.HadoopJob.MainJarFileUri != nil {
				rJobsValHadoopJob["mainJarFileUri"] = *rJobsVal.HadoopJob.MainJarFileUri
			}
			if rJobsVal.HadoopJob.Properties != nil {
				rJobsValHadoopJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.HadoopJob.Properties {
					rJobsValHadoopJobProperties[k] = v
				}
				rJobsValHadoopJob["properties"] = rJobsValHadoopJobProperties
			}
			rJobsObject["hadoopJob"] = rJobsValHadoopJob
		}
		if rJobsVal.HiveJob != nil && rJobsVal.HiveJob != dclService.EmptyWorkflowTemplateJobsHiveJob {
			rJobsValHiveJob := make(map[string]interface{})
			if rJobsVal.HiveJob.ContinueOnFailure != nil {
				rJobsValHiveJob["continueOnFailure"] = *rJobsVal.HiveJob.ContinueOnFailure
			}
			var rJobsValHiveJobJarFileUris []interface{}
			for _, rJobsValHiveJobJarFileUrisVal := range rJobsVal.HiveJob.JarFileUris {
				rJobsValHiveJobJarFileUris = append(rJobsValHiveJobJarFileUris, rJobsValHiveJobJarFileUrisVal)
			}
			rJobsValHiveJob["jarFileUris"] = rJobsValHiveJobJarFileUris
			if rJobsVal.HiveJob.Properties != nil {
				rJobsValHiveJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.HiveJob.Properties {
					rJobsValHiveJobProperties[k] = v
				}
				rJobsValHiveJob["properties"] = rJobsValHiveJobProperties
			}
			if rJobsVal.HiveJob.QueryFileUri != nil {
				rJobsValHiveJob["queryFileUri"] = *rJobsVal.HiveJob.QueryFileUri
			}
			if rJobsVal.HiveJob.QueryList != nil && rJobsVal.HiveJob.QueryList != dclService.EmptyWorkflowTemplateJobsHiveJobQueryList {
				rJobsValHiveJobQueryList := make(map[string]interface{})
				var rJobsValHiveJobQueryListQueries []interface{}
				for _, rJobsValHiveJobQueryListQueriesVal := range rJobsVal.HiveJob.QueryList.Queries {
					rJobsValHiveJobQueryListQueries = append(rJobsValHiveJobQueryListQueries, rJobsValHiveJobQueryListQueriesVal)
				}
				rJobsValHiveJobQueryList["queries"] = rJobsValHiveJobQueryListQueries
				rJobsValHiveJob["queryList"] = rJobsValHiveJobQueryList
			}
			if rJobsVal.HiveJob.ScriptVariables != nil {
				rJobsValHiveJobScriptVariables := make(map[string]interface{})
				for k, v := range rJobsVal.HiveJob.ScriptVariables {
					rJobsValHiveJobScriptVariables[k] = v
				}
				rJobsValHiveJob["scriptVariables"] = rJobsValHiveJobScriptVariables
			}
			rJobsObject["hiveJob"] = rJobsValHiveJob
		}
		if rJobsVal.Labels != nil {
			rJobsValLabels := make(map[string]interface{})
			for k, v := range rJobsVal.Labels {
				rJobsValLabels[k] = v
			}
			rJobsObject["labels"] = rJobsValLabels
		}
		if rJobsVal.PigJob != nil && rJobsVal.PigJob != dclService.EmptyWorkflowTemplateJobsPigJob {
			rJobsValPigJob := make(map[string]interface{})
			if rJobsVal.PigJob.ContinueOnFailure != nil {
				rJobsValPigJob["continueOnFailure"] = *rJobsVal.PigJob.ContinueOnFailure
			}
			var rJobsValPigJobJarFileUris []interface{}
			for _, rJobsValPigJobJarFileUrisVal := range rJobsVal.PigJob.JarFileUris {
				rJobsValPigJobJarFileUris = append(rJobsValPigJobJarFileUris, rJobsValPigJobJarFileUrisVal)
			}
			rJobsValPigJob["jarFileUris"] = rJobsValPigJobJarFileUris
			if rJobsVal.PigJob.LoggingConfig != nil && rJobsVal.PigJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsPigJobLoggingConfig {
				rJobsValPigJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.PigJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValPigJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.PigJob.LoggingConfig.DriverLogLevels {
						rJobsValPigJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValPigJobLoggingConfig["driverLogLevels"] = rJobsValPigJobLoggingConfigDriverLogLevels
				}
				rJobsValPigJob["loggingConfig"] = rJobsValPigJobLoggingConfig
			}
			if rJobsVal.PigJob.Properties != nil {
				rJobsValPigJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.PigJob.Properties {
					rJobsValPigJobProperties[k] = v
				}
				rJobsValPigJob["properties"] = rJobsValPigJobProperties
			}
			if rJobsVal.PigJob.QueryFileUri != nil {
				rJobsValPigJob["queryFileUri"] = *rJobsVal.PigJob.QueryFileUri
			}
			if rJobsVal.PigJob.QueryList != nil && rJobsVal.PigJob.QueryList != dclService.EmptyWorkflowTemplateJobsPigJobQueryList {
				rJobsValPigJobQueryList := make(map[string]interface{})
				var rJobsValPigJobQueryListQueries []interface{}
				for _, rJobsValPigJobQueryListQueriesVal := range rJobsVal.PigJob.QueryList.Queries {
					rJobsValPigJobQueryListQueries = append(rJobsValPigJobQueryListQueries, rJobsValPigJobQueryListQueriesVal)
				}
				rJobsValPigJobQueryList["queries"] = rJobsValPigJobQueryListQueries
				rJobsValPigJob["queryList"] = rJobsValPigJobQueryList
			}
			if rJobsVal.PigJob.ScriptVariables != nil {
				rJobsValPigJobScriptVariables := make(map[string]interface{})
				for k, v := range rJobsVal.PigJob.ScriptVariables {
					rJobsValPigJobScriptVariables[k] = v
				}
				rJobsValPigJob["scriptVariables"] = rJobsValPigJobScriptVariables
			}
			rJobsObject["pigJob"] = rJobsValPigJob
		}
		var rJobsValPrerequisiteStepIds []interface{}
		for _, rJobsValPrerequisiteStepIdsVal := range rJobsVal.PrerequisiteStepIds {
			rJobsValPrerequisiteStepIds = append(rJobsValPrerequisiteStepIds, rJobsValPrerequisiteStepIdsVal)
		}
		rJobsObject["prerequisiteStepIds"] = rJobsValPrerequisiteStepIds
		if rJobsVal.PrestoJob != nil && rJobsVal.PrestoJob != dclService.EmptyWorkflowTemplateJobsPrestoJob {
			rJobsValPrestoJob := make(map[string]interface{})
			var rJobsValPrestoJobClientTags []interface{}
			for _, rJobsValPrestoJobClientTagsVal := range rJobsVal.PrestoJob.ClientTags {
				rJobsValPrestoJobClientTags = append(rJobsValPrestoJobClientTags, rJobsValPrestoJobClientTagsVal)
			}
			rJobsValPrestoJob["clientTags"] = rJobsValPrestoJobClientTags
			if rJobsVal.PrestoJob.ContinueOnFailure != nil {
				rJobsValPrestoJob["continueOnFailure"] = *rJobsVal.PrestoJob.ContinueOnFailure
			}
			if rJobsVal.PrestoJob.LoggingConfig != nil && rJobsVal.PrestoJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsPrestoJobLoggingConfig {
				rJobsValPrestoJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.PrestoJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValPrestoJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.PrestoJob.LoggingConfig.DriverLogLevels {
						rJobsValPrestoJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValPrestoJobLoggingConfig["driverLogLevels"] = rJobsValPrestoJobLoggingConfigDriverLogLevels
				}
				rJobsValPrestoJob["loggingConfig"] = rJobsValPrestoJobLoggingConfig
			}
			if rJobsVal.PrestoJob.OutputFormat != nil {
				rJobsValPrestoJob["outputFormat"] = *rJobsVal.PrestoJob.OutputFormat
			}
			if rJobsVal.PrestoJob.Properties != nil {
				rJobsValPrestoJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.PrestoJob.Properties {
					rJobsValPrestoJobProperties[k] = v
				}
				rJobsValPrestoJob["properties"] = rJobsValPrestoJobProperties
			}
			if rJobsVal.PrestoJob.QueryFileUri != nil {
				rJobsValPrestoJob["queryFileUri"] = *rJobsVal.PrestoJob.QueryFileUri
			}
			if rJobsVal.PrestoJob.QueryList != nil && rJobsVal.PrestoJob.QueryList != dclService.EmptyWorkflowTemplateJobsPrestoJobQueryList {
				rJobsValPrestoJobQueryList := make(map[string]interface{})
				var rJobsValPrestoJobQueryListQueries []interface{}
				for _, rJobsValPrestoJobQueryListQueriesVal := range rJobsVal.PrestoJob.QueryList.Queries {
					rJobsValPrestoJobQueryListQueries = append(rJobsValPrestoJobQueryListQueries, rJobsValPrestoJobQueryListQueriesVal)
				}
				rJobsValPrestoJobQueryList["queries"] = rJobsValPrestoJobQueryListQueries
				rJobsValPrestoJob["queryList"] = rJobsValPrestoJobQueryList
			}
			rJobsObject["prestoJob"] = rJobsValPrestoJob
		}
		if rJobsVal.PysparkJob != nil && rJobsVal.PysparkJob != dclService.EmptyWorkflowTemplateJobsPysparkJob {
			rJobsValPysparkJob := make(map[string]interface{})
			var rJobsValPysparkJobArchiveUris []interface{}
			for _, rJobsValPysparkJobArchiveUrisVal := range rJobsVal.PysparkJob.ArchiveUris {
				rJobsValPysparkJobArchiveUris = append(rJobsValPysparkJobArchiveUris, rJobsValPysparkJobArchiveUrisVal)
			}
			rJobsValPysparkJob["archiveUris"] = rJobsValPysparkJobArchiveUris
			var rJobsValPysparkJobArgs []interface{}
			for _, rJobsValPysparkJobArgsVal := range rJobsVal.PysparkJob.Args {
				rJobsValPysparkJobArgs = append(rJobsValPysparkJobArgs, rJobsValPysparkJobArgsVal)
			}
			rJobsValPysparkJob["args"] = rJobsValPysparkJobArgs
			var rJobsValPysparkJobFileUris []interface{}
			for _, rJobsValPysparkJobFileUrisVal := range rJobsVal.PysparkJob.FileUris {
				rJobsValPysparkJobFileUris = append(rJobsValPysparkJobFileUris, rJobsValPysparkJobFileUrisVal)
			}
			rJobsValPysparkJob["fileUris"] = rJobsValPysparkJobFileUris
			var rJobsValPysparkJobJarFileUris []interface{}
			for _, rJobsValPysparkJobJarFileUrisVal := range rJobsVal.PysparkJob.JarFileUris {
				rJobsValPysparkJobJarFileUris = append(rJobsValPysparkJobJarFileUris, rJobsValPysparkJobJarFileUrisVal)
			}
			rJobsValPysparkJob["jarFileUris"] = rJobsValPysparkJobJarFileUris
			if rJobsVal.PysparkJob.LoggingConfig != nil && rJobsVal.PysparkJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsPysparkJobLoggingConfig {
				rJobsValPysparkJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.PysparkJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValPysparkJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.PysparkJob.LoggingConfig.DriverLogLevels {
						rJobsValPysparkJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValPysparkJobLoggingConfig["driverLogLevels"] = rJobsValPysparkJobLoggingConfigDriverLogLevels
				}
				rJobsValPysparkJob["loggingConfig"] = rJobsValPysparkJobLoggingConfig
			}
			if rJobsVal.PysparkJob.MainPythonFileUri != nil {
				rJobsValPysparkJob["mainPythonFileUri"] = *rJobsVal.PysparkJob.MainPythonFileUri
			}
			if rJobsVal.PysparkJob.Properties != nil {
				rJobsValPysparkJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.PysparkJob.Properties {
					rJobsValPysparkJobProperties[k] = v
				}
				rJobsValPysparkJob["properties"] = rJobsValPysparkJobProperties
			}
			var rJobsValPysparkJobPythonFileUris []interface{}
			for _, rJobsValPysparkJobPythonFileUrisVal := range rJobsVal.PysparkJob.PythonFileUris {
				rJobsValPysparkJobPythonFileUris = append(rJobsValPysparkJobPythonFileUris, rJobsValPysparkJobPythonFileUrisVal)
			}
			rJobsValPysparkJob["pythonFileUris"] = rJobsValPysparkJobPythonFileUris
			rJobsObject["pysparkJob"] = rJobsValPysparkJob
		}
		if rJobsVal.Scheduling != nil && rJobsVal.Scheduling != dclService.EmptyWorkflowTemplateJobsScheduling {
			rJobsValScheduling := make(map[string]interface{})
			if rJobsVal.Scheduling.MaxFailuresPerHour != nil {
				rJobsValScheduling["maxFailuresPerHour"] = *rJobsVal.Scheduling.MaxFailuresPerHour
			}
			if rJobsVal.Scheduling.MaxFailuresTotal != nil {
				rJobsValScheduling["maxFailuresTotal"] = *rJobsVal.Scheduling.MaxFailuresTotal
			}
			rJobsObject["scheduling"] = rJobsValScheduling
		}
		if rJobsVal.SparkJob != nil && rJobsVal.SparkJob != dclService.EmptyWorkflowTemplateJobsSparkJob {
			rJobsValSparkJob := make(map[string]interface{})
			var rJobsValSparkJobArchiveUris []interface{}
			for _, rJobsValSparkJobArchiveUrisVal := range rJobsVal.SparkJob.ArchiveUris {
				rJobsValSparkJobArchiveUris = append(rJobsValSparkJobArchiveUris, rJobsValSparkJobArchiveUrisVal)
			}
			rJobsValSparkJob["archiveUris"] = rJobsValSparkJobArchiveUris
			var rJobsValSparkJobArgs []interface{}
			for _, rJobsValSparkJobArgsVal := range rJobsVal.SparkJob.Args {
				rJobsValSparkJobArgs = append(rJobsValSparkJobArgs, rJobsValSparkJobArgsVal)
			}
			rJobsValSparkJob["args"] = rJobsValSparkJobArgs
			var rJobsValSparkJobFileUris []interface{}
			for _, rJobsValSparkJobFileUrisVal := range rJobsVal.SparkJob.FileUris {
				rJobsValSparkJobFileUris = append(rJobsValSparkJobFileUris, rJobsValSparkJobFileUrisVal)
			}
			rJobsValSparkJob["fileUris"] = rJobsValSparkJobFileUris
			var rJobsValSparkJobJarFileUris []interface{}
			for _, rJobsValSparkJobJarFileUrisVal := range rJobsVal.SparkJob.JarFileUris {
				rJobsValSparkJobJarFileUris = append(rJobsValSparkJobJarFileUris, rJobsValSparkJobJarFileUrisVal)
			}
			rJobsValSparkJob["jarFileUris"] = rJobsValSparkJobJarFileUris
			if rJobsVal.SparkJob.LoggingConfig != nil && rJobsVal.SparkJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsSparkJobLoggingConfig {
				rJobsValSparkJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.SparkJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValSparkJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.SparkJob.LoggingConfig.DriverLogLevels {
						rJobsValSparkJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValSparkJobLoggingConfig["driverLogLevels"] = rJobsValSparkJobLoggingConfigDriverLogLevels
				}
				rJobsValSparkJob["loggingConfig"] = rJobsValSparkJobLoggingConfig
			}
			if rJobsVal.SparkJob.MainClass != nil {
				rJobsValSparkJob["mainClass"] = *rJobsVal.SparkJob.MainClass
			}
			if rJobsVal.SparkJob.MainJarFileUri != nil {
				rJobsValSparkJob["mainJarFileUri"] = *rJobsVal.SparkJob.MainJarFileUri
			}
			if rJobsVal.SparkJob.Properties != nil {
				rJobsValSparkJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.SparkJob.Properties {
					rJobsValSparkJobProperties[k] = v
				}
				rJobsValSparkJob["properties"] = rJobsValSparkJobProperties
			}
			rJobsObject["sparkJob"] = rJobsValSparkJob
		}
		if rJobsVal.SparkRJob != nil && rJobsVal.SparkRJob != dclService.EmptyWorkflowTemplateJobsSparkRJob {
			rJobsValSparkRJob := make(map[string]interface{})
			var rJobsValSparkRJobArchiveUris []interface{}
			for _, rJobsValSparkRJobArchiveUrisVal := range rJobsVal.SparkRJob.ArchiveUris {
				rJobsValSparkRJobArchiveUris = append(rJobsValSparkRJobArchiveUris, rJobsValSparkRJobArchiveUrisVal)
			}
			rJobsValSparkRJob["archiveUris"] = rJobsValSparkRJobArchiveUris
			var rJobsValSparkRJobArgs []interface{}
			for _, rJobsValSparkRJobArgsVal := range rJobsVal.SparkRJob.Args {
				rJobsValSparkRJobArgs = append(rJobsValSparkRJobArgs, rJobsValSparkRJobArgsVal)
			}
			rJobsValSparkRJob["args"] = rJobsValSparkRJobArgs
			var rJobsValSparkRJobFileUris []interface{}
			for _, rJobsValSparkRJobFileUrisVal := range rJobsVal.SparkRJob.FileUris {
				rJobsValSparkRJobFileUris = append(rJobsValSparkRJobFileUris, rJobsValSparkRJobFileUrisVal)
			}
			rJobsValSparkRJob["fileUris"] = rJobsValSparkRJobFileUris
			if rJobsVal.SparkRJob.LoggingConfig != nil && rJobsVal.SparkRJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsSparkRJobLoggingConfig {
				rJobsValSparkRJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.SparkRJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValSparkRJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.SparkRJob.LoggingConfig.DriverLogLevels {
						rJobsValSparkRJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValSparkRJobLoggingConfig["driverLogLevels"] = rJobsValSparkRJobLoggingConfigDriverLogLevels
				}
				rJobsValSparkRJob["loggingConfig"] = rJobsValSparkRJobLoggingConfig
			}
			if rJobsVal.SparkRJob.MainRFileUri != nil {
				rJobsValSparkRJob["mainRFileUri"] = *rJobsVal.SparkRJob.MainRFileUri
			}
			if rJobsVal.SparkRJob.Properties != nil {
				rJobsValSparkRJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.SparkRJob.Properties {
					rJobsValSparkRJobProperties[k] = v
				}
				rJobsValSparkRJob["properties"] = rJobsValSparkRJobProperties
			}
			rJobsObject["sparkRJob"] = rJobsValSparkRJob
		}
		if rJobsVal.SparkSqlJob != nil && rJobsVal.SparkSqlJob != dclService.EmptyWorkflowTemplateJobsSparkSqlJob {
			rJobsValSparkSqlJob := make(map[string]interface{})
			var rJobsValSparkSqlJobJarFileUris []interface{}
			for _, rJobsValSparkSqlJobJarFileUrisVal := range rJobsVal.SparkSqlJob.JarFileUris {
				rJobsValSparkSqlJobJarFileUris = append(rJobsValSparkSqlJobJarFileUris, rJobsValSparkSqlJobJarFileUrisVal)
			}
			rJobsValSparkSqlJob["jarFileUris"] = rJobsValSparkSqlJobJarFileUris
			if rJobsVal.SparkSqlJob.LoggingConfig != nil && rJobsVal.SparkSqlJob.LoggingConfig != dclService.EmptyWorkflowTemplateJobsSparkSqlJobLoggingConfig {
				rJobsValSparkSqlJobLoggingConfig := make(map[string]interface{})
				if rJobsVal.SparkSqlJob.LoggingConfig.DriverLogLevels != nil {
					rJobsValSparkSqlJobLoggingConfigDriverLogLevels := make(map[string]interface{})
					for k, v := range rJobsVal.SparkSqlJob.LoggingConfig.DriverLogLevels {
						rJobsValSparkSqlJobLoggingConfigDriverLogLevels[k] = v
					}
					rJobsValSparkSqlJobLoggingConfig["driverLogLevels"] = rJobsValSparkSqlJobLoggingConfigDriverLogLevels
				}
				rJobsValSparkSqlJob["loggingConfig"] = rJobsValSparkSqlJobLoggingConfig
			}
			if rJobsVal.SparkSqlJob.Properties != nil {
				rJobsValSparkSqlJobProperties := make(map[string]interface{})
				for k, v := range rJobsVal.SparkSqlJob.Properties {
					rJobsValSparkSqlJobProperties[k] = v
				}
				rJobsValSparkSqlJob["properties"] = rJobsValSparkSqlJobProperties
			}
			if rJobsVal.SparkSqlJob.QueryFileUri != nil {
				rJobsValSparkSqlJob["queryFileUri"] = *rJobsVal.SparkSqlJob.QueryFileUri
			}
			if rJobsVal.SparkSqlJob.QueryList != nil && rJobsVal.SparkSqlJob.QueryList != dclService.EmptyWorkflowTemplateJobsSparkSqlJobQueryList {
				rJobsValSparkSqlJobQueryList := make(map[string]interface{})
				var rJobsValSparkSqlJobQueryListQueries []interface{}
				for _, rJobsValSparkSqlJobQueryListQueriesVal := range rJobsVal.SparkSqlJob.QueryList.Queries {
					rJobsValSparkSqlJobQueryListQueries = append(rJobsValSparkSqlJobQueryListQueries, rJobsValSparkSqlJobQueryListQueriesVal)
				}
				rJobsValSparkSqlJobQueryList["queries"] = rJobsValSparkSqlJobQueryListQueries
				rJobsValSparkSqlJob["queryList"] = rJobsValSparkSqlJobQueryList
			}
			if rJobsVal.SparkSqlJob.ScriptVariables != nil {
				rJobsValSparkSqlJobScriptVariables := make(map[string]interface{})
				for k, v := range rJobsVal.SparkSqlJob.ScriptVariables {
					rJobsValSparkSqlJobScriptVariables[k] = v
				}
				rJobsValSparkSqlJob["scriptVariables"] = rJobsValSparkSqlJobScriptVariables
			}
			rJobsObject["sparkSqlJob"] = rJobsValSparkSqlJob
		}
		if rJobsVal.StepId != nil {
			rJobsObject["stepId"] = *rJobsVal.StepId
		}
		rJobs = append(rJobs, rJobsObject)
	}
	u.Object["jobs"] = rJobs
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rParameters []interface{}
	for _, rParametersVal := range r.Parameters {
		rParametersObject := make(map[string]interface{})
		if rParametersVal.Description != nil {
			rParametersObject["description"] = *rParametersVal.Description
		}
		var rParametersValFields []interface{}
		for _, rParametersValFieldsVal := range rParametersVal.Fields {
			rParametersValFields = append(rParametersValFields, rParametersValFieldsVal)
		}
		rParametersObject["fields"] = rParametersValFields
		if rParametersVal.Name != nil {
			rParametersObject["name"] = *rParametersVal.Name
		}
		if rParametersVal.Validation != nil && rParametersVal.Validation != dclService.EmptyWorkflowTemplateParametersValidation {
			rParametersValValidation := make(map[string]interface{})
			if rParametersVal.Validation.Regex != nil && rParametersVal.Validation.Regex != dclService.EmptyWorkflowTemplateParametersValidationRegex {
				rParametersValValidationRegex := make(map[string]interface{})
				var rParametersValValidationRegexRegexes []interface{}
				for _, rParametersValValidationRegexRegexesVal := range rParametersVal.Validation.Regex.Regexes {
					rParametersValValidationRegexRegexes = append(rParametersValValidationRegexRegexes, rParametersValValidationRegexRegexesVal)
				}
				rParametersValValidationRegex["regexes"] = rParametersValValidationRegexRegexes
				rParametersValValidation["regex"] = rParametersValValidationRegex
			}
			if rParametersVal.Validation.Values != nil && rParametersVal.Validation.Values != dclService.EmptyWorkflowTemplateParametersValidationValues {
				rParametersValValidationValues := make(map[string]interface{})
				var rParametersValValidationValuesValues []interface{}
				for _, rParametersValValidationValuesValuesVal := range rParametersVal.Validation.Values.Values {
					rParametersValValidationValuesValues = append(rParametersValValidationValuesValues, rParametersValValidationValuesValuesVal)
				}
				rParametersValValidationValues["values"] = rParametersValValidationValuesValues
				rParametersValValidation["values"] = rParametersValValidationValues
			}
			rParametersObject["validation"] = rParametersValValidation
		}
		rParameters = append(rParameters, rParametersObject)
	}
	u.Object["parameters"] = rParameters
	if r.Placement != nil && r.Placement != dclService.EmptyWorkflowTemplatePlacement {
		rPlacement := make(map[string]interface{})
		if r.Placement.ClusterSelector != nil && r.Placement.ClusterSelector != dclService.EmptyWorkflowTemplatePlacementClusterSelector {
			rPlacementClusterSelector := make(map[string]interface{})
			if r.Placement.ClusterSelector.ClusterLabels != nil {
				rPlacementClusterSelectorClusterLabels := make(map[string]interface{})
				for k, v := range r.Placement.ClusterSelector.ClusterLabels {
					rPlacementClusterSelectorClusterLabels[k] = v
				}
				rPlacementClusterSelector["clusterLabels"] = rPlacementClusterSelectorClusterLabels
			}
			if r.Placement.ClusterSelector.Zone != nil {
				rPlacementClusterSelector["zone"] = *r.Placement.ClusterSelector.Zone
			}
			rPlacement["clusterSelector"] = rPlacementClusterSelector
		}
		if r.Placement.ManagedCluster != nil && r.Placement.ManagedCluster != dclService.EmptyWorkflowTemplatePlacementManagedCluster {
			rPlacementManagedCluster := make(map[string]interface{})
			if r.Placement.ManagedCluster.ClusterName != nil {
				rPlacementManagedCluster["clusterName"] = *r.Placement.ManagedCluster.ClusterName
			}
			if r.Placement.ManagedCluster.Config != nil && r.Placement.ManagedCluster.Config != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfig {
				rPlacementManagedClusterConfig := make(map[string]interface{})
				if r.Placement.ManagedCluster.Config.AutoscalingConfig != nil && r.Placement.ManagedCluster.Config.AutoscalingConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
					rPlacementManagedClusterConfigAutoscalingConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.AutoscalingConfig.Policy != nil {
						rPlacementManagedClusterConfigAutoscalingConfig["policy"] = *r.Placement.ManagedCluster.Config.AutoscalingConfig.Policy
					}
					rPlacementManagedClusterConfig["autoscalingConfig"] = rPlacementManagedClusterConfigAutoscalingConfig
				}
				if r.Placement.ManagedCluster.Config.EncryptionConfig != nil && r.Placement.ManagedCluster.Config.EncryptionConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
					rPlacementManagedClusterConfigEncryptionConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.EncryptionConfig.GcePdKmsKeyName != nil {
						rPlacementManagedClusterConfigEncryptionConfig["gcePdKmsKeyName"] = *r.Placement.ManagedCluster.Config.EncryptionConfig.GcePdKmsKeyName
					}
					rPlacementManagedClusterConfig["encryptionConfig"] = rPlacementManagedClusterConfigEncryptionConfig
				}
				if r.Placement.ManagedCluster.Config.EndpointConfig != nil && r.Placement.ManagedCluster.Config.EndpointConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
					rPlacementManagedClusterConfigEndpointConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.EndpointConfig.EnableHttpPortAccess != nil {
						rPlacementManagedClusterConfigEndpointConfig["enableHttpPortAccess"] = *r.Placement.ManagedCluster.Config.EndpointConfig.EnableHttpPortAccess
					}
					if r.Placement.ManagedCluster.Config.EndpointConfig.HttpPorts != nil {
						rPlacementManagedClusterConfigEndpointConfigHttpPorts := make(map[string]interface{})
						for k, v := range r.Placement.ManagedCluster.Config.EndpointConfig.HttpPorts {
							rPlacementManagedClusterConfigEndpointConfigHttpPorts[k] = v
						}
						rPlacementManagedClusterConfigEndpointConfig["httpPorts"] = rPlacementManagedClusterConfigEndpointConfigHttpPorts
					}
					rPlacementManagedClusterConfig["endpointConfig"] = rPlacementManagedClusterConfigEndpointConfig
				}
				if r.Placement.ManagedCluster.Config.GceClusterConfig != nil && r.Placement.ManagedCluster.Config.GceClusterConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
					rPlacementManagedClusterConfigGceClusterConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.GceClusterConfig.InternalIPOnly != nil {
						rPlacementManagedClusterConfigGceClusterConfig["internalIPOnly"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.InternalIPOnly
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.Metadata != nil {
						rPlacementManagedClusterConfigGceClusterConfigMetadata := make(map[string]interface{})
						for k, v := range r.Placement.ManagedCluster.Config.GceClusterConfig.Metadata {
							rPlacementManagedClusterConfigGceClusterConfigMetadata[k] = v
						}
						rPlacementManagedClusterConfigGceClusterConfig["metadata"] = rPlacementManagedClusterConfigGceClusterConfigMetadata
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.Network != nil {
						rPlacementManagedClusterConfigGceClusterConfig["network"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.Network
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity != nil && r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
						rPlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup != nil {
							rPlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity["nodeGroup"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup
						}
						rPlacementManagedClusterConfigGceClusterConfig["nodeGroupAffinity"] = rPlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.PrivateIPv6GoogleAccess != nil {
						rPlacementManagedClusterConfigGceClusterConfig["privateIPv6GoogleAccess"] = string(*r.Placement.ManagedCluster.Config.GceClusterConfig.PrivateIPv6GoogleAccess)
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity != nil && r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
						rPlacementManagedClusterConfigGceClusterConfigReservationAffinity := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType != nil {
							rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["consumeReservationType"] = string(*r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType)
						}
						if r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Key != nil {
							rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["key"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Key
						}
						var rPlacementManagedClusterConfigGceClusterConfigReservationAffinityValues []interface{}
						for _, rPlacementManagedClusterConfigGceClusterConfigReservationAffinityValuesVal := range r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Values {
							rPlacementManagedClusterConfigGceClusterConfigReservationAffinityValues = append(rPlacementManagedClusterConfigGceClusterConfigReservationAffinityValues, rPlacementManagedClusterConfigGceClusterConfigReservationAffinityValuesVal)
						}
						rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["values"] = rPlacementManagedClusterConfigGceClusterConfigReservationAffinityValues
						rPlacementManagedClusterConfigGceClusterConfig["reservationAffinity"] = rPlacementManagedClusterConfigGceClusterConfigReservationAffinity
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccount != nil {
						rPlacementManagedClusterConfigGceClusterConfig["serviceAccount"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccount
					}
					var rPlacementManagedClusterConfigGceClusterConfigServiceAccountScopes []interface{}
					for _, rPlacementManagedClusterConfigGceClusterConfigServiceAccountScopesVal := range r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccountScopes {
						rPlacementManagedClusterConfigGceClusterConfigServiceAccountScopes = append(rPlacementManagedClusterConfigGceClusterConfigServiceAccountScopes, rPlacementManagedClusterConfigGceClusterConfigServiceAccountScopesVal)
					}
					rPlacementManagedClusterConfigGceClusterConfig["serviceAccountScopes"] = rPlacementManagedClusterConfigGceClusterConfigServiceAccountScopes
					if r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig != nil && r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
						rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring != nil {
							rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableIntegrityMonitoring"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring
						}
						if r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot != nil {
							rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableSecureBoot"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot
						}
						if r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm != nil {
							rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableVtpm"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm
						}
						rPlacementManagedClusterConfigGceClusterConfig["shieldedInstanceConfig"] = rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig
					}
					if r.Placement.ManagedCluster.Config.GceClusterConfig.Subnetwork != nil {
						rPlacementManagedClusterConfigGceClusterConfig["subnetwork"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.Subnetwork
					}
					var rPlacementManagedClusterConfigGceClusterConfigTags []interface{}
					for _, rPlacementManagedClusterConfigGceClusterConfigTagsVal := range r.Placement.ManagedCluster.Config.GceClusterConfig.Tags {
						rPlacementManagedClusterConfigGceClusterConfigTags = append(rPlacementManagedClusterConfigGceClusterConfigTags, rPlacementManagedClusterConfigGceClusterConfigTagsVal)
					}
					rPlacementManagedClusterConfigGceClusterConfig["tags"] = rPlacementManagedClusterConfigGceClusterConfigTags
					if r.Placement.ManagedCluster.Config.GceClusterConfig.Zone != nil {
						rPlacementManagedClusterConfigGceClusterConfig["zone"] = *r.Placement.ManagedCluster.Config.GceClusterConfig.Zone
					}
					rPlacementManagedClusterConfig["gceClusterConfig"] = rPlacementManagedClusterConfigGceClusterConfig
				}
				if r.Placement.ManagedCluster.Config.GkeClusterConfig != nil && r.Placement.ManagedCluster.Config.GkeClusterConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig {
					rPlacementManagedClusterConfigGkeClusterConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget != nil && r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
						rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace != nil {
							rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"] = *r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace
						}
						if r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster != nil {
							rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"] = *r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster
						}
						rPlacementManagedClusterConfigGkeClusterConfig["namespacedGkeDeploymentTarget"] = rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget
					}
					rPlacementManagedClusterConfig["gkeClusterConfig"] = rPlacementManagedClusterConfigGkeClusterConfig
				}
				var rPlacementManagedClusterConfigInitializationActions []interface{}
				for _, rPlacementManagedClusterConfigInitializationActionsVal := range r.Placement.ManagedCluster.Config.InitializationActions {
					rPlacementManagedClusterConfigInitializationActionsObject := make(map[string]interface{})
					if rPlacementManagedClusterConfigInitializationActionsVal.ExecutableFile != nil {
						rPlacementManagedClusterConfigInitializationActionsObject["executableFile"] = *rPlacementManagedClusterConfigInitializationActionsVal.ExecutableFile
					}
					if rPlacementManagedClusterConfigInitializationActionsVal.ExecutionTimeout != nil {
						rPlacementManagedClusterConfigInitializationActionsObject["executionTimeout"] = *rPlacementManagedClusterConfigInitializationActionsVal.ExecutionTimeout
					}
					rPlacementManagedClusterConfigInitializationActions = append(rPlacementManagedClusterConfigInitializationActions, rPlacementManagedClusterConfigInitializationActionsObject)
				}
				rPlacementManagedClusterConfig["initializationActions"] = rPlacementManagedClusterConfigInitializationActions
				if r.Placement.ManagedCluster.Config.LifecycleConfig != nil && r.Placement.ManagedCluster.Config.LifecycleConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
					rPlacementManagedClusterConfigLifecycleConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTime != nil {
						rPlacementManagedClusterConfigLifecycleConfig["autoDeleteTime"] = *r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTime
					}
					if r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTtl != nil {
						rPlacementManagedClusterConfigLifecycleConfig["autoDeleteTtl"] = *r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTtl
					}
					if r.Placement.ManagedCluster.Config.LifecycleConfig.IdleDeleteTtl != nil {
						rPlacementManagedClusterConfigLifecycleConfig["idleDeleteTtl"] = *r.Placement.ManagedCluster.Config.LifecycleConfig.IdleDeleteTtl
					}
					if r.Placement.ManagedCluster.Config.LifecycleConfig.IdleStartTime != nil {
						rPlacementManagedClusterConfigLifecycleConfig["idleStartTime"] = *r.Placement.ManagedCluster.Config.LifecycleConfig.IdleStartTime
					}
					rPlacementManagedClusterConfig["lifecycleConfig"] = rPlacementManagedClusterConfigLifecycleConfig
				}
				if r.Placement.ManagedCluster.Config.MasterConfig != nil && r.Placement.ManagedCluster.Config.MasterConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigMasterConfig {
					rPlacementManagedClusterConfigMasterConfig := make(map[string]interface{})
					var rPlacementManagedClusterConfigMasterConfigAccelerators []interface{}
					for _, rPlacementManagedClusterConfigMasterConfigAcceleratorsVal := range r.Placement.ManagedCluster.Config.MasterConfig.Accelerators {
						rPlacementManagedClusterConfigMasterConfigAcceleratorsObject := make(map[string]interface{})
						if rPlacementManagedClusterConfigMasterConfigAcceleratorsVal.AcceleratorCount != nil {
							rPlacementManagedClusterConfigMasterConfigAcceleratorsObject["acceleratorCount"] = *rPlacementManagedClusterConfigMasterConfigAcceleratorsVal.AcceleratorCount
						}
						if rPlacementManagedClusterConfigMasterConfigAcceleratorsVal.AcceleratorType != nil {
							rPlacementManagedClusterConfigMasterConfigAcceleratorsObject["acceleratorType"] = *rPlacementManagedClusterConfigMasterConfigAcceleratorsVal.AcceleratorType
						}
						rPlacementManagedClusterConfigMasterConfigAccelerators = append(rPlacementManagedClusterConfigMasterConfigAccelerators, rPlacementManagedClusterConfigMasterConfigAcceleratorsObject)
					}
					rPlacementManagedClusterConfigMasterConfig["accelerators"] = rPlacementManagedClusterConfigMasterConfigAccelerators
					if r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig != nil && r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
						rPlacementManagedClusterConfigMasterConfigDiskConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskSizeGb != nil {
							rPlacementManagedClusterConfigMasterConfigDiskConfig["bootDiskSizeGb"] = *r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskSizeGb
						}
						if r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskType != nil {
							rPlacementManagedClusterConfigMasterConfigDiskConfig["bootDiskType"] = *r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskType
						}
						if r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.NumLocalSsds != nil {
							rPlacementManagedClusterConfigMasterConfigDiskConfig["numLocalSsds"] = *r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.NumLocalSsds
						}
						rPlacementManagedClusterConfigMasterConfig["diskConfig"] = rPlacementManagedClusterConfigMasterConfigDiskConfig
					}
					if r.Placement.ManagedCluster.Config.MasterConfig.Image != nil {
						rPlacementManagedClusterConfigMasterConfig["image"] = *r.Placement.ManagedCluster.Config.MasterConfig.Image
					}
					var rPlacementManagedClusterConfigMasterConfigInstanceNames []interface{}
					for _, rPlacementManagedClusterConfigMasterConfigInstanceNamesVal := range r.Placement.ManagedCluster.Config.MasterConfig.InstanceNames {
						rPlacementManagedClusterConfigMasterConfigInstanceNames = append(rPlacementManagedClusterConfigMasterConfigInstanceNames, rPlacementManagedClusterConfigMasterConfigInstanceNamesVal)
					}
					rPlacementManagedClusterConfigMasterConfig["instanceNames"] = rPlacementManagedClusterConfigMasterConfigInstanceNames
					if r.Placement.ManagedCluster.Config.MasterConfig.IsPreemptible != nil {
						rPlacementManagedClusterConfigMasterConfig["isPreemptible"] = *r.Placement.ManagedCluster.Config.MasterConfig.IsPreemptible
					}
					if r.Placement.ManagedCluster.Config.MasterConfig.MachineType != nil {
						rPlacementManagedClusterConfigMasterConfig["machineType"] = *r.Placement.ManagedCluster.Config.MasterConfig.MachineType
					}
					if r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig != nil && r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
						rPlacementManagedClusterConfigMasterConfigManagedGroupConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName != nil {
							rPlacementManagedClusterConfigMasterConfigManagedGroupConfig["instanceGroupManagerName"] = *r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName
						}
						if r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName != nil {
							rPlacementManagedClusterConfigMasterConfigManagedGroupConfig["instanceTemplateName"] = *r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName
						}
						rPlacementManagedClusterConfigMasterConfig["managedGroupConfig"] = rPlacementManagedClusterConfigMasterConfigManagedGroupConfig
					}
					if r.Placement.ManagedCluster.Config.MasterConfig.MinCpuPlatform != nil {
						rPlacementManagedClusterConfigMasterConfig["minCpuPlatform"] = *r.Placement.ManagedCluster.Config.MasterConfig.MinCpuPlatform
					}
					if r.Placement.ManagedCluster.Config.MasterConfig.NumInstances != nil {
						rPlacementManagedClusterConfigMasterConfig["numInstances"] = *r.Placement.ManagedCluster.Config.MasterConfig.NumInstances
					}
					if r.Placement.ManagedCluster.Config.MasterConfig.Preemptibility != nil {
						rPlacementManagedClusterConfigMasterConfig["preemptibility"] = string(*r.Placement.ManagedCluster.Config.MasterConfig.Preemptibility)
					}
					rPlacementManagedClusterConfig["masterConfig"] = rPlacementManagedClusterConfigMasterConfig
				}
				if r.Placement.ManagedCluster.Config.MetastoreConfig != nil && r.Placement.ManagedCluster.Config.MetastoreConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig {
					rPlacementManagedClusterConfigMetastoreConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.MetastoreConfig.DataprocMetastoreService != nil {
						rPlacementManagedClusterConfigMetastoreConfig["dataprocMetastoreService"] = *r.Placement.ManagedCluster.Config.MetastoreConfig.DataprocMetastoreService
					}
					rPlacementManagedClusterConfig["metastoreConfig"] = rPlacementManagedClusterConfigMetastoreConfig
				}
				if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig != nil && r.Placement.ManagedCluster.Config.SecondaryWorkerConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
					rPlacementManagedClusterConfigSecondaryWorkerConfig := make(map[string]interface{})
					var rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators []interface{}
					for _, rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsVal := range r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Accelerators {
						rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsObject := make(map[string]interface{})
						if rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorCount != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsObject["acceleratorCount"] = *rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorCount
						}
						if rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorType != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsObject["acceleratorType"] = *rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorType
						}
						rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators = append(rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators, rPlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsObject)
					}
					rPlacementManagedClusterConfigSecondaryWorkerConfig["accelerators"] = rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig != nil && r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
						rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["bootDiskSizeGb"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb
						}
						if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["bootDiskType"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType
						}
						if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["numLocalSsds"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds
						}
						rPlacementManagedClusterConfigSecondaryWorkerConfig["diskConfig"] = rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig
					}
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Image != nil {
						rPlacementManagedClusterConfigSecondaryWorkerConfig["image"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Image
					}
					var rPlacementManagedClusterConfigSecondaryWorkerConfigInstanceNames []interface{}
					for _, rPlacementManagedClusterConfigSecondaryWorkerConfigInstanceNamesVal := range r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.InstanceNames {
						rPlacementManagedClusterConfigSecondaryWorkerConfigInstanceNames = append(rPlacementManagedClusterConfigSecondaryWorkerConfigInstanceNames, rPlacementManagedClusterConfigSecondaryWorkerConfigInstanceNamesVal)
					}
					rPlacementManagedClusterConfigSecondaryWorkerConfig["instanceNames"] = rPlacementManagedClusterConfigSecondaryWorkerConfigInstanceNames
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.IsPreemptible != nil {
						rPlacementManagedClusterConfigSecondaryWorkerConfig["isPreemptible"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.IsPreemptible
					}
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MachineType != nil {
						rPlacementManagedClusterConfigSecondaryWorkerConfig["machineType"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MachineType
					}
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig != nil && r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
						rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig["instanceGroupManagerName"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName
						}
						if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName != nil {
							rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig["instanceTemplateName"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName
						}
						rPlacementManagedClusterConfigSecondaryWorkerConfig["managedGroupConfig"] = rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig
					}
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MinCpuPlatform != nil {
						rPlacementManagedClusterConfigSecondaryWorkerConfig["minCpuPlatform"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MinCpuPlatform
					}
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.NumInstances != nil {
						rPlacementManagedClusterConfigSecondaryWorkerConfig["numInstances"] = *r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.NumInstances
					}
					if r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Preemptibility != nil {
						rPlacementManagedClusterConfigSecondaryWorkerConfig["preemptibility"] = string(*r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Preemptibility)
					}
					rPlacementManagedClusterConfig["secondaryWorkerConfig"] = rPlacementManagedClusterConfigSecondaryWorkerConfig
				}
				if r.Placement.ManagedCluster.Config.SecurityConfig != nil && r.Placement.ManagedCluster.Config.SecurityConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
					rPlacementManagedClusterConfigSecurityConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig != nil && r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
						rPlacementManagedClusterConfigSecurityConfigKerberosConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustAdminServer"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustKdc"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustRealm"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.EnableKerberos != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["enableKerberos"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.EnableKerberos
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KdcDbKey != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["kdcDbKey"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KdcDbKey
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeyPassword != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keyPassword"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeyPassword
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Keystore != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keystore"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Keystore
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeystorePassword != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keystorePassword"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeystorePassword
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KmsKey != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["kmsKey"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KmsKey
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Realm != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["realm"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Realm
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["rootPrincipalPassword"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["tgtLifetimeHours"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Truststore != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["truststore"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Truststore
						}
						if r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TruststorePassword != nil {
							rPlacementManagedClusterConfigSecurityConfigKerberosConfig["truststorePassword"] = *r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TruststorePassword
						}
						rPlacementManagedClusterConfigSecurityConfig["kerberosConfig"] = rPlacementManagedClusterConfigSecurityConfigKerberosConfig
					}
					rPlacementManagedClusterConfig["securityConfig"] = rPlacementManagedClusterConfigSecurityConfig
				}
				if r.Placement.ManagedCluster.Config.SoftwareConfig != nil && r.Placement.ManagedCluster.Config.SoftwareConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
					rPlacementManagedClusterConfigSoftwareConfig := make(map[string]interface{})
					if r.Placement.ManagedCluster.Config.SoftwareConfig.ImageVersion != nil {
						rPlacementManagedClusterConfigSoftwareConfig["imageVersion"] = *r.Placement.ManagedCluster.Config.SoftwareConfig.ImageVersion
					}
					var rPlacementManagedClusterConfigSoftwareConfigOptionalComponents []interface{}
					for _, rPlacementManagedClusterConfigSoftwareConfigOptionalComponentsVal := range r.Placement.ManagedCluster.Config.SoftwareConfig.OptionalComponents {
						rPlacementManagedClusterConfigSoftwareConfigOptionalComponents = append(rPlacementManagedClusterConfigSoftwareConfigOptionalComponents, string(rPlacementManagedClusterConfigSoftwareConfigOptionalComponentsVal))
					}
					rPlacementManagedClusterConfigSoftwareConfig["optionalComponents"] = rPlacementManagedClusterConfigSoftwareConfigOptionalComponents
					if r.Placement.ManagedCluster.Config.SoftwareConfig.Properties != nil {
						rPlacementManagedClusterConfigSoftwareConfigProperties := make(map[string]interface{})
						for k, v := range r.Placement.ManagedCluster.Config.SoftwareConfig.Properties {
							rPlacementManagedClusterConfigSoftwareConfigProperties[k] = v
						}
						rPlacementManagedClusterConfigSoftwareConfig["properties"] = rPlacementManagedClusterConfigSoftwareConfigProperties
					}
					rPlacementManagedClusterConfig["softwareConfig"] = rPlacementManagedClusterConfigSoftwareConfig
				}
				if r.Placement.ManagedCluster.Config.StagingBucket != nil {
					rPlacementManagedClusterConfig["stagingBucket"] = *r.Placement.ManagedCluster.Config.StagingBucket
				}
				if r.Placement.ManagedCluster.Config.TempBucket != nil {
					rPlacementManagedClusterConfig["tempBucket"] = *r.Placement.ManagedCluster.Config.TempBucket
				}
				if r.Placement.ManagedCluster.Config.WorkerConfig != nil && r.Placement.ManagedCluster.Config.WorkerConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
					rPlacementManagedClusterConfigWorkerConfig := make(map[string]interface{})
					var rPlacementManagedClusterConfigWorkerConfigAccelerators []interface{}
					for _, rPlacementManagedClusterConfigWorkerConfigAcceleratorsVal := range r.Placement.ManagedCluster.Config.WorkerConfig.Accelerators {
						rPlacementManagedClusterConfigWorkerConfigAcceleratorsObject := make(map[string]interface{})
						if rPlacementManagedClusterConfigWorkerConfigAcceleratorsVal.AcceleratorCount != nil {
							rPlacementManagedClusterConfigWorkerConfigAcceleratorsObject["acceleratorCount"] = *rPlacementManagedClusterConfigWorkerConfigAcceleratorsVal.AcceleratorCount
						}
						if rPlacementManagedClusterConfigWorkerConfigAcceleratorsVal.AcceleratorType != nil {
							rPlacementManagedClusterConfigWorkerConfigAcceleratorsObject["acceleratorType"] = *rPlacementManagedClusterConfigWorkerConfigAcceleratorsVal.AcceleratorType
						}
						rPlacementManagedClusterConfigWorkerConfigAccelerators = append(rPlacementManagedClusterConfigWorkerConfigAccelerators, rPlacementManagedClusterConfigWorkerConfigAcceleratorsObject)
					}
					rPlacementManagedClusterConfigWorkerConfig["accelerators"] = rPlacementManagedClusterConfigWorkerConfigAccelerators
					if r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig != nil && r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
						rPlacementManagedClusterConfigWorkerConfigDiskConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskSizeGb != nil {
							rPlacementManagedClusterConfigWorkerConfigDiskConfig["bootDiskSizeGb"] = *r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskSizeGb
						}
						if r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskType != nil {
							rPlacementManagedClusterConfigWorkerConfigDiskConfig["bootDiskType"] = *r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskType
						}
						if r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.NumLocalSsds != nil {
							rPlacementManagedClusterConfigWorkerConfigDiskConfig["numLocalSsds"] = *r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.NumLocalSsds
						}
						rPlacementManagedClusterConfigWorkerConfig["diskConfig"] = rPlacementManagedClusterConfigWorkerConfigDiskConfig
					}
					if r.Placement.ManagedCluster.Config.WorkerConfig.Image != nil {
						rPlacementManagedClusterConfigWorkerConfig["image"] = *r.Placement.ManagedCluster.Config.WorkerConfig.Image
					}
					var rPlacementManagedClusterConfigWorkerConfigInstanceNames []interface{}
					for _, rPlacementManagedClusterConfigWorkerConfigInstanceNamesVal := range r.Placement.ManagedCluster.Config.WorkerConfig.InstanceNames {
						rPlacementManagedClusterConfigWorkerConfigInstanceNames = append(rPlacementManagedClusterConfigWorkerConfigInstanceNames, rPlacementManagedClusterConfigWorkerConfigInstanceNamesVal)
					}
					rPlacementManagedClusterConfigWorkerConfig["instanceNames"] = rPlacementManagedClusterConfigWorkerConfigInstanceNames
					if r.Placement.ManagedCluster.Config.WorkerConfig.IsPreemptible != nil {
						rPlacementManagedClusterConfigWorkerConfig["isPreemptible"] = *r.Placement.ManagedCluster.Config.WorkerConfig.IsPreemptible
					}
					if r.Placement.ManagedCluster.Config.WorkerConfig.MachineType != nil {
						rPlacementManagedClusterConfigWorkerConfig["machineType"] = *r.Placement.ManagedCluster.Config.WorkerConfig.MachineType
					}
					if r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig != nil && r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig != dclService.EmptyWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
						rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig := make(map[string]interface{})
						if r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName != nil {
							rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig["instanceGroupManagerName"] = *r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName
						}
						if r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName != nil {
							rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig["instanceTemplateName"] = *r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName
						}
						rPlacementManagedClusterConfigWorkerConfig["managedGroupConfig"] = rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig
					}
					if r.Placement.ManagedCluster.Config.WorkerConfig.MinCpuPlatform != nil {
						rPlacementManagedClusterConfigWorkerConfig["minCpuPlatform"] = *r.Placement.ManagedCluster.Config.WorkerConfig.MinCpuPlatform
					}
					if r.Placement.ManagedCluster.Config.WorkerConfig.NumInstances != nil {
						rPlacementManagedClusterConfigWorkerConfig["numInstances"] = *r.Placement.ManagedCluster.Config.WorkerConfig.NumInstances
					}
					if r.Placement.ManagedCluster.Config.WorkerConfig.Preemptibility != nil {
						rPlacementManagedClusterConfigWorkerConfig["preemptibility"] = string(*r.Placement.ManagedCluster.Config.WorkerConfig.Preemptibility)
					}
					rPlacementManagedClusterConfig["workerConfig"] = rPlacementManagedClusterConfigWorkerConfig
				}
				rPlacementManagedCluster["config"] = rPlacementManagedClusterConfig
			}
			if r.Placement.ManagedCluster.Labels != nil {
				rPlacementManagedClusterLabels := make(map[string]interface{})
				for k, v := range r.Placement.ManagedCluster.Labels {
					rPlacementManagedClusterLabels[k] = v
				}
				rPlacementManagedCluster["labels"] = rPlacementManagedClusterLabels
			}
			rPlacement["managedCluster"] = rPlacementManagedCluster
		}
		u.Object["placement"] = rPlacement
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.Version != nil {
		u.Object["version"] = *r.Version
	}
	return u
}

func UnstructuredToWorkflowTemplate(u *unstructured.Resource) (*dclService.WorkflowTemplate, error) {
	r := &dclService.WorkflowTemplate{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["dagTimeout"]; ok {
		if s, ok := u.Object["dagTimeout"].(string); ok {
			r.DagTimeout = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DagTimeout: expected string")
		}
	}
	if _, ok := u.Object["jobs"]; ok {
		if s, ok := u.Object["jobs"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rJobs dclService.WorkflowTemplateJobs
					if _, ok := objval["hadoopJob"]; ok {
						if rJobsHadoopJob, ok := objval["hadoopJob"].(map[string]interface{}); ok {
							rJobs.HadoopJob = &dclService.WorkflowTemplateJobsHadoopJob{}
							if _, ok := rJobsHadoopJob["archiveUris"]; ok {
								if s, ok := rJobsHadoopJob["archiveUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.HadoopJob.ArchiveUris = append(rJobs.HadoopJob.ArchiveUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.ArchiveUris: expected []interface{}")
								}
							}
							if _, ok := rJobsHadoopJob["args"]; ok {
								if s, ok := rJobsHadoopJob["args"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.HadoopJob.Args = append(rJobs.HadoopJob.Args, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.Args: expected []interface{}")
								}
							}
							if _, ok := rJobsHadoopJob["fileUris"]; ok {
								if s, ok := rJobsHadoopJob["fileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.HadoopJob.FileUris = append(rJobs.HadoopJob.FileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.FileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsHadoopJob["jarFileUris"]; ok {
								if s, ok := rJobsHadoopJob["jarFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.HadoopJob.JarFileUris = append(rJobs.HadoopJob.JarFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.JarFileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsHadoopJob["loggingConfig"]; ok {
								if rJobsHadoopJobLoggingConfig, ok := rJobsHadoopJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.HadoopJob.LoggingConfig = &dclService.WorkflowTemplateJobsHadoopJobLoggingConfig{}
									if _, ok := rJobsHadoopJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsHadoopJobLoggingConfigDriverLogLevels, ok := rJobsHadoopJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsHadoopJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.HadoopJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.HadoopJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsHadoopJob["mainClass"]; ok {
								if s, ok := rJobsHadoopJob["mainClass"].(string); ok {
									rJobs.HadoopJob.MainClass = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.MainClass: expected string")
								}
							}
							if _, ok := rJobsHadoopJob["mainJarFileUri"]; ok {
								if s, ok := rJobsHadoopJob["mainJarFileUri"].(string); ok {
									rJobs.HadoopJob.MainJarFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.MainJarFileUri: expected string")
								}
							}
							if _, ok := rJobsHadoopJob["properties"]; ok {
								if rJobsHadoopJobProperties, ok := rJobsHadoopJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsHadoopJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.HadoopJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.HadoopJob.Properties: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.HadoopJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["hiveJob"]; ok {
						if rJobsHiveJob, ok := objval["hiveJob"].(map[string]interface{}); ok {
							rJobs.HiveJob = &dclService.WorkflowTemplateJobsHiveJob{}
							if _, ok := rJobsHiveJob["continueOnFailure"]; ok {
								if b, ok := rJobsHiveJob["continueOnFailure"].(bool); ok {
									rJobs.HiveJob.ContinueOnFailure = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rJobs.HiveJob.ContinueOnFailure: expected bool")
								}
							}
							if _, ok := rJobsHiveJob["jarFileUris"]; ok {
								if s, ok := rJobsHiveJob["jarFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.HiveJob.JarFileUris = append(rJobs.HiveJob.JarFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HiveJob.JarFileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsHiveJob["properties"]; ok {
								if rJobsHiveJobProperties, ok := rJobsHiveJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsHiveJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.HiveJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.HiveJob.Properties: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsHiveJob["queryFileUri"]; ok {
								if s, ok := rJobsHiveJob["queryFileUri"].(string); ok {
									rJobs.HiveJob.QueryFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.HiveJob.QueryFileUri: expected string")
								}
							}
							if _, ok := rJobsHiveJob["queryList"]; ok {
								if rJobsHiveJobQueryList, ok := rJobsHiveJob["queryList"].(map[string]interface{}); ok {
									rJobs.HiveJob.QueryList = &dclService.WorkflowTemplateJobsHiveJobQueryList{}
									if _, ok := rJobsHiveJobQueryList["queries"]; ok {
										if s, ok := rJobsHiveJobQueryList["queries"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rJobs.HiveJob.QueryList.Queries = append(rJobs.HiveJob.QueryList.Queries, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rJobs.HiveJob.QueryList.Queries: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.HiveJob.QueryList: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsHiveJob["scriptVariables"]; ok {
								if rJobsHiveJobScriptVariables, ok := rJobsHiveJob["scriptVariables"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsHiveJobScriptVariables {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.HiveJob.ScriptVariables = m
								} else {
									return nil, fmt.Errorf("rJobs.HiveJob.ScriptVariables: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.HiveJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["labels"]; ok {
						if rJobsLabels, ok := objval["labels"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rJobsLabels {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							rJobs.Labels = m
						} else {
							return nil, fmt.Errorf("rJobs.Labels: expected map[string]interface{}")
						}
					}
					if _, ok := objval["pigJob"]; ok {
						if rJobsPigJob, ok := objval["pigJob"].(map[string]interface{}); ok {
							rJobs.PigJob = &dclService.WorkflowTemplateJobsPigJob{}
							if _, ok := rJobsPigJob["continueOnFailure"]; ok {
								if b, ok := rJobsPigJob["continueOnFailure"].(bool); ok {
									rJobs.PigJob.ContinueOnFailure = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.ContinueOnFailure: expected bool")
								}
							}
							if _, ok := rJobsPigJob["jarFileUris"]; ok {
								if s, ok := rJobsPigJob["jarFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PigJob.JarFileUris = append(rJobs.PigJob.JarFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.JarFileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsPigJob["loggingConfig"]; ok {
								if rJobsPigJobLoggingConfig, ok := rJobsPigJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.PigJob.LoggingConfig = &dclService.WorkflowTemplateJobsPigJobLoggingConfig{}
									if _, ok := rJobsPigJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsPigJobLoggingConfigDriverLogLevels, ok := rJobsPigJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsPigJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.PigJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.PigJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPigJob["properties"]; ok {
								if rJobsPigJobProperties, ok := rJobsPigJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsPigJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.PigJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.Properties: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPigJob["queryFileUri"]; ok {
								if s, ok := rJobsPigJob["queryFileUri"].(string); ok {
									rJobs.PigJob.QueryFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.QueryFileUri: expected string")
								}
							}
							if _, ok := rJobsPigJob["queryList"]; ok {
								if rJobsPigJobQueryList, ok := rJobsPigJob["queryList"].(map[string]interface{}); ok {
									rJobs.PigJob.QueryList = &dclService.WorkflowTemplateJobsPigJobQueryList{}
									if _, ok := rJobsPigJobQueryList["queries"]; ok {
										if s, ok := rJobsPigJobQueryList["queries"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rJobs.PigJob.QueryList.Queries = append(rJobs.PigJob.QueryList.Queries, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rJobs.PigJob.QueryList.Queries: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.QueryList: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPigJob["scriptVariables"]; ok {
								if rJobsPigJobScriptVariables, ok := rJobsPigJob["scriptVariables"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsPigJobScriptVariables {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.PigJob.ScriptVariables = m
								} else {
									return nil, fmt.Errorf("rJobs.PigJob.ScriptVariables: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.PigJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["prerequisiteStepIds"]; ok {
						if s, ok := objval["prerequisiteStepIds"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rJobs.PrerequisiteStepIds = append(rJobs.PrerequisiteStepIds, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.PrerequisiteStepIds: expected []interface{}")
						}
					}
					if _, ok := objval["prestoJob"]; ok {
						if rJobsPrestoJob, ok := objval["prestoJob"].(map[string]interface{}); ok {
							rJobs.PrestoJob = &dclService.WorkflowTemplateJobsPrestoJob{}
							if _, ok := rJobsPrestoJob["clientTags"]; ok {
								if s, ok := rJobsPrestoJob["clientTags"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PrestoJob.ClientTags = append(rJobs.PrestoJob.ClientTags, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.ClientTags: expected []interface{}")
								}
							}
							if _, ok := rJobsPrestoJob["continueOnFailure"]; ok {
								if b, ok := rJobsPrestoJob["continueOnFailure"].(bool); ok {
									rJobs.PrestoJob.ContinueOnFailure = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.ContinueOnFailure: expected bool")
								}
							}
							if _, ok := rJobsPrestoJob["loggingConfig"]; ok {
								if rJobsPrestoJobLoggingConfig, ok := rJobsPrestoJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.PrestoJob.LoggingConfig = &dclService.WorkflowTemplateJobsPrestoJobLoggingConfig{}
									if _, ok := rJobsPrestoJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsPrestoJobLoggingConfigDriverLogLevels, ok := rJobsPrestoJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsPrestoJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.PrestoJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.PrestoJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPrestoJob["outputFormat"]; ok {
								if s, ok := rJobsPrestoJob["outputFormat"].(string); ok {
									rJobs.PrestoJob.OutputFormat = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.OutputFormat: expected string")
								}
							}
							if _, ok := rJobsPrestoJob["properties"]; ok {
								if rJobsPrestoJobProperties, ok := rJobsPrestoJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsPrestoJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.PrestoJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.Properties: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPrestoJob["queryFileUri"]; ok {
								if s, ok := rJobsPrestoJob["queryFileUri"].(string); ok {
									rJobs.PrestoJob.QueryFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.QueryFileUri: expected string")
								}
							}
							if _, ok := rJobsPrestoJob["queryList"]; ok {
								if rJobsPrestoJobQueryList, ok := rJobsPrestoJob["queryList"].(map[string]interface{}); ok {
									rJobs.PrestoJob.QueryList = &dclService.WorkflowTemplateJobsPrestoJobQueryList{}
									if _, ok := rJobsPrestoJobQueryList["queries"]; ok {
										if s, ok := rJobsPrestoJobQueryList["queries"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rJobs.PrestoJob.QueryList.Queries = append(rJobs.PrestoJob.QueryList.Queries, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rJobs.PrestoJob.QueryList.Queries: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PrestoJob.QueryList: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.PrestoJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["pysparkJob"]; ok {
						if rJobsPysparkJob, ok := objval["pysparkJob"].(map[string]interface{}); ok {
							rJobs.PysparkJob = &dclService.WorkflowTemplateJobsPysparkJob{}
							if _, ok := rJobsPysparkJob["archiveUris"]; ok {
								if s, ok := rJobsPysparkJob["archiveUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PysparkJob.ArchiveUris = append(rJobs.PysparkJob.ArchiveUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.ArchiveUris: expected []interface{}")
								}
							}
							if _, ok := rJobsPysparkJob["args"]; ok {
								if s, ok := rJobsPysparkJob["args"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PysparkJob.Args = append(rJobs.PysparkJob.Args, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.Args: expected []interface{}")
								}
							}
							if _, ok := rJobsPysparkJob["fileUris"]; ok {
								if s, ok := rJobsPysparkJob["fileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PysparkJob.FileUris = append(rJobs.PysparkJob.FileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.FileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsPysparkJob["jarFileUris"]; ok {
								if s, ok := rJobsPysparkJob["jarFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PysparkJob.JarFileUris = append(rJobs.PysparkJob.JarFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.JarFileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsPysparkJob["loggingConfig"]; ok {
								if rJobsPysparkJobLoggingConfig, ok := rJobsPysparkJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.PysparkJob.LoggingConfig = &dclService.WorkflowTemplateJobsPysparkJobLoggingConfig{}
									if _, ok := rJobsPysparkJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsPysparkJobLoggingConfigDriverLogLevels, ok := rJobsPysparkJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsPysparkJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.PysparkJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.PysparkJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPysparkJob["mainPythonFileUri"]; ok {
								if s, ok := rJobsPysparkJob["mainPythonFileUri"].(string); ok {
									rJobs.PysparkJob.MainPythonFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.MainPythonFileUri: expected string")
								}
							}
							if _, ok := rJobsPysparkJob["properties"]; ok {
								if rJobsPysparkJobProperties, ok := rJobsPysparkJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsPysparkJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.PysparkJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.Properties: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsPysparkJob["pythonFileUris"]; ok {
								if s, ok := rJobsPysparkJob["pythonFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.PysparkJob.PythonFileUris = append(rJobs.PysparkJob.PythonFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.PysparkJob.PythonFileUris: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.PysparkJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["scheduling"]; ok {
						if rJobsScheduling, ok := objval["scheduling"].(map[string]interface{}); ok {
							rJobs.Scheduling = &dclService.WorkflowTemplateJobsScheduling{}
							if _, ok := rJobsScheduling["maxFailuresPerHour"]; ok {
								if i, ok := rJobsScheduling["maxFailuresPerHour"].(int64); ok {
									rJobs.Scheduling.MaxFailuresPerHour = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rJobs.Scheduling.MaxFailuresPerHour: expected int64")
								}
							}
							if _, ok := rJobsScheduling["maxFailuresTotal"]; ok {
								if i, ok := rJobsScheduling["maxFailuresTotal"].(int64); ok {
									rJobs.Scheduling.MaxFailuresTotal = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rJobs.Scheduling.MaxFailuresTotal: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.Scheduling: expected map[string]interface{}")
						}
					}
					if _, ok := objval["sparkJob"]; ok {
						if rJobsSparkJob, ok := objval["sparkJob"].(map[string]interface{}); ok {
							rJobs.SparkJob = &dclService.WorkflowTemplateJobsSparkJob{}
							if _, ok := rJobsSparkJob["archiveUris"]; ok {
								if s, ok := rJobsSparkJob["archiveUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkJob.ArchiveUris = append(rJobs.SparkJob.ArchiveUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.ArchiveUris: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkJob["args"]; ok {
								if s, ok := rJobsSparkJob["args"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkJob.Args = append(rJobs.SparkJob.Args, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.Args: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkJob["fileUris"]; ok {
								if s, ok := rJobsSparkJob["fileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkJob.FileUris = append(rJobs.SparkJob.FileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.FileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkJob["jarFileUris"]; ok {
								if s, ok := rJobsSparkJob["jarFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkJob.JarFileUris = append(rJobs.SparkJob.JarFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.JarFileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkJob["loggingConfig"]; ok {
								if rJobsSparkJobLoggingConfig, ok := rJobsSparkJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.SparkJob.LoggingConfig = &dclService.WorkflowTemplateJobsSparkJobLoggingConfig{}
									if _, ok := rJobsSparkJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsSparkJobLoggingConfigDriverLogLevels, ok := rJobsSparkJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsSparkJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.SparkJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.SparkJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsSparkJob["mainClass"]; ok {
								if s, ok := rJobsSparkJob["mainClass"].(string); ok {
									rJobs.SparkJob.MainClass = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.MainClass: expected string")
								}
							}
							if _, ok := rJobsSparkJob["mainJarFileUri"]; ok {
								if s, ok := rJobsSparkJob["mainJarFileUri"].(string); ok {
									rJobs.SparkJob.MainJarFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.MainJarFileUri: expected string")
								}
							}
							if _, ok := rJobsSparkJob["properties"]; ok {
								if rJobsSparkJobProperties, ok := rJobsSparkJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsSparkJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.SparkJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.SparkJob.Properties: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.SparkJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["sparkRJob"]; ok {
						if rJobsSparkRJob, ok := objval["sparkRJob"].(map[string]interface{}); ok {
							rJobs.SparkRJob = &dclService.WorkflowTemplateJobsSparkRJob{}
							if _, ok := rJobsSparkRJob["archiveUris"]; ok {
								if s, ok := rJobsSparkRJob["archiveUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkRJob.ArchiveUris = append(rJobs.SparkRJob.ArchiveUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkRJob.ArchiveUris: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkRJob["args"]; ok {
								if s, ok := rJobsSparkRJob["args"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkRJob.Args = append(rJobs.SparkRJob.Args, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkRJob.Args: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkRJob["fileUris"]; ok {
								if s, ok := rJobsSparkRJob["fileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkRJob.FileUris = append(rJobs.SparkRJob.FileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkRJob.FileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkRJob["loggingConfig"]; ok {
								if rJobsSparkRJobLoggingConfig, ok := rJobsSparkRJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.SparkRJob.LoggingConfig = &dclService.WorkflowTemplateJobsSparkRJobLoggingConfig{}
									if _, ok := rJobsSparkRJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsSparkRJobLoggingConfigDriverLogLevels, ok := rJobsSparkRJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsSparkRJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.SparkRJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.SparkRJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkRJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsSparkRJob["mainRFileUri"]; ok {
								if s, ok := rJobsSparkRJob["mainRFileUri"].(string); ok {
									rJobs.SparkRJob.MainRFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.SparkRJob.MainRFileUri: expected string")
								}
							}
							if _, ok := rJobsSparkRJob["properties"]; ok {
								if rJobsSparkRJobProperties, ok := rJobsSparkRJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsSparkRJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.SparkRJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.SparkRJob.Properties: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.SparkRJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["sparkSqlJob"]; ok {
						if rJobsSparkSqlJob, ok := objval["sparkSqlJob"].(map[string]interface{}); ok {
							rJobs.SparkSqlJob = &dclService.WorkflowTemplateJobsSparkSqlJob{}
							if _, ok := rJobsSparkSqlJob["jarFileUris"]; ok {
								if s, ok := rJobsSparkSqlJob["jarFileUris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rJobs.SparkSqlJob.JarFileUris = append(rJobs.SparkSqlJob.JarFileUris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkSqlJob.JarFileUris: expected []interface{}")
								}
							}
							if _, ok := rJobsSparkSqlJob["loggingConfig"]; ok {
								if rJobsSparkSqlJobLoggingConfig, ok := rJobsSparkSqlJob["loggingConfig"].(map[string]interface{}); ok {
									rJobs.SparkSqlJob.LoggingConfig = &dclService.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
									if _, ok := rJobsSparkSqlJobLoggingConfig["driverLogLevels"]; ok {
										if rJobsSparkSqlJobLoggingConfigDriverLogLevels, ok := rJobsSparkSqlJobLoggingConfig["driverLogLevels"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rJobsSparkSqlJobLoggingConfigDriverLogLevels {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rJobs.SparkSqlJob.LoggingConfig.DriverLogLevels = m
										} else {
											return nil, fmt.Errorf("rJobs.SparkSqlJob.LoggingConfig.DriverLogLevels: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkSqlJob.LoggingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsSparkSqlJob["properties"]; ok {
								if rJobsSparkSqlJobProperties, ok := rJobsSparkSqlJob["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsSparkSqlJobProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.SparkSqlJob.Properties = m
								} else {
									return nil, fmt.Errorf("rJobs.SparkSqlJob.Properties: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsSparkSqlJob["queryFileUri"]; ok {
								if s, ok := rJobsSparkSqlJob["queryFileUri"].(string); ok {
									rJobs.SparkSqlJob.QueryFileUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rJobs.SparkSqlJob.QueryFileUri: expected string")
								}
							}
							if _, ok := rJobsSparkSqlJob["queryList"]; ok {
								if rJobsSparkSqlJobQueryList, ok := rJobsSparkSqlJob["queryList"].(map[string]interface{}); ok {
									rJobs.SparkSqlJob.QueryList = &dclService.WorkflowTemplateJobsSparkSqlJobQueryList{}
									if _, ok := rJobsSparkSqlJobQueryList["queries"]; ok {
										if s, ok := rJobsSparkSqlJobQueryList["queries"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rJobs.SparkSqlJob.QueryList.Queries = append(rJobs.SparkSqlJob.QueryList.Queries, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rJobs.SparkSqlJob.QueryList.Queries: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rJobs.SparkSqlJob.QueryList: expected map[string]interface{}")
								}
							}
							if _, ok := rJobsSparkSqlJob["scriptVariables"]; ok {
								if rJobsSparkSqlJobScriptVariables, ok := rJobsSparkSqlJob["scriptVariables"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rJobsSparkSqlJobScriptVariables {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rJobs.SparkSqlJob.ScriptVariables = m
								} else {
									return nil, fmt.Errorf("rJobs.SparkSqlJob.ScriptVariables: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rJobs.SparkSqlJob: expected map[string]interface{}")
						}
					}
					if _, ok := objval["stepId"]; ok {
						if s, ok := objval["stepId"].(string); ok {
							rJobs.StepId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rJobs.StepId: expected string")
						}
					}
					r.Jobs = append(r.Jobs, rJobs)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Jobs: expected []interface{}")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["parameters"]; ok {
		if s, ok := u.Object["parameters"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rParameters dclService.WorkflowTemplateParameters
					if _, ok := objval["description"]; ok {
						if s, ok := objval["description"].(string); ok {
							rParameters.Description = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rParameters.Description: expected string")
						}
					}
					if _, ok := objval["fields"]; ok {
						if s, ok := objval["fields"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rParameters.Fields = append(rParameters.Fields, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rParameters.Fields: expected []interface{}")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rParameters.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rParameters.Name: expected string")
						}
					}
					if _, ok := objval["validation"]; ok {
						if rParametersValidation, ok := objval["validation"].(map[string]interface{}); ok {
							rParameters.Validation = &dclService.WorkflowTemplateParametersValidation{}
							if _, ok := rParametersValidation["regex"]; ok {
								if rParametersValidationRegex, ok := rParametersValidation["regex"].(map[string]interface{}); ok {
									rParameters.Validation.Regex = &dclService.WorkflowTemplateParametersValidationRegex{}
									if _, ok := rParametersValidationRegex["regexes"]; ok {
										if s, ok := rParametersValidationRegex["regexes"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rParameters.Validation.Regex.Regexes = append(rParameters.Validation.Regex.Regexes, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rParameters.Validation.Regex.Regexes: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rParameters.Validation.Regex: expected map[string]interface{}")
								}
							}
							if _, ok := rParametersValidation["values"]; ok {
								if rParametersValidationValues, ok := rParametersValidation["values"].(map[string]interface{}); ok {
									rParameters.Validation.Values = &dclService.WorkflowTemplateParametersValidationValues{}
									if _, ok := rParametersValidationValues["values"]; ok {
										if s, ok := rParametersValidationValues["values"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rParameters.Validation.Values.Values = append(rParameters.Validation.Values.Values, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rParameters.Validation.Values.Values: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rParameters.Validation.Values: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rParameters.Validation: expected map[string]interface{}")
						}
					}
					r.Parameters = append(r.Parameters, rParameters)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Parameters: expected []interface{}")
		}
	}
	if _, ok := u.Object["placement"]; ok {
		if rPlacement, ok := u.Object["placement"].(map[string]interface{}); ok {
			r.Placement = &dclService.WorkflowTemplatePlacement{}
			if _, ok := rPlacement["clusterSelector"]; ok {
				if rPlacementClusterSelector, ok := rPlacement["clusterSelector"].(map[string]interface{}); ok {
					r.Placement.ClusterSelector = &dclService.WorkflowTemplatePlacementClusterSelector{}
					if _, ok := rPlacementClusterSelector["clusterLabels"]; ok {
						if rPlacementClusterSelectorClusterLabels, ok := rPlacementClusterSelector["clusterLabels"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rPlacementClusterSelectorClusterLabels {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							r.Placement.ClusterSelector.ClusterLabels = m
						} else {
							return nil, fmt.Errorf("r.Placement.ClusterSelector.ClusterLabels: expected map[string]interface{}")
						}
					}
					if _, ok := rPlacementClusterSelector["zone"]; ok {
						if s, ok := rPlacementClusterSelector["zone"].(string); ok {
							r.Placement.ClusterSelector.Zone = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Placement.ClusterSelector.Zone: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Placement.ClusterSelector: expected map[string]interface{}")
				}
			}
			if _, ok := rPlacement["managedCluster"]; ok {
				if rPlacementManagedCluster, ok := rPlacement["managedCluster"].(map[string]interface{}); ok {
					r.Placement.ManagedCluster = &dclService.WorkflowTemplatePlacementManagedCluster{}
					if _, ok := rPlacementManagedCluster["clusterName"]; ok {
						if s, ok := rPlacementManagedCluster["clusterName"].(string); ok {
							r.Placement.ManagedCluster.ClusterName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Placement.ManagedCluster.ClusterName: expected string")
						}
					}
					if _, ok := rPlacementManagedCluster["config"]; ok {
						if rPlacementManagedClusterConfig, ok := rPlacementManagedCluster["config"].(map[string]interface{}); ok {
							r.Placement.ManagedCluster.Config = &dclService.WorkflowTemplatePlacementManagedClusterConfig{}
							if _, ok := rPlacementManagedClusterConfig["autoscalingConfig"]; ok {
								if rPlacementManagedClusterConfigAutoscalingConfig, ok := rPlacementManagedClusterConfig["autoscalingConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.AutoscalingConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{}
									if _, ok := rPlacementManagedClusterConfigAutoscalingConfig["policy"]; ok {
										if s, ok := rPlacementManagedClusterConfigAutoscalingConfig["policy"].(string); ok {
											r.Placement.ManagedCluster.Config.AutoscalingConfig.Policy = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.AutoscalingConfig.Policy: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.AutoscalingConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["encryptionConfig"]; ok {
								if rPlacementManagedClusterConfigEncryptionConfig, ok := rPlacementManagedClusterConfig["encryptionConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.EncryptionConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{}
									if _, ok := rPlacementManagedClusterConfigEncryptionConfig["gcePdKmsKeyName"]; ok {
										if s, ok := rPlacementManagedClusterConfigEncryptionConfig["gcePdKmsKeyName"].(string); ok {
											r.Placement.ManagedCluster.Config.EncryptionConfig.GcePdKmsKeyName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.EncryptionConfig.GcePdKmsKeyName: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.EncryptionConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["endpointConfig"]; ok {
								if rPlacementManagedClusterConfigEndpointConfig, ok := rPlacementManagedClusterConfig["endpointConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.EndpointConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig{}
									if _, ok := rPlacementManagedClusterConfigEndpointConfig["enableHttpPortAccess"]; ok {
										if b, ok := rPlacementManagedClusterConfigEndpointConfig["enableHttpPortAccess"].(bool); ok {
											r.Placement.ManagedCluster.Config.EndpointConfig.EnableHttpPortAccess = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.EndpointConfig.EnableHttpPortAccess: expected bool")
										}
									}
									if _, ok := rPlacementManagedClusterConfigEndpointConfig["httpPorts"]; ok {
										if rPlacementManagedClusterConfigEndpointConfigHttpPorts, ok := rPlacementManagedClusterConfigEndpointConfig["httpPorts"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rPlacementManagedClusterConfigEndpointConfigHttpPorts {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											r.Placement.ManagedCluster.Config.EndpointConfig.HttpPorts = m
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.EndpointConfig.HttpPorts: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.EndpointConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["gceClusterConfig"]; ok {
								if rPlacementManagedClusterConfigGceClusterConfig, ok := rPlacementManagedClusterConfig["gceClusterConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.GceClusterConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["internalIPOnly"]; ok {
										if b, ok := rPlacementManagedClusterConfigGceClusterConfig["internalIPOnly"].(bool); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.InternalIPOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.InternalIPOnly: expected bool")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["metadata"]; ok {
										if rPlacementManagedClusterConfigGceClusterConfigMetadata, ok := rPlacementManagedClusterConfigGceClusterConfig["metadata"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rPlacementManagedClusterConfigGceClusterConfigMetadata {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											r.Placement.ManagedCluster.Config.GceClusterConfig.Metadata = m
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.Metadata: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["network"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["network"].(string); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.Network = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.Network: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["nodeGroupAffinity"]; ok {
										if rPlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity, ok := rPlacementManagedClusterConfigGceClusterConfig["nodeGroupAffinity"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity = &dclService.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity["nodeGroup"]; ok {
												if s, ok := rPlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity["nodeGroup"].(string); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.NodeGroupAffinity: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["privateIPv6GoogleAccess"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["privateIPv6GoogleAccess"].(string); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.PrivateIPv6GoogleAccess = dclService.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.PrivateIPv6GoogleAccess: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["reservationAffinity"]; ok {
										if rPlacementManagedClusterConfigGceClusterConfigReservationAffinity, ok := rPlacementManagedClusterConfigGceClusterConfig["reservationAffinity"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity = &dclService.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["consumeReservationType"]; ok {
												if s, ok := rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["consumeReservationType"].(string); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType = dclService.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["key"]; ok {
												if s, ok := rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["key"].(string); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Key = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Key: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["values"]; ok {
												if s, ok := rPlacementManagedClusterConfigGceClusterConfigReservationAffinity["values"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Values = append(r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Values, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity.Values: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ReservationAffinity: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["serviceAccount"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["serviceAccount"].(string); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccount = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccount: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["serviceAccountScopes"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["serviceAccountScopes"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccountScopes = append(r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccountScopes, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ServiceAccountScopes: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["shieldedInstanceConfig"]; ok {
										if rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig, ok := rPlacementManagedClusterConfigGceClusterConfig["shieldedInstanceConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableIntegrityMonitoring"]; ok {
												if b, ok := rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableIntegrityMonitoring"].(bool); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring: expected bool")
												}
											}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableSecureBoot"]; ok {
												if b, ok := rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableSecureBoot"].(bool); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot: expected bool")
												}
											}
											if _, ok := rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableVtpm"]; ok {
												if b, ok := rPlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig["enableVtpm"].(bool); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm: expected bool")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.ShieldedInstanceConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["subnetwork"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["subnetwork"].(string); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.Subnetwork = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.Subnetwork: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["tags"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["tags"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Placement.ManagedCluster.Config.GceClusterConfig.Tags = append(r.Placement.ManagedCluster.Config.GceClusterConfig.Tags, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.Tags: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigGceClusterConfig["zone"]; ok {
										if s, ok := rPlacementManagedClusterConfigGceClusterConfig["zone"].(string); ok {
											r.Placement.ManagedCluster.Config.GceClusterConfig.Zone = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig.Zone: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GceClusterConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["gkeClusterConfig"]; ok {
								if rPlacementManagedClusterConfigGkeClusterConfig, ok := rPlacementManagedClusterConfig["gkeClusterConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.GkeClusterConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig{}
									if _, ok := rPlacementManagedClusterConfigGkeClusterConfig["namespacedGkeDeploymentTarget"]; ok {
										if rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget, ok := rPlacementManagedClusterConfigGkeClusterConfig["namespacedGkeDeploymentTarget"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget = &dclService.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
											if _, ok := rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"]; ok {
												if s, ok := rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"].(string); ok {
													r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"]; ok {
												if s, ok := rPlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"].(string); ok {
													r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.GkeClusterConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["initializationActions"]; ok {
								if s, ok := rPlacementManagedClusterConfig["initializationActions"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rPlacementManagedClusterConfigInitializationActions dclService.WorkflowTemplatePlacementManagedClusterConfigInitializationActions
											if _, ok := objval["executableFile"]; ok {
												if s, ok := objval["executableFile"].(string); ok {
													rPlacementManagedClusterConfigInitializationActions.ExecutableFile = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rPlacementManagedClusterConfigInitializationActions.ExecutableFile: expected string")
												}
											}
											if _, ok := objval["executionTimeout"]; ok {
												if s, ok := objval["executionTimeout"].(string); ok {
													rPlacementManagedClusterConfigInitializationActions.ExecutionTimeout = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rPlacementManagedClusterConfigInitializationActions.ExecutionTimeout: expected string")
												}
											}
											r.Placement.ManagedCluster.Config.InitializationActions = append(r.Placement.ManagedCluster.Config.InitializationActions, rPlacementManagedClusterConfigInitializationActions)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.InitializationActions: expected []interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["lifecycleConfig"]; ok {
								if rPlacementManagedClusterConfigLifecycleConfig, ok := rPlacementManagedClusterConfig["lifecycleConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.LifecycleConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{}
									if _, ok := rPlacementManagedClusterConfigLifecycleConfig["autoDeleteTime"]; ok {
										if s, ok := rPlacementManagedClusterConfigLifecycleConfig["autoDeleteTime"].(string); ok {
											r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTime = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTime: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigLifecycleConfig["autoDeleteTtl"]; ok {
										if s, ok := rPlacementManagedClusterConfigLifecycleConfig["autoDeleteTtl"].(string); ok {
											r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTtl = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.LifecycleConfig.AutoDeleteTtl: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigLifecycleConfig["idleDeleteTtl"]; ok {
										if s, ok := rPlacementManagedClusterConfigLifecycleConfig["idleDeleteTtl"].(string); ok {
											r.Placement.ManagedCluster.Config.LifecycleConfig.IdleDeleteTtl = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.LifecycleConfig.IdleDeleteTtl: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigLifecycleConfig["idleStartTime"]; ok {
										if s, ok := rPlacementManagedClusterConfigLifecycleConfig["idleStartTime"].(string); ok {
											r.Placement.ManagedCluster.Config.LifecycleConfig.IdleStartTime = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.LifecycleConfig.IdleStartTime: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.LifecycleConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["masterConfig"]; ok {
								if rPlacementManagedClusterConfigMasterConfig, ok := rPlacementManagedClusterConfig["masterConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.MasterConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigMasterConfig{}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["accelerators"]; ok {
										if s, ok := rPlacementManagedClusterConfigMasterConfig["accelerators"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rPlacementManagedClusterConfigMasterConfigAccelerators dclService.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators
													if _, ok := objval["acceleratorCount"]; ok {
														if i, ok := objval["acceleratorCount"].(int64); ok {
															rPlacementManagedClusterConfigMasterConfigAccelerators.AcceleratorCount = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rPlacementManagedClusterConfigMasterConfigAccelerators.AcceleratorCount: expected int64")
														}
													}
													if _, ok := objval["acceleratorType"]; ok {
														if s, ok := objval["acceleratorType"].(string); ok {
															rPlacementManagedClusterConfigMasterConfigAccelerators.AcceleratorType = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rPlacementManagedClusterConfigMasterConfigAccelerators.AcceleratorType: expected string")
														}
													}
													r.Placement.ManagedCluster.Config.MasterConfig.Accelerators = append(r.Placement.ManagedCluster.Config.MasterConfig.Accelerators, rPlacementManagedClusterConfigMasterConfigAccelerators)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.Accelerators: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["diskConfig"]; ok {
										if rPlacementManagedClusterConfigMasterConfigDiskConfig, ok := rPlacementManagedClusterConfigMasterConfig["diskConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{}
											if _, ok := rPlacementManagedClusterConfigMasterConfigDiskConfig["bootDiskSizeGb"]; ok {
												if i, ok := rPlacementManagedClusterConfigMasterConfigDiskConfig["bootDiskSizeGb"].(int64); ok {
													r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskSizeGb: expected int64")
												}
											}
											if _, ok := rPlacementManagedClusterConfigMasterConfigDiskConfig["bootDiskType"]; ok {
												if s, ok := rPlacementManagedClusterConfigMasterConfigDiskConfig["bootDiskType"].(string); ok {
													r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskType = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.BootDiskType: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigMasterConfigDiskConfig["numLocalSsds"]; ok {
												if i, ok := rPlacementManagedClusterConfigMasterConfigDiskConfig["numLocalSsds"].(int64); ok {
													r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.NumLocalSsds = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig.NumLocalSsds: expected int64")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.DiskConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["image"]; ok {
										if s, ok := rPlacementManagedClusterConfigMasterConfig["image"].(string); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.Image = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.Image: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["instanceNames"]; ok {
										if s, ok := rPlacementManagedClusterConfigMasterConfig["instanceNames"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Placement.ManagedCluster.Config.MasterConfig.InstanceNames = append(r.Placement.ManagedCluster.Config.MasterConfig.InstanceNames, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.InstanceNames: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["isPreemptible"]; ok {
										if b, ok := rPlacementManagedClusterConfigMasterConfig["isPreemptible"].(bool); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.IsPreemptible = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.IsPreemptible: expected bool")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["machineType"]; ok {
										if s, ok := rPlacementManagedClusterConfigMasterConfig["machineType"].(string); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.MachineType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.MachineType: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["managedGroupConfig"]; ok {
										if rPlacementManagedClusterConfigMasterConfigManagedGroupConfig, ok := rPlacementManagedClusterConfigMasterConfig["managedGroupConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{}
											if _, ok := rPlacementManagedClusterConfigMasterConfigManagedGroupConfig["instanceGroupManagerName"]; ok {
												if s, ok := rPlacementManagedClusterConfigMasterConfigManagedGroupConfig["instanceGroupManagerName"].(string); ok {
													r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigMasterConfigManagedGroupConfig["instanceTemplateName"]; ok {
												if s, ok := rPlacementManagedClusterConfigMasterConfigManagedGroupConfig["instanceTemplateName"].(string); ok {
													r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.ManagedGroupConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["minCpuPlatform"]; ok {
										if s, ok := rPlacementManagedClusterConfigMasterConfig["minCpuPlatform"].(string); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.MinCpuPlatform = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.MinCpuPlatform: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["numInstances"]; ok {
										if i, ok := rPlacementManagedClusterConfigMasterConfig["numInstances"].(int64); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.NumInstances = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.NumInstances: expected int64")
										}
									}
									if _, ok := rPlacementManagedClusterConfigMasterConfig["preemptibility"]; ok {
										if s, ok := rPlacementManagedClusterConfigMasterConfig["preemptibility"].(string); ok {
											r.Placement.ManagedCluster.Config.MasterConfig.Preemptibility = dclService.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig.Preemptibility: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MasterConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["metastoreConfig"]; ok {
								if rPlacementManagedClusterConfigMetastoreConfig, ok := rPlacementManagedClusterConfig["metastoreConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.MetastoreConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig{}
									if _, ok := rPlacementManagedClusterConfigMetastoreConfig["dataprocMetastoreService"]; ok {
										if s, ok := rPlacementManagedClusterConfigMetastoreConfig["dataprocMetastoreService"].(string); ok {
											r.Placement.ManagedCluster.Config.MetastoreConfig.DataprocMetastoreService = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MetastoreConfig.DataprocMetastoreService: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.MetastoreConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["secondaryWorkerConfig"]; ok {
								if rPlacementManagedClusterConfigSecondaryWorkerConfig, ok := rPlacementManagedClusterConfig["secondaryWorkerConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.SecondaryWorkerConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["accelerators"]; ok {
										if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["accelerators"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators dclService.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators
													if _, ok := objval["acceleratorCount"]; ok {
														if i, ok := objval["acceleratorCount"].(int64); ok {
															rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators.AcceleratorCount = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators.AcceleratorCount: expected int64")
														}
													}
													if _, ok := objval["acceleratorType"]; ok {
														if s, ok := objval["acceleratorType"].(string); ok {
															rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators.AcceleratorType = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators.AcceleratorType: expected string")
														}
													}
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Accelerators = append(r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Accelerators, rPlacementManagedClusterConfigSecondaryWorkerConfigAccelerators)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Accelerators: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["diskConfig"]; ok {
										if rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["diskConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{}
											if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["bootDiskSizeGb"]; ok {
												if i, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["bootDiskSizeGb"].(int64); ok {
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb: expected int64")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["bootDiskType"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["bootDiskType"].(string); ok {
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["numLocalSsds"]; ok {
												if i, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig["numLocalSsds"].(int64); ok {
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds: expected int64")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.DiskConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["image"]; ok {
										if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["image"].(string); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Image = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Image: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["instanceNames"]; ok {
										if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["instanceNames"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.InstanceNames = append(r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.InstanceNames, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.InstanceNames: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["isPreemptible"]; ok {
										if b, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["isPreemptible"].(bool); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.IsPreemptible = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.IsPreemptible: expected bool")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["machineType"]; ok {
										if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["machineType"].(string); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MachineType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MachineType: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["managedGroupConfig"]; ok {
										if rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["managedGroupConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
											if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig["instanceGroupManagerName"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig["instanceGroupManagerName"].(string); ok {
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig["instanceTemplateName"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig["instanceTemplateName"].(string); ok {
													r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.ManagedGroupConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["minCpuPlatform"]; ok {
										if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["minCpuPlatform"].(string); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MinCpuPlatform = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.MinCpuPlatform: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["numInstances"]; ok {
										if i, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["numInstances"].(int64); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.NumInstances = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.NumInstances: expected int64")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["preemptibility"]; ok {
										if s, ok := rPlacementManagedClusterConfigSecondaryWorkerConfig["preemptibility"].(string); ok {
											r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Preemptibility = dclService.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig.Preemptibility: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecondaryWorkerConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["securityConfig"]; ok {
								if rPlacementManagedClusterConfigSecurityConfig, ok := rPlacementManagedClusterConfig["securityConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.SecurityConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig{}
									if _, ok := rPlacementManagedClusterConfigSecurityConfig["kerberosConfig"]; ok {
										if rPlacementManagedClusterConfigSecurityConfigKerberosConfig, ok := rPlacementManagedClusterConfigSecurityConfig["kerberosConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustAdminServer"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustAdminServer"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustKdc"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustKdc"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustRealm"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustRealm"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["enableKerberos"]; ok {
												if b, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["enableKerberos"].(bool); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.EnableKerberos = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.EnableKerberos: expected bool")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["kdcDbKey"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["kdcDbKey"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KdcDbKey = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KdcDbKey: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keyPassword"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keyPassword"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeyPassword = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeyPassword: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keystore"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keystore"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Keystore = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Keystore: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keystorePassword"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["keystorePassword"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeystorePassword = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KeystorePassword: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["kmsKey"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["kmsKey"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KmsKey = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.KmsKey: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["realm"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["realm"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Realm = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Realm: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["rootPrincipalPassword"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["rootPrincipalPassword"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["tgtLifetimeHours"]; ok {
												if i, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["tgtLifetimeHours"].(int64); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours: expected int64")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["truststore"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["truststore"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Truststore = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.Truststore: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["truststorePassword"]; ok {
												if s, ok := rPlacementManagedClusterConfigSecurityConfigKerberosConfig["truststorePassword"].(string); ok {
													r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TruststorePassword = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig.TruststorePassword: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig.KerberosConfig: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SecurityConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["softwareConfig"]; ok {
								if rPlacementManagedClusterConfigSoftwareConfig, ok := rPlacementManagedClusterConfig["softwareConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.SoftwareConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{}
									if _, ok := rPlacementManagedClusterConfigSoftwareConfig["imageVersion"]; ok {
										if s, ok := rPlacementManagedClusterConfigSoftwareConfig["imageVersion"].(string); ok {
											r.Placement.ManagedCluster.Config.SoftwareConfig.ImageVersion = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SoftwareConfig.ImageVersion: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSoftwareConfig["optionalComponents"]; ok {
										if s, ok := rPlacementManagedClusterConfigSoftwareConfig["optionalComponents"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Placement.ManagedCluster.Config.SoftwareConfig.OptionalComponents = append(r.Placement.ManagedCluster.Config.SoftwareConfig.OptionalComponents, dclService.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(strval))
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SoftwareConfig.OptionalComponents: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigSoftwareConfig["properties"]; ok {
										if rPlacementManagedClusterConfigSoftwareConfigProperties, ok := rPlacementManagedClusterConfigSoftwareConfig["properties"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rPlacementManagedClusterConfigSoftwareConfigProperties {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											r.Placement.ManagedCluster.Config.SoftwareConfig.Properties = m
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SoftwareConfig.Properties: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.SoftwareConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["stagingBucket"]; ok {
								if s, ok := rPlacementManagedClusterConfig["stagingBucket"].(string); ok {
									r.Placement.ManagedCluster.Config.StagingBucket = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.StagingBucket: expected string")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["tempBucket"]; ok {
								if s, ok := rPlacementManagedClusterConfig["tempBucket"].(string); ok {
									r.Placement.ManagedCluster.Config.TempBucket = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.TempBucket: expected string")
								}
							}
							if _, ok := rPlacementManagedClusterConfig["workerConfig"]; ok {
								if rPlacementManagedClusterConfigWorkerConfig, ok := rPlacementManagedClusterConfig["workerConfig"].(map[string]interface{}); ok {
									r.Placement.ManagedCluster.Config.WorkerConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig{}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["accelerators"]; ok {
										if s, ok := rPlacementManagedClusterConfigWorkerConfig["accelerators"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rPlacementManagedClusterConfigWorkerConfigAccelerators dclService.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators
													if _, ok := objval["acceleratorCount"]; ok {
														if i, ok := objval["acceleratorCount"].(int64); ok {
															rPlacementManagedClusterConfigWorkerConfigAccelerators.AcceleratorCount = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rPlacementManagedClusterConfigWorkerConfigAccelerators.AcceleratorCount: expected int64")
														}
													}
													if _, ok := objval["acceleratorType"]; ok {
														if s, ok := objval["acceleratorType"].(string); ok {
															rPlacementManagedClusterConfigWorkerConfigAccelerators.AcceleratorType = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rPlacementManagedClusterConfigWorkerConfigAccelerators.AcceleratorType: expected string")
														}
													}
													r.Placement.ManagedCluster.Config.WorkerConfig.Accelerators = append(r.Placement.ManagedCluster.Config.WorkerConfig.Accelerators, rPlacementManagedClusterConfigWorkerConfigAccelerators)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.Accelerators: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["diskConfig"]; ok {
										if rPlacementManagedClusterConfigWorkerConfigDiskConfig, ok := rPlacementManagedClusterConfigWorkerConfig["diskConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{}
											if _, ok := rPlacementManagedClusterConfigWorkerConfigDiskConfig["bootDiskSizeGb"]; ok {
												if i, ok := rPlacementManagedClusterConfigWorkerConfigDiskConfig["bootDiskSizeGb"].(int64); ok {
													r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskSizeGb: expected int64")
												}
											}
											if _, ok := rPlacementManagedClusterConfigWorkerConfigDiskConfig["bootDiskType"]; ok {
												if s, ok := rPlacementManagedClusterConfigWorkerConfigDiskConfig["bootDiskType"].(string); ok {
													r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskType = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.BootDiskType: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigWorkerConfigDiskConfig["numLocalSsds"]; ok {
												if i, ok := rPlacementManagedClusterConfigWorkerConfigDiskConfig["numLocalSsds"].(int64); ok {
													r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.NumLocalSsds = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig.NumLocalSsds: expected int64")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.DiskConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["image"]; ok {
										if s, ok := rPlacementManagedClusterConfigWorkerConfig["image"].(string); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.Image = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.Image: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["instanceNames"]; ok {
										if s, ok := rPlacementManagedClusterConfigWorkerConfig["instanceNames"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.Placement.ManagedCluster.Config.WorkerConfig.InstanceNames = append(r.Placement.ManagedCluster.Config.WorkerConfig.InstanceNames, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.InstanceNames: expected []interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["isPreemptible"]; ok {
										if b, ok := rPlacementManagedClusterConfigWorkerConfig["isPreemptible"].(bool); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.IsPreemptible = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.IsPreemptible: expected bool")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["machineType"]; ok {
										if s, ok := rPlacementManagedClusterConfigWorkerConfig["machineType"].(string); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.MachineType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.MachineType: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["managedGroupConfig"]; ok {
										if rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig, ok := rPlacementManagedClusterConfigWorkerConfig["managedGroupConfig"].(map[string]interface{}); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig = &dclService.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{}
											if _, ok := rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig["instanceGroupManagerName"]; ok {
												if s, ok := rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig["instanceGroupManagerName"].(string); ok {
													r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName: expected string")
												}
											}
											if _, ok := rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig["instanceTemplateName"]; ok {
												if s, ok := rPlacementManagedClusterConfigWorkerConfigManagedGroupConfig["instanceTemplateName"].(string); ok {
													r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.ManagedGroupConfig: expected map[string]interface{}")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["minCpuPlatform"]; ok {
										if s, ok := rPlacementManagedClusterConfigWorkerConfig["minCpuPlatform"].(string); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.MinCpuPlatform = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.MinCpuPlatform: expected string")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["numInstances"]; ok {
										if i, ok := rPlacementManagedClusterConfigWorkerConfig["numInstances"].(int64); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.NumInstances = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.NumInstances: expected int64")
										}
									}
									if _, ok := rPlacementManagedClusterConfigWorkerConfig["preemptibility"]; ok {
										if s, ok := rPlacementManagedClusterConfigWorkerConfig["preemptibility"].(string); ok {
											r.Placement.ManagedCluster.Config.WorkerConfig.Preemptibility = dclService.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig.Preemptibility: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Placement.ManagedCluster.Config.WorkerConfig: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Placement.ManagedCluster.Config: expected map[string]interface{}")
						}
					}
					if _, ok := rPlacementManagedCluster["labels"]; ok {
						if rPlacementManagedClusterLabels, ok := rPlacementManagedCluster["labels"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rPlacementManagedClusterLabels {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							r.Placement.ManagedCluster.Labels = m
						} else {
							return nil, fmt.Errorf("r.Placement.ManagedCluster.Labels: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Placement.ManagedCluster: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Placement: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["version"]; ok {
		if i, ok := u.Object["version"].(int64); ok {
			r.Version = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Version: expected int64")
		}
	}
	return r, nil
}

func GetWorkflowTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkflowTemplate(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWorkflowTemplate(ctx, r)
	if err != nil {
		return nil, err
	}
	return WorkflowTemplateToUnstructured(r), nil
}

func ListWorkflowTemplate(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWorkflowTemplate(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WorkflowTemplateToUnstructured(r))
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

func ApplyWorkflowTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkflowTemplate(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkflowTemplate(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWorkflowTemplate(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WorkflowTemplateToUnstructured(r), nil
}

func WorkflowTemplateHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkflowTemplate(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkflowTemplate(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWorkflowTemplate(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWorkflowTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkflowTemplate(u)
	if err != nil {
		return err
	}
	return c.DeleteWorkflowTemplate(ctx, r)
}

func WorkflowTemplateID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWorkflowTemplate(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *WorkflowTemplate) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataproc",
		"WorkflowTemplate",
		"alpha",
	}
}

func (r *WorkflowTemplate) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkflowTemplate) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkflowTemplate) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *WorkflowTemplate) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkflowTemplate) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkflowTemplate) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkflowTemplate) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWorkflowTemplate(ctx, config, resource)
}

func (r *WorkflowTemplate) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWorkflowTemplate(ctx, config, resource, opts...)
}

func (r *WorkflowTemplate) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WorkflowTemplateHasDiff(ctx, config, resource, opts...)
}

func (r *WorkflowTemplate) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWorkflowTemplate(ctx, config, resource)
}

func (r *WorkflowTemplate) ID(resource *unstructured.Resource) (string, error) {
	return WorkflowTemplateID(resource)
}

func init() {
	unstructured.Register(&WorkflowTemplate{})
}
