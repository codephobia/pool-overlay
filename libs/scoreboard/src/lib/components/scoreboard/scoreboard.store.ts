import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap } from 'rxjs/operators';

import { IGame } from '@pool-overlay/models';
import { APIService } from '../../services/api.service';

export interface ScoreboardState {
    game: IGame | null;
    hidden: boolean;
}

@Injectable()
export class ScoreboardStore extends ComponentStore<ScoreboardState> {
    constructor(private _apiService: APIService) {
        super({ game: null, hidden: false });
    }

    // updaters
    public readonly setGame = this.updater((state, values: { game: IGame }) => ({
        ...state,
        game: values.game
    }));

    public readonly setHidden = this.updater((state, values: { hidden: boolean }) => ({
        ...state,
        hidden: values.hidden,
    }));

    // selectors
    public readonly game$ = this.select((state) => state.game);
    public readonly hidden$ = this.select((state) => state.hidden);
    public readonly vm$ = this.select(this.game$, this.hidden$, (game, hidden) => ({
        game,
        hidden,
    }));

    // effects
    public readonly getGame = this.effect((trigger$) => {
        return trigger$.pipe(
            switchMap(() =>
                this._apiService.getGame().pipe(
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
}
