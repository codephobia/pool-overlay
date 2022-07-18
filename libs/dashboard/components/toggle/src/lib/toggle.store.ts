import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';
import { withLatestFrom, tap } from 'rxjs/operators';

export interface ToggleState {
    checked: boolean;
}

@Injectable()
export class ToggleStore extends ComponentStore<ToggleState> {
    constructor() {
        super({
            checked: false,
        });
    }

    public readonly setChecked = this.updater<boolean>((state, checked) => ({
        ...state,
        checked,
    }));

    public readonly checked$ = this.select(state => state.checked);
    public readonly vm$ = this.select(
        this.checked$,
        checked => ({ checked })
    );

    public readonly toggle = this.effect(trigger$ => trigger$.pipe(
        withLatestFrom(this.select(state => state.checked)),
        tap(([_, checked]) => this.setChecked(!checked))
    ));
}
