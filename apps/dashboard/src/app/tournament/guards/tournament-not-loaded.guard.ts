import { inject } from '@angular/core';
import { CanActivateFn, Router, UrlTree } from '@angular/router';
import { Observable, of } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { TournamentsService } from '../services/tournament.service';

export function tournamentNotLoadedGuard(redirectUrl: string): CanActivateFn {
    return (): Observable<boolean | UrlTree> => {
        const router = inject(Router);
        const tournamentsService = inject(TournamentsService);

        return tournamentsService.getCurrent().pipe(
            map(() => router.createUrlTree([redirectUrl])),
            catchError(() => {
                return of(true);
            }),
        );
    };
};
