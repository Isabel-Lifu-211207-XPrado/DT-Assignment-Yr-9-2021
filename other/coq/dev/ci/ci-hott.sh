#!/usr/bin/env bash

ci_dir="$(dirname "$0")"
. "${ci_dir}/ci-common.sh"

git_download hott

( cd "${CI_BUILD_DIR}/hott" && ./autogen.sh -skip-submodules && ./configure \
      && sed -i.bak 's/\(HOQC =.*\)/\1 -native-compiler no/' Makefile \
      && make && make validate )
