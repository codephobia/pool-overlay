import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Tournament } from '@pool-overlay/models';
import { TournamentListStore } from './tournament-list.store';

@Component({
    selector: 'pool-overlay-tournament-list',
    templateUrl: 'tournament-list.component.html',
    providers: [TournamentListStore],
})
export class TournamentListComponent implements OnInit {
    readonly vm$ = this.store.vm$;

    @Output()
    public selected = new EventEmitter<{ tournamentId: number }>();

    constructor(private store: TournamentListStore) { }

    public ngOnInit(): void {
        this.store.getTournaments();
    }

    public selectTournament(tournamentId: number): void {
        this.selected.emit({ tournamentId });
    }
}
