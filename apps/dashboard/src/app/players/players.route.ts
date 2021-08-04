import { Route } from '@angular/router';

export const playersRoute: Route = {
    path: 'players',
    loadChildren: () => import('./players.module').then(m => m.PlayersModule),
};
