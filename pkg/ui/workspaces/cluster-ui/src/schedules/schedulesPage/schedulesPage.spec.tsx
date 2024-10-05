// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

import { render } from "@testing-library/react";
import React from "react";
import { MemoryRouter } from "react-router-dom";
import * as H from "history";

import { Schedule } from "src/api/schedulesApi";

import { allSchedulesFixture } from "./schedulesPage.fixture";
import { SchedulesPage, SchedulesPageProps } from "./schedulesPage";

const getMockSchedulesPageProps = (
  schedules: Array<Schedule>,
): SchedulesPageProps => {
  const history = H.createHashHistory();
  return {
    sort: { columnTitle: null, ascending: true },
    status: "",
    show: "50",
    setSort: () => {},
    setStatus: () => {},
    setShow: () => {},
    schedules: schedules,
    schedulesLoading: false,
    schedulesError: null,
    refreshSchedules: () => {},
    location: history.location,
    history,
    match: {
      url: "",
      path: history.location.pathname,
      isExact: false,
      params: {},
    },
  };
};

describe("Schedules", () => {
  it("renders expected schedules table columns", () => {
    const { getByText } = render(
      <MemoryRouter>
        <SchedulesPage {...getMockSchedulesPageProps(allSchedulesFixture)} />
      </MemoryRouter>,
    );
    const expectedColumnTitles = [
      "Label",
      "Status",
      "Schedule ID",
      "Owner",
      "Recurrence",
      "Creation Time (UTC)",
      "Next Execution Time (UTC)",
      "Jobs Running",
    ];

    for (const columnTitle of expectedColumnTitles) {
      getByText(columnTitle);
    }
  });

  it("renders a message when the table is empty", () => {
    const { getByText } = render(
      <MemoryRouter>
        <SchedulesPage {...getMockSchedulesPageProps([])} />
      </MemoryRouter>,
    );
    getByText("No schedules to show");
  });
});
