import { Component } from '@angular/core';

import { GamesListStore } from './games-list.store';

@Component({
    selector: 'pool-overlay-games-list-page',
    templateUrl: './games-list-page.component.html',
    providers: [GamesListStore],
})
export class GamesListPageComponent {
    public readonly vm$ = this.gamesListStore.vm$;

    constructor(
        private gamesListStore: GamesListStore,
    ) { }

    public ngOnInit(): void {
        this.gamesListStore.getGames(1);
    }
}
