load("//tools/build_defs/go:go_library.bzl", "go_library")
load("//tools/build_defs/go:go_test.bzl", "go_test")

go_library(
    name = "bigtable",
    srcs = [
        "appprofile_fuzzer.go",
        "appprofile_mappings.go",
        "authorizedview_fuzzer.go",
        "authorizedview_mapper.go",
        "backup_fuzzer.go",
        "backup_mapper.go",
        "bigtableappprofile_controller.go",
        "cluster_fuzzer.go",
        "cluster_mapper.go",
        "instance_mapper.go",
        "logicalview_controller.go",
        "logicalview_fuzzer.go",
        "mapper.generated.go",
        "utils.go",
    ],
    deps = [
        "//devtools/compliance/common/sets",
        "//google/protobuf:field_mask_go_proto",
    ],
)

go_test(
    name = "bigtable_test",
    srcs = ["roundtrip_test.go"],
    library = ":bigtable",
)
