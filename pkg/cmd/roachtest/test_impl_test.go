// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package main

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/registry"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
)

func TestTeamCityEscape(t *testing.T) {
	require.Equal(t, "|n", TeamCityEscape("\n"))
	require.Equal(t, "|r", TeamCityEscape("\r"))
	require.Equal(t, "||", TeamCityEscape("|"))
	require.Equal(t, "|[", TeamCityEscape("["))
	require.Equal(t, "|]", TeamCityEscape("]"))

	require.Equal(t, "identity", TeamCityEscape("identity"))
	require.Equal(t, "aaa|nbbb", TeamCityEscape("aaa\nbbb"))
	require.Equal(t, "aaa|nbbb||", TeamCityEscape("aaa\nbbb|"))
	require.Equal(t, "||||", TeamCityEscape("||"))
	require.Equal(t, "Connection to 104.196.113.229 port 22: Broken pipe|r|nlost connection: exit status 1",
		TeamCityEscape("Connection to 104.196.113.229 port 22: Broken pipe\r\nlost connection: exit status 1"))

	require.Equal(t,
		"Messages:   	current binary |'24.1|' not found in |'versionToMinSupportedVersion|'",
		TeamCityEscape("Messages:   	current binary '24.1' not found in 'versionToMinSupportedVersion'"),
	)

	// Unicode
	require.Equal(t, "|0x00bf", TeamCityEscape("\u00bf"))
	require.Equal(t, "|0x00bfaaa", TeamCityEscape("\u00bfaaa"))
	require.Equal(t, "bb|0x00bfaaa", TeamCityEscape("bb\u00bfaaa"))
}

func Test_failureSpecifyOwnerAndAddFailureCombination(t *testing.T) {
	ti := testImpl{
		l: nilLogger(),
	}
	ti.addFailure(0, "", errClusterProvisioningFailed(errors.New("oops")))
	errWithOwnership := failuresSpecifyOwner(ti.failures())

	require.NotNil(t, errWithOwnership)
	require.Equal(t, registry.OwnerTestEng, errWithOwnership.Owner)
}
