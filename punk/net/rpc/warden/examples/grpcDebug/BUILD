package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_binary",
)

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "angrymiao-go/punk/net/rpc/warden/examples/grpcDebug",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/jsonpb:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/credentials:go_default_library",
        "//vendor/google.golang.org/grpc/encoding:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "grpcDebug",
    embed = [":go_default_library"],
    tags = ["automanaged"],
)
