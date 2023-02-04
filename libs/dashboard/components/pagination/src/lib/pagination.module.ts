import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { PaginationComponent } from './pagination.component';

const COMPONENTS = [
    PaginationComponent,
];

@NgModule({
    imports: [CommonModule, RouterModule, FontAwesomeModule],
    declarations: [...COMPONENTS],
    exports: [...COMPONENTS],
})
export class PaginationModule { }
