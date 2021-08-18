import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { OutletComponent } from './components/outlet/outlet.component';
import { ConfirmDialogComponent } from './components/confirm-dialog/confirm-dialog.component';

const COMPONENTS = [
    OutletComponent,
    ConfirmDialogComponent,
];

@NgModule({
    declarations: [
        ...COMPONENTS,
    ],
    imports: [
        CommonModule,
        RouterModule,
    ],
    exports: [
        ...COMPONENTS,
    ],
})
export class SharedModule { }
