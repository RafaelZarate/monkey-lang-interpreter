load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/rafaelzarate/monkey-lang-interpreter/interpreter",
    visibility = ["//visibility:private"],
    deps = ["//repl:go_default_library"],
)

go_binary(
    name = "interpreter",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
