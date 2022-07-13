import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { PageEvent } from '@dashboard/components/pagination';
import { faXmark } from '@fortawesome/pro-regular-svg-icons';

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
    public faXmark = faXmark;
    public readonly perPage = 10;
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
            return 'bg-sad-background-active';
        } else if (isEven) {
            return 'bg-sad-table-even hover:bg-sad-background-active';
        }
        return 'bg-sad-table-odd hover:bg-sad-background-active';
    }

    public onPageChange({ page }: PageEvent): void {
        this.store.getPlayers(page);
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
