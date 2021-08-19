import { Route } from '@angular/router';

export const overlayRoute: Route = {
    path: 'overlay',
    loadChildren: () => import('./overlay.module').then(m => m.OverlayModule),
};
