import { Component, OnInit } from '@angular/core';

import { GameType, gameTypeToString } from '../../models/game-type.enum';
import { SocketService } from '../../services/socket.service';
import { ScoreboardStore } from './scoreboard.store';

@Component({
    selector: 'pool-overlay-scoreboard',
    templateUrl: './scoreboard.component.html',
    styleUrls: ['./scoreboard.component.scss'],
    providers: [ScoreboardStore, SocketService],
})
export class ScoreboardComponent implements OnInit {
    public readonly vm$ = this._scoreboardStore.vm$;

    constructor(
        private _scoreboardStore: ScoreboardStore,
        private _socketService: SocketService,
    ) {
        this._scoreboardStore.setState({
            game: null,
            hidden: false,
        });
        this._socketService.bind('GAME_EVENT', this._scoreboardStore.setGame);
        this._socketService.connect();
    }

    public ngOnInit(): void {
        this._scoreboardStore.getGame();
    }

    public gameTypeName(gameType: GameType | null): string {
        if (gameType === null) {
            return '';
        }

        return gameTypeToString(gameType);
    }
}
