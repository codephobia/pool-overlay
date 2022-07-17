import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

import { GameType } from '@pool-overlay/models';
import { ControllerStore } from './controller.store';
import { Direction } from '../../models/direction.model';
import { PlayerModalComponent, PlayerModalData } from '../player-modal/player-modal.component';

@Component({
    selector: 'pool-overlay-controller',
    templateUrl: './controller.component.html',
    providers: [ControllerStore],
})
export class ControllerComponent implements OnInit {
    public readonly vm$ = this.store.vm$;

    constructor(
        private dialog: MatDialog,
        private readonly store: ControllerStore,
    ) { }

    public ngOnInit(): void {
        this.store.getGame();
    }

    public updatePlayer(playerNum: number, playerId: number | undefined): void {
        const dialog = this.dialog.open(PlayerModalComponent, {
            width: '700px',
            height: '726px',
            panelClass: 'no-padding',
            data: {
                playerNum,
                playerId,
            },
        });

        dialog.afterClosed().subscribe((data: PlayerModalData | undefined) => {
            if (!data) {
                return;
            }

            this.store.setPlayer({ playerNum: data.playerNum, playerID: data.playerId });
        });
    }

    public get increment(): Direction {
        return Direction.INCREMENT;
    }

    public get decrement(): Direction {
        return Direction.DECREMENT;
    }

    public get eightBall(): GameType {
        return GameType.EightBall;
    }

    public get nineBall(): GameType {
        return GameType.NineBall;
    }

    public get tenBall(): GameType {
        return GameType.TenBall;
    }

    public toggleOverlay(): void {
        this.store.toggleOverlay();
    }

    public toggleFlags(): void {
        this.store.toggleFlags();
    }

    public toggleFargo(): void {
        this.store.toggleFargo();
    }

    public toggleScore(): void {
        this.store.toggleScore();
    }

    public updateScore(playerNum: number, direction: Direction): void {
        this.store.updateScore({ playerNum, direction });
    }

    public resetScore(): void {
        this.store.resetScore();
    }

    public updateGameType(type: GameType): void {
        this.store.updateGameType(type);
    }

    public updateRaceTo(direction: Direction): void {
        this.store.updateRaceTo(direction);
    }

    public toggleFargoHotHandicap(use: boolean): void {
        this.store.toggleFargoHotHandicap(use);
    }

    public save(): void {
        this.store.save();
    }
}
