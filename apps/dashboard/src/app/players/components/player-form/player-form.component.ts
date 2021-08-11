import { EventEmitter, Input, Output, Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { IFlag, IPlayer } from '@pool-overlay/models';

@Component({
    selector: 'pool-overlay-player-form',
    templateUrl: './player-form.component.html',
})
export class PlayerFormComponent {
    @Input()
    public set player(player: IPlayer | null) {
        if (!player) {
            return;
        }

        this._player = player;
        this.form.addControl('id', this._fb.control(this._player.id, Validators.required));
        this.form.controls.name.patchValue(this._player.name);
        this.form.controls.flag_id.patchValue(this._player.flag_id);
    }
    public get player(): IPlayer | null {
        return this._player;
    }

    @Input()
    public flags: IFlag[] = [
        {
            id: 1,
            country: 'USA',
            image_path: 'us.png',
        },
        {
            id: 2,
            country: 'Canada',
            image_path: 'ca.png',
        },
    ];

    @Output()
    public onSubmit = new EventEmitter<{ player: IPlayer | Omit<IPlayer, 'id'> }>();

    public form: FormGroup;
    private _player: IPlayer | null = null;

    constructor(private _fb: FormBuilder) {
        this.form = this._fb.group({
            name: ['', Validators.required],
            flag_id: [1, Validators.required],
        });
    }

    public get isCreating(): boolean {
        return !this.form.contains('id');
    }

    public get currentFlag(): number {
        return this.form.controls.flag_id.value;
    }

    public setFlag(id: number): void {
        this.form.controls.flag_id.patchValue(id);
    }

    public get canSubmit(): boolean {
        return this.form.valid;
    }

    public submit(): void {
        if (!this.canSubmit) {
            return;
        }

        this.onSubmit.emit({ player: this.form.value });
    }
}
