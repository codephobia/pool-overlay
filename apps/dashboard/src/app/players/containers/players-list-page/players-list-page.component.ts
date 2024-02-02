import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { PageEvent } from '@dashboard/components/pagination';
import { PlayersListStore } from './players-list.store';
import { SearchEvent } from '@dashboard/components/search';

@Component({
    selector: 'pool-overlay-players-list-page',
    templateUrl: './players-list-page.component.html',
})
export class PlayersListPageComponent {
    public readonly perPage = 10;
    public readonly vm$ = this.playersListStore.vm$;

    constructor(
        private router: Router,
        private activatedRoute: ActivatedRoute,
        private playersListStore: PlayersListStore,
    ) { }

    public onPageChange({ page }: PageEvent): void {
        const newPage = page > 1 ? page : undefined;

        this.router.navigate(['.'], {
            relativeTo: this.activatedRoute,
            queryParams: { page: newPage },
            queryParamsHandling: 'merge',
        });
    }

    public onSearch({ search }: SearchEvent): void {
        const newSearch = search?.length ? search : undefined;

        this.router.navigate(['.'], {
            relativeTo: this.activatedRoute,
            queryParams: { search: newSearch },
        });
    }

    public deletePlayerById({ playerId }: { playerId: number }): void {
        this.playersListStore.deletePlayerById(playerId);
    }
}
