import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { OutletComponent } from '../shared/components/outlet/outlet.component';
import { GamesListPageComponent } from './containers/games-list/games-list-page.component';

const routes: Routes = [
    {
        path: '',
        component: OutletComponent,
        children: [
            {
                path: '',
                pathMatch: 'full',
                component: GamesListPageComponent,
            },
        ],
    }
];

@NgModule({
    imports: [
        CommonModule,
        RouterModule.forChild(routes),
    ],
    exports: [RouterModule],
})
export class GamesRoutingModule { }
