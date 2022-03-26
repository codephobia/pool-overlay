import { Component, Input } from '@angular/core';

import { IGame } from '@pool-overlay/models';

@Component({
    selector: 'pool-overlay-games-list',
    templateUrl: './games-list.component.html'
})
export class GamesListComponent {
    @Input()
    public loaded = false;

    @Input()
    public games: Partial<IGame>[] = [];
}
