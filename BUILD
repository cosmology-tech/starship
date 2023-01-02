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

# rule for helm charts
genrule(
    name = "helm-chart",
    srcs = [
        "charts/devnet/",
    ],
    outs = ["manifests"],
    cmd = """
        helm template \
        --values charts/devnet/values.yaml \
        --output $@ \
        charts/devnet/
    """,
)