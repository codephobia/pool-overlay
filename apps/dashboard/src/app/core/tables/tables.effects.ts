import { Injectable } from '@angular/core';
import { Actions, ROOT_EFFECTS_INIT, createEffect, ofType } from '@ngrx/effects';

import { TablesActions } from './tables.actions';
import { TableService } from '../../services/table.service';
import { catchError, map, of, switchMap } from 'rxjs';

@Injectable()
export class TablesEffects {
    constructor(
        private actions$: Actions,
        private tableService: TableService,
    ) { }

    init$ = createEffect(() => this.actions$.pipe(
        ofType(ROOT_EFFECTS_INIT),
        map(() => TablesActions.getCount()),
    ));

    getCount$ = createEffect(() => this.actions$.pipe(
        ofType(TablesActions.getCount),
        switchMap(() => this.tableService.count().pipe(
            map(({ count }) => TablesActions.getCountSuccess({ count })),
            catchError((error) => of(TablesActions.getCountError({ error }))),
        )),
    ));
}
