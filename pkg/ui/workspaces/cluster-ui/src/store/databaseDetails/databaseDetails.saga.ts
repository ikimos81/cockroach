// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

import { all, call, put, takeEvery } from "redux-saga/effects";

import { databaseDetailsReducer } from "./databaseDetails.reducer";
import {
  DatabaseDetailsReqParams,
  ErrorWithKey,
  getDatabaseDetails,
} from "src/api";
import moment from "moment";
import { PayloadAction } from "@reduxjs/toolkit";

const actions = databaseDetailsReducer.actions;
export function* refreshDatabaseDetailsSaga(
  action: PayloadAction<DatabaseDetailsReqParams>,
) {
  yield put(actions.request(action.payload));
}

export function* requestDatabaseDetailsSaga(
  action: PayloadAction<DatabaseDetailsReqParams>,
): any {
  try {
    const result = yield call(
      getDatabaseDetails,
      action.payload,
      moment.duration(10, "m"),
    );
    yield put(
      actions.received({
        key: action.payload.database,
        databaseDetailsResponse: result,
      }),
    );
  } catch (e) {
    const err: ErrorWithKey = {
      err: e,
      key: action.payload.database,
    };
    yield put(actions.failed(err));
  }
}

export function* databaseDetailsSaga() {
  yield all([
    takeEvery(actions.refresh, refreshDatabaseDetailsSaga),
    takeEvery(actions.request, requestDatabaseDetailsSaga),
  ]);
}
