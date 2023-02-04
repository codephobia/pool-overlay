import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Component({
    selector: 'pool-overlay-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent {
    public table$: Observable<number>;

    constructor(
        private route: ActivatedRoute,
    ) {
        this.table$ = this.route.queryParams.pipe(
            map(params => Number(params.table))
        );
    }
}
