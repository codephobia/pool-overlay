import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { IGame, OverlayState, Tournament } from '@pool-overlay/models';
import { switchMap, tap } from 'rxjs';
import { TablesService } from '../../services/tables.service';
import { TournamentsService } from '../../services/tournament.service';
import { Store } from '@ngrx/store';
import * as fromTables from '../../../core/tables';

export enum LoadingState {
    INIT,
    LOADING,
    LOADED,
}

interface TournamentTable {
    game: IGame | null;
    overlay: OverlayState | null;
}

interface TournamentLoadedState {
    callState: LoadingState;
    tournament: Tournament | null;
    tables: Record<number, TournamentTable>;
}

export const initialState: TournamentLoadedState = {
    callState: LoadingState.INIT,
    tournament: null,
    tables: {},
}

@Injectable()
export class TournamentLoadedStore extends ComponentStore<TournamentLoadedState> {
    constructor(
        private router: Router,
        private store: Store,
        private tournamentsService: TournamentsService,
        private tablesService: TablesService,
    ) {
        super(initialState);
    }

    // updaters
    private updateCallState = this.updater<LoadingState>((state, callState) => ({
        ...state,
        callState,
    }));

    private updateTournament = this.updater<Tournament>((state, tournament) => ({
        ...state,
        tournament,
        callState: LoadingState.LOADED,
    }));

    private updateTableGame = this.updater<IGame>((state, game) => ({
        ...state,
        tables: {
            ...state.tables,
            [game.table]: {
                ...state.tables[game.table],
                game,
            },
        },
    }));

    private updateTableOverlay = this.updater<OverlayState>((state, overlay) => ({
        ...state,
        tables: {
            ...state.tables,
            [overlay.table]: {
                ...state.tables[overlay.table],
                overlay,
            },
        },
    }));

    // selectors
    private isLoaded$ = this.select((state) => state.callState === LoadingState.LOADED);
    private tournament$ = this.select((state) => state.tournament);
    private tables$ = this.select((state) => state.tables);
    readonly vm$ = this.select(
        this.isLoaded$,
        this.tournament$,
        this.tables$,
        this.store.select(fromTables.selectTablesCount),
        (isLoaded, tournament, tables, tablesCount) => ({
            isLoaded,
            tournament,
            tables,
            tablesArr: Array.from(new Array(tablesCount), (x, i) => i + 1),
        })
    );

    // effects
    readonly getCurrentTournament = this.effect((trigger$) => trigger$.pipe(
        tap(() => this.updateCallState(LoadingState.LOADING)),
        switchMap(() => this.tournamentsService.getCurrent().pipe(
            tapResponse(
                (tournament) => this.updateTournament(tournament),
                (err) => {
                    console.log(err);
                }
            )
        )),
    ));

    readonly unloadTournament = this.effect((trigger$) => trigger$.pipe(
        switchMap(() => this.tournamentsService.unload().pipe(
            tapResponse(
                () => {
                    void this.router.navigate(['/tournaments']);
                },
                (err) => {
                    console.error(err);
                }
            )
        )),
    ));

    readonly setGame = this.effect<{ game: IGame }>((trigger$) => trigger$.pipe(
        tap(({ game }) => this.updateTableGame(game)),
    ));

    readonly setOverlay = this.effect<OverlayState>((state$) => state$.pipe(
        tap((overlay) => this.updateTableOverlay(overlay)),
    ));

    readonly swapTables = this.effect<{ tableOne: number, tableTwo: number }>((trigger$) => trigger$.pipe(
        switchMap(({ tableOne, tableTwo }) => this.tablesService.swap(tableOne, tableTwo).pipe(
            tapResponse(
                () => { },
                (err) => console.error(err),
            ),
        ))
    ));
}
