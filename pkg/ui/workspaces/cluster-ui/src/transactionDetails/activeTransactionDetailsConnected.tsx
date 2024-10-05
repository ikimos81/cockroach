// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

import { RouteComponentProps, withRouter } from "react-router-dom";
import { connect } from "react-redux";
import { Dispatch } from "redux";

import { actions as sessionsActions } from "src/store/sessions";
import {
  selectActiveTransaction,
  selectContentionDetailsForTransaction,
} from "src/selectors/activeExecutions.selectors";

import { AppState } from "../store";

import {
  ActiveTransactionDetails,
  ActiveTransactionDetailsDispatchProps,
} from "./activeTransactionDetails";

import { ActiveTransactionDetailsStateProps } from ".";

// For tenant cases, we don't show information about node, regions and
// diagnostics.
const mapStateToProps = (
  state: AppState,
  props: RouteComponentProps,
): ActiveTransactionDetailsStateProps => {
  return {
    transaction: selectActiveTransaction(state, props),
    contentionDetails: selectContentionDetailsForTransaction(state, props),
    match: props.match,
  };
};

const mapDispatchToProps = (
  dispatch: Dispatch,
): ActiveTransactionDetailsDispatchProps => ({
  refreshLiveWorkload: () => dispatch(sessionsActions.refresh()),
});

export const RecentTransactionDetailsPageConnected = withRouter(
  connect(mapStateToProps, mapDispatchToProps)(ActiveTransactionDetails),
);

// Prior to 23.1, this component was called
// ActiveTransactionDetailsPageConnected. We keep the alias here to avoid
// breaking the multi-version support in managed-service's console code.
// When managed-service drops support for 22.2 (around the end of 2024?),
// we can remove this code.
export const ActiveTransactionDetailsPageConnected =
  RecentTransactionDetailsPageConnected;
