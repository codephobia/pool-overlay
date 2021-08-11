import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { SharedModule } from '../shared/shared.module';
import { PlayersRoutingModule } from './players-routing.module';
import { PlayersListPageComponent } from './containers/players-list-page/players-list-page.component';
import { PlayersListComponent } from './components/players-list/players-list.component';
import { PlayersService } from './services/players.service';
import { PlayersListStore } from './containers/players-list-page/players-list.store';
import { PlayerCreatePageComponent } from './containers/player-create/player-create-page.component';
import { PlayerEditPageComponent } from './containers/player-edit/player-edit-page.component';
import { PlayerFormComponent } from './components/player-form/player-form.component';

const COMPONENTS = [
    PlayersListPageComponent,
    PlayersListComponent,
    PlayerCreatePageComponent,
    PlayerEditPageComponent,
    PlayerFormComponent,
];

const SERVICES = [
    PlayersService,
    PlayersListStore,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        ReactiveFormsModule,
        SharedModule,
        PlayersRoutingModule,
    ],
    exports: [],
    declarations: [
        ...COMPONENTS,
    ],
    providers: [
        ...SERVICES,
    ],
})
export class PlayersModule { }
