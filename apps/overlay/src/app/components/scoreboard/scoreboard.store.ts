import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { EMPTY } from 'rxjs';
import { catchError, switchMap } from 'rxjs/operators';

import { GameType } from '../../models/game-type.enum';
import { IGame } from '../../models/game.model';
import { IPlayer } from '../../models/player.model';
import { IScore } from '../../models/score.model';
import { VsMode } from '../../models/vs-mode.enum';
import { APIService } from '../../services/api.service';

export interface ScoreboardState {
  game: IGame;
  hidden: boolean;
}

@Injectable()
export class ScoreboardStore extends ComponentStore<ScoreboardState> {
  // updaters

  readonly setGame = this.updater((state, game: IGame) => ({ ...state, game }));

  readonly setGameType = this.updater((state, type: GameType) => ({
    ...state,
    game: {
      ...state.game,
      type,
    },
  }));

  readonly setGameVsMode = this.updater((state, vs_mode: VsMode) => ({
    ...state,
    game: {
      ...state.game,
      vs_mode,
    },
  }));

  readonly setGameRaceTo = this.updater((state, race_to: number) => ({
    ...state,
    game: {
      ...state.game,
      race_to,
    },
  }));

  readonly setGameScore = this.updater((state, score: IScore) => ({
    ...state,
    game: {
      ...state.game,
      score,
    },
  }));

  readonly setGameScoreOne = this.updater((state, one: number) => ({
    ...state,
    game: {
      ...state.game,
      score: {
        ...state.game.score,
        one,
      },
    },
  }));

  readonly setGameScoreTwo = this.updater((state, two: number) => ({
    ...state,
    game: {
      ...state.game,
      score: {
        ...state.game.score,
        two,
      },
    },
  }));

  readonly setGamePlayerOne = this.updater((state, player_one: IPlayer) => ({
    ...state,
    game: {
      ...state.game,
      player_one,
    },
  }));

  readonly setGamePlayerTwo = this.updater((state, player_two: IPlayer) => ({
    ...state,
    game: {
      ...state.game,
      player_two,
    },
  }));

  readonly setHidden = this.updater((state, hidden: boolean) => ({
    ...state,
    hidden,
  }));

  // selectors
  readonly game$ = this.select((state) => state.game);

  readonly hidden$ = this.select((state) => state.hidden);

  readonly vm$ = this.select(this.game$, this.hidden$, (game, hidden) => ({
    game,
    hidden,
  }));

  // effects
  readonly getGame = this.effect((trigger$) => {
    return trigger$.pipe(
      switchMap(() =>
        this._apiService.getGame().pipe(
          tapResponse(
            (game) => {
              this.setGame(game);
              // TODO: REMOVE THIS.
              console.log(game);
            },
            (error) => console.error(error)
          ),
          catchError(() => EMPTY)
        )
      )
    );
  });

  constructor(private _apiService: APIService) {
    super({ game: null, hidden: false });
  }
}
