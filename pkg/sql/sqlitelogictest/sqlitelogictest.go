// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package sqlitelogictest

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
)

// The sqlitelogictestpackage runs logic tests from CockroachDB's fork of sqllogictest:
//    https:www.sqlite.org/sqllogictest/doc/trunk/about.wiki
// This fork contains many generated tests created by the SqlLite project that
// ensure the tested SQL database returns correct statement and query output.
// The logic tests are reasonably independent of the specific dialect of each
// database so that they can be retargeted. In fact, the expected output for
// each test can be generated by one database and then used to verify the output
// of another database.

// By default, these tests are skipped, unless the `bigtest` flag is specified.
// The reason for this is that these tests are contained in another repo that
// must be present on the machine, and because they take a long time to run.

// Globs is a list of globs capturing the sqlite logic tests we run in our
// logictests.
var Globs []string = []string{
	"/test/index/between/*/*.test",
	"/test/index/commute/*/*.test",
	"/test/index/delete/*/*.test",
	"/test/index/in/*/*.test",
	"/test/index/orderby/*/*.test",
	"/test/index/orderby_nosort/*/*.test",
	"/test/index/view/*/*.test",

	"/test/select1.test",
	"/test/select2.test",
	"/test/select3.test",
	"/test/select4.test",
	"/test/select5.test",

	// TODO(pmattis): Incompatibilities in numeric types.
	// For instance, we type SUM(int) as a decimal since all of our ints are
	// int64.
	// "/test/random/expr/*.test",

	// TODO(pmattis): We don't support unary + on strings.
	// "/test/index/random/*/*.test",
	// "/test/random/aggregates/*.test",
	// "/test/random/groupby/*.test",
	// "/test/random/select/*.test",
}

// FindLocalLogicTestClone returns  the path to a local checkout of the
// cockroachdb/sqllogictest repo if one can be found. This only needs to be
// called from non-Bazel test runs, since Bazel test runs can access the
// appropriate content from the runfiles.
func FindLocalLogicTestClone() (string, error) {
	logicTestPath := build.Default.GOPATH + "/src/github.com/cockroachdb/sqllogictest"
	if _, err := os.Stat(logicTestPath); err != nil {
		var fullPath string
		fullPath, err = filepath.Abs(logicTestPath)
		return "", fmt.Errorf("unable to find sqllogictest repo (got error %w): %s\n"+
			"git clone https://github.com/cockroachdb/sqllogictest %s",
			err, logicTestPath, fullPath)
	}
	return logicTestPath, nil
}
