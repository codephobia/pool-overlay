import { HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { IFlag, IPlayer } from '@pool-overlay/models';
import { forkJoin } from 'rxjs';
import { exhaustMap, switchMap, tap } from 'rxjs/operators';
import { FlagsService } from './flags.service';
import { PlayersService } from './players.service';

export enum PlayerFormStatus {
    LOADING,
    LOADED,
    SUBMITTING,
    SUCCESS,
    ERROR,
}

export interface PlayerFormState {
    status: PlayerFormStatus;
    error: string | null;
    player: IPlayer | null;
    flags: IFlag[];
}

@Injectable()
export class PlayerFormStore extends ComponentStore<PlayerFormState> {
    constructor(
        private playersService: PlayersService,
        private flagsService: FlagsService,
    ) {
        super({
            status: PlayerFormStatus.LOADING,
            error: null,
            player: null,
            flags: [],
        });
    }

    public readonly setStatus = this.updater<PlayerFormStatus>((state, status) => ({
        ...state,
        status,
    }));

    public readonly setError = this.updater<string | null>((state, error) => ({
        ...state,
        error,
    }));

    public readonly setPlayer = this.updater<IPlayer | null>((state, player) => ({
        ...state,
        player,
    }));

    public readonly setFlags = this.updater<IFlag[]>((state, flags) => ({
        ...state,
        flags,
    }));

    public readonly status$ = this.select(state => state.status);
    public readonly error$ = this.select(state => state.error);
    public readonly player$ = this.select(state => state.player);
    public readonly flags$ = this.select(state => state.flags);
    public readonly vm$ = this.select(
        this.status$,
        this.error$,
        this.player$,
        this.flags$,
        (status, error, player, flags) => ({
            status,
            error,
            player,
            flags,
        })
    );

    public readonly getFlags = this.effect(trigger$ => trigger$.pipe(
        switchMap(() => this.flagsService.find().pipe(
            tapResponse(
                flags => {
                    this.setFlags(flags);
                    this.setStatus(PlayerFormStatus.LOADED);
                },
                (error: HttpErrorResponse) => {
                    this.setStatus(PlayerFormStatus.ERROR);
                    this.setError(error.message);
                    console.error(error);
                }
            )
        ))
    ));

    public readonly getPlayerAndFlags = this.effect<number>(playerId$ => playerId$.pipe(
        switchMap(playerId => forkJoin({
            player: this.playersService.findByID(playerId),
            flags: this.flagsService.find(),
        }).pipe(
            tapResponse(
                ({ player, flags }) => {
                    this.setPlayer(player);
                    this.setFlags(flags);
                    this.setStatus(PlayerFormStatus.LOADED);
                },
                (error: HttpErrorResponse) => {
                    this.setStatus(PlayerFormStatus.ERROR);
                    this.setError(error.message);
                    console.error(error);
                }
            )
        ))
    ));

    public readonly createPlayer = this.effect<Omit<IPlayer, 'id'>>(player$ => player$.pipe(
        tap(() => {
            this.setStatus(PlayerFormStatus.LOADED);
            this.setError(null);
        }),
        exhaustMap(player => this.playersService.create(player).pipe(
            tapResponse(
                () => {
                    this.setStatus(PlayerFormStatus.SUCCESS);
                },
                (error: HttpErrorResponse) => {
                    this.setStatus(PlayerFormStatus.ERROR);
                    this.setError(error.message);
                    console.error(error)
                }
            )
        ))
    ));

    public readonly updatePlayer = this.effect<IPlayer>(player$ => player$.pipe(
        tap(() => {
            this.setStatus(PlayerFormStatus.LOADED);
            this.setError(null);
        }),
        exhaustMap(player => this.playersService.update(player).pipe(
            tapResponse(
                () => {
                    this.setStatus(PlayerFormStatus.SUCCESS);
                },
                (error: HttpErrorResponse) => {
                    this.setStatus(PlayerFormStatus.ERROR);
                    this.setError(error.message);
                    console.error(error)
                }
            )
        ))
    ));
}
