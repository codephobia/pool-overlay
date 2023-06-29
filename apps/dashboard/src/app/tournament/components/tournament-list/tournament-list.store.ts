import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { Tournament } from '@pool-overlay/models';
import { switchMap, tap } from 'rxjs';
import { TournamentsService } from '../../services/tournament.service';

export enum LoadingState {
    INIT,
    LOADING,
    LOADED,
}

interface TournamentListState {
    callState: LoadingState;
    tournaments: Tournament[];
}

export const initialState: TournamentListState = {
    callState: LoadingState.INIT,
    tournaments: [],
}

@Injectable()
export class TournamentListStore extends ComponentStore<TournamentListState> {
    constructor(
        private tournamentsService: TournamentsService,
    ) {
        super(initialState);
    }

    // updaters
    private updateCallState = this.updater<LoadingState>((state, callState) => ({
        ...state,
        callState,
    }));

    private updateTournaments = this.updater<Tournament[]>((state, tournaments) => ({
        ...state,
        tournaments,
        callState: LoadingState.LOADED,
    }));

    // selectors
    private isLoaded$ = this.select((state) => state.callState === LoadingState.LOADED);
    private tournaments$ = this.select((state) => state.tournaments);
    readonly vm$ = this.select(
        this.isLoaded$,
        this.tournaments$,
        (isLoaded, tournaments) => ({
            isLoaded,
            tournaments,
        })
    );

    // effects
    readonly getTournaments = this.effect((trigger$) => trigger$.pipe(
        tap(() => this.updateCallState(LoadingState.LOADING)),
        switchMap(() => this.tournamentsService.getList().pipe(
            tapResponse(
                (tournaments) => this.updateTournaments(tournaments),
                (err) => {
                    console.log(err);
                }
            )
        )),
    ));
}
