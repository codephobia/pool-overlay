import { Component, Input } from '@angular/core';

import { characterRotater } from './character-rotater.animation';

@Component({
    selector: 'app-character-rotater',
    templateUrl: './character-rotater.component.html',
    styleUrls: ['./character-rotater.component.scss'],
    animations: [characterRotater],
})
export class CharacterRotaterComponent {
    @Input()
    public set characters(characters: string | number | null) {
        if (characters === null) {
            return;
        }

        this._characters = characters;
    }
    public get characters(): string | number {
        return this._characters;
    }

    private _characters: string | number = '';

    public get charArr(): string[] {
        return String(this.characters).split('');
    }

    public trackBy(index: number, value: string): string {
        return `${index}-${value}`;
    }
}
