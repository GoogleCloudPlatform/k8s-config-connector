load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto_grpc//python:defs.bzl", "python_grpc_library")
load("@rules_python//experimental/python:wheel.bzl", "py_package")

def connector(name, resources):
    pkg = "_".join(name.split("/"))
    go_library(
        name = "%s_connector" % pkg,
        srcs = ["register.go"] + ["%s_server.go" % r for r in resources],
        importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/services/%s" % name,
        visibility = ["//visibility:public"],
        deps = [
            "//proto:empty_go_proto",
            "//proto/%s:%s_go_proto" % (name, pkg),
            "@core_dcl//dcl:go_default_library",
            "@core_dcl//hclconv:go_default_library",
            "@core_dcl//services/google/%s:go_default_library" % name,
            "@org_golang_google_grpc//:go_default_library",
        ],
    )

    native.py_library(
        name = pkg + "_py",
        srcs = ["%s.py" % r for r in resources],
        visibility = ["//visibility:public"],
        deps = [
            "//connector:connector_py",
            "//proto:python_proto_library",
        ],
    )

def all_packages(names):
    native.genrule(
        name = "dummy_binary",
        srcs = [],
        outs = ["all_packages.py"],
        cmd = "echo \"\" > \"$@\"",
    )

    native.py_binary(
        name = "all_packages",
        srcs = ["all_packages.py"],
        visibility = ["//visibility:public"],
        deps = [
            "//services/%s:%s_py" % (name, name.replace("/", "_"))
            for name in names
        ],
    )

    py_package(
        name = "pkg",
        deps = [
            ":all_packages",
            "//proto:python_proto_library",
        ],
        packages = ["dcl", "proto", "services", "connector"],
        visibility = ["//visibility:public"],
    )

def proto_package(name, resources):
    pkg = "_".join(name.split("/"))
    proto_library(
        name = "%s_proto" % pkg,
        srcs = ["%s.proto" % r for r in resources],
        visibility = ["//visibility:public"],
        deps = [
            "//proto:empty_proto",
            "//proto/connector:sdk_proto",
        ],
    )

    go_proto_library(
        name = "%s_go_proto" % pkg,
        compilers = ["@io_bazel_rules_go//proto:go_grpc"],
        importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/%s/%s_go_proto" % (name, pkg),
        proto = ":%s_proto" % pkg,
        visibility = ["//visibility:public"],
        deps = [
            "//proto:empty_go_proto",
            "//proto/connector:sdk_go_proto",
        ],
    )

def python_protos(packages):
    python_grpc_library(
        name = "python_proto_library",
        visibility = ["//visibility:public"],
        deps = [
                   ":empty_proto",
                   "//proto/connector:connector_proto",
                   "//proto/connector:sdk_proto",
               ] +
               ["//proto/%s:%s_proto" % (p, p.replace("/", "_")) for p in packages],
    )
