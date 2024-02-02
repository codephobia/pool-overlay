import { Injectable } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store, createSelector } from '@ngrx/store';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { combineLatest, forkJoin } from 'rxjs';
import { concatMap, map, switchMap, takeUntil } from 'rxjs/operators';

import { IPlayer } from '@pool-overlay/models';
import { PlayersService, PlayerFindOptions } from '../../../shared/services/players.service';
import { selectQueryParam } from '../../../shared/utils/router.selectors';

export interface PlayersListState {
    loaded: boolean;
    count: number;
    players: IPlayer[];
}

@Injectable()
export class PlayersListStore extends ComponentStore<PlayersListState> {
    constructor(
        private playersService: PlayersService,
        private store: Store,
    ) {
        super({
            loaded: false,
            count: 0,
            players: []
        });

        combineLatest([
            this.store.select(this.selectPage),
            this.store.select(this.selectSearch),
        ]).pipe(
            takeUntil(this.destroy$),
        ).subscribe(([page, search]) => this.getPlayers({ page, search }));
    }

    public readonly setLoaded = this.updater<boolean>((state, loaded) => ({
        ...state,
        loaded,
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

    private readonly selectPage = createSelector(
        selectQueryParam('page'),
        (page) => Number(page) > 0 ? Number(page) : 1
    );
    private readonly selectSearch = createSelector(
        selectQueryParam('search'),
        search => search ? String(search) : ''
    );

    public readonly loaded$ = this.select(state => state.loaded);
    public readonly count$ = this.select(state => state.count);
    public readonly players$ = this.select(state => state.players);
    public readonly vm$ = this.select(
        this.loaded$,
        this.store.select(this.selectPage),
        this.store.select(this.selectSearch),
        this.count$,
        this.players$,
        (loaded, page, search, count, players) => ({
            loaded,
            page,
            search,
            count,
            players,
        })
    );

    public readonly getPlayers = this.effect<PlayerFindOptions>(options$ => options$.pipe(
        switchMap(({ page, search }) => forkJoin([
            this.playersService.find({ page, search }),
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
