import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap } from 'rxjs/operators';

import { IGame } from '@pool-overlay/models';
import { APIService } from '../../services/api.service';

export interface ScoreboardState {
    game: IGame | null;
    hidden: boolean;
    showFlags: boolean;
    showFargo: boolean;
    showScore: boolean;
}

@Injectable()
export class ScoreboardStore extends ComponentStore<ScoreboardState> {
    constructor(private _apiService: APIService) {
        super({
            game: null,
            hidden: false,
            showFlags: true,
            showFargo: true,
            showScore: true,
        });
    }

    // updaters
    public readonly setGame = this.updater<Pick<ScoreboardState, 'game'>>((state, { game }) => ({
        ...state,
        game,
    }));

    public readonly setHidden = this.updater<Pick<ScoreboardState, 'hidden'>>((state, { hidden }) => ({
        ...state,
        hidden,
    }));

    public readonly setShowFlags = this.updater<Pick<ScoreboardState, 'showFlags'>>((state, { showFlags }) => ({
        ...state,
        showFlags,
    }));

    public readonly setShowFargo = this.updater<Pick<ScoreboardState, 'showFargo'>>((state, { showFargo }) => ({
        ...state,
        showFargo,
    }));

    public readonly setShowScore = this.updater<Pick<ScoreboardState, 'showScore'>>((state, { showScore }) => ({
        ...state,
        showScore,
    }));

    // selectors
    public readonly game$ = this.select((state) => state.game);
    public readonly hidden$ = this.select((state) => state.hidden);
    public readonly showFlags$ = this.select((state) => state.showFlags);
    public readonly showFargo$ = this.select((state) => state.showFargo);
    public readonly showScore$ = this.select((state) => state.showScore);
    public readonly vm$ = this.select(
        this.game$,
        this.hidden$,
        this.showFlags$,
        this.showFargo$,
        this.showScore$,
        (game, hidden, showFlags, showFargo, showScore) => ({
            game,
            hidden,
            showFlags,
            showFargo,
            showScore,
        })
    );

    // effects
    public readonly getGame = this.effect((trigger$) => {
        return trigger$.pipe(
            switchMap(() =>
                this._apiService.getGame().pipe(
                    tapResponse(
                        (game) => {
                            this.setGame({ game });
                            console.log(game);
                        },
                        (error) => console.error(error)
                    ),
                )
            )
        );
    });
}
