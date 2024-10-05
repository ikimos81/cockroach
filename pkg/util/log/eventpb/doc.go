// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Package eventpb defines standardized payloads for cluster-level and
// SQL-level event reporting.
//
// The consumers of structured events include:
//
//   - Instrusion Detection Systems in the CockroachCloud infrastructure.
//     Check in with the security team before changing or evolving
//     any payloads related to privilege changes, user management,
//     authentication or other security-adjacent features.
//
// - The DB Console (ex "Admin UI") code. See ui/util/events*.ts.
//
//   - Customers using automatic event triaging in network log collectors.
//     Check in with the product team before changing or evolving
//     the way event types are classified into categories or channels.
//     Also check in with the product team before changing
//     or evolving "Common" payload types shared across multiple event types.
//
//   - Documentation. These are published APIs. A reference doc is
//     automatically generated (docs/generated/eventlog.md). Significant
//     changes to the event documentation should be reviewed by
//     the doc team.
//
// Maintenance instructions:
//
// Generally, the event payloads must be stable across releases.
//
// They are also part of the external API. Any changes to event
// definitions must be supported by a release note in the commit
// message. The deprecation process applies to remove non-alpha,
// non-reserved published APIs.
//
// It is possible to opt out of the stability guarantee for a new
// payload by documenting an entire event payload or individual fields
// as "Reserved and subject to change without notice". That mention
// must also be placed in a release note.
//
// Note that the JSON compatibility rules apply, not just protobuf:
//
//   - The names of the types and fields are part of the interface. A
//     name change is a breaking change. The casing of the field names
//     matters too.
//
//   - the fields in the proto definition have jsontag starting with the
//     empty string before the comma, e.g. ",omitempty" so as to
//     override the default JSON field name generated by the protobuf
//     compiler.  We want to keep the cased names for compatibility with
//     pre-v21.1 consumers of system.eventlog.
//
//   - likewise, the entire CommonXXXX payloads are marked as embedded
//     and their json tag is removed in every event log message, to
//     ensure that it appears inline in the JSON output. This is because
//     we wish our event structures to be flat and easier to consume.
//
// Beware: because the common structs are marked inline in the
// individual events, care must be taken to not reuse field
// identifiers across the message and common types, otherwise the JSON
// conversions cannot work.
package eventpb
