import { Component, OnInit } from '@angular/core';

import { IPlayer } from '@pool-overlay/models';
import { PlayerFormStore } from '../../services/player-form.store';

@Component({
    selector: 'pool-overlay-player-create-page',
    templateUrl: './player-create-page.component.html',
    providers: [PlayerFormStore]
})
export class PlayerCreatePageComponent implements OnInit {
    public readonly vm$ = this.playerFormStore.vm$;

    constructor(
        private readonly playerFormStore: PlayerFormStore,
    ) { }

    public ngOnInit(): void {
        this.playerFormStore.getFlags();
    }

    public onSubmit(event: { player: IPlayer | Omit<IPlayer, 'id'> }): void {
        this.playerFormStore.createPlayer(event.player as Omit<IPlayer, 'id'>);
    }
}
