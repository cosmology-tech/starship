load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/anmol1696/shuttle
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.work",
        "-prune",
    ],
    command = "update-repos",
)