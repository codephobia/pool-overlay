import { Component } from '@angular/core';
import { faClock, faEllipsisVertical, faLock } from '@fortawesome/pro-regular-svg-icons';
import { GameType, IPlayer, FargoHotHandicap } from '@pool-overlay/models';
import { SocketService } from '@pool-overlay/scoreboard';
import { filter, map, Observable } from 'rxjs';
import { TournamentLoadedStore } from './tournament-loaded.store';

@Component({
    selector: 'pool-overlay-tournament-loaded',
    templateUrl: './tournament-loaded.component.html',
    providers: [TournamentLoadedStore, SocketService],
})
export class TournamentLoadedComponent {
    readonly faClock = faClock;
    readonly faLock = faLock;
    readonly faEllipsisVertical = faEllipsisVertical
    readonly gameType = GameType;
    readonly tables = [1, 2, 3];
    readonly vm$ = this.store.vm$;

    constructor(
        private readonly socketService: SocketService,
        private store: TournamentLoadedStore,
    ) {
        this.store.getCurrentTournament();
        this.socketService.bind('GAME_EVENT', this.store.setGame.bind(this));
        this.socketService.bind('OVERLAY_STATE_EVENT', this.store.setOverlay.bind(this));
        this.socketService.connect();
    }

    public playerOneRaceTo$(table: number): Observable<number> {
        return this.vm$.pipe(
            map((vm) => vm.tables[table].game),
            filter(Boolean),
            filter((game) => !!game.player_one && !!game.player_two && !!game.fargo_hot_handicap),
            map((game) => {
                return (game.player_one as IPlayer).fargo_rating > (game.player_two as IPlayer).fargo_rating ? (game.fargo_hot_handicap as FargoHotHandicap).wins_higher : (game.fargo_hot_handicap as FargoHotHandicap).wins_lower
            }),
        );
    }

    public playerTwoRaceTo$(table: number): Observable<number> {
        return this.vm$.pipe(
            map((vm) => vm.tables[table].game),
            filter(Boolean),
            filter((game) => !!game.player_one && !!game.player_two && !!game.fargo_hot_handicap),
            map((game) => {
                return (game.player_one as IPlayer).fargo_rating < (game.player_two as IPlayer).fargo_rating ? (game.fargo_hot_handicap as FargoHotHandicap).wins_higher : (game.fargo_hot_handicap as FargoHotHandicap).wins_lower
            }),
        );
    }

    public swapTables(tableOne: number, tableTwo: number): void {
        this.store.swapTables({ tableOne, tableTwo });
    }

    public unloadTournament(): void {
        this.store.unloadTournament();
    }
}
