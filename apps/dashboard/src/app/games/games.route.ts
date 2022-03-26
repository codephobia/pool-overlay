import { Route } from '@angular/router';

export const gamesRoute: Route = {
    path: 'games',
    loadChildren: () => import('./games.module').then(m => m.GamesModule),
};
