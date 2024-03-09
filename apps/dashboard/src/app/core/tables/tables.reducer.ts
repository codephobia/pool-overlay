import { createReducer, on } from '@ngrx/store';
import { TablesActions } from './tables.actions';

export const stateKey = 'table-count';

export interface State {
    count: number;
}

export const initialState: State = {
    count: 1,
}

export const reducer = createReducer(
    initialState,
    on(
        TablesActions.getCountSuccess,
        TablesActions.setCount,
        (state, { count }) => ({
            ...state,
            count,
        })
    )
);
