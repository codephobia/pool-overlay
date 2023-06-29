import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { Tournament } from '@pool-overlay/models';
import { switchMap, tap } from 'rxjs';
import { TournamentsService } from '../../services/tournament.service';

export enum LoadingState {
    INIT,
    LOADING,
    LOADED,
}

interface TournamentLoadedState {
    callState: LoadingState;
    tournament: Tournament | null;
}

export const initialState: TournamentLoadedState = {
    callState: LoadingState.INIT,
    tournament: null,
}

@Injectable()
export class TournamentLoadedStore extends ComponentStore<TournamentLoadedState> {
    constructor(
        private router: Router,
        private tournamentsService: TournamentsService,
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

    // selectors
    private isLoaded$ = this.select((state) => state.callState === LoadingState.LOADED);
    private tournament$ = this.select((state) => state.tournament);
    readonly vm$ = this.select(
        this.isLoaded$,
        this.tournament$,
        (isLoaded, tournament) => ({
            isLoaded,
            tournament,
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
}
