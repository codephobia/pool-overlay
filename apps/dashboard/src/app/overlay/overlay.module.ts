import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { MatLegacyDialogModule as MatDialogModule } from '@angular/material/legacy-dialog';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { environment } from '../../environments/environment';
import { SharedModule } from '../shared/shared.module';
import { ScoreboardModule } from '@pool-overlay/scoreboard';
import { OverlayRoutingModule } from './overlay-routing.module';
import { GameService } from './services/game.service';
import { OverlayStateService } from './services/overlay-state.service';
import { HomePageComponent } from './containers/home/home-page.component';
import { ControllerComponent } from './components/controller/controller.component';
import { PlayerModalComponent } from './components/player-modal/player-modal.component';
import { PaginationModule } from '@dashboard/components/pagination';
import { ToggleModule } from '@dashboard/components/toggle';

const COMPONENTS = [
    HomePageComponent,
    ControllerComponent,
    PlayerModalComponent,
];

const SERVICES = [
    GameService,
    OverlayStateService,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        ReactiveFormsModule,
        MatDialogModule,
        FontAwesomeModule,
        SharedModule,
        OverlayRoutingModule,
        ScoreboardModule.withConfig({
            environment,
        }),
        PaginationModule,
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
export class OverlayModule { }
