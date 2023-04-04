import { EventEmitter, Input, Output } from '@angular/core';
import { Component } from '@angular/core';
import { MatLegacyDialog as MatDialog } from '@angular/material/legacy-dialog';

import { IPlayer } from '@pool-overlay/models';
import { ConfirmDialogComponent } from '../../../shared/components/confirm-dialog/confirm-dialog.component';

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

    constructor(
        private dialog: MatDialog,
    ) { }

    public deletePlayerById(player: IPlayer): void {
        const dialogRef = this.dialog.open(ConfirmDialogComponent, {
            width: '400px',
            panelClass: 'no-padding',
            data: {
                player,
            },
        });

        dialogRef.afterClosed().subscribe(confirmation => {
            if (confirmation) {
                this.deletePlayer.emit({ playerId: player.id });
            }
        });
    }
}
