import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { HomePageComponent } from './containers/home/home-page.component';
import { ToggleModule } from '@dashboard/components/toggle';
import { TournamentsRoutingModule } from './tournaments-routing.module';
import { TournamentListComponent } from './components/tournament-list/tournament-list.component';
import { TournamentsService } from './services/tournament.service';
import { TournamentSetupPageComponent } from './containers/tournament-setup/tournament-setup-page.component';
import { TournamentSetupComponent } from './components/tournament-setup/tournament-setup.component';
import { LoadedPageComponent } from './containers/loaded/loaded-page.component';
import { TournamentLoadedComponent } from './components/tournament-loaded/tournament-loaded.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

const COMPONENTS = [
    HomePageComponent,
    TournamentListComponent,
    TournamentSetupPageComponent,
    TournamentSetupComponent,
    LoadedPageComponent,
    TournamentLoadedComponent,
];

const SERVICES = [
    TournamentsService,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        ReactiveFormsModule,
        FontAwesomeModule,
        TournamentsRoutingModule,
        ToggleModule,
    ],
    exports: [],
    declarations: [
        ...COMPONENTS,
    ],
    providers: [
        ...SERVICES,
    ],
})
export class TournamentsModule { }
