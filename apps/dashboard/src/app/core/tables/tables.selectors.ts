import { createFeatureSelector, createSelector } from '@ngrx/store';
import { State, stateKey } from './tables.reducer';

export const selectTablesState = createFeatureSelector<State>(stateKey);

export const selectTablesCount = createSelector(
    selectTablesState,
    (state: State) => state.count
);
