import { Component } from '@angular/core';
import { faPlus } from '@fortawesome/pro-regular-svg-icons';

@Component({
    selector: 'pool-overlay-home-page',
    templateUrl: './home-page.component.html',
})
export class HomePageComponent {
    public faPlus = faPlus;
    public tables: number[] = [1, 2, 3];
    public currentTable = 1;

    public setTable(table: number): void {
        this.currentTable = table;
    }
}
