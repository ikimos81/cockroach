// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_2-0]
	_ = x[Start22_1-1]
	_ = x[PebbleFormatBlockPropertyCollector-2]
	_ = x[ProbeRequest-3]
	_ = x[PublicSchemasWithDescriptors-4]
	_ = x[EnsureSpanConfigReconciliation-5]
	_ = x[EnsureSpanConfigSubscription-6]
	_ = x[EnableSpanConfigStore-7]
	_ = x[SCRAMAuthentication-8]
	_ = x[UnsafeLossOfQuorumRecoveryRangeLog-9]
	_ = x[AlterSystemProtectedTimestampAddColumn-10]
	_ = x[EnableProtectedTimestampsForTenant-11]
	_ = x[DeleteCommentsWithDroppedIndexes-12]
	_ = x[RemoveIncompatibleDatabasePrivileges-13]
	_ = x[AddRaftAppliedIndexTermMigration-14]
	_ = x[PostAddRaftAppliedIndexTermMigration-15]
	_ = x[DontProposeWriteTimestampForLeaseTransfers-16]
	_ = x[EnablePebbleFormatVersionBlockProperties-17]
	_ = x[MVCCIndexBackfiller-18]
	_ = x[EnableLeaseHolderRemoval-19]
	_ = x[LooselyCoupledRaftLogTruncation-20]
	_ = x[ChangefeedIdleness-21]
	_ = x[BackupDoesNotOverwriteLatestAndCheckpoint-22]
	_ = x[EnableDeclarativeSchemaChanger-23]
	_ = x[RowLevelTTL-24]
	_ = x[PebbleFormatSplitUserKeysMarked-25]
	_ = x[IncrementalBackupSubdir-26]
	_ = x[EnableNewStoreRebalancer-27]
	_ = x[ClusterLocksVirtualTable-28]
	_ = x[AutoStatsTableSettings-29]
	_ = x[SuperRegions-30]
	_ = x[EnableNewChangefeedOptions-31]
	_ = x[SpanCountTable-32]
	_ = x[PreSeedSpanCountTable-33]
	_ = x[SeedSpanCountTable-34]
	_ = x[V22_1-35]
	_ = x[Start22_2-36]
	_ = x[LocalTimestamps-37]
	_ = x[PebbleFormatSplitUserKeysMarkedCompacted-38]
	_ = x[EnsurePebbleFormatVersionRangeKeys-39]
	_ = x[EnablePebbleFormatVersionRangeKeys-40]
	_ = x[TrigramInvertedIndexes-41]
	_ = x[RemoveGrantPrivilege-42]
	_ = x[MVCCRangeTombstones-43]
	_ = x[UpgradeSequenceToBeReferencedByID-44]
	_ = x[SampledStmtDiagReqs-45]
	_ = x[AddSSTableTombstones-46]
	_ = x[SystemPrivilegesTable-47]
	_ = x[EnablePredicateProjectionChangefeed-48]
	_ = x[AlterSystemSQLInstancesAddLocality-49]
	_ = x[SystemExternalConnectionsTable-50]
}

const _Key_name = "V21_2Start22_1PebbleFormatBlockPropertyCollectorProbeRequestPublicSchemasWithDescriptorsEnsureSpanConfigReconciliationEnsureSpanConfigSubscriptionEnableSpanConfigStoreSCRAMAuthenticationUnsafeLossOfQuorumRecoveryRangeLogAlterSystemProtectedTimestampAddColumnEnableProtectedTimestampsForTenantDeleteCommentsWithDroppedIndexesRemoveIncompatibleDatabasePrivilegesAddRaftAppliedIndexTermMigrationPostAddRaftAppliedIndexTermMigrationDontProposeWriteTimestampForLeaseTransfersEnablePebbleFormatVersionBlockPropertiesMVCCIndexBackfillerEnableLeaseHolderRemovalLooselyCoupledRaftLogTruncationChangefeedIdlenessBackupDoesNotOverwriteLatestAndCheckpointEnableDeclarativeSchemaChangerRowLevelTTLPebbleFormatSplitUserKeysMarkedIncrementalBackupSubdirEnableNewStoreRebalancerClusterLocksVirtualTableAutoStatsTableSettingsSuperRegionsEnableNewChangefeedOptionsSpanCountTablePreSeedSpanCountTableSeedSpanCountTableV22_1Start22_2LocalTimestampsPebbleFormatSplitUserKeysMarkedCompactedEnsurePebbleFormatVersionRangeKeysEnablePebbleFormatVersionRangeKeysTrigramInvertedIndexesRemoveGrantPrivilegeMVCCRangeTombstonesUpgradeSequenceToBeReferencedByIDSampledStmtDiagReqsAddSSTableTombstonesSystemPrivilegesTableEnablePredicateProjectionChangefeedAlterSystemSQLInstancesAddLocalitySystemExternalConnectionsTable"

var _Key_index = [...]uint16{0, 5, 14, 48, 60, 88, 118, 146, 167, 186, 220, 258, 292, 324, 360, 392, 428, 470, 510, 529, 553, 584, 602, 643, 673, 684, 715, 738, 762, 786, 808, 820, 846, 860, 881, 899, 904, 913, 928, 968, 1002, 1036, 1058, 1078, 1097, 1130, 1149, 1169, 1190, 1225, 1259, 1289}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
