import { Action } from '@ngrx/store';
import { HttpErrorResponse } from '@angular/common/http';

import { DailyCheckInCountCollection,
  TopErrorsCollection, CountedDurationCollection } from './desktop.model';

export enum DesktopActionTypes {
  GET_DAILY_CHECK_IN_TIME_SERIES              = 'DESKTOP::GET::DAILY_CHECK_IN_TIME_SERIES',
  GET_DAILY_CHECK_IN_TIME_SERIES_SUCCESS      = 'DESKTOP::GET::DAILY_CHECK_IN_TIME_SERIES::SUCCESS',
  GET_DAILY_CHECK_IN_TIME_SERIES_FAILURE      = 'DESKTOP::GET::DAILY_CHECK_IN_TIME_SERIES::FAILURE',
  SET_DAYS_AGO_SELECTED                       = 'DESKTOP::SET::DAYS_AGO_SELECTED',
  GET_TOP_ERRORS_COLLECTION                   = 'DESKTOP::GET::TOP_ERRORS_COLLECTION',
  GET_TOP_ERRORS_COLLECTION_SUCCESS           = 'DESKTOP::GET::TOP_ERRORS_COLLECTION::SUCCESS',
  GET_TOP_ERRORS_COLLECTION_FAILURE           = 'DESKTOP::GET::TOP_ERRORS_COLLECTION::FAILURE',
  GET_UNKNOWN_DESKTOP_DURATION_COUNTS         = 'DESKTOP::GET::UNKNOWN_DESKTOP_DURATION_COUNT',
  GET_UNKNOWN_DESKTOP_DURATION_COUNTS_SUCCESS = 'DESKTOP::GET::UNKNOWN_DESKTOP_DURATION_COUNT::SUCCESS',
  GET_UNKNOWN_DESKTOP_DURATION_COUNTS_FAILURE = 'DESKTOP::GET::UNKNOWN_DESKTOP_DURATION_COUNT::FAILURE'
}

export class SetDaysAgoSelected implements Action {
  readonly type = DesktopActionTypes.SET_DAYS_AGO_SELECTED;
  constructor(public payload: {daysAgo: number}) { }
}

export class GetDailyCheckInTimeSeries implements Action {
  readonly type = DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES;
}

export class GetDailyCheckInTimeSeriesSuccess implements Action {
  readonly type = DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES_SUCCESS;
  constructor(public payload: DailyCheckInCountCollection) { }
}

export class GetDailyCheckInTimeSeriesFailure implements Action {
  readonly type = DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class GetTopErrorsCollection implements Action {
  readonly type = DesktopActionTypes.GET_TOP_ERRORS_COLLECTION;
}

export class GetTopErrorsCollectionSuccess implements Action {
  readonly type = DesktopActionTypes.GET_TOP_ERRORS_COLLECTION_SUCCESS;
  constructor(public payload: TopErrorsCollection) { }
}

export class GetTopErrorsCollectionFailure implements Action {
  readonly type = DesktopActionTypes.GET_TOP_ERRORS_COLLECTION_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class GetUnknownDesktopDurationCounts implements Action {
  readonly type = DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS;
}

export class GetUnknownDesktopDurationCountsSuccess implements Action {
  readonly type = DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS_SUCCESS;
  constructor(public payload: CountedDurationCollection) { }
}

export class GetUnknownDesktopDurationCountsFailure implements Action {
  readonly type = DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export type DesktopActions =
  | SetDaysAgoSelected
  | GetDailyCheckInTimeSeries
  | GetDailyCheckInTimeSeriesSuccess
  | GetDailyCheckInTimeSeriesFailure
  | GetTopErrorsCollection
  | GetTopErrorsCollectionSuccess
  | GetTopErrorsCollectionFailure
  | GetUnknownDesktopDurationCounts
  | GetUnknownDesktopDurationCountsSuccess
  | GetUnknownDesktopDurationCountsFailure;
