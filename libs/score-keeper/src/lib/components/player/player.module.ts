import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { NativeScriptCommonModule } from '@nativescript/angular';
import { PlayerComponent } from './player.component';

const COMPONENTS = [
    PlayerComponent,
];

@NgModule({
    imports: [NativeScriptCommonModule],
    declarations: [...COMPONENTS],
    exports: [...COMPONENTS],
    schemas: [NO_ERRORS_SCHEMA],
})
export class PlayerModule { }
