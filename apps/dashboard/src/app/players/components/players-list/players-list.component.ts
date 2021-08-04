import { Input } from '@angular/core';
import { Component } from '@angular/core';

import { IPlayer } from '@pool-overlay/models';

@Component({
    selector: 'pool-overlay-players-list',
    templateUrl: './players-list.component.html',
})
export class PlayersListComponent {
    @Input()
    public loaded = false;

    @Input()
    public players: IPlayer[] = [];
}
