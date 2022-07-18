import { Component, EventEmitter, Input, Output } from '@angular/core';

import { ToggleStore } from './toggle.store';

const TOGGLE_ID = 'dashboard_toggle_';
let TOGGLE_ID_COUNTER = 0;

@Component({
    selector: 'dashboard-toggle',
    templateUrl: 'toggle.component.html',
    providers: [ToggleStore],
})
export class ToggleComponent {
    @Input()
    public label = '';

    @Input()
    public set checked(value: boolean) {
        console.log(`input: ${value}`);
        this.store.setChecked(value);
    }

    @Output()
    public toggled = new EventEmitter<void>();

    public vm$ = this.store.vm$;
    public id: string;

    constructor(
        private readonly store: ToggleStore,
    ) {
        this.id = `${TOGGLE_ID}${++TOGGLE_ID_COUNTER}`;
    }

    public toggle(): void {
        this.store.toggle();
        this.toggled.emit();
    }
}
