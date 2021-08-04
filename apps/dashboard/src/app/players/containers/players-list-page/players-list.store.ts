import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap } from 'rxjs/operators';

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

    // TODO: ADD PAGE NUMBER TO GETPLAYERS
    // TODO: ADD PLAYER COUNT TO STATE AND GETPLAYERS CALL

    public readonly getPlayers = this.effect(trigger$ => trigger$.pipe(
        switchMap(() => this.playersService.find(1).pipe(
            tapResponse(
                players => {
                    this.setPlayers(players)
                },
                error => console.error(error)
            ),
        )),
    ));
}
