# Copyright 2024 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "empty_proto",
    srcs = ["empty.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "empty_go_proto",
    importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto",
    proto = ":empty_proto",
    visibility = ["//visibility:public"],
)

load("//:connector_rules.bzl", "python_protos")

python_protos(packages = ["ud/graphite/mmv2/connector","apikeys","apikeys/beta","apikeys/alpha","apigee","apigee/beta","apigee/alpha","assuredworkloads","assuredworkloads/beta","assuredworkloads/alpha","billingbudgets","billingbudgets/beta","billingbudgets/alpha","bigquery","bigquery/beta","bigquery/alpha","bigqueryreservation","bigqueryreservation/beta","bigqueryreservation/alpha","binaryauthorization","binaryauthorization/beta","binaryauthorization/alpha","cloudbuild","cloudbuild/beta","cloudbuild/alpha","cloudbuildv2","cloudbuildv2/beta","cloudbuildv2/alpha","clouddeploy","clouddeploy/beta","clouddeploy/alpha","cloudfunctions","cloudfunctions/beta","cloudfunctions/alpha","cloudidentity","cloudidentity/beta","cloudidentity/alpha","cloudkms","cloudkms/beta","cloudkms/alpha","cloudresourcemanager","cloudresourcemanager/beta","cloudresourcemanager/alpha","cloudscheduler","cloudscheduler/beta","cloudscheduler/alpha","compute","compute/beta","compute/alpha","configcontroller/alpha","containeranalysis","containeranalysis/beta","containeranalysis/alpha","containeraws","containeraws/beta","containeraws/alpha","containerazure","containerazure/beta","containerazure/alpha","datafusion/alpha","datafusion/beta","dataplex","dataplex/beta","dataplex/alpha","dataproc","dataproc/beta","dataproc/alpha","dlp","dlp/beta","dlp/alpha","eventarc","eventarc/beta","eventarc/alpha","filestore","filestore/beta","filestore/alpha","firebase/alpha","firebase/beta","firebaserules","firebaserules/beta","firebaserules/alpha","gkehub","gkehub/beta","gkehub/alpha","iam","iam/beta","iam/alpha","iap","iap/beta","iap/alpha","identitytoolkit","identitytoolkit/beta","identitytoolkit/alpha","logging","logging/beta","logging/alpha","monitoring","monitoring/beta","monitoring/alpha","networkconnectivity","networkconnectivity/beta","networkconnectivity/alpha","networksecurity/alpha","networksecurity/beta","networkservices","networkservices/beta","networkservices/alpha","orgpolicy","orgpolicy/beta","orgpolicy/alpha","osconfig","osconfig/beta","osconfig/alpha","pubsub","pubsub/beta","pubsub/alpha","run/alpha","storage","storage/beta","storage/alpha","privateca","privateca/beta","privateca/alpha","vpcaccess","vpcaccess/beta","vpcaccess/alpha","recaptchaenterprise","recaptchaenterprise/beta","recaptchaenterprise/alpha"])
