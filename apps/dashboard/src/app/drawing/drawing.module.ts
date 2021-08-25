import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { SharedModule } from '../shared/shared.module';
import { DrawingRoutingModule } from './drawing-routing.module';
import { HomePageComponent } from './containers/home/home-page.component';
import { ScreenComponent } from './components/screen/screen.component';
import { CanvasComponent } from './components/canvas/canvas.component';
import { DrawingDirective } from './directives/drawing.directive';

const COMPONENTS = [
    HomePageComponent,
    ScreenComponent,
    CanvasComponent,
];

const DIRECTIVES = [
    DrawingDirective,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        SharedModule,
        DrawingRoutingModule,
    ],
    exports: [],
    declarations: [
        ...COMPONENTS,
        ...DIRECTIVES,
    ],
    providers: [],
})
export class DrawingModule { }
