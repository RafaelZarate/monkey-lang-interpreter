load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["parser.go"],
    importpath = "github.com/rafaelzarate/monkey-lang-interpreter/parser",
    visibility = ["//visibility:public"],
    deps = [
        "//ast:go_default_library",
        "//lexer:go_default_library",
        "//token:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["parser_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//ast:go_default_library",
        "//lexer:go_default_library",
    ],
)
