import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Direction } from '@pool-overlay/models';
import { PlayerUpdateEvent } from './player.types';

@Component({
    moduleId: module.id,
    selector: 'app-player',
    templateUrl: './player.component.html',
    styleUrls: ['./player.component.scss'],
})
export class PlayerComponent {
    @Input()
    public playerNum = 1;

    @Input()
    public name = '';

    @Input()
    public score = 0;

    @Input()
    public raceTo = 0;

    @Input()
    public showRaceTo = false;

    @Input()
    public color: 'blue' | 'orange' = 'blue';

    @Output()
    public update = new EventEmitter<PlayerUpdateEvent>();

    public get increment(): Direction {
        return Direction.INCREMENT;
    }

    public get decrement(): Direction {
        return Direction.DECREMENT;
    }

    public updateScore(direction: Direction) {
        const directionNum = direction === Direction.INCREMENT ? 1 : -1;

        if (this.score + directionNum >= 0 && this.score + directionNum <= this.raceTo) {
            this.update.emit({
                playerNum: this.playerNum,
                direction,
            });
        }
    }
}
