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

package contexts

func init() {
	resourceContextMap["batchdataflowjob"] = ResourceContext{
		ResourceKind: "DataflowJob",
		// Batch jobs don't support updates unlike streaming jobs
		SkipUpdate: true,
	}
	resourceContextMap["batchdataflowflextemplatejob"] = ResourceContext{
		ResourceKind: "DataflowFlexTemplateJob",
		// Flex Template Jobs don't support updates
		SkipUpdate: true,
	}
	resourceContextMap["streamingdataflowflextemplatejob"] = ResourceContext{
		ResourceKind: "DataflowFlexTemplateJob",
		// Flex Template Jobs don't support updates
		SkipUpdate: true,
	}

	resourceContextMap["streamingdataflowjobupdateparameters"] = ResourceContext{
		ResourceKind: "DataflowJob",
		// The streamingdataflowjob has been flaky, which seems to be attributed
		// to the dataflow jobs themselves. see b/166669646
		SkipUpdate: true,
	}

	resourceContextMap["streamingdataflowjobupdatetemplate"] = ResourceContext{
		ResourceKind: "DataflowJob",
		// The streamingdataflowjob has been flaky, which seems to be attributed
		// to the dataflow jobs themselves. see b/166669646
		SkipUpdate: true,
	}
}
