import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap } from 'rxjs/operators';

import { IGame } from '@pool-overlay/models';
import { GamesService } from '../../services/games.service';

interface GamesListState {
    loaded: boolean;
    page: number;
    games: Partial<IGame>[];
}

@Injectable()
export class GamesListStore extends ComponentStore<GamesListState> {
    constructor(
        private gamesService: GamesService,
    ) {
        super({
            loaded: false,
            page: 1,
            games: [],
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

    public readonly setGames = this.updater<Partial<IGame>[]>((state, games) => ({
        ...state,
        loaded: true,
        games,
    }));

    public readonly loaded$ = this.select(state => state.loaded);
    public readonly page$ = this.select(state => state.page);
    public readonly games$ = this.select(state => state.games);
    public readonly vm$ = this.select(
        this.loaded$,
        this.page$,
        this.games$,
        (loaded, page, games) => ({
            loaded,
            page,
            games,
        })
    );

    public readonly getGames = this.effect<number>(page$ => page$.pipe(
        switchMap(page => this.gamesService.find(page).pipe(
            tapResponse(
                games => {
                    this.setGames(games)
                },
                error => console.error(error)
            ),
        )),
    ));
}
