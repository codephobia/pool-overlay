import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

import { overlayRoute } from './overlay/overlay.route';
import { playersRoute } from './players/players.route';

const routes: Routes = [
    overlayRoute,
    playersRoute,
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
