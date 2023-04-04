import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { NativeScriptCommonModule } from '@nativescript/angular';
import { RaceToComponent } from './race-to.component';

const COMPONENTS = [
    RaceToComponent,
];

@NgModule({
    imports: [NativeScriptCommonModule],
    declarations: [...COMPONENTS],
    exports: [...COMPONENTS],
    schemas: [NO_ERRORS_SCHEMA],
})
export class RaceToModule { }
