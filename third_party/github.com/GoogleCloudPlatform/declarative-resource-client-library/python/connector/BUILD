load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "connector_go",
    srcs = [
        "connector.go",
        "server.go",
    ],
    cgo = True,
    importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/connector",
    linkmode = "c-archive",
    deps = [
        "//proto:empty_go_proto",
        "//proto/connector:connector_go_proto",
        "//services/redis:redis_connector",
        "@com_github_golang_glog//:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_genproto//googleapis/rpc/status:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
    ],
)

cc_import(
    name = "connector_go_lib",
    # This file isn't generated, it's copied in manually, but it should
    # ~never change.
    hdrs = ["connector.cgo.h"],
    static_library = ":connector_go",
    alwayslink = 1,
)

cc_library(
    name = "connector_lib",
    srcs = [
        "connector.cc",
    ],
    hdrs = ["connector.h"],
    deps = [
        ":connector_go_lib",
        "@absl_cpp//absl/strings",
    ],
    alwayslink = 1,
)

load("//:py_extension.bzl", "py_extension")

py_extension(
    name = "connector_ext",
    srcs = ["connector_ext.cc"],
    deps = [
        ":connector_lib",
        "@absl_cpp//absl/strings",
        "@pybind11",
        "@python//:python-lib",
    ],
)

py_library(
    name = "connector_py",
    srcs = ["channel.py"],
    visibility = ["//visibility:public"],
    deps = [
        ":connector_ext",
        "//proto:python_proto_library",
    ],
)
