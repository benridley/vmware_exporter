workspace(name = "vmware_exporter")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7be7dc01f1e0afdba6c8eb2b43d2fa01c743be1b9273ab1eaf6c233df078d705",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.5/rules_go-0.16.5.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "bc653d3e058964a5a26dcad02b6c72d7d63e6bb88d94704990b908a1445b8758",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.13.0/bazel-gazelle-0.13.0.tar.gz"],
)

http_archive(
    name = "bazel_toolchains",
    sha256 = "109a99384f9d08f9e75136d218ebaebc68cc810c56897aea2224c57932052d30",
    strip_prefix = "bazel-toolchains-94d31935a2c94fe7e7c7379a0f3393e181928ff7",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-toolchains/archive/94d31935a2c94fe7e7c7379a0f3393e181928ff7.tar.gz",
        "https://github.com/bazelbuild/bazel-toolchains/archive/94d31935a2c94fe7e7c7379a0f3393e181928ff7.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_parnurzeal_gorequest",
    commit = "9433d8c3b4c5f3bbeca487a6a410008eab3ce5ee",
    importpath = "github.com/parnurzeal/gorequest",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "059132a15dd08d6704c67711dae0cf35ab991756",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_moul_http2curl",
    commit = "faeffb3553568c6ecaa9c103c09dea941ca9c570",
    importpath = "github.com/moul/http2curl",
)

go_repository(
    name = "com_github_alecthomas_template",
    commit = "a0175ee3bccc567396460bf5acd36800cb10c49c",
    importpath = "github.com/alecthomas/template",
)

go_repository(
    name = "com_github_alecthomas_units",
    commit = "2efee857e7cfd4f3d0138cc3cbb1b4966962b93a",
    importpath = "github.com/alecthomas/units",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    commit = "d2ead25884778582e740573999f7b07f47e171b4",
    importpath = "github.com/prometheus/client_golang",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "b1a0a9a36d7453ba0f62578b99712f3a6c5f82d1",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "f287a105a20ec685d797f65cd0ce8fbeaef42da1",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "2998b132700a7d019ff618c06a234b47c1f3f681",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973f24aa725d07868b467d1ddfceafb",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    commit = "c182affec369e30f25d3eb8cd8a478dee585ae7d",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    commit = "947dcec5ba9c011838740e680966fd7087a71d0d",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    commit = "78fb3852d92683dc28da6cc3d5f965100677c27d",
    importpath = "github.com/sirupsen/logrus",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "ff983b9c42bc9fbf91556e191cc8efb585c16908",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "com_github_alecthomas_kingpin",
    commit = "e7f8ee3d9b4b47000af44d34904a9e4f598e577f",
    importpath = "github.com/alecthomas/kingpin",
)

go_repository(
    name = "com_github_vmware_govmomi",
    commit = "c94f5f3aed1c44b3c977bd44a9679a3dd1733616",
    importpath = "github.com/vmware/govmomi",
)
