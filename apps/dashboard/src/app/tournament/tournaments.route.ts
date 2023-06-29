import { Route } from '@angular/router';

export const tournamentsRoute: Route = {
    path: 'tournaments',
    loadChildren: () => import('./tournaments.module').then(m => m.TournamentsModule),
};
