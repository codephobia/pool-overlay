import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';

export interface PaginationState {
    page: number;
    count: number;
    perPage: number;
}

@Injectable()
export class PaginationStore extends ComponentStore<PaginationState> {
    constructor() {
        super({
            page: 1,
            count: 0,
            perPage: 10,
        });
    }

    public readonly setPage = this.updater<number>((state, page) => ({
        ...state,
        page,
    }));

    public readonly setCount = this.updater<number>((state, count) => ({
        ...state,
        count,
    }));

    public readonly setPerPage = this.updater<number>((state, perPage) => ({
        ...state,
        perPage,
    }));

    public readonly page$ = this.select(state => state.page);
    public readonly count$ = this.select(state => state.count);
    public readonly perPage$ = this.select(state => state.perPage);
    public readonly totalPages$ = this.select(
        this.count$,
        this.perPage$,
        (count, perPage) => Math.ceil(count / perPage)
    );
    public readonly hasPrevPage$ = this.select(state => state.page > 1);
    public readonly hasNextPage$ = this.select(
        this.page$,
        this.totalPages$,
        (page, totalPages) => page < totalPages
    );
    public readonly prevPage$ = this.select(({ page }) => (page > 1) ? page - 1 : page);
    public readonly nextPage$ = this.select(
        this.page$,
        this.totalPages$,
        (page, totalPages) => (page < totalPages) ? page + 1 : page
    );
    public readonly vm$ = this.select(
        this.hasPrevPage$,
        this.hasNextPage$,
        this.prevPage$,
        this.nextPage$,
        (hasPrevPage, hasNextPage, prevPage, nextPage) => ({
            hasPrevPage,
            hasNextPage,
            prevPage,
            nextPage,
        })
    );
}
