// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Code generated by "stringer"; DO NOT EDIT.

package schemachange

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ColumnConversionDangerous - -1]
	_ = x[ColumnConversionImpossible-0]
	_ = x[ColumnConversionTrivial-1]
	_ = x[ColumnConversionValidate-2]
	_ = x[ColumnConversionGeneral-3]
}

func (i ColumnConversionKind) String() string {
	switch i {
	case ColumnConversionDangerous:
		return "Dangerous"
	case ColumnConversionImpossible:
		return "Impossible"
	case ColumnConversionTrivial:
		return "Trivial"
	case ColumnConversionValidate:
		return "Validate"
	case ColumnConversionGeneral:
		return "General"
	default:
		return "ColumnConversionKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
