load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "sdk_proto",
    srcs = ["sdk.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "sdk_go_proto",
    importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sdk_go_proto",
    proto = ":sdk_proto",
    visibility = ["//visibility:public"],
)

proto_library(
    name = "connector_proto",
    srcs = ["connector.proto"],
    visibility = ["//visibility:public"],
    deps = ["@com_google_googleapis//google/rpc:status_proto"],
)

go_proto_library(
    name = "connector_go_proto",
    importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/connector_go_proto",
    proto = ":connector_proto",
    visibility = ["//visibility:public"],
    deps = ["@go_googleapis//google/rpc:status_go_proto"],
)
