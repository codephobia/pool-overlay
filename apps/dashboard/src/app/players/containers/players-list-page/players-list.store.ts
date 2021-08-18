import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { concatMap, switchMap } from 'rxjs/operators';

import { IPlayer } from '@pool-overlay/models';
import { PlayersService } from '../../services/players.service';

export interface PlayersListState {
    loaded: boolean;
    page: number;
    players: IPlayer[];
}

@Injectable()
export class PlayersListStore extends ComponentStore<PlayersListState> {
    constructor(
        private playersService: PlayersService,
    ) {
        super({
            loaded: false,
            page: 1,
            players: []
        });
    }

    public readonly setLoaded = this.updater<boolean>((state, loaded) => ({
        ...state,
        loaded,
    }));

    public readonly setPage = this.updater<number>((state, page) => ({
        ...state,
        page,
    }));

    public readonly setPlayers = this.updater<IPlayer[]>((state, players) => ({
        ...state,
        loaded: true,
        players,
    }));

    public readonly deletePlayer = this.updater<number>((state, playerId) => {
        const playerIndex = state.players.findIndex(player => player.id === playerId);
        const newPlayers = [...state.players];
        newPlayers.splice(playerIndex, 1);

        return {
            ...state,
            players: newPlayers,
        };
    });

    public readonly loaded$ = this.select(state => state.loaded);
    public readonly page$ = this.select(state => state.page);
    public readonly players$ = this.select(state => state.players);
    public readonly vm$ = this.select(
        this.loaded$,
        this.page$,
        this.players$,
        (loaded, page, players) => ({
            loaded,
            page,
            players,
        })
    );

    // TODO: ADD PLAYER COUNT TO STATE AND GETPLAYERS CALL

    public readonly getPlayers = this.effect<number>(page$ => page$.pipe(
        switchMap(page => this.playersService.find(page).pipe(
            tapResponse(
                players => {
                    this.setPlayers(players)
                },
                error => console.error(error)
            ),
        )),
    ));

    public readonly deletePlayerById = this.effect<number>(playerId$ => playerId$.pipe(
        concatMap(playerId => this.playersService.delete(playerId).pipe(
            tapResponse(
                () => {
                    this.deletePlayer(playerId);
                },
                error => console.error(error)
            ),
        )),
    ));
}
