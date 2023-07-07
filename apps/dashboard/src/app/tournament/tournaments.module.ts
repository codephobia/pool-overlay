import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { CdkMenuModule } from '@angular/cdk/menu';

import { TournamentsService } from './services/tournament.service';
import { TablesService } from './services/tables.service';
import { HomePageComponent } from './containers/home/home-page.component';
import { ToggleModule } from '@dashboard/components/toggle';
import { TournamentsRoutingModule } from './tournaments-routing.module';
import { TournamentListComponent } from './components/tournament-list/tournament-list.component';
import { TournamentSetupPageComponent } from './containers/tournament-setup/tournament-setup-page.component';
import { TournamentSetupComponent } from './components/tournament-setup/tournament-setup.component';
import { LoadedPageComponent } from './containers/loaded/loaded-page.component';
import { TournamentLoadedComponent } from './components/tournament-loaded/tournament-loaded.component';

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
    TablesService,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        ReactiveFormsModule,
        FontAwesomeModule,
        CdkMenuModule,
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
