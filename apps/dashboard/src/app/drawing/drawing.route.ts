import { Route } from '@angular/router';

export const drawingRoute: Route = {
    path: 'drawing',
    loadChildren: () => import('./drawing.module').then(m => m.DrawingModule),
};
