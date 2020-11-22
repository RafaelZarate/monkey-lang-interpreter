load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["repl.go"],
    importpath = "github.com/rafaelzarate/monkey-lang-interpreter/repl",
    visibility = ["//visibility:public"],
    deps = [
        "//lexer:go_default_library",
        "//token:go_default_library",
    ],
)
