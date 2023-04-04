import { Component } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ScoreKeeperStore } from '../../services/score-keeper.store';

@Component({
    moduleId: module.id,
    selector: 'app-settings',
    templateUrl: './settings.component.html',
    styleUrls: ['./settings.component.scss'],
})
export class SettingsComponent {
    public tables: number[] = [1, 2];
    public table$ = this._store.table$;
    public serverAddress$ = this._store.serverAddress$;
    public form: FormGroup;

    constructor(
        private _store: ScoreKeeperStore,
    ) {
        this.form = new FormGroup({})
    }

    public updateTable(table: number) {
        this._store.setTable(table);
    }

    public saveServerAddress(serverAddress: string) {
        console.log(serverAddress);
    }
}
