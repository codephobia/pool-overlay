import { Component } from '@angular/core';
import { GameType } from '@pool-overlay/models';
import { TournamentLoadedStore } from './tournament-loaded.store';

@Component({
    selector: 'pool-overlay-tournament-loaded',
    templateUrl: './tournament-loaded.component.html',
    providers: [TournamentLoadedStore],
})
export class TournamentLoadedComponent {
    readonly gameType = GameType;
    readonly vm$ = this.store.vm$;

    constructor(
        private store: TournamentLoadedStore,
    ) {
        this.store.getCurrentTournament();
    }

    public unloadTournament(): void {
        this.store.unloadTournament();
    }
}
