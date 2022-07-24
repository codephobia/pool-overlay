import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap } from 'rxjs/operators';

import { IGame, OverlayState } from '@pool-overlay/models';
import { APIService } from '../../services/api.service';

export interface ScoreboardState {
    game: IGame | null;
    overlay: OverlayState;
}

@Injectable()
export class ScoreboardStore extends ComponentStore<ScoreboardState> {
    constructor(private _apiService: APIService) {
        super({
            game: null,
            overlay: {
                hidden: false,
                showFlags: true,
                showFargo: true,
                showScore: true,
            },
        });
    }

    // updaters
    public readonly setGame = this.updater<Pick<ScoreboardState, 'game'>>((state, { game }) => ({
        ...state,
        game,
    }));

    public readonly setOverlay = this.updater<OverlayState>((state, overlay) => ({
        ...state,
        overlay,
    }));

    // selectors
    public readonly game$ = this.select((state) => state.game);
    public readonly overlay$ = this.select((state) => state.overlay);
    public readonly vm$ = this.select(
        this.game$,
        this.overlay$,
        (game, overlay) => ({
            game,
            overlay,
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
                        },
                        (error) => console.error(error)
                    ),
                )
            )
        );
    });
}
