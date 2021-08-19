import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

import { PlayerModalStore } from './player-modal.store';

export interface PlayerModalData {
    playerNum: number;
    playerId: number;
}

@Component({
    selector: 'pool-overlay-player-modal',
    templateUrl: './player-modal.component.html',
    providers: [PlayerModalStore],
})
export class PlayerModalComponent implements OnInit {
    public playerNum: number;
    public currentPlayerId: number | undefined;
    public vm$ = this.store.vm$;

    constructor(
        private dialogRef: MatDialogRef<PlayerModalComponent, PlayerModalData>,
        @Inject(MAT_DIALOG_DATA) data: PlayerModalData,
        private readonly store: PlayerModalStore,
    ) {
        this.playerNum = data.playerNum;
        this.currentPlayerId = data.playerId;
    }

    public ngOnInit(): void {
        this.store.getPlayers(1);
    }

    public bgColor(playerId: number, isEven: boolean): string {
        if (playerId === this.currentPlayerId) {
            return 'bg-blue-700 hover:bg-blue-600';
        } else if (isEven) {
            return 'bg-gray-500 hover:bg-gray-400';
        }
        return 'bg-gray-700 hover:bg-gray-600';
    }

    public selectPlayer(playerId: number): void {
        this.dialogRef.close({
            playerNum: this.playerNum,
            playerId,
        });
    }

    public close(): void {
        this.dialogRef.close();
    }
}
