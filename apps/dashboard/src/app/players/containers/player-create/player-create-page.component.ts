import { Component } from '@angular/core';

import { IPlayer } from '@pool-overlay/models';

@Component({
    selector: 'pool-overlay-player-create-page',
    templateUrl: './player-create-page.component.html',
})
export class PlayerCreatePageComponent {
    public onSubmit(event: { player: IPlayer | Omit<IPlayer, 'id'> }): void {
        console.log(event);
    }
}
