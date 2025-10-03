load("@bazel_gazelle//:def.bzl", "gazelle")
load("@rules_python//experimental/python:wheel.bzl", "py_wheel")

# gazelle:prefix github.com/GoogleCloudPlatform/declarative-resource-client-library/python
gazelle(name = "gazelle")

py_wheel(
    name = "wheel",
    distribution = "dcl",
    python_tag = "py3",
    requires = [
        "grpcio",
        "protobuf",
    ],
    strip_path_prefixes = ["proto/python_proto_library_pb"],
    version = "0.0.1",
    deps = [
        "//services:pkg",
    ],
)
