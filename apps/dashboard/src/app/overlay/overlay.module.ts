import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';

import { environment } from '../../environments/environment';
import { SharedModule } from '../shared/shared.module';
import { ScoreboardModule } from '@pool-overlay/scoreboard';
import { OverlayRoutingModule } from './overlay-routing.module';
import { GameService } from './services/game.service';
import { HomePageComponent } from './containers/home/home-page.component';
import { ControllerComponent } from './components/controller/controller.component';
import { PlayerModalComponent } from './components/player-modal/player-modal.component';

const COMPONENTS = [
    HomePageComponent,
    ControllerComponent,
    PlayerModalComponent,
];

const SERVICES = [
    GameService,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        ReactiveFormsModule,
        MatDialogModule,
        SharedModule,
        OverlayRoutingModule,
        ScoreboardModule.withConfig({
            environment,
        }),
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
