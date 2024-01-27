import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { GameType, Tournament } from '@pool-overlay/models';
import { switchMap, tap, withLatestFrom } from 'rxjs';
import { TournamentsService } from '../../services/tournament.service';

export enum LoadingState {
    INIT,
    LOADING,
    LOADED,
}

interface TournamentSetupState {
    callState: LoadingState;
    tournament: Tournament | null;
    maxTables: number;
    isHandicapped: boolean;
    showOverlay: boolean;
    showFlags: boolean;
    showFargo: boolean;
    showScore: boolean;
    gameType: GameType;
    aSideRaceTo: number;
    bSideRaceTo: number;
}

export const initialState: TournamentSetupState = {
    callState: LoadingState.INIT,
    tournament: null,
    maxTables: 3,
    isHandicapped: true,
    showOverlay: true,
    showFlags: false,
    showFargo: true,
    showScore: true,
    gameType: GameType.EightBall,
    aSideRaceTo: 3,
    bSideRaceTo: 2,
}

@Injectable()
export class TournamentSetupStore extends ComponentStore<TournamentSetupState> {
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

    readonly updateMaxTables = this.updater<number>((state, maxTables) => ({
        ...state,
        maxTables,
    }));

    readonly updateGameType = this.updater<GameType>((state, gameType) => ({
        ...state,
        gameType,
    }));

    readonly updateShowOverlay = this.updater<boolean>((state, showOverlay) => ({
        ...state,
        showOverlay,
    }));

    readonly updateShowFlags = this.updater<boolean>((state, showFlags) => ({
        ...state,
        showFlags,
    }));

    readonly updateShowFargo = this.updater<boolean>((state, showFargo) => ({
        ...state,
        showFargo,
    }));

    readonly updateShowScore = this.updater<boolean>((state, showScore) => ({
        ...state,
        showScore,
    }));

    readonly updateIsHandicapped = this.updater<boolean>((state, isHandicapped) => ({
        ...state,
        isHandicapped,
    }));

    readonly updateASideRaceTo = this.updater<number>((state, aSideRaceTo) => ({
        ...state,
        aSideRaceTo,
    }));

    readonly updateBSideRaceTo = this.updater<number>((state, bSideRaceTo) => ({
        ...state,
        bSideRaceTo,
    }));

    // selectors
    private isLoaded$ = this.select((state) => state.callState === LoadingState.LOADED);
    private tournament$ = this.select((state) => state.tournament);
    private maxTables$ = this.select((state) => state.maxTables);
    private isHandicapped$ = this.select((state) => state.isHandicapped);
    private showOverlay$ = this.select((state) => state.showOverlay);
    private showFlags$ = this.select((state) => state.showFlags);
    private showFargo$ = this.select((state) => state.showFargo);
    private showScore$ = this.select((state) => state.showScore);
    private gameType$ = this.select((state) => state.gameType);
    private aSideRaceTo$ = this.select((state) => state.aSideRaceTo);
    private bSideRaceTo$ = this.select((state) => state.bSideRaceTo);
    readonly vm$ = this.select(
        this.isLoaded$,
        this.tournament$,
        this.maxTables$,
        this.isHandicapped$,
        this.showOverlay$,
        this.showFlags$,
        this.showFargo$,
        this.showScore$,
        this.gameType$,
        this.aSideRaceTo$,
        this.bSideRaceTo$,
        (isLoaded, tournament, maxTables, isHandicapped, showOverlay, showFlags, showFargo, showScore, gameType, aSideRaceTo, bSideRaceTo) => ({
            isLoaded,
            tournament,
            maxTables,
            isHandicapped,
            showOverlay,
            showFlags,
            showFargo,
            showScore,
            gameType,
            aSideRaceTo,
            bSideRaceTo,
        })
    );

    // effects
    readonly getTournamentById = this.effect<number>((tournamentId$) => tournamentId$.pipe(
        tap(() => this.updateCallState(LoadingState.LOADING)),
        switchMap((tournamentId) => this.tournamentsService.getById(tournamentId).pipe(
            tapResponse(
                (tournament) => this.updateTournament(tournament),
                (err) => {
                    console.log(err);
                }
            )
        )),
    ));

    readonly loadTournament = this.effect((trigger$) => trigger$.pipe(
        withLatestFrom(
            this.tournament$,
            this.maxTables$,
            this.gameType$,
            this.showOverlay$,
            this.showFlags$,
            this.showFargo$,
            this.showScore$,
            this.isHandicapped$,
            this.aSideRaceTo$,
            this.bSideRaceTo$,
        ),
        switchMap(([, tournament, max_tables, game_type, show_overlay, show_flags, show_fargo, show_score, is_handicapped, a_side_race_to, b_side_race_to]) => this.tournamentsService.load(tournament!.id, {
            max_tables,
            game_type,
            show_overlay,
            show_flags,
            show_fargo,
            show_score,
            is_handicapped,
            a_side_race_to,
            b_side_race_to,
        }).pipe(
            tapResponse(
                () => {
                    void this.router.navigate(['/tournaments/loaded']);
                },
                (err) => {
                    console.error(err);
                }
            )
        )),
    ));
}
