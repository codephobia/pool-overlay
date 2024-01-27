import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { GameType } from '@pool-overlay/models';
import { TournamentSetupStore } from './tournament-setup.store';

@Component({
    selector: 'pool-overlay-tournament-setup',
    templateUrl: './tournament-setup.component.html',
    providers: [TournamentSetupStore],
})
export class TournamentSetupComponent {
    readonly gameType = GameType;
    readonly vm$ = this.store.vm$;

    constructor(
        route: ActivatedRoute,
        private store: TournamentSetupStore,
    ) {
        const tournamentId = Number(route.snapshot.paramMap.get('tournamentId'));
        this.store.getTournamentById(tournamentId);
    }

    public updateMaxTables(maxTables: number): void {
        this.store.updateMaxTables(maxTables);
    }

    public updateGameType(gameType: GameType): void {
        this.store.updateGameType(gameType);
    }

    public updateShowOverlay(show: boolean): void {
        this.store.updateShowOverlay(show);
    }

    public updateShowFlags(show: boolean): void {
        this.store.updateShowFlags(show);
    }

    public updateShowFargo(show: boolean): void {
        this.store.updateShowFargo(show);
    }

    public updateShowScore(show: boolean): void {
        this.store.updateShowScore(show);
    }

    public updateIsHandicapped(isHandicapped: boolean): void {
        this.store.updateIsHandicapped(isHandicapped);
    }

    public updateASideRaceTo(raceTo: number): void {
        this.store.updateASideRaceTo(raceTo);
    }

    public updateBSideRaceTo(raceTo: number): void {
        this.store.updateBSideRaceTo(raceTo);
    }

    public loadTournament(): void {
        this.store.loadTournament();
    }
}
