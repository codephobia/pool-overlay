import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { OutletComponent } from './components/outlet/outlet.component';

const COMPONENTS = [
    OutletComponent,
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
