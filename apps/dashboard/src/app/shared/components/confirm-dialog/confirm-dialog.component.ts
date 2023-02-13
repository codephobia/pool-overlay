import { Component, Inject } from '@angular/core';
import { MatLegacyDialogRef as MatDialogRef, MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA } from '@angular/material/legacy-dialog';

import { IPlayer } from '@pool-overlay/models';

export interface IConfirmDialogResponse {
    confirmed: boolean;
}

@Component({
    selector: 'pool-overlay-confirm-dialog',
    templateUrl: './confirm-dialog.component.html',
})
export class ConfirmDialogComponent {
    public playerName = '';

    constructor(
        public dialogRef: MatDialogRef<boolean>,
        @Inject(MAT_DIALOG_DATA) data: { player: IPlayer },
    ) {
        this.playerName = data.player.name;
    }

    public close(confirmation: boolean): void {
        this.dialogRef.close(confirmation);
    }
}
