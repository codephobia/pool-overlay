import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { OutletComponent } from '../shared/components/outlet/outlet.component';
import { PlayerCreatePageComponent } from './containers/player-create/player-create-page.component';
import { PlayerEditPageComponent } from './containers/player-edit/player-edit-page.component';
import { PlayersListPageComponent } from './containers/players-list-page/players-list-page.component';

const routes: Routes = [
    {
        path: '',
        component: OutletComponent,
        children: [
            {
                path: '',
                pathMatch: 'full',
                component: PlayersListPageComponent,
            },
            {
                path: 'create',
                component: PlayerCreatePageComponent,
            },
            {
                path: 'edit/:playerId',
                component: PlayerEditPageComponent,
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
export class PlayersRoutingModule { }
