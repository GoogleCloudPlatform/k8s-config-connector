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
#ifndef CLOUD_GRAPHITE_MMV2_CONNECTOR_CONNECTOR_H_
#define CLOUD_GRAPHITE_MMV2_CONNECTOR_CONNECTOR_H_

#include <string>

#include "absl/strings/string_view.h"

class Connector {
 public:
  static void Initialize(absl::string_view request, std::string *response);
  static void Call(absl::string_view request, std::string *response);
};

#endif  // CLOUD_GRAPHITE_MMV2_CONNECTOR_CONNECTOR_H_
