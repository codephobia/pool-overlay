import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap, withLatestFrom } from 'rxjs/operators';

import { IGame } from '@pool-overlay/models';
import { APIService } from './api.service';
import { PlayerUpdateEvent } from '@pool-overlay/score-keeper';

export interface Settings {
    table: number;
    serverAddress: string;
};

export interface ScoreKeeperState {
    game: IGame | null;
    settings: Settings;
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
                table: 2,
                serverAddress: '192.168.0.26',
            },
        });
    }

    // updaters
    public readonly setGame = this.updater<{ game: IGame }>((state, { game }) => ({
        ...state,
        game,
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
    public readonly table$ = this.select((state) => state.settings.table);
    public readonly serverAddress$ = this.select((state) => state.settings.serverAddress);
    public readonly vm$ = this.select(
        this.game$,
        this.table$,
        this.serverAddress$,
        (game, table, serverAddress) => ({
            game,
            table,
            serverAddress,
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
        withLatestFrom(this.select(state => state.settings.table)),
        switchMap(([{ playerNum, direction }, table]) => this._apiService.updateScore(table, playerNum, direction).pipe(
            tapResponse(
                () => { },
                error => console.error(error),
            )
        ))
    ));
}
