import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { filter, switchMap, tap, withLatestFrom } from 'rxjs/operators';
import { Dialogs } from '@nativescript/core';

import { Direction, IGame, OverlayState } from '@pool-overlay/models';
import { APIService } from './api.service';
import { PlayerUpdateEvent } from '@pool-overlay/score-keeper';

export interface Settings {
    table: number;
    serverAddress: string;
};

export interface ScoreKeeperState {
    game: IGame | null;
    settings: Settings;
    overlay: OverlayState;
}

@Injectable({
    providedIn: 'root',
})
export class ScoreKeeperStore extends ComponentStore<ScoreKeeperState> {
    constructor(
        private _apiService: APIService,
    ) {
        super({
            game: null,
            settings: {
                table: 1,
                serverAddress: '192.168.0.26',
            },
            overlay: {
                table: 1,
                hidden: false,
                showFlags: false,
                showFargo: true,
                showScore: true,
                waitingForPlayers: false,
                waitingForTournamentStart: false,
                tableNoLongerInUse: false,
            },
        });
    }

    // updaters
    public readonly setGame = this.updater<{ game: IGame }>((state, { game }) => ({
        ...state,
        game,
    }));

    public readonly setOverlay = this.updater<OverlayState>((state, overlay) => ({
        ...state,
        overlay: {
            ...state.overlay,
            waitingForPlayers: overlay.waitingForPlayers,
            waitingForTournamentStart: overlay.waitingForTournamentStart,
            tableNoLongerInUse: overlay.tableNoLongerInUse,
        },
    }));

    public readonly setTable = this.updater<number>((state, table) => ({
        ...state,
        settings: {
            ...state.settings,
            table,
        },
    }));

    public readonly setServerAddress = this.updater<string>((state, serverAddress) => ({
        ...state,
        settings: {
            ...state.settings,
            serverAddress,
        },
    }));

    // selectors
    public readonly game$ = this.select((state) => state.game);
    public readonly playerOneRaceTo$ = this.select(
        this.game$,
        (game) => {
            if (!game) {
                return 1;
            }

            if (!game.use_fargo_hot_handicap) {
                return game.race_to;
            }

            const playerOneFargo = game.player_one?.fargo_rating ?? 0;
            const playerTwoFargo = game.player_two?.fargo_rating ?? 0;
            const winsHigher = game.fargo_hot_handicap?.wins_higher ?? 0;
            const winsLower = game.fargo_hot_handicap?.wins_lower ?? 0;

            return playerOneFargo > playerTwoFargo ? winsHigher : winsLower;
        }
    );
    public readonly playerTwoRaceTo$ = this.select(
        this.game$,
        (game) => {
            if (!game) {
                return 1;
            }

            if (!game.use_fargo_hot_handicap) {
                return game.race_to;
            }

            const playerOneFargo = game.player_one?.fargo_rating ?? 0;
            const playerTwoFargo = game.player_two?.fargo_rating ?? 0;
            const winsHigher = game.fargo_hot_handicap?.wins_higher ?? 0;
            const winsLower = game.fargo_hot_handicap?.wins_lower ?? 0;

            return playerTwoFargo > playerOneFargo ? winsHigher : winsLower;
        }
    );
    public readonly showRaceTo$ = this.select(
        this.game$,
        (game) => {
            if (!game?.use_fargo_hot_handicap || !game?.fargo_hot_handicap) {
                return false;
            }

            const { race_to, wins_higher, wins_lower } = game.fargo_hot_handicap;

            return wins_higher !== race_to || wins_lower !== race_to;
        }
    );
    public readonly table$ = this.select((state) => state.settings.table);
    public readonly serverAddress$ = this.select((state) => state.settings.serverAddress);
    public readonly overlay$ = this.select((state) => state.overlay);
    public readonly vm$ = this.select(
        this.game$,
        this.playerOneRaceTo$,
        this.playerTwoRaceTo$,
        this.showRaceTo$,
        this.table$,
        this.serverAddress$,
        this.overlay$,
        (game, playerOneRaceTo, playerTwoRaceTo, showRaceTo, table, serverAddress, overlay) => ({
            game,
            playerOneRaceTo,
            playerTwoRaceTo,
            showRaceTo,
            table,
            serverAddress,
            overlay,
        })
    );

    // effects
    public readonly getGame = this.effect((trigger$) => {
        return trigger$.pipe(
            withLatestFrom(this.select(state => state.settings.table)),
            switchMap(([_, table]) =>
                this._apiService.getGame(table).pipe(
                    tapResponse(
                        (game) => {
                            this.setGame({ game });
                        },
                        (error) => console.error(error)
                    ),
                )
            )
        );
    });

    public readonly updateScore = this.effect<PlayerUpdateEvent>(trigger$ => trigger$.pipe(
        withLatestFrom(
            this.table$,
            this.game$,
            this.playerOneRaceTo$,
            this.playerTwoRaceTo$,
        ),
        switchMap(([{ playerNum, direction }, table, game, playerOneRaceTo, playerTwoRaceTo]) => this._apiService.updateScore(
            table,
            playerNum,
            direction,
        ).pipe(
            tapResponse(
                ({ scoreOne, scoreTwo }) => {
                    let winner = false;
                    let playerName = '';

                    if (playerNum === 1 && scoreOne === playerOneRaceTo) {
                        winner = true;
                        playerName = game.player_one.name;
                    } else if (playerNum === 2 && scoreTwo === playerTwoRaceTo) {
                        winner = true;
                        playerName = game.player_two.name;
                    }

                    if (winner) {
                        Dialogs.confirm({
                            title: 'Confirm Game Winner',
                            message: `Confirm ${playerName} won the game.`,
                            okButtonText: 'Confirm',
                            cancelButtonText: 'Cancel',
                        }).then((confirmed) => {
                            if (confirmed) {
                                this.saveGame();
                            } else {
                                const reverseDirection = (direction === Direction.INCREMENT) ? Direction.DECREMENT : Direction.INCREMENT;
                                this.updateScore({ playerNum, direction: reverseDirection });
                            }
                        });
                    }
                },
                error => console.error(error),
            ),
        ))
    ));

    public readonly updateGame = this.effect<IGame>((game$) => game$.pipe(
        withLatestFrom(this.select(state => state.settings.table)),
        filter(([game, table]) => game.table === table),
        tap(([game]) => this.setGame({ game })),
    ));

    public readonly updateOverlay = this.effect<OverlayState>((state$) => state$.pipe(
        withLatestFrom(this.select(state => state.settings.table)),
        filter(([state, table]) => state.table === table),
        tap(([state]) => this.setOverlay(state)),
    ));

    public readonly saveGame = this.effect((trigger$) => trigger$.pipe(
        withLatestFrom(this.select(state => state.settings.table)),
        switchMap(([, table]) => this._apiService.saveGame(table).pipe(
            tapResponse(
                () => { },
                error => console.error(error),
            ),
        )),
    ));
}
