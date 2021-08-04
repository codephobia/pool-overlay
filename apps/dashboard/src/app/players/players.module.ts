import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { SharedModule } from '../shared/shared.module';
import { PlayersRoutingModule } from './players-routing.module';
import { PlayersListPageComponent } from './containers/players-list-page/players-list-page.component';
import { PlayersListComponent } from './components/players-list/players-list.component';
import { PlayersService } from './services/players.service';
import { PlayersListStore } from './containers/players-list-page/players-list.store';

const COMPONENTS = [
    PlayersListPageComponent,
    PlayersListComponent,
];

const SERVICES = [
    PlayersService,
    PlayersListStore,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
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
