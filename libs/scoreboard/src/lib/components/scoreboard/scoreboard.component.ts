import { Component, Input, OnInit } from '@angular/core';

import { GameType, gameTypeToString, IGame, OverlayState } from '@pool-overlay/models';
import { SocketService } from '../../services/socket.service';
import { handicapTransition } from './handicap.animation';
import { scoreboardTransition } from './scoreboard.animation';
import { ScoreboardStore } from './scoreboard.store';

@Component({
    selector: 'pool-overlay-scoreboard',
    templateUrl: './scoreboard.component.html',
    styleUrls: ['./scoreboard.component.scss'],
    providers: [ScoreboardStore, SocketService],
    animations: [scoreboardTransition, handicapTransition],
})
export class ScoreboardComponent implements OnInit {
    @Input()
    public table: number = 1;

    public readonly vm$ = this._scoreboardStore.vm$;

    constructor(
        private _scoreboardStore: ScoreboardStore,
        private _socketService: SocketService,
    ) {

    }

    public ngOnInit(): void {
        this._socketService.bind('GAME_EVENT', (res: { game: IGame }) => {
            if (res.game.table === this.table) {
                this._scoreboardStore.setGame(res);
            }
        });
        this._socketService.bind('OVERLAY_STATE_EVENT', (res: OverlayState) => {
            if (res.table === this.table) {
                this._scoreboardStore.setOverlay(res);
            }
        });
        this._socketService.connect();

        this._scoreboardStore.setOverlayTable(this.table);
        this._scoreboardStore.getGame();
    }

    public gameTypeName(gameType: GameType | null | undefined): string {
        if (gameType === null || gameType === undefined) {
            return '';
        }

        return gameTypeToString(gameType);
    }

    public playerOneRaceTo(game: IGame | null): number {
        if (!game) {
            return 0;
        }

        const playerOneFargo = game?.player_one?.fargo_rating ?? 0;
        const playerTwoFargo = game?.player_two?.fargo_rating ?? 0;
        const winsHigher = game?.fargo_hot_handicap?.wins_higher ?? 0;
        const winsLower = game?.fargo_hot_handicap?.wins_lower ?? 0;

        return playerOneFargo > playerTwoFargo ? winsHigher : winsLower;
    }

    public playerTwoRaceTo(game: IGame | null): number {
        if (!game) {
            return 0;
        }

        const playerOneFargo = game?.player_one?.fargo_rating ?? 0;
        const playerTwoFargo = game?.player_two?.fargo_rating ?? 0;
        const winsHigher = game?.fargo_hot_handicap?.wins_higher ?? 0;
        const winsLower = game?.fargo_hot_handicap?.wins_lower ?? 0;

        return playerTwoFargo > playerOneFargo ? winsHigher : winsLower;
    }
}
