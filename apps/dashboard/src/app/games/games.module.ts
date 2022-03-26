import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { GamesRoutingModule } from './games-routing.module';
import { GamesService } from './services/games.service';
import { GamesListPageComponent } from './containers/games-list/games-list-page.component';
import { GamesListComponent } from './components/games-list/games-list.component';

const COMPONENTS = [
    GamesListPageComponent,
    GamesListComponent,
];

const SERVICES = [
    GamesService,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        GamesRoutingModule,
    ],
    exports: [],
    declarations: [
        ...COMPONENTS,
    ],
    providers: [
        ...SERVICES,
    ],
})
export class GamesModule { }
