import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { concatMap, switchMap, tap, withLatestFrom } from 'rxjs/operators';

import { GameService } from '../../services/game.service';
import { OverlayStateService } from '../../services/overlay-state.service';
import { GameType, IGame, OverlayState } from '@pool-overlay/models';
import { Direction } from '../../models/direction.model';

export interface ControllerState {
    pending: boolean;
    game: IGame | null;
    overlay: OverlayState;
}

@Injectable()
export class ControllerStore extends ComponentStore<ControllerState> {
    constructor(
        private gameService: GameService,
        private overlayStateService: OverlayStateService,
    ) {
        super({
            pending: false,
            game: null,
            overlay: {
                table: 1,
                hidden: false,
                showFlags: false,
                showFargo: true,
                showScore: true,
            },
        });
    }

    public readonly setPending = this.updater<boolean>((state, pending) => ({
        ...state,
        pending
    }));

    public readonly setGame = this.updater<IGame>((state, game) => ({
        ...state,
        game,
    }));

    public readonly setRaceTo = this.updater<{ race_to: number, use_fargo_hot_handicap: boolean }>((state, { race_to, use_fargo_hot_handicap }) => ({
        ...state,
        game: {
            ...(state.game as IGame),
            race_to,
            use_fargo_hot_handicap,
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

    public readonly setFargoHotHandicap = this.updater<boolean>((state, use_fargo_hot_handicap) => ({
        ...state,
        game: {
            ...(state.game as IGame),
            use_fargo_hot_handicap,
        },
    }));

    public readonly setOverlay = this.updater<OverlayState>((state, overlay) => ({
        ...state,
        overlay,
    }));

    public readonly setHidden = this.updater<boolean>((state, hidden) => ({
        ...state,
        overlay: {
            ...(state.overlay as OverlayState),
            hidden,
        },
    }));

    public readonly setShowFlags = this.updater<boolean>((state, showFlags) => ({
        ...state,
        overlay: {
            ...(state.overlay as OverlayState),
            showFlags,
        },
    }));

    public readonly setShowFargo = this.updater<boolean>((state, showFargo) => ({
        ...state,
        overlay: {
            ...(state.overlay as OverlayState),
            showFargo,
        },
    }));

    public readonly setShowScore = this.updater<boolean>((state, showScore) => ({
        ...state,
        overlay: {
            ...(state.overlay as OverlayState),
            showScore,
        },
    }));

    public readonly setOverlayTable = this.updater<number>((state, table) => ({
        ...state,
        overlay: {
            ...(state.overlay as OverlayState),
            table,
        },
    }));

    public readonly pending$ = this.select(state => state.pending);
    public readonly game$ = this.select(state => state.game);
    public readonly overlay$ = this.select(state => state.overlay);
    public readonly vm$ = this.select(
        this.pending$,
        this.game$,
        this.overlay$,
        (pending, game, overlay) => ({
            pending,
            game,
            overlay,
        }),
    );

    public readonly getGame = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([_, table]) => this.gameService.getGame(table).pipe(
            tapResponse(
                game => { this.setGame(game); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly unsetPlayer = this.effect<number>(playerNum$ => playerNum$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([playerNum, table]) => this.gameService.unsetPlayer(table, playerNum).pipe(
            tapResponse(
                game => { this.setGame(game); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly setPlayer = this.effect<{ playerNum: number, playerID: number }>(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([{ playerNum, playerID }, table]) => this.gameService.setPlayer(table, playerNum, playerID).pipe(
            tapResponse(
                game => { this.setGame(game); },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly updateRaceTo = this.effect<Direction>(direction$ => direction$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([direction, table]) => this.gameService.updateRaceTo(table, direction).pipe(
            tapResponse(
                ({ raceTo, useFargoHotHandicap }) => {
                    this.setRaceTo({
                        race_to: raceTo,
                        use_fargo_hot_handicap: useFargoHotHandicap
                    });
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly resetScore = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([_, table]) => this.gameService.resetScore(table).pipe(
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
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([{ playerNum, direction }, table]) => this.gameService.updateScore(table, playerNum, direction).pipe(
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
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([gameType, table]) => this.gameService.setGameType(table, gameType).pipe(
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
        withLatestFrom(this.select(state => state.overlay)),
        switchMap(([_, overlay]) => this.overlayStateService.toggle(overlay.table).pipe(
            tapResponse(
                ({ hidden }) => {
                    this.setHidden(hidden);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly toggleFlags = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay)),
        switchMap(([_, overlay]) => this.overlayStateService.toggleFlags(overlay.table).pipe(
            tapResponse(
                ({ showFlags }) => {
                    this.setShowFlags(showFlags);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly toggleFargo = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay)),
        switchMap(([_, overlay]) => this.overlayStateService.toggleFargo(overlay.table).pipe(
            tapResponse(
                ({ showFargo }) => {
                    this.setShowFargo(showFargo);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly toggleScore = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay)),
        switchMap(([_, overlay]) => this.overlayStateService.toggleScore(overlay.table).pipe(
            tapResponse(
                ({ showScore }) => {
                    this.setShowScore(showScore);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly toggleFargoHotHandicap = this.effect<boolean>(useFargoHotHandicap$ => useFargoHotHandicap$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        switchMap(([useFargoHotHandicap, table]) => this.gameService.setFargoHotHandicap(table, useFargoHotHandicap).pipe(
            tapResponse(
                ({ useFargoHotHandicap }) => {
                    this.setFargoHotHandicap(useFargoHotHandicap);
                },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        ))
    ));

    public readonly save = this.effect(trigger$ => trigger$.pipe(
        tap(() => { this.setPending(true); }),
        withLatestFrom(this.select(state => state.overlay.table)),
        concatMap(([_, table]) => this.gameService.save(table).pipe(
            tapResponse(
                () => { },
                error => console.error(error),
                () => { this.setPending(false); },
            )
        )),
    ));
}
