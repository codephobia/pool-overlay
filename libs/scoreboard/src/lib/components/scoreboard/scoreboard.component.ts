import { Component, OnInit } from '@angular/core';

import { GameType, gameTypeToString, IGame } from '@pool-overlay/models';
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
    public readonly vm$ = this._scoreboardStore.vm$;

    constructor(
        private _scoreboardStore: ScoreboardStore,
        private _socketService: SocketService,
    ) {
        this._socketService.bind('GAME_EVENT', this._scoreboardStore.setGame);
        this._socketService.bind('OVERLAY_TOGGLE', this._scoreboardStore.setHidden);
        this._socketService.bind('OVERLAY_TOGGLE_FLAGS', this._scoreboardStore.setShowFlags);
        this._socketService.bind('OVERLAY_TOGGLE_FARGO', this._scoreboardStore.setShowFargo);
        this._socketService.bind('OVERLAY_TOGGLE_SCORE', this._scoreboardStore.setShowScore);
        this._socketService.connect();
    }

    public ngOnInit(): void {
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
