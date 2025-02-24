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
#include "connector/connector.h"

#include <cstdlib>
#include <string>

#include "connector/connector.cgo.h"
#include "absl/strings/string_view.h"

void Connector::Initialize(absl::string_view request, std::string *response) {
  GoSlice go_request = {
    const_cast<char*>(request.data()),
    static_cast<GoInt>(request.size()),
    static_cast<GoInt>(request.size()),
  };
  auto go_response = ::Initialize(go_request);
  if (go_response.r0) {
    *response = std::string(reinterpret_cast<char*>(go_response.r0),
                            go_response.r1);
    std::free(go_response.r0);
  } else {
    *response = "";
  }
}

void Connector::Call(absl::string_view request, std::string *response) {
  GoSlice go_request = {
    const_cast<char*>(request.data()),
    static_cast<GoInt>(request.size()),
    static_cast<GoInt>(request.size()),
  };
  auto go_response = ::Call(go_request);
  if (go_response.r0) {
    *response = std::string(reinterpret_cast<char*>(go_response.r0),
                            go_response.r1);
    std::free(go_response.r0);
  } else {
    *response = "";
  }
}
