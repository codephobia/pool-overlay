import { Component, NgZone, OnInit } from '@angular/core';
import { IGame } from '@pool-overlay/models';
import { PlayerUpdateEvent } from '@pool-overlay/score-keeper';
import { filter, map, Observable } from 'rxjs';
import { ScoreKeeperStore } from '../../services/score-keeper.store';
import { SocketService } from '../../services/socket.service';

@Component({
    moduleId: module.id,
    selector: 'app-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
})
export class HomeComponent implements OnInit {
    public table = 2;
    public vm$ = this._store.vm$;

    constructor(
        private _store: ScoreKeeperStore,
        private _socketService: SocketService,
        private ngZone: NgZone,
    ) {
        this._socketService.bind('GAME_EVENT', this.updateGame.bind(this));
        this._socketService.connect();
    }

    public ngOnInit(): void {
        this._store.getGame();
    }

    public get showRaceTo$(): Observable<boolean> {
        return this._store.game$.pipe(
            filter(Boolean),
            map(game => {
                if (!game.use_fargo_hot_handicap) {
                    return false;
                }

                const { race_to, wins_higher, wins_lower } = game.fargo_hot_handicap;

                return wins_higher !== race_to || wins_lower !== race_to;
            }),
        );
    }

    public get playerOneRaceTo$(): Observable<number> {
        return this._store.game$.pipe(
            filter(Boolean),
            map(game => {
                if (!game.use_fargo_hot_handicap) {
                    return game.race_to;
                }

                const playerOneFargo = game.player_one?.fargo_rating ?? 0;
                const playerTwoFargo = game.player_two?.fargo_rating ?? 0;
                const winsHigher = game.fargo_hot_handicap?.wins_higher ?? 0;
                const winsLower = game.fargo_hot_handicap?.wins_lower ?? 0;

                return playerOneFargo > playerTwoFargo ? winsHigher : winsLower;
            }),
        );
    }

    public get playerTwoRaceTo$(): Observable<number> {
        return this._store.game$.pipe(
            filter(Boolean),
            map(game => {
                if (!game.use_fargo_hot_handicap) {
                    return game.race_to;
                }

                const playerOneFargo = game.player_one?.fargo_rating ?? 0;
                const playerTwoFargo = game.player_two?.fargo_rating ?? 0;
                const winsHigher = game.fargo_hot_handicap?.wins_higher ?? 0;
                const winsLower = game.fargo_hot_handicap?.wins_lower ?? 0;

                return playerTwoFargo > playerOneFargo ? winsHigher : winsLower;
            }),
        );
    }

    public updateGame(res: { game: IGame }) {
        this.ngZone.run(() => {
            if (res.game.table === this.table) {
                this._store.setGame(res);
            }
        });
    }

    public updatePlayer(update: PlayerUpdateEvent) {
        this._store.updateScore(update);
    }
}
