import { Component, OnInit } from '@angular/core';
import { PlayersListStore } from './players-list.store';

@Component({
    selector: 'pool-overlay-players-list-page',
    templateUrl: './players-list-page.component.html',
})
export class PlayersListPageComponent implements OnInit {
    public readonly vm$ = this.playersListStore.vm$;

    constructor(
        private playersListStore: PlayersListStore,
    ) { }

    public ngOnInit(): void {
        this.playersListStore.getPlayers();
    }
}
