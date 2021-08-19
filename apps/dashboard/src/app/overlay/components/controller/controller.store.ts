import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { switchMap, tap } from 'rxjs/operators';

import { GameService } from '../../services/game.service';
import { GameType, IGame } from '@pool-overlay/models';
import { Direction } from '../../models/direction.model';

export interface ControllerState {
    pending: boolean;
    game: IGame | null;
    hidden: boolean;
}

@Injectable()
export class ControllerStore extends ComponentStore<ControllerState> {
    constructor(
        private gameService: GameService,
    ) {
        super({ pending: false, game: null, hidden: false });
    }

    public readonly setPending = this.updater<boolean>((state, pending) => ({
        ...state,
        pending
    }));

    public readonly setGame = this.updater<IGame>((state, game) => ({
        ...state,
        game,
    }));

    public readonly setRaceTo = this.updater<number>((state, race_to) => ({
        ...state,
        game: {
            ...(state.game as IGame),
            race_to,
        },
    }));

    public readonly setScore = this.updater<{ score_one: number, score_two: number }>((state, { score_one, score_two }) => ({
        ...state,
        game: {
            ...(state.game as IGame),
            score_one,
            score_two,
        },
    }));

    public readonly setGameType = this.updater<GameType>((state, type) => ({
        ...state,
        game: {
            ...(state.game as IGame),
            type,
        },
    }));

    public readonly setHidden = this.updater<boolean>((state, hidden) => ({
        ...state,
        hidden,
    }));

    public readonly pending$ = this.select(state => state.pending);
    public readonly game$ = this.select(state => state.game);
    public readonly hidden$ = this.select(state => state.hidden);
    public readonly vm$ = this.select(
        this.pending$,
        this.game$,
        this.hidden$,
        (pending, game, hidden) => ({
            pending,
            game,
            hidden,
        }),
    );

    public readonly getGame = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(() => this.gameService.getGame().pipe(
            tapResponse(
                game => { this.setGame(game); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly unsetPlayer = this.effect<number>(playerNum$ => playerNum$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(playerNum => this.gameService.unsetPlayer(playerNum).pipe(
            tapResponse(
                game => { this.setGame(game); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly setPlayer = this.effect<{ playerNum: number, playerID: number }>(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(({ playerNum, playerID }) => this.gameService.setPlayer(playerNum, playerID).pipe(
            tapResponse(
                game => { this.setGame(game); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly updateRaceTo = this.effect<Direction>(direction$ => direction$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(direction => this.gameService.updateRaceTo(direction).pipe(
            tapResponse(
                ({ raceTo }) => { this.setRaceTo(raceTo); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly resetScore = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(() => this.gameService.resetScore().pipe(
            tapResponse(
                ({ scoreOne, scoreTwo }) => {
                    this.setScore({
                        score_one: scoreOne,
                        score_two: scoreTwo,
                    });
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly updateScore = this.effect<{ playerNum: number, direction: Direction }>(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(({ playerNum, direction }) => this.gameService.updateScore(playerNum, direction).pipe(
            tapResponse(
                ({ scoreOne, scoreTwo }) => {
                    this.setScore({
                        score_one: scoreOne,
                        score_two: scoreTwo,
                    });
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly updateGameType = this.effect<GameType>(gameType$ => gameType$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(gameType => this.gameService.setGameType(gameType).pipe(
            tapResponse(
                ({ type }) => {
                    this.setGameType(type);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly toggleOverlay = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        switchMap(() => this.gameService.toggleOverlay().pipe(
            tapResponse(
                ({ hidden }) => {
                    this.setHidden(hidden);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));
}
