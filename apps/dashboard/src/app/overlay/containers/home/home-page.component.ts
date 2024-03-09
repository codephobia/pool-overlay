import { Component } from '@angular/core';
import { faPlus } from '@fortawesome/pro-regular-svg-icons';
import { Store } from '@ngrx/store';
import * as fromTables from '../../../core/tables';
import { map } from 'rxjs';

@Component({
    selector: 'pool-overlay-home-page',
    templateUrl: './home-page.component.html',
})
export class HomePageComponent {
    public faPlus = faPlus;
    public currentTable = 1;
    public tables$ = this.store.select(fromTables.selectTablesCount).pipe(
        map((count) => Array.from(new Array(count), (x, i) => i + 1)),
    );

    constructor(private store: Store) { }

    public setTable(table: number): void {
        this.currentTable = table;
    }
}
