package(default_visibility=["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
  name = "dockerized",
  srcs = [
    "src/dockerexec.go",
  ],
  data = ["resources/compose.test.yml"],
  importpath = "dockerized",
)