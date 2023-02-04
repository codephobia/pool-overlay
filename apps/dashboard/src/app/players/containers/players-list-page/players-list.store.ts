import { Injectable } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { forkJoin } from 'rxjs';
import { concatMap, map, switchMap, takeUntil } from 'rxjs/operators';

import { IPlayer } from '@pool-overlay/models';
import { PlayersService } from '../../../shared/services/players.service';

export interface PlayersListState {
    loaded: boolean;
    page: number;
    count: number;
    players: IPlayer[];
}

@Injectable()
export class PlayersListStore extends ComponentStore<PlayersListState> {
    constructor(
        route: ActivatedRoute,
        private playersService: PlayersService,
    ) {
        super({
            loaded: false,
            page: 1,
            count: 0,
            players: []
        });

        route.queryParamMap.pipe(
            map(params => Number(params.get('page'))),
            takeUntil(this.destroy$),
        ).subscribe(page => {
            const newPage = page ? page : 1;
            this.setPage(newPage);
            this.getPlayers(newPage);
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
                    this.setCount(count);
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
