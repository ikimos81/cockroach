#!/usr/bin/env bash

# Copyright 2023 The Cockroach Authors.
#
# Use of this software is governed by the CockroachDB Software License
# included in the /LICENSE file.


set -xeuo pipefail

bazel build //pkg/cmd/bazci --config=ci

$(bazel info bazel-bin --config=ci)/pkg/cmd/bazci/bazci_/bazci -- \
  coverage \
  --config=ci --config=use_ci_timeouts -c fastbuild \
  --@io_bazel_rules_go//go/config:cover_format=lcov --combined_report=lcov \
  --instrumentation_filter="//pkg/..." \
  --profile=/artifacts/profile.gz \
  --flaky_test_attempts=4 \
  //pkg:small_non_ccl_tests //pkg:medium_non_ccl_tests //pkg:large_non_ccl_tests

lcov_file="$(bazel info output_path)/_coverage/_coverage_report.dat"
if [ ! -f "${lcov_file}" ]; then
  echo "Coverage file ${lcov_file} does not exist"
  exit 1
fi

cp "${lcov_file}" /artifacts/unit_tests_nonccl.lcov
