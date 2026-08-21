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
#include "absl/strings/string_view.h"
#include "third_party/pybind11/include/pybind11/detail/common.h"
#include "include/pybind11/pybind11.h"
#include "third_party/pybind11/include/pybind11/pytypes.h"

PYBIND11_MODULE(connector_ext, m) {
  namespace py = pybind11;

  py::class_<Connector>(m, "Connector")
      .def_static("Initialize",
                  [](std::string request) {
                    auto response = std::string();
                    Connector::Initialize(request, &response);
                    return py::bytes(response);
                  })
      .def_static("Call", [](std::string request) {
        auto response = std::string();
        Connector::Call(request, &response);
        return py::bytes(response);
      });
}
