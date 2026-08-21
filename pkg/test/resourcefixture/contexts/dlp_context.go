// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the"License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an"AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contexts

func init() {
	resourceContextMap["bigqueryfieldstoredinfotype"] = ResourceContext{
		ResourceKind: "DLPStoredInfoType",
		// There is no update method for this resource.
		SkipUpdate: true,
	}
	resourceContextMap["cloudstoragefilesetstoredinfotype"] = ResourceContext{
		ResourceKind: "DLPStoredInfoType",
		// There is no update method for this resource.
		SkipUpdate: true,
	}
	resourceContextMap["cloudstoragepathstoredinfotype"] = ResourceContext{
		ResourceKind: "DLPStoredInfoType",
		// There is no update method for this resource.
		SkipUpdate: true,
	}
	resourceContextMap["regexstoredinfotype"] = ResourceContext{
		ResourceKind: "DLPStoredInfoType",
		// There is no update method for this resource.
		SkipUpdate: true,
	}
	resourceContextMap["wordliststoredinfotype"] = ResourceContext{
		ResourceKind: "DLPStoredInfoType",
		// There is no update method for this resource.
		SkipUpdate: true,
	}
}
