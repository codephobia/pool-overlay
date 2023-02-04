import { Component, EventEmitter, Input, Output } from '@angular/core';
import { faChevronLeft, faChevronRight } from '@fortawesome/pro-regular-svg-icons';

import { PaginationStore } from './pagination.store';

export interface PageEvent {
    page: number;
}

@Component({
    selector: 'dashboard-pagination',
    templateUrl: './pagination.component.html',
    providers: [PaginationStore],
})
export class PaginationComponent {
    @Input()
    public set page(page: number) {
        this.paginationStore.setPage(page);
    }

    @Input()
    public set count(count: number) {
        this.paginationStore.setCount(count);
    }

    @Input()
    public set perPage(perPage: number) {
        this.paginationStore.setPerPage(perPage);
    }

    @Output()
    public pageChange = new EventEmitter<PageEvent>();

    public faChevronLeft = faChevronLeft;
    public faChevronRight = faChevronRight;
    public vm$ = this.paginationStore.vm$;

    constructor(
        private paginationStore: PaginationStore,
    ) { }

    public goToPage(page: number): void {
        this.pageChange.emit({ page });
    }
}
