HEADER_GLOBS = ["include/**/*.h"]

EXCLUDE_HEADERS = [
    "include/pybind11/common.h",
]

# These flags are needed for parse_headers feature.
COPTS = [
    "-fexceptions",
    "-Wno-c++98-c++11-c++14-compat",
    "-Wno-c++98-c++11-compat",
    "-Wno-google3-inline-namespace",
    "-Wno-google3-literal-operator",
    "-Wno-undefined-inline",
]

INCLUDES = ["include"]

cc_library(
    name = "pybind11",
    hdrs = glob(
        HEADER_GLOBS,
        exclude = EXCLUDE_HEADERS + [
            "include/pybind11/eigen.h",
        ],
    ),
    copts = COPTS,
    includes = INCLUDES,
)
