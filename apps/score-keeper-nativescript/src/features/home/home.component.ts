import { Component, NgZone } from '@angular/core';
import { IGame, OverlayState } from '@pool-overlay/models';
import { PlayerUpdateEvent } from '@pool-overlay/score-keeper';
import { ScoreKeeperStore } from '../../services/score-keeper.store';
import { SocketService } from '../../services/socket.service';

@Component({
    moduleId: module.id,
    selector: 'app-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
})
export class HomeComponent {
    public vm$ = this._store.vm$;

    constructor(
        private _store: ScoreKeeperStore,
        private _socketService: SocketService,
        private ngZone: NgZone,
    ) {
        this._socketService.bind('GAME_EVENT', this.updateGame.bind(this));
        this._socketService.bind('OVERLAY_STATE_EVENT', this.updateOverlay.bind(this));
        this._socketService.connect();
        this._store.getGame();
    }

    public updateGame({ game }: { game: IGame }) {
        this.ngZone.run(() => {
            this._store.updateGame(game);
        });
    }

    public updateOverlay(res: OverlayState): void {
        this.ngZone.run(() => {
            this._store.updateOverlay(res);
        });
    }

    public updatePlayer(update: PlayerUpdateEvent) {
        this._store.updateScore(update);
    }
}
