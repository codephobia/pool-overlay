import { EventEmitter, Input, Output } from '@angular/core';
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

    @Output()
    public deletePlayer = new EventEmitter<{ playerId: number }>();

    public deletePlayerById(playerId: number | undefined): void {
        if (!playerId) {
            return;
        }

        this.deletePlayer.emit({ playerId });
    }
}
