import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { forkJoin } from 'rxjs';
import { switchMap } from 'rxjs/operators';

import { IPlayer } from '@pool-overlay/models';
import { PlayersService } from '../../../shared/services/players.service';

export interface PlayerModalState {
    loaded: boolean;
    page: number;
    count: number;
    players: IPlayer[];
}

@Injectable()
export class PlayerModalStore extends ComponentStore<PlayerModalState> {
    constructor(
        private playersService: PlayersService,
    ) {
        super({
            loaded: false,
            page: 1,
            count: 0,
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

    public readonly setCount = this.updater<number>((state, count) => ({
        ...state,
        count,
    }));

    public readonly setPlayers = this.updater<IPlayer[]>((state, players) => ({
        ...state,
        loaded: true,
        players,
    }));

    public readonly loaded$ = this.select(state => state.loaded);
    public readonly page$ = this.select(state => state.page);
    public readonly count$ = this.select(state => state.count);
    public readonly players$ = this.select(state => state.players);
    public readonly vm$ = this.select(
        this.loaded$,
        this.page$,
        this.count$,
        this.players$,
        (loaded, page, count, players) => ({
            loaded,
            page,
            count,
            players,
        })
    );

    public readonly getPlayers = this.effect<number>(page$ => page$.pipe(
        switchMap(page => forkJoin([
            this.playersService.find(page),
            this.playersService.count(),
        ]).pipe(
            tapResponse(
                ([players, { count }]) => {
                    this.setPlayers(players);
                    this.setPage(page);
                    this.setCount(count);
                },
                error => console.error(error)
            ),
        )),
    ));
}
