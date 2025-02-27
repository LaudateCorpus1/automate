#shellcheck disable=SC2034
#shellcheck disable=SC2154

pkg_name=automate-workflow-nginx
pkg_origin=chef
pkg_version=2.8.61
pkg_maintainer="Chef Software, Inc. <support@chef.io>"
pkg_license=('Chef-MLSA')
vendor_origin=${vendor_origin:-"chef"}

pkg_deps=(
  # WARNING: Version pin managed by .expeditor/update_chef_server.sh
  "${vendor_origin}/openresty-noroot/1.13.6.2/20210225010004"
  "${vendor_origin}/automate-workflow-web"
  chef/mlsa
  core/bash
  core/curl
  core/coreutils
  core/libossp-uuid
)

pkg_exposes=(port ssl-port)
pkg_exports=(
  [port]=port
  [ssl-port]=ssl_port
)

pkg_binds=(
  [automate-workflow-server]="port"
)

pkg_svc_user="root"

do_download() {
  return 0
}

do_build() {
  return 0
}

do_install() {
    mkdir -pv "$pkg_prefix/www/workflow" "$pkg_prefix/www/loading"
    cp -Rv "$(pkg_path_for "${vendor_origin}/automate-workflow-web")"/dist/* "${pkg_prefix}/www/workflow"
    cp -v "$SRC_PATH/loading.html" "$pkg_prefix/www/loading/index.html"
}

do_strip() {
  return 0
}
