import { Component, OnInit } from '@angular/core';

import { GameType, gameTypeToString } from '@pool-overlay/models';
import { SocketService } from '../../services/socket.service';
import { scoreboardTransition } from './scoreboard.animation';
import { ScoreboardStore } from './scoreboard.store';

@Component({
    selector: 'pool-overlay-scoreboard',
    templateUrl: './scoreboard.component.html',
    styleUrls: ['./scoreboard.component.scss'],
    providers: [ScoreboardStore, SocketService],
    animations: [scoreboardTransition],
})
export class ScoreboardComponent implements OnInit {
    public readonly vm$ = this._scoreboardStore.vm$;

    constructor(
        private _scoreboardStore: ScoreboardStore,
        private _socketService: SocketService,
    ) {
        this._socketService.bind('GAME_EVENT', this._scoreboardStore.setGame);
        this._socketService.bind('OVERLAY_TOGGLE', this._scoreboardStore.setHidden);
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
}
