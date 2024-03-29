import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

import { overlayRoute } from './overlay/overlay.route';
import { tournamentsRoute } from './tournament/tournaments.route';
import { playersRoute } from './players/players.route';
import { gamesRoute } from './games/games.route';
import { drawingRoute } from './drawing/drawing.route';

const routes: Routes = [
    overlayRoute,
    tournamentsRoute,
    playersRoute,
    gamesRoute,
    drawingRoute,
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {
            preloadingStrategy: PreloadAllModules,
            initialNavigation: 'enabledBlocking',
        }),
    ],
    exports: [RouterModule],
})
export class AppRoutingModule { }
