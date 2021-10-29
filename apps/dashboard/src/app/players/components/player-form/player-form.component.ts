import { EventEmitter, Input, Output, Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { faSyncAlt } from '@fortawesome/free-solid-svg-icons';

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
        this.form.controls.fargo_id.patchValue(this._player.fargo_id);
        this.form.controls.fargo_rating.patchValue(this._player.fargo_rating);
    }
    public get player(): IPlayer | null {
        return this._player;
    }

    @Input()
    public flags: IFlag[] = [];

    @Output()
    public onSubmit = new EventEmitter<{ player: IPlayer | Omit<IPlayer, 'id'> }>();

    public faSyncAlt = faSyncAlt;
    public form: FormGroup;
    public isUpdatingFargoRating = false;
    private _player: IPlayer | null = null;
    private fargoURL = 'https://dashboard.fargorate.com/api/indexsearch?q=';

    constructor(private _fb: FormBuilder) {
        this.form = this._fb.group({
            name: ['', Validators.required],
            flag_id: [1, Validators.required],
            fargo_id: [0, Validators.required],
            fargo_rating: [0, Validators.required],
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

    public async updateFargoRating(): Promise<void> {
        const fargoID = this.form.controls.fargo_id.value;
        if (!fargoID) {
            return;
        }

        this.isUpdatingFargoRating = true;

        try {
            const response = await fetch(this.fargoURL + fargoID);
            const body = await response.json();
            const rating = parseInt(body?.value?.[0]?.effectiveRating);

            if (rating) {
                this.form.controls.fargo_rating.patchValue(rating);
            }
        } catch (err) {
            console.error(err);
        }

        this.isUpdatingFargoRating = false;
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
