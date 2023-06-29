import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { OutletComponent } from '../shared/components/outlet/outlet.component';
import { HomePageComponent } from './containers/home/home-page.component';
import { LoadedPageComponent } from './containers/loaded/loaded-page.component';
import { TournamentSetupPageComponent } from './containers/tournament-setup/tournament-setup-page.component';
import { tournamentLoadedGuard } from './guards/tournament-loaded.guard';
import { tournamentNotLoadedGuard } from './guards/tournament-not-loaded.guard';

const routes: Routes = [
    {
        path: '',
        component: OutletComponent,
        children: [
            {
                path: '',
                pathMatch: 'full',
                component: HomePageComponent,
                canActivate: [tournamentNotLoadedGuard('/tournaments/loaded')],
            },
            {
                path: 'loaded',
                component: LoadedPageComponent,
                canActivate: [tournamentLoadedGuard('/tournaments')],
            },
            {
                path: ':tournamentId',
                component: TournamentSetupPageComponent,
                canActivate: [tournamentNotLoadedGuard('/tournaments/loaded')],
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
export class TournamentsRoutingModule { }
