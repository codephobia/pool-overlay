import { Component, Input } from '@angular/core';

import { flagTrigger } from './flag.animation';

type PlayerClass = 'player-one' | 'player-two';

interface Trigger {
    value: string | null;
    params: {
        offsetEnter: string;
        offsetLeave: string;
    };
}

@Component({
    selector: 'app-flag',
    templateUrl: './flag.component.html',
    styleUrls: ['./flag.component.scss'],
    animations: [flagTrigger],
})
export class FlagComponent {
    @Input()
    public imagePath: string | null = null;

    @Input()
    public class: PlayerClass = 'player-one';

    public get trigger(): Trigger {
        const isPlayerOne = this.class === 'player-one';
        const offsetEnter = isPlayerOne ? 'calc(100% - 20px)' : 'calc(-100% + 20px)';
        const offsetLeave = isPlayerOne ? 'calc(-100% + 20px)' : 'calc(100% - 20px)';

        return {
            value: this.imagePath,
            params: {
                offsetEnter,
                offsetLeave,
            },
        };
    }
}
