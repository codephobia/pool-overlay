import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { OutletComponent } from './components/outlet/outlet.component';
import { ConfirmDialogComponent } from './components/confirm-dialog/confirm-dialog.component';
import { PlayersService } from './services/players.service';

const COMPONENTS = [
    OutletComponent,
    ConfirmDialogComponent,
];

const SERVICES = [
    PlayersService,
];

@NgModule({
    declarations: [
        ...COMPONENTS,
    ],
    imports: [
        CommonModule,
        RouterModule,
    ],
    providers: [
        ...SERVICES,
    ],
    exports: [
        ...COMPONENTS,
    ],
})
export class SharedModule { }
