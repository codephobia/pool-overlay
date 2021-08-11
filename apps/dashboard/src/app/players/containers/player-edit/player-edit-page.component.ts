import { Component } from '@angular/core';

import { IPlayer } from '@pool-overlay/models';

@Component({
    selector: 'pool-overlay-player-edit-page',
    templateUrl: './player-edit-page.component.html',
})
export class PlayerEditPageComponent {
    public onSubmit({ player }: { player: IPlayer | Omit<IPlayer, 'id'> }): void {
        console.log({ player });
    }
}
